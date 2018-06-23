# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//parameters"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to set DB instance parameters.

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := parameters.Create(client, versionId, projectId, instanceId, parameters.CreateOpts{
        Values: map[string]interface{}{
            "connect_timeout": 17,
            "sync_binlog": 1,
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to obtain all the parameters that can be modified of the current database version.

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    datastoreId := "9dae9226-5f8e-4b10-bd41-adae2e693e89"
    result, err := parameters.List(client, versionId, projectId, datastoreId).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to obtain the parameter information that can be modified of a specified database version.

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    datastoreId := "9dae9226-5f8e-4b10-bd41-adae2e693e89"
    parameterName := "auto_increment_offset"
    result, err := parameters.Get(client, versionId, projectId, datastoreId, parameterName).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to restore a DB instance's parameter group to the default parameter group.

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := parameters.Restore(client, versionId, projectId, instanceId).Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func Create(*golangsdk.ServiceClient, string, string, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Get(*golangsdk.ServiceClient, string, string, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, string, string) (ListResult)](#func-list)**  
**[func Restore(*golangsdk.ServiceClient, string, string, string) (RestoreResult)](#func-restore)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|rds|func Create(*golangsdk.ServiceClient, string, string, string, CreateOptsBuilder) (CreateResult)|PUT /rds/{versionId}/{project_id}/instances/{instanceId}/parameters|
|rds|func Get(*golangsdk.ServiceClient, string, string, string, string) (GetResult)|GET /rds/{versionId}/{project_id}/datastores/versions/{datastore_version_id}/parameters/{parameter_name}|
|rds|func List(*golangsdk.ServiceClient, string, string, string) (ListResult)|GET /rds/{versionId}/{project_id}/datastores/versions/{datastore_version_id}/parameters|
|rds|func Restore(*golangsdk.ServiceClient, string, string, string) (RestoreResult)|PUT /rds/{versionId}/{project_id}/instances/{instanceId}/parameters/default|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, string, string, string, CreateOptsBuilder) (CreateResult)  
This interface is used to set DB instance parameters.
## func Get
    func Get(*golangsdk.ServiceClient, string, string, string, string) (GetResult)  
This interface is used to obtain the parameter information that can be modified of a specified database version.
## func List
    func List(*golangsdk.ServiceClient, string, string, string) (ListResult)  
This interface is used to obtain all the parameters that can be modified of the current database version.
## func Restore
    func Restore(*golangsdk.ServiceClient, string, string, string) (RestoreResult)  
This interface is used to restore a DB instance's parameter group to the default parameter group.
