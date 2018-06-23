package policylogs

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

	// Specifies the logs of AS policy execution.
	ScalingPolicyExecuteLog []struct {

		// Specifies the AS policy execution status.SUCCESS: indicates that the AS policy is successfully executed.FAIL: indicates that the AS policy failed to be executed.EXECUTING: The task is in process.
		Status string `json:"status,omitempty"`

		// Specifies the AS policy execution failure.
		FailedReason string `json:"failed_reason,omitempty"`

		// Specifies the AS policy execution type.SCHEDULE: automatically triggered at a specified time point,RECURRENCE: automatically triggered at a specified time period,ALARM: alarm-triggered,MANUAL: manually triggered
		ExecuteType string `json:"execute_type,omitempty"`

		// Specifies the time when an AS policy is executed. The time format complies with UTC.
		ExecuteTime string `json:"execute_time,omitempty"`

		// Specifies the ID of an AS policy execution log.
		Id string `json:"id,omitempty"`

		// Specifies the tenant ID.
		TenantId string `json:"tenant_id,omitempty"`

		// Specifies the AS policy ID.
		ScalingPolicyId string `json:"scaling_policy_id,omitempty"`

		// Specifies the scaling resource type.AS group: SCALING_GROUP,Bandwidth: BANDWIDTH
		ScalingResourceType string `json:"scaling_resource_type,omitempty"`

		// Specifies the scaling resource ID.
		ScalingResourceId string `json:"scaling_resource_id,omitempty"`

		// Specifies the source value.
		OldValue string `json:"old_value,omitempty"`

		// Specifies the target value.
		DesireValue string `json:"desire_value,omitempty"`

		// Specifies the AS policy execution type.ADD: indicates adding instances.REMOVE: indicates reducing instances.SET: indicates setting the number of instances to a specified value.
		Type string `json:"type,omitempty"`

		// Specifies the tasks contained in a scaling action based on an AS policy.
		JobRecords []struct {

			// Specifies the task name.
			JobName string `json:"job_name,omitempty"`

			// Specifies the record type.API: API calling type,MEG: message type
			RecordType string `json:"record_type,omitempty"`

			// Specifies the record time.
			RecordTime string `json:"record_time,omitempty"`

			// Specifies the request body. This parameter is valid only if record_type is set to API.
			Request string `json:"request,omitempty"`

			// Specifies the response body. This parameter is valid only if record_type is set to API.
			Response string `json:"response,omitempty"`

			// Specifies the returned code. This parameter is valid only if record_type is set to API.
			Code int `json:"code,omitempty"`

			// Specifies the message. This parameter is valid only if record_type is set to MEG.
			Message string `json:"message,omitempty"`

			// Specifies the execution status of the task.SUCCESS: The task is successfully executed.FAIL: The task failed to be executed.
			JobStatus string `json:"job_status,omitempty"`
		} `json:"job_records,omitempty"`
	} `json:"scaling_policy_execute_log,omitempty"`
}
