package subnets

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "subnets")
}

func DeleteURL(c *golangsdk.ServiceClient, subnetId string) string {
	return c.ServiceURL("v2.0", "subnets", subnetId)
}

func GetURL(c *golangsdk.ServiceClient, subnetId string) string {
	return c.ServiceURL("v2.0", "subnets", subnetId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("v2.0", "subnets")
}

func UpdateURL(c *golangsdk.ServiceClient, subnetId string) string {
	return c.ServiceURL("v2.0", "subnets", subnetId)
}
