package configures

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {

	// Specifies the AS configuration ID.
	ScalingConfigurationId string `json:"scaling_configuration_id,omitempty"`
}

type DeleteResult struct {
	commonResult
}

func (r DeleteResult) Extract() (*DeleteResponse, error) {
	var response DeleteResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteResponse struct {
}

type DeleteWithBatchResult struct {
	commonResult
}

func (r DeleteWithBatchResult) Extract() (*DeleteWithBatchResponse, error) {
	var response DeleteWithBatchResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteWithBatchResponse struct {
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

	// Provides AS configuration details.
	ScalingConfiguration ScalingConfiguration `json:"scaling_configuration,omitempty"`
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

	// Specifies the total number of query records.
	TotalNumber int `json:"total_number,omitempty"`

	// Specifies the start line number.
	StartNumber int `json:"start_number,omitempty"`

	// Specifies the number of query records.
	Limit int `json:"limit,omitempty"`

	// Specifies the AS configuration list.
	ScalingConfigurations []ScalingConfiguration `json:"scaling_configurations,omitempty"`
}
