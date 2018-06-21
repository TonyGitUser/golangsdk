# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/securitygroups"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to query network resource quotas for the VPC service of a tenant. The network resources include VPCs, subnets, security groups, security group rules, elastic IP addresses, and VPNs.

Sample Code, This interface is used to create a security group.

    
    result, err := securitygroups.Create(client, securitygroups.CreateOpts{
        SecurityGroup: securitygroups.CreateSecurityGroup{
            Name:        "EricSG",
            Description: "Test SecurityGroup",
            TenantId: tenantID,
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query details about a security group.

    
    result, err := securitygroups.Get(client, "a988649d-cfbf-4c2a-bd91-3b84d2403747").Extract()
    
    if err != nil {
      panic(err)
    }
    
Sample Code, This interface is used to query security groups using search criteria and to display the security groups in a list.

    
    allPages, err := securitygroups.List(client, securitygroups.ListOpts{
        Limit: 2,
    }).AllPages()
    
    result, err := securitygroups.ExtractList(allPages.(securitygroups.ListPage))
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a security group.

    
    result := securitygroups.Delete(client, tenantID, "2465d913-1084-4a6a-91e7-2fd6f490ecb3")
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to update a security group.

    
    result, err := securitygroups.Update(client, "7af80d49-0a43-462d-aed8-a1e12ac91af6", securitygroups.UpdateOpts{
        SecurityGroup: securitygroups.UpdateSecurityGroup{
            Name:        "EricSG",
            Description: "Modified SecurityGroup",
        },
    }).Extract()
    
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
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/security-groups|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/security-groups/{security_group_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/security-groups/{security_group_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/security-groups|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/security-groups/{security_group_id}|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
This interface is used to create a security group.
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
This interface is used to delete a security group.
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
This interface is used to query details about a security group.
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
This interface is used to query security groups using search criteria and to display the security groups in a list.
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to update a security group.
