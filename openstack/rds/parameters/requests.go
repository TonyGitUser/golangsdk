package parameters

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOpts struct {

	// Specifies the parameter value list.
	Values map[string]interface{} `json:"values,omitempty"`
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts CreateOptsBuilder) (r CreateResult) {
	b, _ := opts.ToCreateMap()

	_, r.Err = client.Put(CreateURL(client, versionId, projectId, instanceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func Get(client *golangsdk.ServiceClient, versionId string, projectId string, datastoreVersionId string, parameterName string) (r GetResult) {
	url := GetURL(client, versionId, projectId, datastoreVersionId, parameterName)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func List(client *golangsdk.ServiceClient, versionId string, projectId string, datastoreVersionId string) (r ListResult) {
	url := ListURL(client, versionId, projectId, datastoreVersionId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func Restore(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) (r RestoreResult) {
	_, r.Err = client.Put(RestoreURL(client, versionId, projectId, instanceId), nil, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}
