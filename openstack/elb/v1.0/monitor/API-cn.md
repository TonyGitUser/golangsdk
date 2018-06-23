# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/monitor"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
查询普通租户下所有负载均衡LVS 四层和七层流量的监控指标。本功能只能由普通租户操作。

示例代码, 查询普通租户下所有负载均衡LVS 四层和七层流量的监控指标。本功能只能由普通租户操作。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := monitor.List(client, tenantID).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func List(*golangsdk.ServiceClient, string) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|elb|func List(*golangsdk.ServiceClient, string) (ListResult)|GET /v1.0/{tenant_id}/elbaas/monitor|
## 开始
## func List
    func List(*golangsdk.ServiceClient, string) (ListResult)  
查询普通租户下所有负载均衡LVS 四层和七层流量的监控指标。本功能只能由普通租户操作。
