package lifecyclehooks

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type CallBackResult struct {
	commonResult
}

func (r CallBackResult) Extract() (*CallBackResponse, error) {
	var response CallBackResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CallBackResponse struct {
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

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name,omitempty"`

	// Specifies the lifecycle hook type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type,omitempty"`

	// Specifies the default lifecycle hook callback operation.ABANDON,CONTINUE
	DefaultResult string `json:"default_result,omitempty"`

	// Specifies the lifecycle hook timeout duration in the unit of second.
	DefaultTimeout int `json:"default_timeout,omitempty"`

	// Specifies a unique topic in SMN.
	NotificationTopicUrn string `json:"notification_topic_urn,omitempty"`

	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name,omitempty"`

	// Specifies the customized notification.
	NotificationMetadata string `json:"notification_metadata,omitempty"`
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

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name,omitempty"`

	// Specifies the lifecycle hook type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type,omitempty"`

	// Specifies the default lifecycle hook callback operation.ABANDON,CONTINUE
	DefaultResult string `json:"default_result,omitempty"`

	// Specifies the lifecycle hook timeout duration in the unit of second.
	DefaultTimeout int `json:"default_timeout,omitempty"`

	// Specifies a unique topic in SMN.
	NotificationTopicUrn string `json:"notification_topic_urn,omitempty"`

	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name,omitempty"`

	// Specifies the customized notification.
	NotificationMetadata string `json:"notification_metadata,omitempty"`

	// Specifies the time when the lifecycle hook is created. The time is UTC-compliant.
	CreateTime string `json:"create_time,omitempty"`
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

	// Specifies lifecycle hooks.
	LifecycleHooks []struct {

		// Specifies the lifecycle hook name.
		LifecycleHookName string `json:"lifecycle_hook_name,omitempty"`

		// Specifies the lifecycle hook type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
		LifecycleHookType string `json:"lifecycle_hook_type,omitempty"`

		// Specifies the default lifecycle hook callback operation.ABANDON,CONTINUE
		DefaultResult string `json:"default_result,omitempty"`

		// Specifies the lifecycle hook timeout duration in the unit of second.
		DefaultTimeout int `json:"default_timeout,omitempty"`

		// Specifies a unique topic in SMN.
		NotificationTopicUrn string `json:"notification_topic_urn,omitempty"`

		// Specifies the topic name in SMN.
		NotificationTopicName string `json:"notification_topic_name,omitempty"`

		// Specifies the customized notification.
		NotificationMetadata string `json:"notification_metadata,omitempty"`

		// Specifies the time when the lifecycle hook is created. The time is UTC-compliant.
		CreateTime string `json:"create_time,omitempty"`
	} `json:"lifecycle_hooks,omitempty"`
}

type ListWithSuspensionResult struct {
	commonResult
}

func (r ListWithSuspensionResult) Extract() (*ListWithSuspensionResponse, error) {
	var response ListWithSuspensionResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListWithSuspensionResponse struct {

	// Specifies lifecycle hook information about an AS instance.
	InstanceHangingInfo []struct {

		// Specifies the lifecycle hook name.
		LifecycleHookName string `json:"lifecycle_hook_name,omitempty"`

		// Specifies the lifecycle action key, which determines the lifecycle callback object.
		LifecycleActionKey string `json:"lifecycle_action_key,omitempty"`

		// Specifies the AS instance ID.
		InstanceId string `json:"instance_id,omitempty"`

		// Specifies the AS group ID.
		ScalingGroupId string `json:"scaling_group_id,omitempty"`

		// Specifies the lifecycle hook status.HANGING: suspends the instance.CONTINUE: continues the instance.ABANDON: terminates the instance.
		LifecycleHookStatus string `json:"lifecycle_hook_status,omitempty"`

		// Specifies the timeout duration in the format of "YYYY-MM-DDThh:mm:ssZ". The time is UTC-compliant.
		Timeout string `json:"timeout,omitempty"`

		// Specifies the default lifecycle hook callback operation.
		DefaultResult string `json:"default_result,omitempty"`
	} `json:"instance_hanging_info,omitempty"`
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

	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name,omitempty"`

	// Specifies the lifecycle hook type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type,omitempty"`

	// Specifies the default lifecycle hook callback operation.ABANDON,CONTINUE
	DefaultResult string `json:"default_result,omitempty"`

	// Specifies the lifecycle hook timeout duration in the unit of second.
	DefaultTimeout int `json:"default_timeout,omitempty"`

	// Specifies a unique topic in SMN.
	NotificationTopicUrn string `json:"notification_topic_urn,omitempty"`

	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name,omitempty"`

	// Specifies the customized notification.
	NotificationMetadata string `json:"notification_metadata,omitempty"`

	// Specifies the time when the lifecycle hook is created. The time is UTC-compliant.
	CreateTime string `json:"create_time,omitempty"`
}
