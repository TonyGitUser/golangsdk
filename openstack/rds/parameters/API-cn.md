# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//parameters"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 设置实例参数。

    
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
    
示例代码, 获取数据库版本的所有可修改参数信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    datastoreId := "9dae9226-5f8e-4b10-bd41-adae2e693e89"
    result, err := parameters.List(client, versionId, projectId, datastoreId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 获取指定数据库版本下可修改的参数信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    datastoreId := "9dae9226-5f8e-4b10-bd41-adae2e693e89"
    parameterName := "auto_increment_offset"
    result, err := parameters.Get(client, versionId, projectId, datastoreId, parameterName).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 该接口用于将实例参数组恢复为默认参数组。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := parameters.Restore(client, versionId, projectId, instanceId).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, string, string, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Get(*golangsdk.ServiceClient, string, string, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, string, string) (ListResult)](#func-list)**  
**[func Restore(*golangsdk.ServiceClient, string, string, string) (RestoreResult)](#func-restore)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|rds|func Create(*golangsdk.ServiceClient, string, string, string, CreateOptsBuilder) (CreateResult)|PUT /rds/{versionId}/{project_id}/instances/{instanceId}/parameters|
|rds|func Get(*golangsdk.ServiceClient, string, string, string, string) (GetResult)|GET /rds/{versionId}/{project_id}/datastores/versions/{datastore_version_id}/parameters/{parameter_name}|
|rds|func List(*golangsdk.ServiceClient, string, string, string) (ListResult)|GET /rds/{versionId}/{project_id}/datastores/versions/{datastore_version_id}/parameters|
|rds|func Restore(*golangsdk.ServiceClient, string, string, string) (RestoreResult)|PUT /rds/{versionId}/{project_id}/instances/{instanceId}/parameters/default|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, string, string, CreateOptsBuilder) (CreateResult)  
设置实例参数。
## func Get
    func Get(*golangsdk.ServiceClient, string, string, string, string) (GetResult)  
获取指定数据库版本下可修改的参数信息。
## func List
    func List(*golangsdk.ServiceClient, string, string, string) (ListResult)  
获取数据库版本的所有可修改参数信息。
## func Restore
    func Restore(*golangsdk.ServiceClient, string, string, string) (RestoreResult)  
该接口用于将实例参数组恢复为默认参数组。
