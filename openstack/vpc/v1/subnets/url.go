package subnets

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "subnets")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, vpcId string, subnetId string) string {
	return c.ServiceURL("v1", tenantId, "vpcs", vpcId, "subnets", subnetId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, subnetId string) string {
	return c.ServiceURL("v1", tenantId, "subnets", subnetId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "subnets")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, vpcId string, subnetId string) string {
	return c.ServiceURL("v1", tenantId, "vpcs", vpcId, "subnets", subnetId)
}
