# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/healthcheck"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to configure a health check task for a backend ECS.

    
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
    
Sample Code, This interface is used to delete a health check task.

    
    healthcheckId := "c74c17f02c5142b3befc973056f220d3"
    healthcheck.Delete(client, healthcheckId)
    
Sample Code, This interface is used to query details about a health check task.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    healthcheckId := "c74c17f02c5142b3befc973056f220d3"
    result, err := healthcheck.Get(client, tenantId, healthcheckId).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to modify information about a health check task.

    
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
## Index
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|elb|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1.0/{tenant_id}/elbaas/healthcheck|
|elb|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v1.0/elbaas/healthcheck/{healthcheck_id}|
|elb|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1.0/{tenant_id}/elbaas/healthcheck/{healthcheck_id}|
|elb|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1.0/{tenant_id}/elbaas/healthcheck/{healthcheck_id}|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
This interface is used to configure a health check task for a backend ECS.
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
This interface is used to delete a health check task.
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
This interface is used to query details about a health check task.
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to modify information about a health check task.
