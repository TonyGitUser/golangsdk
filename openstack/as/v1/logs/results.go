package logs

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
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

	// Specifies the scaling action log list.
	ScalingActivityLog []struct {

		// Specifies the status of the scaling action.SUCCESS: indicates the scaling action is successfully performed.FAIL: indicates the action failed to be performed.DOING: indicates the scaling action is being performed.
		Status string `json:"status,omitempty"`

		// Specifies the start time of the scaling action. The time format must comply with UTC.
		StartTime string `json:"start_time,omitempty"`

		// Specifies the end time of the scaling action. The time format must comply with UTC.
		EndTime string `json:"end_time,omitempty"`

		// Specifies the scaling action log ID.
		Id string `json:"id,omitempty"`

		// Specifies the name list of the instances removed from the AS group after the scaling action is complete. The instance names are separated by commas (,).
		InstanceRemovedList string `json:"instance_removed_list,omitempty"`

		// Specifies the name list of the instances deleted from the AS group and deleted after the scaling action is complete. The instance names are separated by commas (,).
		InstanceDeletedList string `json:"instance_deleted_list,omitempty"`

		// Specifies the name list of the instances added to the AS group after the scaling action is complete. The instance names are separated by commas (,).
		InstanceAddedList string `json:"instance_added_list,omitempty"`

		// Specifies the number of added or deleted instances during the scaling.
		ScalingValue int `json:"scaling_value,omitempty"`

		// Specifies the description of the scaling action.
		Description string `json:"description,omitempty"`

		// Specifies the number of instances in the AS group.
		InstanceValue int `json:"instance_value,omitempty"`

		// Specifies the expected number of instances in the scaling action.
		DesireValue int `json:"desire_value,omitempty"`
	} `json:"scaling_activity_log,omitempty"`
}
