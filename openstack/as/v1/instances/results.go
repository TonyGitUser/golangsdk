package instances

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type ActionResult struct {
	commonResult
}

func (r ActionResult) Extract() (*ActionResponse, error) {
	var response ActionResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ActionResponse struct {
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

	// Specifies the total number of query records.
	TotalNumber int `json:"total_number,omitempty"`

	// Specifies the start line number.
	StartNumber int `json:"start_number,omitempty"`

	// Specifies the number of query records.
	Limit int `json:"limit,omitempty"`

	// Specifies details about the instances in the AS group.
	ScalingGroupInstances []struct {

		// Specifies the instance ID.
		InstanceId string `json:"instance_id,omitempty"`

		// Specifies the instance name.
		InstanceName string `json:"instance_name,omitempty"`

		// Specifies the ID of the AS group to which the instance belongs.
		ScalingGroupId string `json:"scaling_group_id,omitempty"`

		// Specifies the name of the AS group to which the instance belongs.
		ScalingGroupName string `json:"scaling_group_name,omitempty"`

		// Specifies the instance lifecycle status in the AS group.INSERVICE: The instance in the AS group is in use.PENDING: The instance is being added to the AS group.PENDING_WAIT: The instance is waiting to be added to the AS group.REMOVING: The instance is being removed from the AS group.REMOVING_WAIT: The instance is waiting to be removed from the AS group.
		LifeCycleState string `json:"life_cycle_state,omitempty"`

		// Specifies the instance health status.The status can be NORMAL or ERROR.
		HealthStatus string `json:"health_status,omitempty"`

		// Specifies the AS configuration name.If the AS configuration has been deleted, no information is displayed.If the instance is manually added to the AS group, MANNUAL_ADD is returned.
		ScalingConfigurationName string `json:"scaling_configuration_name,omitempty"`

		// Specifies the AS configuration ID.
		ScalingConfigurationId string `json:"scaling_configuration_id,omitempty"`

		// Specifies the time when the instance is added to the AS group. The time format complies with UTC.
		CreateTime string `json:"create_time,omitempty"`

		// Specifies the instance protection status.
		ProtectFromScalingDown bool `json:"protect_from_scaling_down,omitempty"`
	} `json:"scaling_group_instances,omitempty"`
}
