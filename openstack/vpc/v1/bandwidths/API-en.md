# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/bandwidths"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
You can allocate bandwidth when assigning an EIP so that the ECS bound with the EIP can use the bandwidth to access the Internet. he bandwidth displays network resource usage and can be used for service metering.

Sample Code, This interface is used to update information about a bandwidth.

    
Sample Code, This interface is used to query details about a bandwidth.

    
Sample Code, This interface is used to query bandwidths using search criteria and to display the bandwidths in a list.

    
## Index
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/bandwidths/{bandwidth_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/bandwidths|
|vpc|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/{tenant_id}/bandwidths/{bandwidth_id}|
## Content
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
This interface is used to query details about a bandwidth.
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
This interface is used to query bandwidths using search criteria and to display the bandwidths in a list.
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to update information about a bandwidth.
