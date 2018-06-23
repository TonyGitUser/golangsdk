package logs

import (
	"github.com/huaweicloud/golangsdk"
)

func ListURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_activity_log", scalingGroupId)
}
