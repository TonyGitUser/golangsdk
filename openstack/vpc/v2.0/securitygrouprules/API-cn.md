# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/securitygrouprules"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
查询单租户在VPC服务下的网络资源配额,包括vpc配额、子网配额、安全组配额、安全组规则配额、安全组规则配额，弹性IP配额，vpn等。

示例代码, 创建安全组规则。

    
    result, err := securitygrouprules.Create(client, securitygrouprules.CreateOpts{
       SecurityGroupRule: securitygrouprules.CreateSecurityGroupRule {
           Description: "Test SecurityGroup",
           TenantId: tenantID,
           SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
           Direction: "egress",
           Protocol: "tcp",
           RemoteIpPrefix: "10.10.0.0/24",
       },
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
示例代码, 查询单个安全组规则。

    
    result, err := securitygrouprules.Get(client, "26243298-ae79-46a3-bad9-34395762e033").Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询安全组规则列表。

    
    allPages, err := securitygrouprules.List(client, securitygrouprules.ListOpts{
        Limit: 2,
        Protocol: "tcp",
    }).AllPages()
    
    result, err := securitygrouprules.ExtractList(allPages.(securitygrouprules.ListPage))
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除安全组规则。

    
    result = securitygrouprules.Delete(client, "26243298-ae79-46a3-bad9-34395762e033")
## 目录
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/security-group-rules|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/security-group-rules/{security-groups-rules-id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/security-group-rules/{security-groups-rules-id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/security-group-rules|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建安全组规则。
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除安全组规则。
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询单个安全组规则。
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询安全组规则列表。
