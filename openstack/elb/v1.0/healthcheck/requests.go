package healthcheck

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOpts struct {

	// Specifies the ID of the listener to which the health check task belongs.
	ListenerId string `json:"listener_id,omitempty"`

	// Specifies the protocol used for the health check.The value can be HTTP or TCP (case-insensitive).
	HealthcheckProtocol string `json:"healthcheck_protocol,omitempty"`

	// Specifies the URI for health check. This parameter is valid when healthcheck_ protocol is HTTP.The value is a string of 1 to 80 characters that must start with a slash (/). Only contain letters, digits, and special characters such as -/.%?#&_= can be contained.
	HealthcheckPath string `json:"healthcheck_path,omitempty"`

	// Specifies the port used for the health check.The value contains 1 to 65,535 characters.
	HealthcheckConnectPort int `json:"healthcheck_connect_port,omitempty"`

	// Specifies the threshold at which the health check result is success, that is, the number of consecutive successful health checks when the health check result of the backend ECS changes from fail to success.The value contains 1 to 10 characters.
	HealthyThreshold int `json:"healthy_threshold,omitempty"`

	// Specifies the threshold at which the health check result is fail, that is, the number of consecutive failed health checks when the health check result of the backend ECS changes from success to fail.The value contains 1 to 10 characters.
	UnhealthyThreshold int `json:"unhealthy_threshold,omitempty"`

	// Specifies the maximum timeout duration (s) for the health check.The value ranges from 1 to 50.
	HealthcheckTimeout int `json:"healthcheck_timeout,omitempty"`

	// Specifies the maximum interval (s) for health check.The value ranges from 1 to 50.
	HealthcheckInterval int `json:"healthcheck_interval,omitempty"`
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

func Delete(client *golangsdk.ServiceClient, healthcheckId string) (r DeleteResult) {
	url := DeleteURL(client, healthcheckId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, healthcheckId string) (r GetResult) {
	url := GetURL(client, tenantId, healthcheckId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the protocol used for the health check.The value can be HTTP or TCP (case-insensitive).
	HealthcheckProtocol string `json:"healthcheck_protocol,omitempty"`

	// Specifies the URI for health check. This parameter is valid when healthcheck_ protocol is HTTP.The value is a string of 1 to 80 characters that must start with a slash (/). Only contain letters, digits, and special characters such as -/.%?#&_= can be contained.
	HealthcheckPath string `json:"healthcheck_path,omitempty"`

	// Specifies the port used for the health check.The value contains 1 to 65,535 characters.
	HealthcheckConnectPort int `json:"healthcheck_connect_port,omitempty"`

	// Specifies the threshold at which the health check result is success, that is, the number of consecutive successful health checks when the health check result of the backend ECS changes from fail to success.The value contains 1 to 10 characters.
	HealthyThreshold int `json:"healthy_threshold,omitempty"`

	// Specifies the threshold at which the health check result is fail, that is, the number of consecutive failed health checks when the health check result of the backend ECS changes from success to fail.The value contains 1 to 10 characters.
	UnhealthyThreshold int `json:"unhealthy_threshold,omitempty"`

	// Specifies the maximum timeout duration (s) for the health check.The value ranges from 1 to 50.
	HealthcheckTimeout int `json:"healthcheck_timeout,omitempty"`

	// Specifies the maximum interval (s) for health check.The value ranges from 1 to 50.
	HealthcheckInterval int `json:"healthcheck_interval,omitempty"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, healthcheckId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, healthcheckId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
