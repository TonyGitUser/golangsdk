package backups

import (
	"github.com/huaweicloud/golangsdk"
)

func BackupURL(c *golangsdk.ServiceClient, versionId string, projectId string) string {
	return c.ServiceURL("rds", versionId, projectId, "backups")
}

func CreatePolicyURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "backups", "policy")
}

func DeleteURL(c *golangsdk.ServiceClient, versionId string, projectId string, backupId string) string {
	return c.ServiceURL("rds", versionId, projectId, "backups", backupId)
}

func GetPolicyURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "backups", "policy")
}

func ListURL(c *golangsdk.ServiceClient, versionId string, projectId string) string {
	return c.ServiceURL("rds", versionId, projectId, "backups")
}

func RestoreURL(c *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances", instanceId, "action")
}

func RestoreWithNewURL(c *golangsdk.ServiceClient, versionId string, projectId string) string {
	return c.ServiceURL("rds", versionId, projectId, "instances")
}
