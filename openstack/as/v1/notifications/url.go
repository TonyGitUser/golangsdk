package notifications

import (
	"github.com/huaweicloud/golangsdk"
)

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string, topicUrn string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_notification", scalingGroupId, topicUrn)
}

func EnableURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_notification", scalingGroupId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_notification", scalingGroupId)
}
