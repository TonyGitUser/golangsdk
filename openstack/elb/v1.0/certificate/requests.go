package certificate

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOpts struct {

	// Specifies the certificate name.The value is a string of 0 to 64 characters that consist of Chinese characters, letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the certificate.The value contains a maximum of 128 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`

	// Specifies the domain name of the certificate signed on the server.The parameter value is a string that contains a maximum of 254 characters, can only contain letters, digits, hyphens (-), and periods (.), and must start with uppercase letters or digits.
	Domain string `json:"domain,omitempty"`

	// Specifies the content of the certificate used by HTTPS.The value is in PEM coding format.
	Certificate string `json:"certificate,omitempty"`

	// Specifies the content of the certificate used by HTTPS.The value is in PEM coding format.
	PrivateKey string `json:"private_key,omitempty"`
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, tenantId string, opts CreateOptsBuilder) (r CreateResult) {
	b, _ := opts.ToCreateMap()

	_, r.Err = client.Post(CreateURL(client, tenantId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, tenantId string, certificateId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, certificateId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func List(client *golangsdk.ServiceClient, tenantId string) (r ListResult) {
	url := ListURL(client, tenantId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the certificate name.The value is a string of 0 to 64 characters that consist of Chinese characters, letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the certificate.The value contains a maximum of 128 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`
}

type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *golangsdk.ServiceClient, tenantId string, certificateId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, certificateId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
