package versions

import (
	"github.com/huaweicloud/golangsdk"
)

type Link struct {

	// Indicates the API URL and the value is .
	Href string `json:"href,omitempty"`

	// Its value is , indicating that  is a local link.
	Rel string `json:"rel,omitempty"`
}

type Version struct {

	// Indicates the API version.
	Id string `json:"id,omitempty"`

	// Indicates the API version link information. Its value is empty.
	Links []Link `json:"links,omitempty"`

	// Indicates the version status.
	Status string `json:"status,omitempty"`

	// Indicates the version update time.
	Updated string `json:"updated,omitempty"`
}

func Get(client *golangsdk.ServiceClient, versionId string) (r GetResult) {
	url := GetURL(client, versionId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func List(client *golangsdk.ServiceClient) (r ListResult) {
	url := ListURL(client)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
