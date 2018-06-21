package securitygroups

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "security-groups")
}

func DeleteURL(c *golangsdk.ServiceClient, securityGroupId string) string {
	return c.ServiceURL("v2.0", "security-groups", securityGroupId)
}

func GetURL(c *golangsdk.ServiceClient, securityGroupId string) string {
	return c.ServiceURL("v2.0", "security-groups", securityGroupId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "security-groups")
}

func UpdateURL(c *golangsdk.ServiceClient, securityGroupId string) string {
	return c.ServiceURL("v2.0", "security-groups", securityGroupId)
}
