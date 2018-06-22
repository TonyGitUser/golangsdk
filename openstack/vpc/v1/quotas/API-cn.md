# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/quotas"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
查询单租户在VPC服务下的网络资源配额,包括vpc配额、子网配额、安全组配额、安全组规则配额、安全组规则配额，弹性IP配额，vpn等。

示例代码, 查询单租户在VPC服务下的网络资源配额,包括vpc配额、子网配额、安全组配额、安全组规则配额、安全组规则配额，弹性IP配额，vpn等。

    
    result, err := quotas.List(client.ServiceClient(), tenantID, quotas.ListOpts{
       Type: "vpc",
    }).Extract()
## 目录
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/quotas|
## 开始
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
查询单租户在VPC服务下的网络资源配额,包括vpc配额、子网配额、安全组配额、安全组规则配额、安全组规则配额，弹性IP配额，vpn等。
