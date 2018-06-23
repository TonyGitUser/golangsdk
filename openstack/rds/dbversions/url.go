package dbversions

import (
	"github.com/huaweicloud/golangsdk"
)

func ListURL(c *golangsdk.ServiceClient, versionId string, projectId string, datastoreName string) string {
	return c.ServiceURL("rds", versionId, projectId, "datastores", datastoreName, "versions")
}
