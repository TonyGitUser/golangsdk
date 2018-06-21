package securitygrouprules

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "security-group-rules")
}

func DeleteURL(c *golangsdk.ServiceClient, securityGroupsRulesId string) string {
	return c.ServiceURL("v2.0", "security-group-rules", securityGroupsRulesId)
}

func GetURL(c *golangsdk.ServiceClient, securityGroupsRulesId string) string {
	return c.ServiceURL("v2.0", "security-group-rules", securityGroupsRulesId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "security-group-rules")
}
