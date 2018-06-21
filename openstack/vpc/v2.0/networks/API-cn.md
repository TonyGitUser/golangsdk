# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/networks"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
对网络资源进行管理和操作，包括查询网络列表、创建网络、查询指定网络、删除网络以及更新网络等接口。

示例代码, 创建网络

    
    result, err := networks.Create(client, networks.CreateOpts{
        Network: networks.CreateNetwork{
            Name:         "NetWork Test",
            Shared:       false,
            AdminStateUp: true,
            TenantId:     "57e98940a77f4bb988a21a7d0603a626",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 更新网络

    
    result, err := networks.Update(client, "1c6af92e-bd06-4350-8f51-5ec32167814f", networks.UpdateOpts{
        Network: networks.UpdateNetwork{
            Name:   "Test Net Works Updated",
            Shared: true,
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询网络详情。

    
    result, err := networks.Get(client, "1c6af92e-bd06-4350-8f51-5ec32167814f").Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除网络

    
    result := networks.Delete(client, "1c6af92e-bd06-4350-8f51-5ec32167814f")
    
示例代码, 查询提交请求的租户的所有网络

    
    allPages, err := networks.List(client, networks.ListOpts{
        Limit: 2,
    }).AllPages()
    
    result, err := networks.ExtractList(allPages.(networks.ListPage))
    
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
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/networks|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/networks/{network-id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/networks/{network-id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/networks|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/networks/{port_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建网络
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除网络
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询网络详情。
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询提交请求的租户的所有网络
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
更新网络
