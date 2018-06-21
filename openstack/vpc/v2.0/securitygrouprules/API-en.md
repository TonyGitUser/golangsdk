# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/securitygrouprules"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to query network resource quotas for the VPC service of a tenant. The network resources include VPCs, subnets, security groups, security group rules, elastic IP addresses, and VPNs.

Sample Code, This interface is used to create a security group rule.

    
    result, err := securitygrouprules.Create(client, securitygrouprules.CreateOpts{
       SecurityGroupRule: securitygrouprules.CreateSecurityGroupRule {
           Description: "Test SecurityGroup",
           TenantId: tenantID,
           SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
           Direction: "egress",
           Protocol: "tcp",
           RemoteIpPrefix: "10.10.0.0/24",
       },
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
Sample Code, This interface is used to query details about a security group rule.

    
    result, err := securitygrouprules.Get(client, "26243298-ae79-46a3-bad9-34395762e033").Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query security group rules using search criteria and to display the security group rules in a list.

    
    allPages, err := securitygrouprules.List(client, securitygrouprules.ListOpts{
        Limit: 2,
        Protocol: "tcp",
    }).AllPages()
    
    result, err := securitygrouprules.ExtractList(allPages.(securitygrouprules.ListPage))
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a security group rule.

    
    result = securitygrouprules.Delete(client, "26243298-ae79-46a3-bad9-34395762e033")
## Index
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/security-group-rules|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/security-group-rules/{security-groups-rules-id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/security-group-rules/{security-groups-rules-id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/security-group-rules|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
This interface is used to create a security group rule.
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
This interface is used to delete a security group rule.
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
This interface is used to query details about a security group rule.
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
This interface is used to query security group rules using search criteria and to display the security group rules in a list.
