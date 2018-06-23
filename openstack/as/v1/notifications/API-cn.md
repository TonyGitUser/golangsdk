# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/notifications"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 查询弹性伸缩组通知列表。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := notifications.List(client, tenantId, scalingGroupId).Extract()
    if err != nil {
        panic(err)
    }
    
示例代码, 删除指定的弹性伸缩组中指定的通知。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    topicUrnId := "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1"
    result := notifications.Delete(client, tenantId, scalingGroupId, topicUrnId)
    
    if err != nil {
        panic(err)
    }
    
示例代码, 给弹性伸缩组配置通知功能。每调用一次该接口，伸缩组即配置一个通知主题及其通知场景，每个伸缩组最多可以增加5个主题。通知主题由用户事先在SMN创建，当通知主题对应的通知场景出现时，伸缩组会向用户发送通知。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := notifications.Enable(client, tenantId, scalingGroupId, notifications.EnableOpts{
        TopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
        TopicScene: [] string{
            "SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)](#func-delete)**  
**[func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)](#func-enable)**  
**[func List(*golangsdk.ServiceClient, string, string) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_notification/{scaling_group_id}/{topic_urn}|
|as|func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_notification/{scaling_group_id}|
|as|func List(*golangsdk.ServiceClient, string, string) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_notification/{scaling_group_id}|
## 开始
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)  
删除指定的弹性伸缩组中指定的通知。
## func Enable
    func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)  
给弹性伸缩组配置通知功能。每调用一次该接口，伸缩组即配置一个通知主题及其通知场景，每个伸缩组最多可以增加5个主题。通知主题由用户事先在SMN创建，当通知主题对应的通知场景出现时，伸缩组会向用户发送通知。
## func List
    func List(*golangsdk.ServiceClient, string, string) (ListResult)  
查询弹性伸缩组通知列表。
