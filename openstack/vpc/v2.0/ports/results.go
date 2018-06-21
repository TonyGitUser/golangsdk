package ports

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
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

	// Specifies the port object.
	Port ListPort `json:"port,"`
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

	// Specifies the port object.
	Port ListPort `json:"port,"`
}

type ListPage struct {
	postpagination.LinkedPageBase
}

func (r ListPage) IsEmpty() (bool, error) {
	response, err := ExtractList(r)
	return len(response.Ports) == 0, err
}

func ExtractList(r ListPage) (*ListResponse, error) {
	var list ListResponse
	r.ExtractInto(&list)
	return &list, r.Err
}

type ListResponse struct {

	// Specifies the port objects.
	Ports []ListPort `json:"ports,"`
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

	// Specifies the port object.
	Port ListPort `json:"port,"`
}
