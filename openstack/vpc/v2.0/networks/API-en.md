# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/networks"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to manage and perform operations on network resources, including querying networks, creating a network, querying a specified network, deleting a network, and updating a network.

Sample Code, Creates a network.

    
    result, err := networks.Create(client, networks.CreateOpts{
        Network: networks.CreateNetwork{
            Name:         "NetWork Test",
            Shared:       false,
            AdminStateUp: true,
            TenantId:     "57e98940a77f4bb988a21a7d0603a626",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, Updates a network.

    
    result, err := networks.Update(client, "1c6af92e-bd06-4350-8f51-5ec32167814f", networks.UpdateOpts{
        Network: networks.UpdateNetwork{
            Name:   "Test Net Works Updated",
            Shared: true,
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, Queries details about the specified network.

    
    result, err := networks.Get(client, "1c6af92e-bd06-4350-8f51-5ec32167814f").Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, Deletes all networks to which the specified tenant has access.

    
    result := networks.Delete(client, "1c6af92e-bd06-4350-8f51-5ec32167814f")
    
Sample Code, Queries all networks accessible to the tenant submitting the request. A maximum of 2000 records can be returned for each query operation. 

    
    allPages, err := networks.List(client, networks.ListOpts{
        Limit: 2,
    }).AllPages()
    
    result, err := networks.ExtractList(allPages.(networks.ListPage))
    
    if err != nil {
        panic(err)
    }
## Index
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/networks|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/networks/{network-id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/networks/{network-id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/networks|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/networks/{port_id}|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
Creates a network.
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
Deletes all networks to which the specified tenant has access.
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
Queries details about the specified network.
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
Queries all networks accessible to the tenant submitting the request. A maximum of 2000 records can be returned for each query operation. 
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
Updates a network.
