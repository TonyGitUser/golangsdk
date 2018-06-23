package policies

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("autoscaling-api", "v2", tenantId, "scaling_policy")
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) string {
	return c.ServiceURL("autoscaling-api", "v2", tenantId, "scaling_policy", scalingPolicyId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string, scalingResourceId string) string {
	return c.ServiceURL("autoscaling-api", "v2", tenantId, "scaling_policy", scalingResourceId, "list")
}

func UpdateURL(c *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) string {
	return c.ServiceURL("autoscaling-api", "v2", tenantId, "scaling_policy", scalingPolicyId)
}
