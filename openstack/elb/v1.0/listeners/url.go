package listeners

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, listenerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners", listenerId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, listenerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners", listenerId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, listenerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners", listenerId)
}
