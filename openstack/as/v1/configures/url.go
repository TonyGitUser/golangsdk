package configures

import (
	"github.com/huaweicloud/golangsdk"
)

func CreateURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_configuration")
}

func DeleteURL(c *golangsdk.ServiceClient, tenantId string, scalingConfigurationId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_configuration", scalingConfigurationId)
}

func DeleteWithBatchURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_configurations")
}

func GetURL(c *golangsdk.ServiceClient, tenantId string, scalingConfigurationId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_configuration", scalingConfigurationId)
}

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("autoscaling-api", "v1", tenantId, "scaling_configuration")
}
