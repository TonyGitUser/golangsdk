# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/logs"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 查询伸缩活动日志。分页查询，根据输入的条件过滤。

    
    tenantId  := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := logs.List(client, tenantId, groupId, logs.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_activity_log/{scaling_group_id}|
## 开始
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
查询伸缩活动日志。分页查询，根据输入的条件过滤。
