package routers

import (
	"github.com/huaweicloud/golangsdk"
)

func AddInterfaceURL(c *golangsdk.ServiceClient, routerId string) string {
	return c.ServiceURL("v2.0", "routers", routerId, "add_router_interface")
}

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "routers")
}

func DeleteURL(c *golangsdk.ServiceClient, routerId string) string {
	return c.ServiceURL("v2.0", "routers", routerId)
}

func GetURL(c *golangsdk.ServiceClient, routerId string) string {
	return c.ServiceURL("v2.0", "routers", routerId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "routers")
}

func RemoveInterfaceURL(c *golangsdk.ServiceClient, routerId string) string {
	return c.ServiceURL("v2.0", "routers", routerId, "remove_router_interface")
}

func UpdateURL(c *golangsdk.ServiceClient, routerId string) string {
	return c.ServiceURL("v2.0", "routers", routerId)
}
