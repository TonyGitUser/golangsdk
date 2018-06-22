package publicips

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

	// Specifies the elastic IP address objects.
	Publicip PublicIP `json:"publicip,"`
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

	// Specifies the elastic IP address objects.
	Publicip struct {

		// Specifies the ID of the elastic IP address, which uniquely identifies the elastic IP address.
		Id string `json:"id,"`

		// Specifies the status of the elastic IP address.
		Status string `json:"status,"`

		// Specifies the type of the elastic IP address.
		Type string `json:"type,"`

		// Specifies the obtained elastic IP address.
		PublicIpAddress string `json:"public_ip_address,"`

		// Specifies the private IP address bound to the elastic IP address.
		PrivateIpAddress string `json:"private_ip_address,omitempty"`

		// Specifies the port ID.
		PortId string `json:"port_id,omitempty"`

		// Specifies the tenant ID of the operator.
		TenantId string `json:"tenant_id,"`

		// Specifies the time for applying for the elastic IP address.
		CreateTime string `json:"create_time,"`

		// Specifies the bandwidth ID of the elastic IP address.
		BandwidthId string `json:"bandwidth_id,"`

		// Specifies the bandwidth size.
		BandwidthSize int `json:"bandwidth_size,"`

		// Specifies whether the bandwidth is shared or exclusive.
		BandwidthShareType string `json:"bandwidth_share_type,"`

		// Specifies the bandwidth name.
		BandwidthName string `json:"bandwidth_name,"`
	} `json:"publicip,"`
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

	// Specifies the elastic IP address list objects.
	Publicips []PublicIP `json:"publicips,"`
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

	// Specifies the elastic IP address objects.
	Publicip PublicIP `json:"publicip,"`
}
