package privateips

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "privateips")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, privateipId string) string {
	return c.ServiceURL("v1", tenantId, "privateips", privateipId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, privateipId string) string {
	return c.ServiceURL("v1", tenantId, "privateips", privateipId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string, subnetId string) string {
	return c.ServiceURL("v1", tenantId, "subnets", subnetId, "privateips")
}
