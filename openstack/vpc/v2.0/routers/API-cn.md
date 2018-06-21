# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/routers"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
对路由器资源进行管理和操作，包括查询、路由器列表、创建路由器、查询指定路由器、删除路由器以及更新路由器等接口。

示例代码, 创建路由器

    
示例代码, 更新路由器。

    
示例代码, 查询提交请求的租户指定的路由器详细信息

    
示例代码, 查询提交请求的租户有权限操作的所有路由器信息。

    
    
示例代码, 删除租户指定的路由器。

    
## 目录
**[func AddInterface(*golangsdk.ServiceClient, string, AddInterfaceOptsBuilder) (AddInterfaceResult)](#func-addinterface)**  
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func RemoveInterface(*golangsdk.ServiceClient, string, RemoveInterfaceOptsBuilder) (RemoveInterfaceResult)](#func-removeinterface)**  
**[func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func AddInterface(*golangsdk.ServiceClient, string, AddInterfaceOptsBuilder) (AddInterfaceResult)|PUT /v2.0/routers/{router_id}/add_router_interface|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/routers|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/routers/{router_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/routers/{router_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/routers|
|vpc|func RemoveInterface(*golangsdk.ServiceClient, string, RemoveInterfaceOptsBuilder) (RemoveInterfaceResult)|PUT /v2.0/routers/{router_id}/remove_router_interface|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/routers/{router_id}|
## 开始
## func AddInterface
    func AddInterface(*golangsdk.ServiceClient, string, AddInterfaceOptsBuilder) (AddInterfaceResult)  
添加路由器接口
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建路由器
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除租户指定的路由器。
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询提交请求的租户指定的路由器详细信息
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询提交请求的租户有权限操作的所有路由器信息。
## func RemoveInterface
    func RemoveInterface(*golangsdk.ServiceClient, string, RemoveInterfaceOptsBuilder) (RemoveInterfaceResult)  
删除路由器接口。
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
更新路由器。
