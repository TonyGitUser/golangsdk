package healthcheck

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

	// Specifies the maximum interval (s) for health check.
	HealthcheckInterval int `json:"healthcheck_interval,omitempty"`

	// Specifies the ID of the listener to which the health check task belongs.
	ListenerId string `json:"listener_id,omitempty"`

	// Specifies the health check task ID.
	Id string `json:"id,omitempty"`

	// Specifies the health check protocol.
	HealthcheckProtocol string `json:"healthcheck_protocol,omitempty"`

	// Specifies the threshold at which the health check result is fail, that is, the number of consecutive failed health checks when the health check result of the backend ECS changes from success to fail.
	UnhealthyThreshold int `json:"unhealthy_threshold,omitempty"`

	// Specifies the time when information about the health check task was updated.
	UpdateTime string `json:"update_time,omitempty"`

	// Specifies the time when the health check task was created.
	CreateTime string `json:"create_time,omitempty"`

	// Specifies the port for health check.
	HealthcheckConnectPort int `json:"healthcheck_connect_port,omitempty"`

	// Specifies the maximum timeout duration for health check.
	HealthcheckTimeout int `json:"healthcheck_timeout,omitempty"`

	// Specifies the URI for health check. This parameter is valid when healthcheck_ protocol is HTTP.
	HealthcheckPath string `json:"healthcheck_path,omitempty"`

	// Specifies the threshold at which the health check result is success, that is, the number of consecutive successful health checks when the health check result of the backend ECS changes from fail to success.
	HealthyThreshold int `json:"healthy_threshold,omitempty"`
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

	// Specifies the maximum interval (s) for health check.
	HealthcheckInterval int `json:"healthcheck_interval,omitempty"`

	// Specifies the ID of the listener to which the health check task belongs.
	ListenerId string `json:"listener_id,omitempty"`

	// Specifies the health check task ID.
	Id string `json:"id,omitempty"`

	// Specifies the health check protocol.
	HealthcheckProtocol string `json:"healthcheck_protocol,omitempty"`

	// Specifies the threshold at which the health check result is fail, that is, the number of consecutive failed health checks when the health check result of the backend ECS changes from success to fail.
	UnhealthyThreshold int `json:"unhealthy_threshold,omitempty"`

	// Specifies the time when information about the health check task was updated.
	UpdateTime string `json:"update_time,omitempty"`

	// Specifies the time when the health check task was created.
	CreateTime string `json:"create_time,omitempty"`

	// Specifies the port for health check.
	HealthcheckConnectPort int `json:"healthcheck_connect_port,omitempty"`

	// Specifies the maximum timeout duration for health check.
	HealthcheckTimeout int `json:"healthcheck_timeout,omitempty"`

	// Specifies the URI for health check. This parameter is valid when healthcheck_ protocol is HTTP.
	HealthcheckPath string `json:"healthcheck_path,omitempty"`

	// Specifies the threshold at which the health check result is success, that is, the number of consecutive successful health checks when the health check result of the backend ECS changes from fail to success.
	HealthyThreshold int `json:"healthy_threshold,omitempty"`
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

	// Specifies the maximum interval (s) for health check.
	HealthcheckInterval int `json:"healthcheck_interval,omitempty"`

	// Specifies the ID of the listener to which the health check task belongs.
	ListenerId string `json:"listener_id,omitempty"`

	// Specifies the health check task ID.
	Id string `json:"id,omitempty"`

	// Specifies the health check protocol.
	HealthcheckProtocol string `json:"healthcheck_protocol,omitempty"`

	// Specifies the threshold at which the health check result is fail, that is, the number of consecutive failed health checks when the health check result of the backend ECS changes from success to fail.
	UnhealthyThreshold int `json:"unhealthy_threshold,omitempty"`

	// Specifies the time when information about the health check task was updated.
	UpdateTime string `json:"update_time,omitempty"`

	// Specifies the time when the health check task was created.
	CreateTime string `json:"create_time,omitempty"`

	// Specifies the port for health check.
	HealthcheckConnectPort int `json:"healthcheck_connect_port,omitempty"`

	// Specifies the maximum timeout duration for health check.
	HealthcheckTimeout int `json:"healthcheck_timeout,omitempty"`

	// Specifies the URI for health check. This parameter is valid when healthcheck_ protocol is HTTP.
	HealthcheckPath string `json:"healthcheck_path,omitempty"`

	// Specifies the threshold at which the health check result is success, that is, the number of consecutive successful health checks when the health check result of the backend ECS changes from fail to success.
	HealthyThreshold int `json:"healthy_threshold,omitempty"`
}
