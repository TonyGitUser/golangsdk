package backendces

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

	// Specifies the URI of the task for adding a backend ECS. It is returned by Combined API.
	Uri string `json:"uri,omitempty"`

	// Specifies the unique ID assigned to the task for adding a backend ECS in Combined API.
	JobId string `json:"job_id,omitempty"`
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

	// Specifies the URI of the task for removing a backend ECS. It is returned by Combined API.
	Uri string `json:"uri,omitempty"`

	// Specifies the unique ID assigned to the task for removing a backend ECS in Combined API.
	JobId string `json:"job_id,omitempty"`
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response.Items)
	return &response, err
}

type ListResponse struct {
	Items []struct {

		// Specifies the private IP address of the backend ECS.
		ServerAddress string `json:"server_address,omitempty"`

		// Specifies the backend ECS ID.
		Id string `json:"id,omitempty"`

		// Specifies the floating IP address assigned to the backend ECS.
		Address string `json:"address,omitempty"`

		// Specifies the backend ECS status. The value is ACTIVE, PENDING, or ERROR.
		Status string `json:"status,omitempty"`

		// Specifies the health check status. The value is NORMAL, ABNORMAL, or UNAVAILABLE.
		HealthStatus string `json:"health_status,omitempty"`

		// Specifies the time when information about the backend ECS was updated.
		UpdateTime string `json:"update_time,omitempty"`

		// Specifies the time when the backend ECS was created.
		CreateTime string `json:"create_time,omitempty"`

		// Specifies the backend ECS name.
		ServerName string `json:"server_name,omitempty"`

		// Specifies the original back member ID.
		ServerId string `json:"server_id,omitempty"`

		// Specifies the listener to which the backend ECS belongs.
		Listeners []struct {

			// Specifies the ID of the listener to which the backend ECS belongs.
			Id string `json:"id,omitempty"`
		} `json:"listeners,omitempty"`
	}
}
