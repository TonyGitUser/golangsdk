package lifecyclehooks

import (
	"github.com/huaweicloud/golangsdk"
)

func CallBackURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_instance_hook", scalingGroupId, "callback")
}

func CreateURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_lifecycle_hook", scalingGroupId)
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string, lifecycleHookName string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_lifecycle_hook", scalingGroupId, lifecycleHookName)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string, lifecycleHookName string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_lifecycle_hook", scalingGroupId, lifecycleHookName)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_lifecycle_hook", scalingGroupId, "list")
}

func ListWithSuspensionURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_instance_hook", scalingGroupId, "list")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string, lifecycleHookName string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_lifecycle_hook", scalingGroupId, lifecycleHookName)
}
