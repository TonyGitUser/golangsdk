# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//dbversions"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to obtain version information about a specified type of a database.

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    datastoreName := "MySQL"
    result, err := dbversions.List(client, versionId, projectId, datastoreName).Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func List(*golangsdk.ServiceClient, string, string, string) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|rds|func List(*golangsdk.ServiceClient, string, string, string) (ListResult)|GET /rds/{versionId}/{project_id}/datastores/{datastore_name}/versions|
## Content
## func List
    func List(*golangsdk.ServiceClient, string, string, string) (ListResult)  
This interface is used to obtain version information about a specified type of a database.
