# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/lifecyclehooks"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 查询生命周期挂钩列表。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := lifecyclehooks.List(client, tenantId, groupId).Extract()
    if err != nil {
        panic(err)
    }
    
示例代码, 通过生命周期操作令牌或者通过实例ID和生命周期挂钩名称对伸缩实例指定的挂钩进行回调操作。1、如果在超时时间结束前已完成自定义操作，选择终止或继续完成生命周期操作。2、如果需要更多时间完成自定义操作，选择延长超时时间，实例保持等待状态的时间将增加1小时。3、只有实例的生命周期挂钩状态为 HANGING 时才可以进行回调操作。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := lifecyclehooks.ListWithSuspension(client, tenantId, groupId, lifecyclehooks.ListWithSuspensionOpts{
        InstanceId: "6abadacf-87af-4225-b762-4a56853aec02",
    }).Extract()
    if err != nil {
        panic(err)
    }
    
示例代码, 通过生命周期操作令牌或者通过实例ID和生命周期挂钩名称对伸缩实例指定的挂钩进行回调操作。1、如果在超时时间结束前已完成自定义操作，选择终止或继续完成生命周期操作。2、如果需要更多时间完成自定义操作，选择延长超时时间，实例保持等待状态的时间将增加1小时。3、只有实例的生命周期挂钩状态为 HANGING 时才可以进行回调操作。

    
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
    
示例代码, 删除一个指定生命周期挂钩。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    lifecycleHookName := "test-hook1"
    result := lifecyclehooks.Delete(client, tenantId, groupId, lifecycleHookName)
    
    if err != nil {
        panic(err)
    }
    
示例代码, 创建生命周期挂钩，可为伸缩组添加一个或多个生命周期挂钩，最多添加5个。添加生命周期挂钩后，当伸缩组进行伸缩活动时，实例将被生命周期挂钩挂起并置于等待状态（正在加入伸缩组或正在移出伸缩组），实例将保持此状态直至超时时间结束或者用户手动回调。用户能够在实例保持等待状态的时间段内执行自定义操作，例如，用户可以在新启动的实例上安装或配置软件，也可以在实例终止前从实例中下载日志文件。

    
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
    
示例代码, 修改一个指定生命周期挂钩中的信息。

    
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
    
示例代码, 查询一个指定生命周期挂钩详情。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    lifecycleHookName := "test-hook1"
    result, err := lifecyclehooks.Get(client, tenantId, groupId, lifecycleHookName).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func CallBack(*golangsdk.ServiceClient, string, string, CallBackOptsBuilder) (CallBackResult)](#func-callback)**  
**[func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, string) (ListResult)](#func-list)**  
**[func ListWithSuspension(*golangsdk.ServiceClient, string, string, ListWithSuspensionOptsBuilder) (ListWithSuspensionResult)](#func-listwithsuspension)**  
**[func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func CallBack(*golangsdk.ServiceClient, string, string, CallBackOptsBuilder) (CallBackResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_instance_hook/{scaling_group_id}/callback|
|as|func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}|
|as|func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}|
|as|func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}|
|as|func List(*golangsdk.ServiceClient, string, string) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/list|
|as|func ListWithSuspension(*golangsdk.ServiceClient, string, string, ListWithSuspensionOptsBuilder) (ListWithSuspensionResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_instance_hook/{scaling_group_id}/list|
|as|func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_lifecycle_hook/{scaling_group_id}/{lifecycle_hook_name}|
## 开始
## func CallBack
    func CallBack(*golangsdk.ServiceClient, string, string, CallBackOptsBuilder) (CallBackResult)  
通过生命周期操作令牌或者通过实例ID和生命周期挂钩名称对伸缩实例指定的挂钩进行回调操作。1、如果在超时时间结束前已完成自定义操作，选择终止或继续完成生命周期操作。2、如果需要更多时间完成自定义操作，选择延长超时时间，实例保持等待状态的时间将增加1小时。3、只有实例的生命周期挂钩状态为 HANGING 时才可以进行回调操作。
## func Create
    func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)  
创建生命周期挂钩，可为伸缩组添加一个或多个生命周期挂钩，最多添加5个。添加生命周期挂钩后，当伸缩组进行伸缩活动时，实例将被生命周期挂钩挂起并置于等待状态（正在加入伸缩组或正在移出伸缩组），实例将保持此状态直至超时时间结束或者用户手动回调。用户能够在实例保持等待状态的时间段内执行自定义操作，例如，用户可以在新启动的实例上安装或配置软件，也可以在实例终止前从实例中下载日志文件。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)  
删除一个指定生命周期挂钩。
## func Get
    func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)  
查询一个指定生命周期挂钩详情。
## func List
    func List(*golangsdk.ServiceClient, string, string) (ListResult)  
查询生命周期挂钩列表。
## func ListWithSuspension
    func ListWithSuspension(*golangsdk.ServiceClient, string, string, ListWithSuspensionOptsBuilder) (ListWithSuspensionResult)  
通过生命周期操作令牌或者通过实例ID和生命周期挂钩名称对伸缩实例指定的挂钩进行回调操作。1、如果在超时时间结束前已完成自定义操作，选择终止或继续完成生命周期操作。2、如果需要更多时间完成自定义操作，选择延长超时时间，实例保持等待状态的时间将增加1小时。3、只有实例的生命周期挂钩状态为 HANGING 时才可以进行回调操作。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)  
修改一个指定生命周期挂钩中的信息。
