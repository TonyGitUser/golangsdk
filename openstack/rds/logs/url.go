package logs

import (
	"github.com/huaweicloud/golangsdk"
)

func ListErrorsURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "errorlog")
}

func ListInfosURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "slowlog")
}
