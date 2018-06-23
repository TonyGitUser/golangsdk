package instances

import (
	"github.com/huaweicloud/golangsdk"
)

func ActionURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_group_instance", scalingGroupId, "action")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, instanceId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_group_instance", instanceId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_group_instance", scalingGroupId, "list")
}
