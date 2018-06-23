# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/healthcheck"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 为后端云服务器添加健康检查配置。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := healthcheck.Create(client, tenantId, healthcheck.CreateOpts{
        HealthcheckConnectPort: 80,
        HealthcheckInterval: 5,
        HealthcheckProtocol: "HTTP",
        HealthcheckTimeout: 10,
        HealthyThreshold: 3,
        ListenerId: "d5f3197c6bd8491ca1dfc905350d85ea",
        UnhealthyThreshold: 3,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除健康检查。

    
    healthcheckId := "c74c17f02c5142b3befc973056f220d3"
    healthcheck.Delete(client, healthcheckId)
    
示例代码, 查询健康检查详细信息。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    healthcheckId := "c74c17f02c5142b3befc973056f220d3"
    result, err := healthcheck.Get(client, tenantId, healthcheckId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 修改健康检查信息。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    healthcheckId := "c74c17f02c5142b3befc973056f220d3"
    result, err := healthcheck.Update(client, tenantId, healthcheckId, healthcheck.UpdateOpts{
        HealthcheckConnectPort: 81,
        HealthcheckInterval: 6,
        HealthcheckProtocol: "HTTP",
        HealthcheckTimeout: 20,
        HealthyThreshold: 4,
        UnhealthyThreshold: 4,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|elb|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1.0/{tenant_id}/elbaas/healthcheck|
|elb|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v1.0/elbaas/healthcheck/{healthcheck_id}|
|elb|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1.0/{tenant_id}/elbaas/healthcheck/{healthcheck_id}|
|elb|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1.0/{tenant_id}/elbaas/healthcheck/{healthcheck_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
为后端云服务器添加健康检查配置。
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除健康检查。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询健康检查详细信息。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
修改健康检查信息。
