# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/quotas"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


## 目录
**[func List(*golangsdk.ServiceClient, string) (ListResult)](#func-list)**  
**[func ListWithInstances(*golangsdk.ServiceClient, string, string) (ListWithInstancesResult)](#func-listwithinstances)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func List(*golangsdk.ServiceClient, string) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/quotas|
|as|func ListWithInstances(*golangsdk.ServiceClient, string, string) (ListWithInstancesResult)|GET /autoscaling-api/v1/{tenant_id}/quotas/{scaling_group_id}|
## 开始
## func List
    func List(*golangsdk.ServiceClient, string) (ListResult)  
查询指定租户的弹性伸缩组和伸缩配置的配额总数及已使用配额数。查询指定租户的弹性伸缩策略和伸缩实例的配额总数。
## func ListWithInstances
    func ListWithInstances(*golangsdk.ServiceClient, string, string) (ListWithInstancesResult)  
查询一个指定弹性伸缩组下的弹性伸缩策略和弹性伸缩实例的配额总数及已使用配额数。
