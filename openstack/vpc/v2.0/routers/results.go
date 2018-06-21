package routers

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type commonResult struct {
	golangsdk.Result
}

type AddInterfaceResult struct {
	commonResult
}

func (r AddInterfaceResult) Extract() (*AddInterfaceResponse, error) {
	var response AddInterfaceResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type AddInterfaceResponse struct {

	// Specifies the subnet ID.
	SubnetId string `json:"subnet_id,"`

	// Specifies the tenant ID.
	TenantId string `json:"tenant_id,"`

	// Specifies the port ID.
	PortId string `json:"port_id,"`

	// Specifies the router ID.
	Id string `json:"id,"`
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

	// Specifies the router list.
	Router ListRouter `json:"router,"`
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

	// Specifies the router list.
	Router ListRouter `json:"router,"`
}

type ListPage struct {
	postpagination.LinkedPageBase
}

func (r ListPage) IsEmpty() (bool, error) {
	response, err := ExtractList(r)
	return len(response.Routers) == 0, err
}

func ExtractList(r ListPage) (*ListResponse, error) {
	var list ListResponse
	r.ExtractInto(&list)
	return &list, r.Err
}

type ListResponse struct {

	// Specifies the router list.
	Routers []ListRouter `json:"routers,"`
}

type RemoveInterfaceResult struct {
	commonResult
}

func (r RemoveInterfaceResult) Extract() (*RemoveInterfaceResponse, error) {
	var response RemoveInterfaceResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RemoveInterfaceResponse struct {

	// Specifies the subnet ID.
	SubnetId string `json:"subnet_id,"`

	// Specifies the tenant ID.
	TenantId string `json:"tenant_id,"`

	// Specifies the port ID.
	PortId string `json:"port_id,"`

	// Specifies the router ID.
	Id string `json:"id,"`
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

	// Specifies the router list.
	Router ListRouter `json:"router,"`
}
