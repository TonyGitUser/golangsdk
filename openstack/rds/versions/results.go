package versions

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

	// Indicates the list of detailed API version information.
	Versions Version `json:"versions,omitempty"`

	// Indicates the list of detailed API version information.
	Version Version `json:"version,omitempty"`
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

	// Indicates the list of detailed API version information.
	Versions []Version `json:"versions,omitempty"`
}
