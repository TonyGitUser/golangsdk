package versions

import (
	"github.com/huaweicloud/golangsdk"
)

func GetURL(c *golangsdk.ServiceClient, versionId string) string {
	return c.ServiceURL("rds", versionId)
}

func ListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("rds")
}
