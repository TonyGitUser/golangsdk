package bandwidths

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

	// Specifies the bandwidth objects.
	Bandwidth BandWidth `json:"bandwidth,"`
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

	// Specifies the bandwidth list objects.
	Bandwidths []BandWidth `json:"bandwidths,"`
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

	// Specifies the bandwidth objects.
	Bandwidth BandWidth `json:"bandwidth,"`
}
