# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/routers"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to manage and perform operations on router resources, including querying routers, creating a router, querying a specified router, deleting a router, and updating a router.

Sample Code, Creates a router.

    
Sample Code, Updates router information.

    
Sample Code, Queries details about a specific router accessible to the tenant submitting the request.

    
Sample Code, Queries all routers accessible to the tenant submitting the request. 

    
    
Sample Code, Deletes a specified router.

    
## Index
**[func AddInterface(*golangsdk.ServiceClient, string, AddInterfaceOptsBuilder) (AddInterfaceResult)](#func-addinterface)**  
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func RemoveInterface(*golangsdk.ServiceClient, string, RemoveInterfaceOptsBuilder) (RemoveInterfaceResult)](#func-removeinterface)**  
**[func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func AddInterface(*golangsdk.ServiceClient, string, AddInterfaceOptsBuilder) (AddInterfaceResult)|PUT /v2.0/routers/{router_id}/add_router_interface|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/routers|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/routers/{router_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/routers/{router_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/routers|
|vpc|func RemoveInterface(*golangsdk.ServiceClient, string, RemoveInterfaceOptsBuilder) (RemoveInterfaceResult)|PUT /v2.0/routers/{router_id}/remove_router_interface|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/routers/{router_id}|
## Content
## func AddInterface
    func AddInterface(*golangsdk.ServiceClient, string, AddInterfaceOptsBuilder) (AddInterfaceResult)  
Adds an interface to a router.
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
Creates a router.
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
Deletes a specified router.
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
Queries details about a specific router accessible to the tenant submitting the request.
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
Queries all routers accessible to the tenant submitting the request. 
## func RemoveInterface
    func RemoveInterface(*golangsdk.ServiceClient, string, RemoveInterfaceOptsBuilder) (RemoveInterfaceResult)  
Removes an interface from a router.
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
Updates router information.
