# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/policylogs"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


## 目录
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_policy_execute_log/{scaling_policy_id}|
## 开始
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
查询策略执行的历史记录。分页查询，根据输入的条件过滤。
