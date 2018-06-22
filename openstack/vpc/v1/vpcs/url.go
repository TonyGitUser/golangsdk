package vpcs

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "vpcs")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, vpcId string) string {
	return c.ServiceURL("v1", tenantId, "vpcs", vpcId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, vpcId string) string {
	return c.ServiceURL("v1", tenantId, "vpcs", vpcId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "vpcs")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, vpcId string) string {
	return c.ServiceURL("v1", tenantId, "vpcs", vpcId)
}
