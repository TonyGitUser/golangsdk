package bandwidths

import (
	"github.com/huaweicloud/golangsdk"
)

func GetURL(c *golangsdk.ServiceClient, tenantId string, bandwidthId string) string {
	return c.ServiceURL("v1", tenantId, "bandwidths", bandwidthId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1", tenantId, "bandwidths")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, bandwidthId string) string {
	return c.ServiceURL("v1", tenantId, "bandwidths", bandwidthId)
}
