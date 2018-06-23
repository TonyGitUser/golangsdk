package loadbalancers

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "loadbalancers")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, loadbalancerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "loadbalancers", loadbalancerId)
}

func DeleteKeepEIPURL(c *golangsdk.ServiceClient, tenantId string, loadbalancerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "loadbalancers", loadbalancerId, "keep-eip")
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, loadbalancerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "loadbalancers", loadbalancerId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "loadbalancers")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, loadbalancerId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "loadbalancers", loadbalancerId)
}
