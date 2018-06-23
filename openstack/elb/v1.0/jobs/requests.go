package jobs

import (
	"github.com/huaweicloud/golangsdk"
)

func Get(client *golangsdk.ServiceClient, tenantId string, jobId string) (r GetResult) {
	url := GetURL(client, tenantId, jobId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
