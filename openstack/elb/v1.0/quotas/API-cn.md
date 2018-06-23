# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/quotas"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
查询负载均衡器的配额。

示例代码, 查询负载均衡器的配额。

## 目录
**[func List(*golangsdk.ServiceClient, string) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|elb|func List(*golangsdk.ServiceClient, string) (ListResult)|GET /v1.0/{tenant_id}/elbaas/quotas|
## 开始
## func List
    func List(*golangsdk.ServiceClient, string) (ListResult)  
查询负载均衡器的配额。
