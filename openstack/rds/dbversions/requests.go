package dbversions

import (
	"github.com/huaweicloud/golangsdk"
)

func List(client *golangsdk.ServiceClient, versionId string, projectId string, datastoreName string) (r ListResult) {
	url := ListURL(client, versionId, projectId, datastoreName)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}
