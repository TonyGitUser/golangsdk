package ports

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "ports")
}

func DeleteURL(c *golangsdk.ServiceClient, portId string) string {
	return c.ServiceURL("v2.0", "ports", portId)
}

func GetURL(c *golangsdk.ServiceClient, portId string) string {
	return c.ServiceURL("v2.0", "ports", portId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "ports")
}

func UpdateURL(c *golangsdk.ServiceClient, portId string) string {
	return c.ServiceURL("v2.0", "ports", portId)
}
