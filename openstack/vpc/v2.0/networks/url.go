package networks

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "networks")
}

func DeleteURL(c *golangsdk.ServiceClient, networkId string) string {
	return c.ServiceURL("v2.0", "networks", networkId)
}

func GetURL(c *golangsdk.ServiceClient, networkId string) string {
	return c.ServiceURL("v2.0", "networks", networkId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "networks")
}

func UpdateURL(c *golangsdk.ServiceClient, portId string) string {
	return c.ServiceURL("v2.0", "networks", portId)
}
