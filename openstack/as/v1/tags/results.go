package tags

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type ListResourceTagsResult struct {
	commonResult
}

func (r ListResourceTagsResult) Extract() (*ListResourceTagsResponse, error) {
	var response ListResourceTagsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResourceTagsResponse struct {

	// Specifies the resource tag.
	Tags []struct {

		// Specifies the resource tag key.
		Key string `json:"key,omitempty"`

		// Specifies the resource tag values.
		Value string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
}

type ListTenantTagsResult struct {
	commonResult
}

func (r ListTenantTagsResult) Extract() (*ListTenantTagsResponse, error) {
	var response ListTenantTagsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListTenantTagsResponse struct {

	// Specifies the resource tag.
	Tags []struct {

		// Specifies the resource tag key.
		Key string `json:"key,omitempty"`

		// Specifies the resource tag values.
		Values []string `json:"values,omitempty"`
	} `json:"tags,omitempty"`
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*UpdateResponse, error) {
	var response UpdateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateResponse struct {
}
