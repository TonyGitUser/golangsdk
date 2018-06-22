# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/vpcs"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
虚拟私有云（Virtual Private Cloud，以下简称VPC），为弹性云服务器构建隔离的、用户自主配置和管理的虚拟网络环境，提升用户云中资源的安全性，简化用户的网络部署。

示例代码, 创建虚拟私有云。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.Create(client, tenantID, vpcs.CreateOpts{
        Vpc: vpcs.VPC{
            Name: "ABC",
            Cidr: "192.168.0.0/16",
        },
    }).Extract()
    
示例代码, 更新虚拟私有云。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.Update(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0", vpcs.UpdateOpts{
        Vpc: vpcs.VPC{
            Name: "ABC-back",
            Cidr: "192.168.0.0/24",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询虚拟私有云。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
        result, err := vpcs.Get(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0").Extract()
    
示例代码, 查询虚拟私有云列表。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.List(client, tenantID, vpcs.ListOpts{
       Limit: 2,
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
示例代码, 删除虚拟私有云。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := vpcs.Delete(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0")
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/vpcs|
|vpc|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/vpcs/{vpc_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/vpcs/{vpc_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/vpcs|
|vpc|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/{tenant_id}/vpcs/{vpc_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
创建虚拟私有云。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除虚拟私有云。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询虚拟私有云。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
查询虚拟私有云列表。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
更新虚拟私有云。
