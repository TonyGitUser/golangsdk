# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/ports"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
对端口资源进行管理和操作，包括查询端口列表、创建端口、查询指定端口、删除端口以及更新端口等接口。

示例代码, 创建端口。

    
    result, err := ports.Create(client, ports.CreateOpts{
        Port: ports.CreatePort{
            Name:      "EricTestPort",
            NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 更新端口

    
    result, err := ports.Update(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986", ports.UpdateOpts{
        Port: ports.UpdatePort{
            Name: "ModifiedPort",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询单个端口

    
    result, err := ports.Get(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986").Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询端口列表

    
    allPages, err := ports.List(client, ports.ListOpts {
        Limit: 2,
        Name: "EricTestPort",
    }).AllPages()
    
    result, err := ports.ExtractList(allPages.(ports.ListPage))
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除端口

    result := ports.Delete(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986")
    
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
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/ports|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/ports/{port_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/ports/{port_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/ports|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/ports/{port_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建端口。
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除端口
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询单个端口
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询端口列表
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
更新端口
