package notifications

import (
	"github.com/huaweicloud/golangsdk"
)

func Delete(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, topicUrn string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, scalingGroupId, topicUrn)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

type EnableOpts struct {

	// Specifies a unified topic in SMN.
	TopicUrn string `json:"topic_urn,omitempty"`

	// Specifies a notification scenario, which can be one of the following:SCALING_UP: indicates that the capacity is expanded.SCALING_UP_FAIL: indicates that the capacity expansion failed.SCALING_DOWN: indicates that the capacity is reduced.SCALING_DOWN_FAIL: indicates that the capacity reduction failed.SCALING_GROUP_ABNORMAL: indicates that an exception has occurred in the AS group.
	TopicScene []string `json:"topic_scene,omitempty"`
}

type EnableOptsBuilder interface {
	ToEnableMap() (map[string]interface{}, error)
}

func (opts EnableOpts) ToEnableMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Enable(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts EnableOptsBuilder) (r EnableResult) {
	b, _ := opts.ToEnableMap()

	_, r.Err = client.Put(EnableURL(client, tenantId, scalingGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func List(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string) (r ListResult) {
	url := ListURL(client, tenantId, scalingGroupId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
