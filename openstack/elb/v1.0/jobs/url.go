package jobs

import (
	"github.com/huaweicloud/golangsdk"
)

func GetURL(c *golangsdk.ServiceClient, tenantId string, jobId string) string {
	return c.ServiceURL("v1.0", tenantId, "jobs", jobId)
}
