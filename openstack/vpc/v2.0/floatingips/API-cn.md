# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/floatingips"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
对浮动IP资源进行管理和操作，包括查询浮动IP、创建浮动IP、查询指定浮动IP、删除浮动IP以及更新浮动IP等接口。

示例代码, 创建一个浮动IP，并与一个内部端口关联

    
示例代码, 更新指定的浮动ip和相关联的端口

    
示例代码, 查询提交请求的浮动IP详情

    
示例代码, 查询提交请求的租户有权限操作的所有浮动IP地址。

    
示例代码, 删除指定的浮动IP

## 目录
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/floatingips|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/floatingips/{floatingip_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/floatingips/{floatingip_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/floatingips|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/floatingips/{floatingip_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建一个浮动IP，并与一个内部端口关联
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除指定的浮动IP
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询提交请求的浮动IP详情
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询提交请求的租户有权限操作的所有浮动IP地址。
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
更新指定的浮动ip和相关联的端口
