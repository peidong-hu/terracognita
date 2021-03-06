// Code generated by "enumer -type ResourceType -addprefix aws_ -transform snake -linecomment"; DO NOT EDIT.

package aws

import (
	"fmt"
)

const _ResourceTypeName = "aws_instanceaws_vpcaws_vpc_peering_connectionaws_key_pairaws_security_groupaws_subnetaws_ebs_volumeaws_elasticache_clusteraws_elbaws_albaws_alb_listeneraws_alb_listener_ruleaws_alb_listener_certificateaws_alb_target_groupaws_lbaws_lb_listeneraws_lb_listener_ruleaws_lb_listener_certificateaws_lb_target_groupaws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_s3_bucketaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_zoneaws_route53_zone_associationaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_ses_active_receipt_rule_setaws_ses_domain_identityaws_ses_domain_identity_verificationaws_ses_domain_dkimaws_ses_domain_mail_fromaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_configuration_setaws_ses_identity_notification_topicaws_ses_templateaws_launch_configurationaws_launch_templateaws_autoscaling_groupaws_autoscaling_policy"

var _ResourceTypeIndex = [...]uint16{0, 12, 19, 45, 57, 75, 85, 99, 122, 129, 136, 152, 173, 201, 221, 227, 242, 262, 289, 308, 323, 345, 364, 377, 404, 441, 466, 493, 511, 532, 563, 576, 600, 620, 651, 675, 706, 720, 732, 751, 781, 802, 828, 840, 869, 888, 918, 938, 964, 988, 1009, 1027, 1043, 1071, 1100, 1137, 1168, 1191, 1227, 1246, 1270, 1292, 1312, 1336, 1361, 1396, 1412, 1436, 1455, 1476, 1498}

const _ResourceTypeLowerName = "aws_instanceaws_vpcaws_vpc_peering_connectionaws_key_pairaws_security_groupaws_subnetaws_ebs_volumeaws_elasticache_clusteraws_elbaws_albaws_alb_listeneraws_alb_listener_ruleaws_alb_listener_certificateaws_alb_target_groupaws_lbaws_lb_listeneraws_lb_listener_ruleaws_lb_listener_certificateaws_lb_target_groupaws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_s3_bucketaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_zoneaws_route53_zone_associationaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_ses_active_receipt_rule_setaws_ses_domain_identityaws_ses_domain_identity_verificationaws_ses_domain_dkimaws_ses_domain_mail_fromaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_configuration_setaws_ses_identity_notification_topicaws_ses_templateaws_launch_configurationaws_launch_templateaws_autoscaling_groupaws_autoscaling_policy"

func (i ResourceType) String() string {
	i -= 1
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i+1)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

