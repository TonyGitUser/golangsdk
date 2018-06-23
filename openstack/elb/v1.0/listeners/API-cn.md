# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/listeners"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
承担ELB具体的协议和端口配置；云服务器协议和端口配置；监听策略配置。

示例代码, 给负载均衡器下创建监听器，用于监听后端云服务器的状态。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := listeners.Create(client, tenantID, listeners.CreateOpts{
        Name:              "TestLS",
        Description:       "LS TESTING",
        LoadbalancerId:    "ebe982b8afe04851bbc26ad4609003bf",
        Protocol:          "HTTP",
        Port:              80,
        BackendProtocol:   "HTTP",
        BackendPort:       81,
        LbAlgorithm:       "roundrobin",
        SessionSticky:     true,
        StickySessionType: "insert",
        CookieTimeout:     1,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除监听器。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
    listeners.Delete(client, tenantID, listenerId)
    
示例代码, 查询监听器的详细信息。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
    result, err := listeners.Get(client, tenantID, listenerId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询监听器列表信息。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := listeners.List(client, tenantID, listeners.ListOpts{
       LoadbalancerId: "ebe982b8afe04851bbc26ad4609003bf",
    }).Extract()
    
示例代码, 修改监听器的信息，目前支持修改名称、描述、管理状态。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
    result, err := listeners.Update(client, tenantID, listenerId, listeners.UpdateOpts{
        Name:        "ModifiedTestLS",
        Description: "LS TESTING",
        Port:        80,
        BackendPort: 81,
        LbAlgorithm: "roundrobin",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|elb|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1.0/{tenant_id}/elbaas/listeners|
|elb|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1.0/{tenant_id}/elbaas/listeners/{listener_id}|
|elb|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1.0/{tenant_id}/elbaas/listeners/{listener_id}|
|elb|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1.0/{tenant_id}/elbaas/listeners|
|elb|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1.0/{tenant_id}/elbaas/listeners/{listener_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
给负载均衡器下创建监听器，用于监听后端云服务器的状态。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除监听器。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询监听器的详细信息。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
查询监听器列表信息。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
修改监听器的信息，目前支持修改名称、描述、管理状态。
