# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/securitygroups"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to query network resource quotas for the VPC service of a tenant. The network resources include VPCs, subnets, security groups, security group rules, elastic IP addresses, and VPNs.

Sample Code, This interface is used to create a security group.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := securitygroups.Create(client, tenantID, securitygroups.CreateOpts{
        SecurityGroup: securitygroups.CreateSecurityGroup{
            Name:        "EricSG",
            Description: "Test SecurityGroup",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query details about a security group.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := securitygroups.Get(client, tenantID, "f7616338-fa30-42b8-bf6b-754c0701aab8").Extract()
    
    if err != nil {
      panic(err)
    }
    
Sample Code, This interface is used to query security groups using search criteria and to display the security groups in a list.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := securitygroups.List(client, tenantID, securitygroups.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a security group.

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := securitygroups.Delete(client, tenantID, "2465d913-1084-4a6a-91e7-2fd6f490ecb3")
## Index
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/security-groups|
|vpc|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/security-groups/{security_group_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/security-groups/{security_group_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/security-groups|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
This interface is used to create a security group.
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
This interface is used to delete a security group.
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
This interface is used to query details about a security group.
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
This interface is used to query security groups using search criteria and to display the security groups in a list.
