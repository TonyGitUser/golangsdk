# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/ports"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 创建端口。

    
    result, err := ports.Create(client, ports.CreateOpts{
        Port: ports.CreatePort{
            Name:         "EricTestPort",
            NetworkId:    "5ae24488-454f-499c-86c4-c0355704005d",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 更新端口

    
    result, err := ports.Update(client,"5e56a480-f337-4985-8ca4-98546cb4fdae", ports.UpdateOpts{
      Port: ports.UpdatePort{
          Name: "ModifiedPort",
      },
    }).Extract()
    
    if err != nil {
      panic(err)
    }
    
示例代码, 查询单个端口

    
    result, err := ports.Get(client, "5e56a480-f337-4985-8ca4-98546cb4fdae").Extract()
    
    if err != nil {
      panic(err)
    }
    
示例代码, 查询端口列表

    
    result, err := ports.List(client, ports.ListOpts{
        Limit: 3,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除端口

    result := ports.Delete(client, "5e56a480-f337-4985-8ca4-98546cb4fdae")
    
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
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v1/ports|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v1/ports/{port_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v1/ports/{port_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v1/ports|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/ports/{port_id}|
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
