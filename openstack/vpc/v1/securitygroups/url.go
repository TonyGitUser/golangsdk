package securitygroups

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "security-groups")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, securityGroupId string) string {
	return c.ServiceURL("v1", tenantId, "security-groups", securityGroupId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, securityGroupId string) string {
	return c.ServiceURL("v1", tenantId, "security-groups", securityGroupId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "security-groups")
}
