# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/securitygroups"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
查询单租户在VPC服务下的网络资源配额,包括vpc配额、子网配额、安全组配额、安全组规则配额、安全组规则配额，弹性IP配额，vpn等。

示例代码, 创建安全组。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := securitygroups.Create(client, tenantID, securitygroups.CreateOpts{
        SecurityGroup: securitygroups.CreateSecurityGroup{
            Name:        "EricSG",
            Description: "Test SecurityGroup",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询单个安全组。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := securitygroups.Get(client, tenantID, "f7616338-fa30-42b8-bf6b-754c0701aab8").Extract()
    
    if err != nil {
      panic(err)
    }
    
示例代码, 查询安全组列表。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := securitygroups.List(client, tenantID, securitygroups.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除安全组。

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := securitygroups.Delete(client, tenantID, "2465d913-1084-4a6a-91e7-2fd6f490ecb3")
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/security-groups|
|vpc|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/security-groups/{security_group_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/security-groups/{security_group_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/security-groups|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
创建安全组。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除安全组。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询单个安全组。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
查询安全组列表。
