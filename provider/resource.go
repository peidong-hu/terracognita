package provider

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/chr4/pwgen"
	"github.com/cycloidio/terracognita/errcode"
	"github.com/cycloidio/terracognita/filter"
	"github.com/cycloidio/terracognita/tag"
	"github.com/cycloidio/terracognita/util"
	"github.com/cycloidio/terracognita/writer"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=../mock/resource.go -mock_names=Resource=Resource -package mock github.com/cycloidio/terracognita/provider Resource

// Resource represents the minimal information needed to
// define a Provider resource
type Resource interface {
	// ID is the ID of the Resource
	ID() string

	// Type is the type of resource (ex: aws_instance)
	Type() string

	// TFResource is the definition of that resource
	TFResource() *schema.Resource

	// Data is the actual data of the Resource
	Data() *schema.ResourceData

	// Provider is the Provider of that Resource
	Provider() Provider

	// Read read the remote information of the Resource
	Read(f *filter.Filter) error

	// State calculates the state of the Resource and
	// writes it to w
	State(w writer.Writer) error

	// HCL returns the HCL configuration of the Resource and
	// writes it to HCL
	HCL(w writer.Writer) error
}

// resources is a general implementation of Resource interface
// that fulfills all the usecases for Terracognita
type resource struct {
	id string

	// resourceType is the type of resource (ex: aws_instance)
	// as type is a reserved word resourceType is used on private
	resourceType string

	tfResource *schema.Resource

	data *schema.ResourceData

	provider Provider

	// The name it has on the config
	// so it can be the same on the HCL
	// and State
	configName string
}

var (
	autogeneratedAWSResourcesRe = regexp.MustCompile(`^aws:(?:autoscaling|cloudformation)`)
)

// NewResource returns an implementation of the Resource
func NewResource(id, tp string, tfResource *schema.Resource, data *schema.ResourceData, p Provider) Resource {
	return &resource{
		id:           id,
		resourceType: tp,
		tfResource:   tfResource,
		data:         data,
		provider:     p,
	}
}

func (r *resource) ID() string { return r.id }

func (r *resource) Type() string { return r.resourceType }

func (r *resource) TFResource() *schema.Resource { return r.tfResource }

func (r *resource) Data() *schema.ResourceData { return r.data }

func (r *resource) Provider() Provider { return r.provider }

func (r *resource) Read(f *filter.Filter) error {
	// Retry if any error happen
	err := util.RetryDefault(func() error {
		return r.tfResource.Read(r.data, r.provider.TFClient())
	})
	if err != nil {
		return errors.Wrapf(err, "while reading on type %q", r.resourceType)
	}

	// TODO: Extreme case, it should be on an "AfterRead" function
	// but for now we'll do it like this
	if r.resourceType == "aws_iam_user_group_membership" {
		gps := r.data.Get("groups").(*schema.Set).List()
		var gpsKey string
		for _, gp := range gps {
			if gpsKey == "" {
				gpsKey = gp.(string)
			}
			gpsKey = fmt.Sprintf("%s/%s", gpsKey, gp)
		}
		if len(gps) != 0 {
			r.data.SetId(fmt.Sprintf("%s/%s", r.data.Get("user"), gpsKey))
		}
	}

	// For some reason it failed to fetch the Resource, it should not be an error
	// because it could be an account related resource that it's not delcared or
	// is default.
	// Therefore we just continue
	if r.data.Id() == "" {
		return errors.Wrapf(errcode.ErrProviderResourceNotRead, "the resource %q with ID %q did not return an ID", r.resourceType, r.id)
	}

	// Some resources can not be filtered by tags,
	// so we have to do it manually
	// it's not all of them though
	for _, t := range f.Tags {
		if v, ok := r.data.GetOk(fmt.Sprintf("%s.%s", r.Provider().TagKey(), t.Name)); ok && v.(string) != t.Value {
			return errors.WithStack(errcode.ErrProviderResourceDoNotMatchTag)
		}
	}

	// Filter out autogenerated resources from AWS
	if v, ok := r.data.GetOk(r.Provider().TagKey()); ok {
		for k := range v.(map[string]interface{}) {
			if autogeneratedAWSResourcesRe.MatchString(k) {
				return errors.WithStack(errcode.ErrProviderResourceAutogenerated)
			}
		}
	}

	return nil
}

