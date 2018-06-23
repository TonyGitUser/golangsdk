package policies

import (
	"github.com/huaweicloud/golangsdk"
)

func ActionURL(c *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_policy", scalingPolicyId, "action")
}

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_policy")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_policy", scalingPolicyId)
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_policy", scalingPolicyId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string, scalingGroupId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_policy", scalingGroupId, "list")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_policy", scalingPolicyId)
}
