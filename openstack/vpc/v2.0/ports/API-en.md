# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/ports"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to manage and perform operations on ports, including querying ports, creating a port, querying a specified port, deleting a port, and updating a port.

Sample Code, This interface is used to create a port.

    
    result, err := ports.Create(client, ports.CreateOpts{
        Port: ports.CreatePort{
            Name:      "EricTestPort",
            NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to update a port.

    
    result, err := ports.Update(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986", ports.UpdateOpts{
        Port: ports.UpdatePort{
            Name: "ModifiedPort",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query a single port.

    
    result, err := ports.Get(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986").Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query ports and to display the ports in a list.

    
    allPages, err := ports.List(client, ports.ListOpts {
        Limit: 2,
        Name: "EricTestPort",
    }).AllPages()
    
    result, err := ports.ExtractList(allPages.(ports.ListPage))
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a port.

    result := ports.Delete(client, "d6d0258f-7bf2-4504-845b-ab6b8513a986")
    
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
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/ports|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/ports/{port_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/ports/{port_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/ports|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/ports/{port_id}|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
This interface is used to create a port.
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
This interface is used to delete a port.
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
This interface is used to query a single port.
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
This interface is used to query ports and to display the ports in a list.
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to update a port.
