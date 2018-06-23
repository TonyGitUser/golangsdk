package certificate

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

	// Specifies the tenant ID.
	TenantId string `json:"tenant_id,omitempty"`

	// Specifies the certificate ID.
	Id string `json:"id,omitempty"`

	// Specifies the certificate name.
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the certificate.
	Description string `json:"description,omitempty"`

	// Specifies the domain name of the certificate signed on the server.
	Domain string `json:"domain,omitempty"`

	// Specifies the content of the certificate.
	Certificate string `json:"certificate,omitempty"`

	// Specifies the private key of the certificate.
	PrivateKey string `json:"private_key,omitempty"`

	// Specifies the time when the certificate was created.
	CreateTime string `json:"create_time,omitempty"`

	// Specifies the time when the certificate was updated.
	UpdateTime string `json:"update_time,omitempty"`
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

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResponse struct {

	// Specifies the list of certificates.
	Certificates []struct {

		// Specifies the certificate ID.
		Id string `json:"id,omitempty"`

		// Specifies the certificate name.
		Name string `json:"name,omitempty"`

		// Provides supplementary information about the certificate.
		Description string `json:"description,omitempty"`

		// Specifies the domain name of the certificate signed on the server.
		Domain string `json:"domain,omitempty"`

		// Specifies the content of the certificate.
		Certificate string `json:"certificate,omitempty"`

		// Specifies the private key of the certificate.
		PrivateKey string `json:"private_key,omitempty"`

		// Specifies the time when the certificate was created.
		CreateTime string `json:"create_time,omitempty"`

		// Specifies the time when the certificate was updated.
		UpdateTime string `json:"update_time,omitempty"`
	} `json:"certificates,omitempty"`

	// Specifies the number of certificates.
	InstanceNum string `json:"instance_num,omitempty"`
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

	// Specifies the certificate ID.
	Id string `json:"id,omitempty"`

	// Specifies the certificate name.
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the certificate.
	Description string `json:"description,omitempty"`

	// Specifies the domain name of the certificate signed on the server.
	Domain string `json:"domain,omitempty"`

	// Specifies the content of the certificate.
	Certificate string `json:"certificate,omitempty"`

	// Specifies the private key of the certificate.
	PrivateKey string `json:"private_key,omitempty"`

	// Specifies the time when the certificate was created.
	CreateTime string `json:"create_time,omitempty"`

	// Specifies the time when the certificate was updated.
	UpdateTime string `json:"update_time,omitempty"`
}
