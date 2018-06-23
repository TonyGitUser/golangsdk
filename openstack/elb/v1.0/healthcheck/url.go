package healthcheck

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "healthcheck")
}

func DeleteURL(c *golangsdk.ServiceClient, healthcheckId string) string {
	return c.ServiceURL("v1.0", "elbaas", "healthcheck", healthcheckId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, healthcheckId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "healthcheck", healthcheckId)
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, healthcheckId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "healthcheck", healthcheckId)
}
