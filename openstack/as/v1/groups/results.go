package groups

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

	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id,omitempty"`
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

type EnableResult struct {
	commonResult
}

func (r EnableResult) Extract() (*EnableResponse, error) {
	var response EnableResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type EnableResponse struct {
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

	// Specifies details about the AS group.
	ScalingGroup ScalingGroup `json:"scaling_group,omitempty"`
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

	// Specifies the start number of query records.
	StartNumber int `json:"start_number,omitempty"`

	// Specifies the number of query records.
	Limit int `json:"limit,omitempty"`

	// Specifies the scaling group list.
	ScalingGroups []ScalingGroup `json:"scaling_groups,omitempty"`
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

	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id,omitempty"`
}
