package tags

import (
	"github.com/huaweicloud/golangsdk"
)

func ListResourceTags(client *golangsdk.ServiceClient, tenantId string, resourceType string, resourceId string) (r ListResourceTagsResult) {
	url := ListResourceTagsURL(client, tenantId, resourceType, resourceId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func ListTenantTags(client *golangsdk.ServiceClient, tenantId string, resourceType string) (r ListTenantTagsResult) {
	url := ListTenantTagsURL(client, tenantId, resourceType)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the tag list.If action is set to delete, the tag structure cannot be missing, and the key cannot be left blank or an empty string.
	Tags []struct {

		// Specifies the resource tag key.It contains a maximum of 36 Unicode characters and cannot be left blank. It cannot contain the following characters: ASCII (0-31), equal signs (=), asterisks (*), left angle brackets (<), right angle brackets (>), backslashes (\), commas (,), vertical bars (|), and slashes (/).Tag keys of a resource must be unique.If action is set to delete, the tag character set is not verified. The key contains a maximum of 127 Unicode characters.
		Key string `json:"key,omitempty"`

		// Specifies the resource tag values.Each value contains a maximum of 43 Unicode characters and can be left blank. It cannot contain the following characters: ASCII (0-31), equal signs (=), asterisks (*), left angle brackets (<), right angle brackets (>), backslashes (\), commas (,), vertical bars (|), and slashes (/).If action is set to delete, the tag character set is not verified. Each value contains a maximum of 255 Unicode characters. If value is specified, tags are deleted by key and value. If value is not specified, tags are deleted by key.
		Value string `json:"value,omitempty"`
	} `json:"tags,omitempty"`

	// Operation ID (case sensitive).update: indicates updating a tag. If the same key value exists, it will be overwritten. If no same key value exists, a new tag will be created.delete: indicates deleting a tag.create: indicates creating a tag. If the same key value already exists, it will be overwritten.
	Action string `json:"action,omitempty"`
}

type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *golangsdk.ServiceClient, tenantId string, resourceType string, resourceId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Post(UpdateURL(client, tenantId, resourceType, resourceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
