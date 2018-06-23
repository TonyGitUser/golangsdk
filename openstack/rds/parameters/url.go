package parameters

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "parameters")
}

func GetURL(c *golangsdk.ServiceClient, versionId string, projectId string, datastoreVersionId string, parameterName string) string {
	return c.ServiceURL("rds", versionId, projectId, "datastores", "versions", datastoreVersionId, "parameters", parameterName)
}

func ListURL(c *golangsdk.ServiceClient, versionId string, projectId string, datastoreVersionId string) string {
	return c.ServiceURL("rds", versionId, projectId, "datastores", "versions", datastoreVersionId, "parameters")
}

func RestoreURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "parameters", "default")
}
