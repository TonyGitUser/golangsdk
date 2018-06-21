package floatingips

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "floatingips")
}

func DeleteURL(c *golangsdk.ServiceClient, floatingipId string) string {
	return c.ServiceURL("v2.0", "floatingips", floatingipId)
}

func GetURL(c *golangsdk.ServiceClient, floatingipId string) string {
	return c.ServiceURL("v2.0", "floatingips", floatingipId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "floatingips")
}

func UpdateURL(c *golangsdk.ServiceClient, floatingipId string) string {
	return c.ServiceURL("v2.0", "floatingips", floatingipId)
}
