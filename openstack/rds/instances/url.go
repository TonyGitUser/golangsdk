package instances

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, versionId string, projectId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances")
}

func DeleteURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId)
}

func GetURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId)
}

func GetFlavorURL(c *golangsdk.ServiceClient, versionId string, projectId string, flavorId string) string {
	return c.ServiceURL("rds", versionId, projectId, "flavors", flavorId)
}

func ListURL(c *golangsdk.ServiceClient, versionId string, projectId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances")
}

func ListFlavorsURL(c *golangsdk.ServiceClient, versionId string, projectId string) string {
	return c.ServiceURL("rds", versionId, projectId, "flavors")
}

func RebootURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "action")
}

func UpdateURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "action")
}
