# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/floatingips"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
Manage and perform operations on floating IP addresses, including querying floating IP addresses, creating floating IP addresses, querying a specified floating IP address, deleting a floating IP address, and updating a floating IP address.

Sample Code, Creates a floating IP address and associates it with an internal port.

    
Sample Code, Updates a specific floating IP address and the port associated with the address.

    
Sample Code, Queries details about a specific floating IP address accessible to the tenant submitting the request.

    
Sample Code, Queries all floating IP addresses accessible to the tenant submitting the request.

    
Sample Code, Deletes a specific floating IP address.

## Index
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/floatingips|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/floatingips/{floatingip_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/floatingips/{floatingip_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/floatingips|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/floatingips/{floatingip_id}|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
Creates a floating IP address and associates it with an internal port.
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
Deletes a specific floating IP address.
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
Queries details about a specific floating IP address accessible to the tenant submitting the request.
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
Queries all floating IP addresses accessible to the tenant submitting the request.
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
Updates a specific floating IP address and the port associated with the address.
