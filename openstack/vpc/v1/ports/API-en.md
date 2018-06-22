# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/ports"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to create a port.

    
    result, err := ports.Create(client, ports.CreateOpts{
        Port: ports.CreatePort{
            Name:         "EricTestPort",
            NetworkId:    "5ae24488-454f-499c-86c4-c0355704005d",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to update a port.

    
    result, err := ports.Update(client,"5e56a480-f337-4985-8ca4-98546cb4fdae", ports.UpdateOpts{
      Port: ports.UpdatePort{
          Name: "ModifiedPort",
      },
    }).Extract()
    
    if err != nil {
      panic(err)
    }
    
Sample Code, This interface is used to query a single port.

    
    result, err := ports.Get(client, "5e56a480-f337-4985-8ca4-98546cb4fdae").Extract()
    
    if err != nil {
      panic(err)
    }
    
Sample Code, This interface is used to query ports and to display the ports in a list.

    
    result, err := ports.List(client, ports.ListOpts{
        Limit: 3,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a port.

    result := ports.Delete(client, "5e56a480-f337-4985-8ca4-98546cb4fdae")
    
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
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v1/ports|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v1/ports/{port_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v1/ports/{port_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v1/ports|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/ports/{port_id}|
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
