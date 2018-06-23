package policylogs

import (
	"github.com/huaweicloud/golangsdk"
)

func ListURL(c *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_policy_execute_log", scalingPolicyId)
}
