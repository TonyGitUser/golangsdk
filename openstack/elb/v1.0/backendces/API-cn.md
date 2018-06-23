# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/backendces"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
添加后端云服务器之前首先要检查后端云服务器所在安全组规则是否配置放行100.125.0.0/16，并配置ELB用于健康检查的协议和端口，否则无法对已添加的后端云服务器执行健康检查。

示例代码, 给指定负载均衡器下的监听器里添加被监听的后端云服务器。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "d5f3197c6bd8491ca1dfc905350d85ea"
    result, err := backendces.Create(client, tenantId, listenerId, backendces.CreateOpts{
        Items: []struct {
            ServerId string `json:"server_id,omitempty"`
            Address  string `json:"address,omitempty"`
        }{
            {
                ServerId: "e1040c67-b130-41ee-9405-07c6c6c20208",
                Address:  "192.168.1.215",
            },
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 将后端云服务器从监听器中移除，停止对该云服务器的监听。支持同时移除多个后端云服务器。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "ca082827b61d4902bfaf32e8174e244a"
    backendces.Delete(client, tenantId, listenerId, backendces.DeleteOpts{
        RemoveMember: []struct {
            Id string `json:"id,omitempty"`
        }{
            {
                Id: "e1040c67-b130-41ee-9405-07c6c6c20208",
            },
        },
    })
    
    if err != nil {
        panic(err)
    }
    
示例代码, 普通租户可查询指定监听器下后端服务器列表。管理员租户查询后端服务器列表为空列表。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    listenerId := "d5f3197c6bd8491ca1dfc905350d85ea"
    result, err := backendces.List(client, tenantId, listenerId, backendces.ListOpts{
        Limit: "1",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)](#func-delete)**  
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|elb|func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)|POST /v1.0/{tenant_id}/elbaas/listeners/{listener_id}/members|
|elb|func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)|POST /v1.0/{tenant_id}/elbaas/listeners/{listener_id}/members/action|
|elb|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /v1.0/{tenant_id}/elbaas/listeners/{listener_id}/members|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)  
给指定负载均衡器下的监听器里添加被监听的后端云服务器。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)  
将后端云服务器从监听器中移除，停止对该云服务器的监听。支持同时移除多个后端云服务器。
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
普通租户可查询指定监听器下后端服务器列表。管理员租户查询后端服务器列表为空列表。