var _ResourceTypeValues = []ResourceType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:12]:           1,
	_ResourceTypeLowerName[0:12]:      1,
	_ResourceTypeName[12:19]:          2,
	_ResourceTypeLowerName[12:19]:     2,
	_ResourceTypeName[19:45]:          3,
	_ResourceTypeLowerName[19:45]:     3,
	_ResourceTypeName[45:57]:          4,
	_ResourceTypeLowerName[45:57]:     4,
	_ResourceTypeName[57:75]:          5,
	_ResourceTypeLowerName[57:75]:     5,
	_ResourceTypeName[75:85]:          6,
	_ResourceTypeLowerName[75:85]:     6,
	_ResourceTypeName[85:99]:          7,
	_ResourceTypeLowerName[85:99]:     7,
	_ResourceTypeName[99:122]:         8,
	_ResourceTypeLowerName[99:122]:    8,
	_ResourceTypeName[122:129]:        9,
	_ResourceTypeLowerName[122:129]:   9,
	_ResourceTypeName[129:136]:        10,
	_ResourceTypeLowerName[129:136]:   10,
	_ResourceTypeName[136:152]:        11,
	_ResourceTypeLowerName[136:152]:   11,
	_ResourceTypeName[152:173]:        12,
	_ResourceTypeLowerName[152:173]:   12,
	_ResourceTypeName[173:201]:        13,
	_ResourceTypeLowerName[173:201]:   13,
	_ResourceTypeName[201:221]:        14,
	_ResourceTypeLowerName[201:221]:   14,
	_ResourceTypeName[221:227]:        15,
	_ResourceTypeLowerName[221:227]:   15,
	_ResourceTypeName[227:242]:        16,
	_ResourceTypeLowerName[227:242]:   16,
	_ResourceTypeName[242:262]:        17,
	_ResourceTypeLowerName[242:262]:   17,
	_ResourceTypeName[262:289]:        18,
	_ResourceTypeLowerName[262:289]:   18,
	_ResourceTypeName[289:308]:        19,
	_ResourceTypeLowerName[289:308]:   19,
	_ResourceTypeName[308:323]:        20,
	_ResourceTypeLowerName[308:323]:   20,
	_ResourceTypeName[323:345]:        21,
	_ResourceTypeLowerName[323:345]:   21,
	_ResourceTypeName[345:364]:        22,
	_ResourceTypeLowerName[345:364]:   22,
	_ResourceTypeName[364:377]:        23,
	_ResourceTypeLowerName[364:377]:   23,
	_ResourceTypeName[377:404]:        24,
	_ResourceTypeLowerName[377:404]:   24,
	_ResourceTypeName[404:441]:        25,
	_ResourceTypeLowerName[404:441]:   25,
	_ResourceTypeName[441:466]:        26,
	_ResourceTypeLowerName[441:466]:   26,
	_ResourceTypeName[466:493]:        27,
	_ResourceTypeLowerName[466:493]:   27,
	_ResourceTypeName[493:511]:        28,
	_ResourceTypeLowerName[493:511]:   28,
	_ResourceTypeName[511:532]:        29,
	_ResourceTypeLowerName[511:532]:   29,
	_ResourceTypeName[532:563]:        30,
	_ResourceTypeLowerName[532:563]:   30,
	_ResourceTypeName[563:576]:        31,
	_ResourceTypeLowerName[563:576]:   31,
	_ResourceTypeName[576:600]:        32,
	_ResourceTypeLowerName[576:600]:   32,
	_ResourceTypeName[600:620]:        33,
	_ResourceTypeLowerName[600:620]:   33,
	_ResourceTypeName[620:651]:        34,
	_ResourceTypeLowerName[620:651]:   34,
	_ResourceTypeName[651:675]:        35,
	_ResourceTypeLowerName[651:675]:   35,
	_ResourceTypeName[675:706]:        36,
	_ResourceTypeLowerName[675:706]:   36,
	_ResourceTypeName[706:720]:        37,
	_ResourceTypeLowerName[706:720]:   37,
	_ResourceTypeName[720:732]:        38,
	_ResourceTypeLowerName[720:732]:   38,
	_ResourceTypeName[732:751]:        39,
	_ResourceTypeLowerName[732:751]:   39,
	_ResourceTypeName[751:781]:        40,
	_ResourceTypeLowerName[751:781]:   40,
	_ResourceTypeName[781:802]:        41,
	_ResourceTypeLowerName[781:802]:   41,
	_ResourceTypeName[802:828]:        42,
	_ResourceTypeLowerName[802:828]:   42,
	_ResourceTypeName[828:840]:        43,
	_ResourceTypeLowerName[828:840]:   43,
	_ResourceTypeName[840:869]:        44,
	_ResourceTypeLowerName[840:869]:   44,
	_ResourceTypeName[869:888]:        45,
	_ResourceTypeLowerName[869:888]:   45,
	_ResourceTypeName[888:918]:        46,
	_ResourceTypeLowerName[888:918]:   46,
	_ResourceTypeName[918:938]:        47,
	_ResourceTypeLowerName[918:938]:   47,
	_ResourceTypeName[938:964]:        48,
	_ResourceTypeLowerName[938:964]:   48,
	_ResourceTypeName[964:988]:        49,
	_ResourceTypeLowerName[964:988]:   49,
	_ResourceTypeName[988:1009]:       50,
	_ResourceTypeLowerName[988:1009]:  50,
	_ResourceTypeName[1009:1027]:      51,
	_ResourceTypeLowerName[1009:1027]: 51,
	_ResourceTypeName[1027:1043]:      52,
	_ResourceTypeLowerName[1027:1043]: 52,
	_ResourceTypeName[1043:1071]:      53,
	_ResourceTypeLowerName[1043:1071]: 53,
	_ResourceTypeName[1071:1100]:      54,
	_ResourceTypeLowerName[1071:1100]: 54,
	_ResourceTypeName[1100:1137]:      55,
	_ResourceTypeLowerName[1100:1137]: 55,
	_ResourceTypeName[1137:1168]:      56,
	_ResourceTypeLowerName[1137:1168]: 56,
	_ResourceTypeName[1168:1191]:      57,
	_ResourceTypeLowerName[1168:1191]: 57,
	_ResourceTypeName[1191:1227]:      58,
	_ResourceTypeLowerName[1191:1227]: 58,
	_ResourceTypeName[1227:1246]:      59,
	_ResourceTypeLowerName[1227:1246]: 59,
	_ResourceTypeName[1246:1270]:      60,
	_ResourceTypeLowerName[1246:1270]: 60,
	_ResourceTypeName[1270:1292]:      61,
	_ResourceTypeLowerName[1270:1292]: 61,
	_ResourceTypeName[1292:1312]:      62,
	_ResourceTypeLowerName[1292:1312]: 62,
	_ResourceTypeName[1312:1336]:      63,
	_ResourceTypeLowerName[1312:1336]: 63,
	_ResourceTypeName[1336:1361]:      64,
	_ResourceTypeLowerName[1336:1361]: 64,
	_ResourceTypeName[1361:1396]:      65,
	_ResourceTypeLowerName[1361:1396]: 65,
	_ResourceTypeName[1396:1412]:      66,
	_ResourceTypeLowerName[1396:1412]: 66,
	_ResourceTypeName[1412:1436]:      67,
	_ResourceTypeLowerName[1412:1436]: 67,
	_ResourceTypeName[1436:1455]:      68,
	_ResourceTypeLowerName[1436:1455]: 68,
	_ResourceTypeName[1455:1476]:      69,
	_ResourceTypeLowerName[1455:1476]: 69,
	_ResourceTypeName[1476:1498]:      70,
	_ResourceTypeLowerName[1476:1498]: 70,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:12],
	_ResourceTypeName[12:19],
	_ResourceTypeName[19:45],
	_ResourceTypeName[45:57],
	_ResourceTypeName[57:75],
	_ResourceTypeName[75:85],
	_ResourceTypeName[85:99],
	_ResourceTypeName[99:122],
	_ResourceTypeName[122:129],
	_ResourceTypeName[129:136],
	_ResourceTypeName[136:152],
	_ResourceTypeName[152:173],
	_ResourceTypeName[173:201],
	_ResourceTypeName[201:221],
	_ResourceTypeName[221:227],
	_ResourceTypeName[227:242],
	_ResourceTypeName[242:262],
	_ResourceTypeName[262:289],
	_ResourceTypeName[289:308],
	_ResourceTypeName[308:323],
	_ResourceTypeName[323:345],
	_ResourceTypeName[345:364],
	_ResourceTypeName[364:377],
	_ResourceTypeName[377:404],
	_ResourceTypeName[404:441],
	_ResourceTypeName[441:466],
	_ResourceTypeName[466:493],
	_ResourceTypeName[493:511],
	_ResourceTypeName[511:532],
	_ResourceTypeName[532:563],
	_ResourceTypeName[563:576],
	_ResourceTypeName[576:600],
	_ResourceTypeName[600:620],
	_ResourceTypeName[620:651],
	_ResourceTypeName[651:675],
	_ResourceTypeName[675:706],
	_ResourceTypeName[706:720],
	_ResourceTypeName[720:732],
	_ResourceTypeName[732:751],
	_ResourceTypeName[751:781],
	_ResourceTypeName[781:802],
	_ResourceTypeName[802:828],
	_ResourceTypeName[828:840],
	_ResourceTypeName[840:869],
	_ResourceTypeName[869:888],
	_ResourceTypeName[888:918],
	_ResourceTypeName[918:938],
	_ResourceTypeName[938:964],
	_ResourceTypeName[964:988],
	_ResourceTypeName[988:1009],
	_ResourceTypeName[1009:1027],
	_ResourceTypeName[1027:1043],
	_ResourceTypeName[1043:1071],
	_ResourceTypeName[1071:1100],
	_ResourceTypeName[1100:1137],
	_ResourceTypeName[1137:1168],
	_ResourceTypeName[1168:1191],
	_ResourceTypeName[1191:1227],
	_ResourceTypeName[1227:1246],
	_ResourceTypeName[1246:1270],
	_ResourceTypeName[1270:1292],
	_ResourceTypeName[1292:1312],
	_ResourceTypeName[1312:1336],
	_ResourceTypeName[1336:1361],
	_ResourceTypeName[1361:1396],
	_ResourceTypeName[1396:1412],
	_ResourceTypeName[1412:1436],
	_ResourceTypeName[1436:1455],
	_ResourceTypeName[1455:1476],
	_ResourceTypeName[1476:1498],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
