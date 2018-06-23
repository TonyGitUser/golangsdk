package groups

import (
	"github.com/huaweicloud/golangsdk"
)

type ScalingGroup struct {

	// Specifies the name of the AS group.
	ScalingGroupName string `json:"scaling_group_name,omitempty"`

	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id,omitempty"`

	// Specifies the status of the AS group.
	ScalingGroupStatus string `json:"scaling_group_status,omitempty"`

	// Specifies the AS configuration ID.
	ScalingConfigurationId string `json:"scaling_configuration_id,omitempty"`

	// Specifies the AS configuration name.
	ScalingConfigurationName string `json:"scaling_configuration_name,omitempty"`

	// Specifies the number of current instances in the AS group.
	CurrentInstanceNumber int `json:"current_instance_number,omitempty"`

	// Specifies the expected number of instances in the AS group.
	DesireInstanceNumber int `json:"desire_instance_number,omitempty"`

	// Specifies the minimum number of instances in the AS group.
	MinInstanceNumber int `json:"min_instance_number,omitempty"`

	// Specifies the maximum number of instances in the AS group.
	MaxInstanceNumber int `json:"max_instance_number,omitempty"`

	// Specifies the cooling duration (s).
	CoolDownTime int `json:"cool_down_time,omitempty"`

	// Specifies the ID of a typical ELB listener. ELB listener IDs are separated using a comma (,).
	LbListenerId string `json:"lb_listener_id,omitempty"`

	// This field is reserved.
	LbaasListeners []string `json:"lbaas_listeners,omitempty"`

	// Specifies the AZ information.
	AvailableZones []string `json:"available_zones,omitempty"`

	// Specifies network information.
	Networks []struct {

		// Specifies the network ID.
		Id string `json:"id,omitempty"`
	} `json:"networks,omitempty"`

	// Specifies security group information.
	SecurityGroups []struct {

		// Specifies the ID of the security group.
		Id string `json:"id,omitempty"`
	} `json:"security_groups,omitempty"`

	// Specifies the time when an AS group was created. The time format complies with UTC.
	CreateTime string `json:"create_time,omitempty"`

	// Specifies the ID of the VPC to which the AS group belongs.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies details about the AS group.
	Detail string `json:"detail,omitempty"`

	// Specifies the scaling flag of the AS group.
	IsScaling bool `json:"is_scaling,omitempty"`

	// Specifies the health check method.
	HealthPeriodicAuditMethod string `json:"health_periodic_audit_method,omitempty"`

	// Specifies the health check interval.
	HealthPeriodicAuditTime int `json:"health_periodic_audit_time,omitempty"`

	// Specifies the instance removal policy.
	InstanceTerminatePolicy string `json:"instance_terminate_policy,omitempty"`

	// Specifies the notification mode.
	Notifications []string `json:"notifications,omitempty"`

	// Specifies whether to delete the EIP bound to the ECS when deleting the ECS.
	DeletePublicip bool `json:"delete_publicip,omitempty"`

	// This field is reserved.
	CloudLocationId string `json:"cloud_location_id,omitempty"`
}

