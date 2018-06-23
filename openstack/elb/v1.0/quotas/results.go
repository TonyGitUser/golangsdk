package quotas

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResponse struct {

	// Specifies the resource quotas.
	Quotas struct {

		// Specifies the resource quotas.
		Resources []struct {

			// Specifies the resource type, such as elb and listener.
			Type string `json:"type,omitempty"`

			// Specifies the quantity of assigned resources.
			Used int `json:"used,omitempty"`

			// Specifies the total resource quotas.
			Quota int `json:"quota,omitempty"`

			// Specifies the maximum number of resources.
			Max int `json:"max,omitempty"`

			// Specifies the minimum number of resources.
			Min int `json:"min,omitempty"`
		} `json:"resources,omitempty"`
	} `json:"quotas,omitempty"`
}
