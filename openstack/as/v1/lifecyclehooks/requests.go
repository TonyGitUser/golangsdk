package lifecyclehooks

import (
	"github.com/huaweicloud/golangsdk"
)

type CallBackOpts struct {

	// Specifies the lifecycle action key.When specifying a lifecycle callback object, this field is mandatory if the instance_id field is not used. If both this field and the instance_id field are used, preferentially use this field for callback.
	LifecycleActionKey string `json:"lifecycle_action_key,omitempty"`

	// Specifies the instance ID.When specifying a lifecycle callback object, this field is mandatory if the lifecycle_action_key field is not used.
	InstanceId string `json:"instance_id,omitempty"`

	// Specifies the lifecycle hook name.When specifying a lifecycle callback object, this field is mandatory if the lifecycle_action_key field is not used.
	LifecycleHookName string `json:"lifecycle_hook_name,omitempty"`

	// Specifies the lifecycle callback action.ABANDON: terminates the instance.CONTINUE: continues the instance.EXTEND: extends the timeout duration, one hour each time.
	LifecycleActionResult string `json:"lifecycle_action_result,omitempty"`
}

type CallBackOptsBuilder interface {
	ToCallBackMap() (map[string]interface{}, error)
}

func (opts CallBackOpts) ToCallBackMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CallBack(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts CallBackOptsBuilder) (r CallBackResult) {
	b, _ := opts.ToCallBackMap()

	_, r.Err = client.Put(CallBackURL(client, tenantId, scalingGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type CreateOpts struct {

	// Specifies the lifecycle hook name. The name can contain letters, digits, underscores (_), and hyphens (-), and must be between 1 and 32 characters in length.
	LifecycleHookName string `json:"lifecycle_hook_name,omitempty"`

	// Specifies the lifecycle hook type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING,The INSTANCE_TERMINATING hook suspends an instance when the instance terminates. The INSTANCE_LAUNCHING hook suspends an instance when the instance starts.
	LifecycleHookType string `json:"lifecycle_hook_type,omitempty"`

	// Specifies the default lifecycle hook callback operation. By default, this operation is performed when the timeout duration expires.ABANDON,CONTINUE,If an instance is starting, CONTINUE indicates that your customized operations are successful and the instance can be used. ABANDON indicates that your customized operations failed, and the instance will be terminated. In such a case, the scaling action fails, and you must create a new instance.If an instance is stopping, both ABANDON and CONTINUE allow instance termination. The difference between the two states is as follows: ABANDON stops other lifecycle hooks, but CONTINUE allows the completion of other lifecycle hooks.The default value of this field is ABANDON.
	DefaultResult string `json:"default_result,omitempty"`

	// Specifies the lifecycle hook timeout duration, which ranges from 300 to 86400 in the unit of second. The default value is 3600.By default, this parameter specifies the instance waiting duration. You can prolong the timeout duration or perform the CONTINUE or ABANDON operation before the timeout duration expires.
	DefaultTimeout int `json:"default_timeout,omitempty"`

	// Specifies a unique topic in SMN.This parameter specifies a notification object for a lifecycle hook. When an instance is suspended by the lifecycle hook, the SMN service sends a notification to the object. This notification contains the basic instance information, your customized notification content, and the token for controlling lifecycle operations.
	NotificationTopicUrn string `json:"notification_topic_urn,omitempty"`

	// Specifies a customized notification, which contains no more than 256 characters in length. The message cannot contain the following characters: < > & ' ( ).After a notification object is configured, the SMN service sends your customized notification to the object.
	NotificationMetadata string `json:"notification_metadata,omitempty"`
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

func Create(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts CreateOptsBuilder) (r CreateResult) {
	b, _ := opts.ToCreateMap()

	_, r.Err = client.Post(CreateURL(client, tenantId, scalingGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, lifecycleHookName string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, scalingGroupId, lifecycleHookName)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, lifecycleHookName string) (r GetResult) {
	url := GetURL(client, tenantId, scalingGroupId, lifecycleHookName)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func List(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string) (r ListResult) {
	url := ListURL(client, tenantId, scalingGroupId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListWithSuspensionOpts struct {

	// Specifies the AS instance ID.
	InstanceId string `q:"instance_id"`
}

type ListWithSuspensionOptsBuilder interface {
	ToListWithSuspensionQuery() (string, error)
}

func (opts ListWithSuspensionOpts) ToListWithSuspensionQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func ListWithSuspension(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts ListWithSuspensionOptsBuilder) (r ListWithSuspensionResult) {
	url := ListWithSuspensionURL(client, tenantId, scalingGroupId)
	if opts != nil {
		query, _ := opts.ToListWithSuspensionQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the lifecycle hook type.INSTANCE_TERMINATING,INSTANCE_LAUNCHING,The INSTANCE_TERMINATING hook suspends an instance when the instance terminates. The INSTANCE_LAUNCHING hook suspends an instance when the instance starts.
	LifecycleHookType string `json:"lifecycle_hook_type,omitempty"`

	// Specifies the default lifecycle hook callback operation. By default, this operation is performed when the timeout duration expires.ABANDON,CONTINUE,If an instance is starting, CONTINUE indicates that your customized operations are successful and the instance can be used. ABANDON indicates that your customized operations failed, and the instance will be terminated. In such a case, the scaling action fails, and you must create a new instance.If an instance is stopping, both ABANDON and CONTINUE allow instance termination. The difference between the two states is as follows: ABANDON stops other lifecycle hooks, but CONTINUE allows the completion of other lifecycle hooks.The default value of this field is ABANDON.
	DefaultResult string `json:"default_result,omitempty"`

	// Specifies the lifecycle hook timeout duration, which ranges from 300 to 86400 in the unit of second. The default value is 3600.By default, this parameter specifies the instance waiting duration. You can prolong the timeout duration or perform the CONTINUE or ABANDON operation before the timeout duration expires.
	DefaultTimeout int `json:"default_timeout,omitempty"`

	// Specifies a unique topic in SMN.This parameter specifies a notification object for a lifecycle hook. When an instance is suspended by the lifecycle hook, the SMN service sends a notification to the object. This notification contains the basic instance information, your customized notification content, and the token for controlling lifecycle operations.
	NotificationTopicUrn string `json:"notification_topic_urn,omitempty"`

	// Specifies a customized notification, which contains no more than 256 characters in length. The message cannot contain the following characters: < > & ' ( ).After a notification object is configured, the SMN service sends your customized notification to the object.
	NotificationMetadata string `json:"notification_metadata,omitempty"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, lifecycleHookName string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, scalingGroupId, lifecycleHookName), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
