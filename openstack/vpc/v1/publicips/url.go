package publicips

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "publicips")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, publicipId string) string {
	return c.ServiceURL("v1", tenantId, "publicips", publicipId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, publicipId string) string {
	return c.ServiceURL("v1", tenantId, "publicips", publicipId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "publicips")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, publicipId string) string {
	return c.ServiceURL("v1", tenantId, "publicips", publicipId)
}
