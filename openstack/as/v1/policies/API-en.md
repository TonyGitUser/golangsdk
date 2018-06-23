# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/policies"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to immediately execute, enable, or disable a specified AS policy.An AS policy can be executed only when the AS group and AS policy are in the INSERVICE state. Otherwise, the execution fails.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
    result, err := policies.Action(client, tenantId, policyId, policies.ActionOpts{
        Action: "resume",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to create an AS policy.If you set scaling actions to be triggered by alarms, the selected or created alarm can associate with only one AS group.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := policies.Create(client, tenantId, policies.CreateOpts{
        ScalingPolicyName: "as-policy-7a75",
        ScalingPolicyAction: struct {
            Operation          string `json:"operation,omitempty"`
            InstanceNumber     int    `json:"instance_number,omitempty"`
            InstancePercentage int    `json:"instance_percentage,omitempty"`
        }{
            Operation:      "ADD",
            InstanceNumber: 1,
        },
        CoolDownTime: 900,
        ScheduledPolicy: struct {
            LaunchTime      string `json:"launch_time,omitempty"`
            RecurrenceType  string `json:"recurrence_type,omitempty"`
            RecurrenceValue string `json:"recurrence_value,omitempty"`
            StartTime       string `json:"start_time,omitempty"`
            EndTime         string `json:"end_time,omitempty"`
        }{
            LaunchTime:      "16:00",
            RecurrenceType:  "Daily",
            RecurrenceValue: "",
            StartTime:       "2015-12-14T03:34Z",
            EndTime:         "2019-12-27T03:34Z",
        },
        ScalingPolicyType: "RECURRENCE",
        ScalingGroupId:    "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a specified AS policy.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
    result := policies.Delete(client, tenantId, policyId)
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query details about a specified AS policy.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
    result, err := policies.Get(client, tenantId, policyId).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query information about AS policies. The results are filtered based on the conditions you input and displayed by page.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := policies.List(client, tenantId, scalingGroupId, policies.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
    
Sample Code, This interface is used to modify a specified AS policy.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
    result, err := policies.Update(client, tenantId, policyId, policies.UpdateOpts{
        ScalingPolicyType: "RECURRENCE",
        ScalingPolicyName: "policy_01",
        ScheduledPolicy: struct {
            LaunchTime      string `json:"launch_time,omitempty"`
            RecurrenceType  string `json:"recurrence_type,omitempty"`
            RecurrenceValue string `json:"recurrence_value,omitempty"`
            StartTime       string `json:"start_time,omitempty"`
            EndTime         string `json:"end_time,omitempty"`
        }{
            LaunchTime:      "16:00",
            RecurrenceType:  "Daily",
            RecurrenceValue: "",
            EndTime:         "2020-02-08T17:31Z",
            StartTime:       "2019-01-08T17:31Z",
        },
        CoolDownTime: 300,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)](#func-action)**  
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|as|func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}/action|
|as|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_policy|
|as|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}|
|as|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}|
|as|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_group_id}/list|
|as|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}|
## Content
## func Action
    func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)  
This interface is used to immediately execute, enable, or disable a specified AS policy.An AS policy can be executed only when the AS group and AS policy are in the INSERVICE state. Otherwise, the execution fails.
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
This interface is used to create an AS policy.If you set scaling actions to be triggered by alarms, the selected or created alarm can associate with only one AS group.
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
This interface is used to delete a specified AS policy.
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
This interface is used to query details about a specified AS policy.
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
This interface is used to query information about AS policies. The results are filtered based on the conditions you input and displayed by page.
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to modify a specified AS policy.
