package quotas

import (
	"github.com/huaweicloud/golangsdk"
)

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "quotas")
}

func ListWithInstancesURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "quotas", scalingGroupId)
}