// State calculates the state of the Resource and
// writes it to w
func (r *resource) State(w writer.Writer) error {

	if importer := r.tfResource.Importer; importer != nil {
		resourceDatas, err := importer.State(r.data, r.provider.TFClient())
		if err != nil {
			return err
		}

		// TODO: The multple return could potentially be the `depends_on` of the
		// terraform.ResourceState
		// Investigate on SG
		for i, rd := range resourceDatas {
			if i != 0 {
				// for now we scape all the other ones
				// the firs one is the one we need and
				// for what've see the others are
				// 'aws_security_group_rules' (in aws)
				// and are not imported by default by
				// Terraform
				continue
			}

			tis := rd.State()
			if tis == nil {
				// IDK why some times it does not have
				// the ID (the only case in tis it's nil)
				// so if nil we don't need it
				continue
			}

			trs := &terraform.ResourceState{
				Type:     r.resourceType,
				Primary:  tis,
				Provider: r.Provider().String(),
			}

			// If it does not have any configName we will generate one
			// and store it, so net time it'll use that one on any config
			if r.configName == "" {
				configName := tag.GetNameFromTag(r.provider.TagKey(), rd, r.id)
				err := w.Write(fmt.Sprintf("%s.%s", tis.Ephemeral.Type, configName), trs)
				if err != nil {
					if errors.Cause(err) == errcode.ErrWriterAlreadyExistsKey {
						configName = pwgen.Alpha(5)
						err = w.Write(fmt.Sprintf("%s.%s", tis.Ephemeral.Type, configName), trs)
						if err != nil {
							return err
						}
						r.configName = configName
						return nil
					}
					return err
				}
				r.configName = configName
			} else {
				err := w.Write(fmt.Sprintf("%s.%s", tis.Ephemeral.Type, r.configName), trs)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}
	return nil
}

// HCL returns the HCL configuration of the Resource and
// writes it to HCL
func (r *resource) HCL(w writer.Writer) error {
	cfg := mergeFullConfig(r.data, r.tfResource.Schema, "")

	if r.configName == "" {
		configName := tag.GetNameFromTag(r.provider.TagKey(), r.data, r.id)
		err := w.Write(fmt.Sprintf("%s.%s", r.resourceType, configName), cfg)
		if err != nil {
			if errors.Cause(err) == errcode.ErrWriterAlreadyExistsKey {
				configName = pwgen.Alpha(5)
				err = w.Write(fmt.Sprintf("%s.%s", r.resourceType, configName), cfg)
				if err != nil {
					return err
				}
				r.configName = configName
				return nil
			}
			return err
		}
		r.configName = configName
	} else {
		err := w.Write(fmt.Sprintf("%s.%s", r.resourceType, r.configName), cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

// mergeFullConfig creates the key to the map and if it had a value before set it, if
func mergeFullConfig(cfgr *schema.ResourceData, sch map[string]*schema.Schema, key string) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range sch {
		// If it's just a Computed value, do not add it to the output
		if !isConfig(v) {
			continue
		}

		// Basically calculates the needed
		// key to the current access
		var kk string
		if key != "" {
			kk = key + "." + k
		} else {
			kk = k
		}

		// schema.Resource means that it has nested fields
		if sr, ok := v.Elem.(*schema.Resource); ok {
			// Example would be aws_security_group
			if v.Type == schema.TypeSet {
				s, ok := cfgr.GetOk(kk)
				if !ok {
					continue
				}

				res[k] = normalizeSetList(sr.Schema, s.(*schema.Set).List())
			} else if v.Type == schema.TypeList {
				var ar interface{} = make([]interface{}, 0)

				l, ok := cfgr.GetOk(kk)
				if !ok {
					continue
				}

				list := l.([]interface{})
				for i := range list {
					ar = append(ar.([]interface{}), mergeFullConfig(cfgr, sr.Schema, fmt.Sprintf("%s.%d", kk, i)))
				}

				res[k] = ar
			} else {
				res[k] = mergeFullConfig(cfgr, sr.Schema, kk)
			}
			// As it's a nested element it does not require any of
			// the other code as it's for singel value schemas
			continue
		}

		// This sets the singel values that we see on the
		// end result

		vv, ok := cfgr.GetOk(kk)
		// If the value is Required we need to add it
		// even if it's not send
		if (!ok || vv == nil) && !v.Required {
			continue
		}

		// A value in which this one conflicts has been set before
		// so we should no set this one as it'll raise an error of
		// `conflicts with *` on Terraform
		if hasConflict(res, v.ConflictsWith) {
			continue
		}

		if s, ok := vv.(*schema.Set); ok {
			res[k] = s.List()
		} else {
			res[k] = normalizeInterpolation(normalizeValue(vv))
		}
	}
	return res
}

// hasConflict checks if any of the keys is present on the res
func hasConflict(res map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		if _, ok := res[key]; ok {
			return true
		}
	}
	return false
}

// normalizeValue removes the \n from the value now
func normalizeValue(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		return strings.Replace(s, "\n", "", -1)
	}
	return v
}

var iamInternpolationRe = regexp.MustCompile(`(\$\{[^}]+\})`)

// normalizeInterpolation fixes the https://github.com/hashicorp/terraform/issues/18937
// on reading
func normalizeInterpolation(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		return iamInternpolationRe.ReplaceAllString(s, `$$$1`)
	}
	return v
}

// normalizeSetList returns the normalization of a schema.Set.List
// it could be a simple list or a embedded structure.
// The sch it's used to also add required values if needed
func normalizeSetList(sch map[string]*schema.Schema, list []interface{}) interface{} {
	var ar interface{} = make([]interface{}, 0)

	for _, set := range list {
		switch val := set.(type) {
		case map[string]interface{}:
			// This case it's when a TypeSet has
			// a nested structure,
			// example: aws_security_group.ingress
			res := make(map[string]interface{})
			for k, v := range val {
				switch vv := v.(type) {
				case *schema.Set:
					nsch := make(map[string]*schema.Schema)
					if sc, ok := sch[k]; ok {
						if rs, ok := sc.Elem.(*schema.Resource); ok {
							nsch = rs.Schema
						}
					}
					ns := normalizeSetList(nsch, vv.List())
					if !isDefault(sch[k], ns) {
						res[k] = ns
					}
				case []interface{}:
					nsch := make(map[string]*schema.Schema)
					if sc, ok := sch[k]; ok {
						if rs, ok := sc.Elem.(*schema.Resource); ok {
							nsch = rs.Schema
						}
					}
					ns := normalizeSetList(nsch, vv)
					if !isDefault(sch[k], ns) {
						res[k] = ns
					}
				case interface{}:
					if !isDefault(sch[k], v) {
						res[k] = v
					}
				}
			}
			ar = append(ar.([]interface{}), res)
		case []interface{}:
			ns := normalizeSetList(sch, val)
			if !isDefault(nil, ns) {
				ar = append(ar.([]interface{}), ns)
			}
		case interface{}:
			// This case is normally for the
			// "Type: schema.TypeSet, Elm: schema.Schema{Type: schema.TypeString}"
			// definitions on TF,
			// example: aws_security_group.ingress.security_groups
			if !isDefault(nil, val) {
				ar = append(ar.([]interface{}), val)
			}
		}
	}

	return ar
}

var (
	// Ideally this could be generated using "enumer", it
	// would be a better idea as then we do not have
	// to maintain this list
	tfTypes = []schema.ValueType{
		schema.TypeBool,
		schema.TypeInt,
		schema.TypeFloat,
		schema.TypeString,
		schema.TypeList,
		schema.TypeMap,
		schema.TypeSet,
	}
)

// isDefault is used on normalizSet as the Sets do not use the normal
// TF strucure (access by key) and are stored as raw maps with some
// default values that we don't want on the HCL output.
// example: [], false, "", 0 ...
func isDefault(sch *schema.Schema, v interface{}) bool {
	if sch != nil {
		if sch.Required {
			return false
		}

		// This way values that are not suppose
		// to be on the config are also not added
		if !isConfig(sch) {
			return true
		}
	}

	for _, t := range tfTypes {
		if reflect.DeepEqual(t.Zero(), v) {
			// If it has a default value which is different
			// than the one v then v has to be setted.
			// Example: Default => true, v => false
			// the v = false has to be setted
			if sch.Default != nil {
				if v != sch.Default {
					return false
				}
			}
			return true
		}
	}
	return false
}

// isConfig  checks if the sch has to be
// set to a config opt or not
func isConfig(sch *schema.Schema) bool {
	if (sch.Computed && !sch.Optional && !sch.Required) || sch.Deprecated != "" {
		return false
	}
	return true
}
