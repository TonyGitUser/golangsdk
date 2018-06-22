# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/vpcs"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
The Virtual Private Cloud (VPC) service enables you to provision logically isolated, configurable, and manageable virtual networks for Elastic Cloud Servers (ECSs), improving the security of resources in the cloud system and simplifying network deployment.

Sample Code, This interface is used to create a VPC.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.Create(client, tenantID, vpcs.CreateOpts{
        Vpc: vpcs.VPC{
            Name: "ABC",
            Cidr: "192.168.0.0/16",
        },
    }).Extract()
    
Sample Code, This interface is used to update information about a VPC.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.Update(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0", vpcs.UpdateOpts{
        Vpc: vpcs.VPC{
            Name: "ABC-back",
            Cidr: "192.168.0.0/24",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query details about a VPC.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
        result, err := vpcs.Get(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0").Extract()
    
Sample Code, This interface is used to query VPCs using search criteria and to display the VPCs in a list.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.List(client, tenantID, vpcs.ListOpts{
       Limit: 2,
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
Sample Code, This interface is used to delete a VPC.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := vpcs.Delete(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0")
## Index
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/vpcs|
|vpc|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/vpcs/{vpc_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/vpcs/{vpc_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/vpcs|
|vpc|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/{tenant_id}/vpcs/{vpc_id}|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
This interface is used to create a VPC.
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
This interface is used to delete a VPC.
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
This interface is used to query details about a VPC.
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
This interface is used to query VPCs using search criteria and to display the VPCs in a list.
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to update information about a VPC.
