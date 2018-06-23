# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/lifecyclehooks"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to query lifecycle hooks.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := lifecyclehooks.List(client, tenantId, groupId).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to call back the lifecycle hook specified by a scaling instance based on the lifecycle action key or based on the instance ID and lifecycle hook name.If you have completed customized operations before the timeout duration expires, you can terminate or continue the lifecycle operations.If you require more time for customized operations, you can prolong the timeout duration. Then, the instance waiting duration will be prolonged by one hour each time.The callback operation can be performed only when the lifecycle hook of the target instance is in HANGING state.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := lifecyclehooks.ListWithSuspension(client, tenantId, groupId, lifecyclehooks.ListWithSuspensionOpts{
        InstanceId: "6abadacf-87af-4225-b762-4a56853aec02",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to call back the lifecycle hook specified by a scaling instance based on the lifecycle action key or based on the instance ID and lifecycle hook name.If you have completed customized operations before the timeout duration expires, you can terminate or continue the lifecycle operations.If you require more time for customized operations, you can prolong the timeout duration. Then, the instance waiting duration will be prolonged by one hour each time.The callback operation can be performed only when the lifecycle hook of the target instance is in HANGING state.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result := lifecyclehooks.CallBack(client, tenantId, groupId, lifecyclehooks.CallBackOpts{
        LifecycleActionResult: "ABANDON",
        InstanceId:            "e1040c67-b130-41ee-9405-07c6c6c20208",
        LifecycleHookName:     "test-hook1",
    })
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a specified lifecycle hook.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    lifecycleHookName := "test-hook1"
    result := lifecyclehooks.Delete(client, tenantId, groupId, lifecycleHookName)
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to create a lifecycle hook for an AS group. Up to five lifecycle hooks can be created for one AS group.After the creation, when the AS group performs a scaling action, the lifecycle hook suspends the target instance and sets it to be in Wait (Adding to AS group) or Wait (Removing from AS group) status. This status retains until the timeout duration expires or you manually call back this status.During the instance waiting duration, you can perform customized operations. For example, you can install or configure software on a newly started instance, or download the log file from the instance before the instance terminates.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := lifecyclehooks.Create(client, tenantId, groupId, lifecyclehooks.CreateOpts{
        LifecycleHookName:    "test-hook1",
        DefaultResult:        "ABANDON",
        DefaultTimeout:       3600,
        NotificationTopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
        LifecycleHookType:    "INSTANCE_LAUNCHING",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to modify the information about a specified lifecycle hook.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    lifecyclehookName := "test-hook1"
    result, err := lifecyclehooks.Update(client, tenantId, groupId, lifecyclehookName, lifecyclehooks.UpdateOpts{
        DefaultResult:        "CONTINUE",
        DefaultTimeout:       1800,
        NotificationTopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
        LifecycleHookType:    "INSTANCE_LAUNCHING",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query details about a specified lifecycle hook.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    lifecycleHookName := "test-hook1"
    result, err := lifecyclehooks.Get(client, tenantId, groupId, lifecycleHookName).Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func CallBack(*golangsdk.ServiceClient, string, string, CallBackOptsBuilder) (CallBackResult)](#func-callback)**  
**[func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, string) (ListResult)](#func-list)**  
**[func ListWithSuspension(*golangsdk.ServiceClient, string, string, ListWithSuspensionOptsBuilder) (ListWithSuspensionResult)](#func-listwithsuspension)**  
**[func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|as|func CallBack(*golangsdk.ServiceClient, string, string, CallBackOptsBuilder) (CallBackResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_instance_hook/{scaling_group_id}/callback|
|as|func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}|
|as|func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}|
|as|func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}|
|as|func List(*golangsdk.ServiceClient, string, string) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/list|
|as|func ListWithSuspension(*golangsdk.ServiceClient, string, string, ListWithSuspensionOptsBuilder) (ListWithSuspensionResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_instance_hook/{scaling_group_id}/list|
|as|func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}|
## Content
## func CallBack
    func CallBack(*golangsdk.ServiceClient, string, string, CallBackOptsBuilder) (CallBackResult)  
This interface is used to call back the lifecycle hook specified by a scaling instance based on the lifecycle action key or based on the instance ID and lifecycle hook name.If you have completed customized operations before the timeout duration expires, you can terminate or continue the lifecycle operations.If you require more time for customized operations, you can prolong the timeout duration. Then, the instance waiting duration will be prolonged by one hour each time.The callback operation can be performed only when the lifecycle hook of the target instance is in HANGING state.
## func Create
    func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)  
This interface is used to create a lifecycle hook for an AS group. Up to five lifecycle hooks can be created for one AS group.After the creation, when the AS group performs a scaling action, the lifecycle hook suspends the target instance and sets it to be in Wait (Adding to AS group) or Wait (Removing from AS group) status. This status retains until the timeout duration expires or you manually call back this status.During the instance waiting duration, you can perform customized operations. For example, you can install or configure software on a newly started instance, or download the log file from the instance before the instance terminates.
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)  
This interface is used to delete a specified lifecycle hook.
## func Get
    func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)  
This interface is used to query details about a specified lifecycle hook.
## func List
    func List(*golangsdk.ServiceClient, string, string) (ListResult)  
This interface is used to query lifecycle hooks.
## func ListWithSuspension
    func ListWithSuspension(*golangsdk.ServiceClient, string, string, ListWithSuspensionOptsBuilder) (ListWithSuspensionResult)  
This interface is used to call back the lifecycle hook specified by a scaling instance based on the lifecycle action key or based on the instance ID and lifecycle hook name.If you have completed customized operations before the timeout duration expires, you can terminate or continue the lifecycle operations.If you require more time for customized operations, you can prolong the timeout duration. Then, the instance waiting duration will be prolonged by one hour each time.The callback operation can be performed only when the lifecycle hook of the target instance is in HANGING state.
## func Update
    func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to modify the information about a specified lifecycle hook.
