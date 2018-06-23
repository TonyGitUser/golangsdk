# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//versions"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to query the currently supported API version list.

    
    result, err := versions.List(client).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query the specified API version.

    
    result, err := versions.Get(client, "v1").Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|rds|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /rds/{versionId}|
|rds|func List(*golangsdk.ServiceClient) (ListResult)|GET /rds/|
## Content
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
This interface is used to query the specified API version.
## func List
    func List(*golangsdk.ServiceClient) (ListResult)  
This interface is used to query the currently supported API version list.
