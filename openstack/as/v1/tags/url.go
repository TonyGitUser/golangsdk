package tags

import (
	"github.com/huaweicloud/golangsdk"
)

func ListResourceTagsURL(c *golangsdk.ServiceClient, tenantId string, resourceType string, resourceId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, resourceType, resourceId, "tags")
}

func ListTenantTagsURL(c *golangsdk.ServiceClient, tenantId string, resourceType string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, resourceType, "tags")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, resourceType string, resourceId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, resourceType, resourceId, "tags", "action")
}
