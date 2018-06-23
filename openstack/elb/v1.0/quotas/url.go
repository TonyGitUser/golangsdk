package quotas

import (
	"github.com/huaweicloud/golangsdk"
)

func ListURL(c *golangsdk.ServiceClient, tenantId string) string {
	return c.ServiceURL("v1.0", tenantId, "elbaas", "quotas")
}
