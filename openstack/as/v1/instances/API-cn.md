# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/instances"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 批量移出伸缩组中的实例或批量添加伸缩组外的实例，或者批量对伸缩组中的实例设置或取消其实例保护属性。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := instances.Action(client, tenantId, groupId, instances.ActionOpts{
        InstancesId:     []string{
            "e1040c67-b130-41ee-9405-07c6c6c20208",
        },
        InstanceDelete: "no",
        Action:         "ADD",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询弹性伸缩组中实例信息。分页查询，根据输入的条件过滤。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := instances.List(client, tenantId, groupId, instances.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 移出一个指定弹性伸缩组中实例。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "e1040c67-b130-41ee-9405-07c6c6c20208"
    result := instances.Delete(client, tenantId, instanceId, instances.DeleteOpts{
       InstanceDelete: "no",
    })
    
    if err != nil {
       panic(err)
    }
## 目录
**[func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)](#func-action)**  
**[func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)](#func-delete)**  
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_group_instance/{scaling_group_id}/action|
|as|func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_group_instance/{instance_id}|
|as|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_group_instance/{scaling_group_id}/list|
## 开始
## func Action
    func Action(*golangsdk.ServiceClient, string, string, ActionOptsBuilder) (ActionResult)  
批量移出伸缩组中的实例或批量添加伸缩组外的实例，或者批量对伸缩组中的实例设置或取消其实例保护属性。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)  
移出一个指定弹性伸缩组中实例。
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
查询弹性伸缩组中实例信息。分页查询，根据输入的条件过滤。
