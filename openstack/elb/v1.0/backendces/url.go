package backendces

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string, listenerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners", listenerId, "members")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, listenerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners", listenerId, "members", "action")
}

func ListURL(c *golangsdk.ServiceClient, tenantId string, listenerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "listeners", listenerId, "members")
}
