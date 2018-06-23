package policies

import (
	"github.com/huaweicloud/golangsdk"
)

type ActionOpts struct {

	// Specifies the operation flag for an AS group.execute: executes the AS group.resume: enables the AS group.pause: disables the AS group.
	Action string `json:"action,omitempty"`
}

type ActionOptsBuilder interface {
	ToActionMap() (map[string]interface{}, error)
}

func (opts ActionOpts) ToActionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Action(client *golangsdk.ServiceClient, tenantId string, scalingPolicyId string, opts ActionOptsBuilder) (r ActionResult) {
	b, _ := opts.ToActionMap()

	_, r.Err = client.Post(ActionURL(client, tenantId, scalingPolicyId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type CreateOpts struct {

	// Specifies the AS policy name. The name can contain letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	ScalingPolicyName string `json:"scaling_policy_name,omitempty"`

	// Specifies the AS group ID. You can obtain its value from the API used to query AS groups. For details, see Querying AS Groups.https://support.huaweicloud.com/en-us/api-as/en-us_topic_0043063030.html
	ScalingGroupId string `json:"scaling_group_id,omitempty"`

	// Specifies the AS policy type.ALARM (corresponding to alarm_id): indicates that the scaling action is triggered by an alarm.SCHEDULED (corresponding to scheduled_policy): indicates that the scaling action is triggered as scheduled.RECURRENCE (corresponding to scheduled_policy): indicates that the scaling action is triggered periodically.
	ScalingPolicyType string `json:"scaling_policy_type,omitempty"`

	// Specifies the alarm rule ID. This parameter is mandatory when scaling_policy_type is set to ALARM. After this parameter is specified, the value of scheduled_policy does not take effect.After you create an alarm policy, the system automatically adds an alarm triggering activity of the autoscaling type to the alarm_actions field in the alarm rule specified by the parameter value.You can obtain the parameter value by querying the CES alarm rule list. For details, see section Querying Alarms in the Cloud Eye API Reference.
	AlarmId string `json:"alarm_id,omitempty"`

	// Specifies the periodic or scheduled AS policy. This parameter is mandatory when scaling_policy_type is set to SCHEDULED or RECURRENCE. After this parameter is specified, the value of alarm_id does not take effect.
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

	// Specifies the action of the AS policy.
	ScalingPolicyAction struct {

		// Specifies the operation to be performed. The default operation is ADD.ADD: adds specified number of instances to the AS group.REMOVE: removes specified number of instances from the AS group.SET: sets the number of instances in the AS group.
		Operation string `json:"operation,omitempty"`

		// Specifies the number of instances to be operated. The default number is 1.Either instance_number or instance_percentage is required.
		InstanceNumber int `json:"instance_number,omitempty"`

		// Indicates the percentage of instances to be operated. You can increase/decrease or set the number of instances in an AS group to the specified percentage of the current number of instances.If neither instance_number nor instance_percentage is specified, the number of instances to be operated is 1.Either instance_number or instance_percentage is required.
		InstancePercentage int `json:"instance_percentage,omitempty"`
	} `json:"scaling_policy_action,omitempty"`

	// Specifies the cooling duration (in seconds), and is 900 by default.
	CoolDownTime int `json:"cool_down_time,omitempty"`
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

func Delete(client *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, scalingPolicyId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, scalingPolicyId string) (r GetResult) {
	url := GetURL(client, tenantId, scalingPolicyId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies the AS policy name.
	ScalingPolicyName string `q:"scaling_policy_name"`

	// Specifies the AS policy type.
	ScalingPolicyType string `q:"scaling_policy_type"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default value is 20 and the maximum value is 100.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId, scalingGroupId)
	if opts != nil {
		query, _ := opts.ToListQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the AS policy name. The name can contain letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	ScalingPolicyName string `json:"scaling_policy_name,omitempty"`

	// Specifies the AS policy type.ALARM (corresponding to alarm_id): indicates that the scaling action is triggered by an alarm.SCHEDULED (corresponding to scheduled_policy): indicates that the scaling action is triggered as scheduled.RECURRENCE (corresponding to scheduled_policy): indicates that the scaling action is triggered periodically.
	ScalingPolicyType string `json:"scaling_policy_type,omitempty"`

	// Specifies the alarm rule ID. This parameter is mandatory when scaling_policy_type is set to ALARM. After this parameter is specified, the value of scheduled_policy does not take effect.After you create an alarm policy, the system automatically adds an alarm triggering activity of the autoscaling type to the alarm_actions field in the alarm rule specified by the parameter value.You can obtain the parameter value by querying the CES alarm rule list. For details, see section Querying Alarms in the Cloud Eye API Reference.
	AlarmId string `json:"alarm_id,omitempty"`

	// Specifies the periodic or scheduled AS policy. This parameter is mandatory when scaling_policy_type is set to SCHEDULED or RECURRENCE. After this parameter is specified, the value of alarm_id does not take effect.
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

		// Specifies the number of instances to be operated. The default number is 1.Either instance_number or instance_percentage is required.
		InstanceNumber int `json:"instance_number,omitempty"`

		// Indicates the percentage of instances to be operated. You can increase/decrease or set the number of instances in an AS group to the specified percentage of the current number of instances.If neither instance_number nor instance_percentage is specified, the number of instances to be operated is 1.Either instance_number or instance_percentage is required.
		InstancePercentage int `json:"instance_percentage,omitempty"`
	} `json:"scaling_policy_action,omitempty"`

	// Specifies the cooling duration (in seconds), which is 900 by default.
	CoolDownTime int `json:"cool_down_time,omitempty"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, scalingPolicyId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, scalingPolicyId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
