# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/securitygroups"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
查询单租户在VPC服务下的网络资源配额,包括vpc配额、子网配额、安全组配额、安全组规则配额、安全组规则配额，弹性IP配额，vpn等。

示例代码, 创建安全组。

    
    result, err := securitygroups.Create(client, securitygroups.CreateOpts{
        SecurityGroup: securitygroups.CreateSecurityGroup{
            Name:        "EricSG",
            Description: "Test SecurityGroup",
            TenantId: tenantID,
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询单个安全组。

    
    result, err := securitygroups.Get(client, "a988649d-cfbf-4c2a-bd91-3b84d2403747").Extract()
    
    if err != nil {
      panic(err)
    }
    
示例代码, 查询安全组列表。

    
    allPages, err := securitygroups.List(client, securitygroups.ListOpts{
        Limit: 2,
    }).AllPages()
    
    result, err := securitygroups.ExtractList(allPages.(securitygroups.ListPage))
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除安全组。

    
    result := securitygroups.Delete(client, tenantID, "2465d913-1084-4a6a-91e7-2fd6f490ecb3")
    
    if err != nil {
        panic(err)
    }
    
示例代码, 更新安全组。

    
    result, err := securitygroups.Update(client, "7af80d49-0a43-462d-aed8-a1e12ac91af6", securitygroups.UpdateOpts{
        SecurityGroup: securitygroups.UpdateSecurityGroup{
            Name:        "EricSG",
            Description: "Modified SecurityGroup",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/security-groups|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/security-groups/{security_group_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/security-groups/{security_group_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/security-groups|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/security-groups/{security_group_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建安全组。
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除安全组。
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询单个安全组。
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询安全组列表。
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
更新安全组。