type CreateOpts struct {

	// Specifies the AS group name. The name can contain letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	ScalingGroupName string `json:"scaling_group_name,omitempty"`

	// Specifies the AS configuration ID. You can obtain its value from the API used to query AS configurations. For details, see Querying AS Configurations.https://support.huaweicloud.com/en-us/api-as/en-us_topic_0043063056.html
	ScalingConfigurationId string `json:"scaling_configuration_id,omitempty"`

	// Specifies the expected number of instances. The default value is the minimum number of instances.The value ranges from the minimum number of instances to the maximum number of instances.
	DesireInstanceNumber int `json:"desire_instance_number,omitempty"`

	// Specifies the minimum number of instances. The default value is 0.
	MinInstanceNumber int `json:"min_instance_number,omitempty"`

	// Specifies the maximum number of instances. The default value is 0.
	MaxInstanceNumber int `json:"max_instance_number,omitempty"`

	// Specifies the cooling duration (in seconds). The value ranges from 0 to 86400, and is 900 by default.The cooling duration takes effect only for scaling actions triggered by alarms. Scaling actions executed manually as well as those scheduled to trigger periodically or at a particular time will not be affected.
	CoolDownTime int `json:"cool_down_time,omitempty"`

	// Specifies the ID of a typical ELB listener. The system supports the binding of up to three ELB listeners, the IDs of which are separated using a comma (,). You can use vpc_id to obtain a load balancer ID from the API used to query an ELB list. For details, see section "Querying Load Balancers" in the Elastic Load Balance API Reference. Then, use the load balancer ID to query the ELB listener list to obtain the ELB listener ID. For details, see section "Querying Listeners" in the Elastic Load Balance API Reference.
	LbListenerId string `json:"lb_listener_id,omitempty"`

	// Specifies the AZ information. The ECS associated with the scaling action will be created in the specified AZ. If you do not specify an AZ, the system automatically specifies one.For details, see section Regions and Endpoints.https://developer.huaweicloud.com/endpoint
	AvailableZones []string `json:"available_zones,omitempty"`

	// Specifies network information. The system supports up to five subnets. The first subnet transferred serves as the primary NIC of the ECS by default. You can use vpc_id to obtain the parameter value from the API used to query VPC subnets. For details, see section Subnet > Querying Subnets in the Virtual Private Cloud API Reference.
	Networks []struct {

		// Specifies the network ID.
		Id string `json:"id,omitempty"`
	} `json:"networks,omitempty"`

	// Specifies security group information. You can use vpc_id to obtain the parameter value from the API used to query VPC security groups. For details, see section "Querying Security Groups" in the Virtual Private Cloud API Reference.
	SecurityGroups []struct {

		// Specifies the ID of the security group.
		Id string `json:"id,omitempty"`
	} `json:"security_groups,omitempty"`

	// Specifies the VPC information. You can obtain its value from the API used to query VPCs. For details, see section "Querying VPCs" in the Virtual Private Cloud API Reference.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies the health check method for instances in the AS group. The health check methods include ELB_AUDIT and NOVA_AUDIT. If load balancing is configured, the default value of this parameter is ELB_AUDIT. Otherwise, the default value is NOVA_AUDIT.ELB_AUDIT refers to the ELB health check, which takes effect in an AS group that has a listener.NOVA_AUDIT refers to the health check delivered with AS.
	HealthPeriodicAuditMethod string `json:"health_periodic_audit_method,omitempty"`

	// Specifies the health check period for instances. The period has four options: 5 minutes (default), 15 minutes, 60 minutes, and 180 minutes.
	HealthPeriodicAuditTime int `json:"health_periodic_audit_time,omitempty"`

	// Specifies the instance removal policy.(Default) OLD_CONFIG_OLD_INSTANCE: The earlier-created instances that were selected from the instances created based on earlier-created configurations are removed first.OLD_CONFIG_NEW_INSTANCE: The later-created instances that were selected from the instances created based on earlier-created configurations are removed first.OLD_INSTANCE: The earlier-created instances are removed first.NEW_INSTANCE: The later-created instances are removed first.
	InstanceTerminatePolicy string `json:"instance_terminate_policy,omitempty"`

	// Specifies the notification mode.EMAIL refers to notification by email.This notification mode is going to be canceled. You are advised to configure the notification function for the AS group. See Notifications for details.https://support.huaweicloud.com/en-us/api-as/en-us_topic_0043063034.html
	Notifications []string `json:"notifications,omitempty"`

	// Specifies whether to delete the EIP bound to the ECS when deleting the ECS. If you do not want to delete the EIP, set this parameter to false. Then, the system only unbinds the EIP from the ECS and reserves the EIP.The value can be true or false. The default value is false.true: deletes the EIP bound to the ECS when deleting the ECS. When the ECS is charged in Yearly/Monthly mode, the EIP bound to the ECS will not be deleted when the ECS is deleted.false: only unbinds the EIP bound to the ECS when deleting the ECS
	DeletePublicip bool `json:"delete_publicip,omitempty"`
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

type DeleteOpts struct {

	// Specifies whether to forcibly delete an AS group. The value can be yes or no(default).
	ForceDelete string `q:"force_delete"`
}

type DeleteOptsBuilder interface {
	ToDeleteQuery() (string, error)
}

func (opts DeleteOpts) ToDeleteQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func Delete(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := DeleteURL(client, tenantId, scalingGroupId)
	if opts != nil {
		query, _ := opts.ToDeleteQuery()
		url += query
	}

	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

type EnableOpts struct {

	// Specifies a flag for enabling or disabling an AS group.resume: enables the AS group.pause: disables the AS group.
	Action string `json:"action,omitempty"`
}

type EnableOptsBuilder interface {
	ToEnableMap() (map[string]interface{}, error)
}

func (opts EnableOpts) ToEnableMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Enable(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts EnableOptsBuilder) (r EnableResult) {
	b, _ := opts.ToEnableMap()

	_, r.Err = client.Post(EnableURL(client, tenantId, scalingGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string) (r GetResult) {
	url := GetURL(client, tenantId, scalingGroupId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies the AS group name.
	ScalingGroupName string `q:"scaling_group_name"`

	// Specifies the AS configuration ID. You can obtain its value from the API used to query AS configurations. For details, see Querying AS Configurations.https://support.huaweicloud.com/en-us/api-as/en-us_topic_0043063056.html
	ScalingConfigurationId string `q:"scaling_configuration_id"`

	// Specifies the AS group status. The value can be INSERVICE, PAUSED, ERROR, or DELETING.
	ScalingGroupStatus string `q:"scaling_group_status"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the number of query records. The default value is 20.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, tenantId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId)
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

	// Specifies the AS group name. The name can contain letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	ScalingGroupName string `json:"scaling_group_name,omitempty"`

	// Specifies the expected number of instances.The value ranges from the minimum number of instances to the maximum number of instances.
	DesireInstanceNumber int `json:"desire_instance_number,omitempty"`

	// Specifies the minimum number of instances.
	MinInstanceNumber int `json:"min_instance_number,omitempty"`

	// Specifies the maximum number of instances, which is greater than or equal to the minimum number of instances.
	MaxInstanceNumber int `json:"max_instance_number,omitempty"`

	// Specifies the cooling duration (in seconds). The value ranges from 0 to 86400.
	CoolDownTime int `json:"cool_down_time,omitempty"`

	// Specifies the AZ information. The ECS associated with the scaling action will be created in the specified AZ. If you do not specify an AZ, the system automatically specifies one.For details, see section Regions and Endpoints.The value of this parameter can be changed only when all the following conditions are met:No scaling actions are triggered in the AS group.The number of instances in the AS group is 0.The AS group is not in service.
	AvailabilityZones []string `json:"availability_zones,omitempty"`

	// Specifies network information. The system supports up to five subnets. The first subnet transferred serves as the primary NIC of the ECS by default. You can use vpc_id to obtain the parameter value from the API used to query VPC subnets. For details, see section Subnet > Querying Subnets in the Virtual Private Cloud API Reference. The value of this parameter can be changed only when all the following conditions are met:No scaling actions are triggered in the AS group.The number of instances in the AS group is 0.The AS group is not in service.
	Networks []struct {

		// Specifies the network ID.
		Id string `json:"id,omitempty"`
	} `json:"networks,omitempty"`

	// Specifies security group information. You can use vpc_id to obtain the parameter value from the API used to query VPC security groups. For details, see section "Querying Security Groups" in the Virtual Private Cloud API Reference.The value of this parameter can be changed only when all the following conditions are met:No scaling actions are triggered in the AS group.The number of instances in the AS group is 0.The AS group is not in service.
	SecurityGroups []struct {

		// Specifies the ID of the security group.
		Id string `json:"id,omitempty"`
	} `json:"security_groups,omitempty"`

	// Specifies the ID of a typical ELB listener. The system supports the binding of up to three ELB listeners, the IDs of which are separated using a comma (,). You can use vpc_id to obtain a load balancer ID from the API used to query an ELB list. For details, see section "Querying Load Balancers" in the Elastic Load Balance API Reference. Then, use the load balancer ID to query the ELB listener list to obtain the ELB listener ID. For details, see section "Querying Listeners" in the Elastic Load Balance API Reference.The value of this parameter can be changed only when all the following conditions are met:No scaling actions are triggered in the AS group.The number of instances in the AS group is 0.The AS group is not in service.
	LbListenerId string `json:"lb_listener_id,omitempty"`

	// Specifies the health check method for instances in the AS group. The health check methods include ELB_AUDIT and NOVA_AUDIT.ELB_AUDIT refers to the ELB health check, which takes effect in an AS group that has a listener.NOVA_AUDIT refers to the health check delivered with AS.
	HealthPeriodicAuditMethod string `json:"health_periodic_audit_method,omitempty"`

	// Specifies the health check period for the instances in the AS group. The value can be 5 minutes, 15 minutes, 60 minutes, or 180 minutes.
	HealthPeriodicAuditTime int `json:"health_periodic_audit_time,omitempty"`

	// Specifies the instance removal policy.(Default) OLD_CONFIG_OLD_INSTANCE: The earlier-created instances that were selected from the instances created based on earlier-created configurations are removed first.OLD_CONFIG_NEW_INSTANCE: The later-created instances that were selected from the instances created based on earlier-created configurations are removed first.OLD_INSTANCE: The earlier-created instances are removed first.NEW_INSTANCE: The later-created instances are removed first.
	InstanceTerminatePolicy string `json:"instance_terminate_policy,omitempty"`

	// Specifies the AS configuration ID. You can obtain its value from the API used to query AS configurations. For details, see Querying AS Configurations.https://support.huaweicloud.com/en-us/api-as/en-us_topic_0043063056.html
	ScalingConfigurationId string `json:"scaling_configuration_id,omitempty"`

	// Specifies the notification mode.EMAIL refers to notification by email.This notification mode is going to be canceled. You are advised to configure the notification function for the AS group. See Notifications for details.https://support.huaweicloud.com/en-us/api-as/en-us_topic_0043063034.html
	Notifications []string `json:"notifications,omitempty"`

	// Specifies whether to delete the EIP bound to the ECS when deleting the ECS. If you do not want to delete the EIP, set this parameter to false. Then, the system only unbinds the EIP from the ECS and reserves the EIP.The value can be true or false. The default value is false.true: deletes the EIP bound to the ECS when deleting the ECS. When the ECS is charged in Yearly/Monthly mode, the EIP bound to the ECS will not be deleted when the ECS is deleted.false: only unbinds the EIP bound to the ECS when deleting the ECS
	DeletePublicip bool `json:"delete_publicip,omitempty"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, scalingGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
