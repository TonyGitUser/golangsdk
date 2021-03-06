package quotas

import (
	"github.com/huaweicloud/golangsdk"
)

func List(client *golangsdk.ServiceClient, tenantId string) (r ListResult) {
	url := ListURL(client, tenantId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
