package certificate

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "certificate")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, certificateId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "certificate", certificateId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "certificate")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, certificateId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "certificate", certificateId)
}
