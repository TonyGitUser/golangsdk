# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/policies"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 立即执行或启用或停止一个指定弹性伸缩策略。当伸缩组、伸缩策略状态处于INSERVICE时，伸缩策略才能被正确执行，否则会执行失败。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
    result, err := policies.Action(client, tenantId, policyId, policies.ActionOpts{
        Action: "resume",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
示例代码, 创建弹性伸缩策略。如果选择告警策略，则选择或者创建的告警只能关联一个弹性伸缩组。

    
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
    
示例代码, 删除一个指定弹性伸缩策略。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
    result := policies.Delete(client, tenantId, policyId)
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询指定弹性伸缩策略信息。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
    result, err := policies.Get(client, tenantId, policyId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询弹性伸缩策略信息。分页查询，根据输入的条件过滤。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := policies.List(client, tenantId, scalingGroupId, policies.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
    
示例代码, 修改指定弹性伸缩策略。

    
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
## 目录
**[func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)](#func-action)**  
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}/action|
|as|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_policy|
|as|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}|
|as|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}|
|as|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_group_id}/list|
|as|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_policy/{scaling_policy_id}|
## 开始
## func Action
    func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)  
立即执行或启用或停止一个指定弹性伸缩策略。当伸缩组、伸缩策略状态处于INSERVICE时，伸缩策略才能被正确执行，否则会执行失败。
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
创建弹性伸缩策略。如果选择告警策略，则选择或者创建的告警只能关联一个弹性伸缩组。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除一个指定弹性伸缩策略。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询指定弹性伸缩策略信息。
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
查询弹性伸缩策略信息。分页查询，根据输入的条件过滤。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
修改指定弹性伸缩策略。
