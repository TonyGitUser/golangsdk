package vpcs

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

	// Specifies the VPC objects.
	Vpc VPC `json:"vpc,"`
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

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

	// Specifies the VPC objects.
	Vpc VPC `json:"vpc,"`
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

	// Specifies the VPC list objects.
	Vpcs []VPC `json:"vpcs,"`
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

	// Specifies the VPC objects.
	Vpc VPC `json:"vpc,"`
}
