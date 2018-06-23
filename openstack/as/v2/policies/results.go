package policies

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

	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id,omitempty"`
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

	// Specifies details about the AS policy.
	ScalingPolicy struct {

		// Specifies the scaling resource ID.
		ScalingResourceId string `json:"scaling_resource_id,omitempty"`

		// Specifies the scaling resource type.AS group: SCALING_GROUP,Bandwidth: BANDWIDTH
		ScalingResourceType string `json:"scaling_resource_type,omitempty"`

		// Specifies the AS policy name.
		ScalingPolicyName string `json:"scaling_policy_name,omitempty"`

		// Specifies the AS policy ID.
		ScalingPolicyId string `json:"scaling_policy_id,omitempty"`

		// Specifies the AS policy status.INSERVICE: indicates that the AS policy is in use.PAUSED: indicates that the AS policy is disabled.EXECUTING: The task is in process.
		PolicyStatus string `json:"policy_status,omitempty"`

		// Specifies the AS policy type.ALARM: indicates that the scaling action is triggered by an alarm. A value is returned for alarm_id, and no value is returned for scheduled_policy.SCHEDULED: indicates that the scaling action is triggered as scheduled. A value is returned for scheduled_policy, and no value is returned for alarm_id, recurrence_type, recurrence_value, start_time, or end_time.RECURRENCE: indicates that the scaling action is triggered periodically. Values are returned for scheduled_policy, recurrence_type, recurrence_value, start_time, and end_time, and no value is returned for alarm_id.
		ScalingPolicyType string `json:"scaling_policy_type,omitempty"`

		// Specifies the alarm ID.
		AlarmId string `json:"alarm_id,omitempty"`

		// Specifies the periodic or scheduled AS policy.
		ScheduledPolicy struct {

			// Specifies the time when the scaling action is triggered. The time format must comply with UTC.If scaling_policy_type is set to SCHEDULED, the time format is YYYY-MM-DDThh:mmZ.If scaling_policy_type is set to RECURRENCE, the time format is hh:mm.
			LaunchTime string `json:"launch_time,omitempty"`

			// Specifies the periodic triggering type. This parameter is mandatory when scaling_policy_type is set to RECURRENCE.Daily: indicates that the scaling action is triggered once a day.Weekly: indicates that the scaling action is triggered once a week.Monthly indicates that the scaling action is triggered once a month.
			RecurrenceType string `json:"recurrence_type,omitempty"`

			// Specifies the frequency at which scaling actions are triggered.If recurrence_type is set to Daily, the value is null, indicating that the scaling action is triggered once a day.If recurrence_type is set to Weekly, the value ranges from 1 (Sunday) to 7 (Saturday). The digits refer to dates in each week and separated by a comma. For example, 1,3,5.If recurrence_type is set to Monthly, the value ranges from 1 to 31. The digits refer to the dates in each month and separated by a comma, such as 1,10,13,28.
			RecurrenceValue string `json:"recurrence_value,omitempty"`

			// Specifies the start time of the scaling action triggered periodically. The time format complies with UTC.The current time is used by default.The time format is YYYY-MM-DDThh:mmZ.
			StartTime string `json:"start_time,omitempty"`

			// Specifies the end time of the scaling action triggered periodically. The time format complies with UTC. This parameter is mandatory when scaling_policy_type is set to RECURRENCE.When the scaling action is triggered periodically, the end time cannot be earlier than the current and start time.The time format is YYYY-MM-DDThh:mmZ
			EndTime string `json:"end_time,omitempty"`
		} `json:"scheduled_policy,omitempty"`

		// Specifies the scaling action of the AS policy.
		ScalingPolicyAction struct {

			// Specifies the operation to be performed. The default operation is ADD.ADD: adds specified number of instances to the AS group.REMOVE: removes specified number of instances from the AS group.SET: sets the number of instances in the AS group.
			Operation string `json:"operation,omitempty"`

			// Specifies the number of instances or the bandwidth.
			Size int `json:"size,omitempty"`

			// Specifies the percentage of instances to be operated.
			Percentage int `json:"percentage,omitempty"`

			// Specifies the operation restrictions.
			Limits int `json:"limits,omitempty"`
		} `json:"scaling_policy_action,omitempty"`

		// Specifies the cooling duration (s).
		CoolDownTime int `json:"cool_down_time,omitempty"`

		// Specifies the time when an AS policy is created. The time format complies with UTC.
		CreateTime string `json:"create_time,omitempty"`
	} `json:"scaling_policy,omitempty"`
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

	// Specifies the total number of query records.
	Limit int `json:"limit,omitempty"`

	// Specifies the AS policy list.
	ScalingPolicies []struct {

		// Specifies the AS policy name.
		ScalingPolicyName string `json:"scaling_policy_name,omitempty"`

		// Specifies the AS policy ID.
		ScalingPolicyId string `json:"scaling_policy_id,omitempty"`

		// Specifies the AS policy ID.
		ScalingResourceId string `json:"scaling_resource_id,omitempty"`

		// Specifies the scaling resource type.AS group: SCALING_GROUP.Bandwidth: BANDWIDTH
		ScalingResourceType string `json:"scaling_resource_type,omitempty"`

		// Specifies the AS policy status.INSERVICE: indicates that the AS policy is in use.PAUSED: indicates that the AS policy is disabled.
		PolicyStatus string `json:"policy_status,omitempty"`

		// Specifies the AS policy type.ALARM: indicates that the scaling action is triggered by an alarm. A value is returned for alarm_id, and no value is returned for scheduled_policy.SCHEDULED: indicates that the scaling action is triggered as scheduled. A value is returned for scheduled_policy, and no value is returned for alarm_id, recurrence_type, recurrence_value, start_time, and end_time.RECURRENCE: indicates that the scaling action is triggered periodically. Values are returned for scheduled_policy, recurrence_type, recurrence_value, start_time, and end_time, and no value is returned for alarm_id.
		ScalingPolicyType string `json:"scaling_policy_type,omitempty"`

		// Specifies the alarm ID.
		AlarmId string `json:"alarm_id,omitempty"`

		// Specifies the periodic or scheduled AS policy.
		ScheduledPolicy string `json:"scheduled_policy,omitempty"`

		// Specifies the action of the AS policy.
		ScalingPolicyAction string `json:"scaling_policy_action,omitempty"`

		// Specifies the cooling duration (s).
		CoolDownTime string `json:"cool_down_time,omitempty"`

		// Specifies the time when an AS policy is created. The time format must comply with UTC.
		CreateTime string `json:"create_time,omitempty"`
	} `json:"scaling_policies,omitempty"`
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

	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id,omitempty"`
}
