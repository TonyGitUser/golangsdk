# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//logs"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 查询数据库错误日志信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := logs.ListErrors(client, versionId, projectId, instanceId, logs.ListErrorsOpts{
        StartDate: "2018-01-29",
        EndDate: "2018-08-29",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询数据库慢日志信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := logs.ListInfos(client, versionId, projectId, instanceId, logs.ListInfosOpts{
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func ListErrors(*golangsdk.ServiceClient, string, string, string, ListErrorsOptsBuilder) (ListErrorsResult)](#func-listerrors)**  
**[func ListInfos(*golangsdk.ServiceClient, string, string, string, ListInfosOptsBuilder) (ListInfosResult)](#func-listinfos)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|rds|func ListErrors(*golangsdk.ServiceClient, string, string, string, ListErrorsOptsBuilder) (ListErrorsResult)|GET /rds/{versionId}/{project_id}/instances/{instanceId}/errorlog|
|rds|func ListInfos(*golangsdk.ServiceClient, string, string, string, ListInfosOptsBuilder) (ListInfosResult)|GET /rds/{versionId}/{project_id}/instances/{instanceId}/slowlog|
## 开始
## func ListErrors
    func ListErrors(*golangsdk.ServiceClient, string, string, string, ListErrorsOptsBuilder) (ListErrorsResult)  
查询数据库错误日志信息。
## func ListInfos
    func ListInfos(*golangsdk.ServiceClient, string, string, string, ListInfosOptsBuilder) (ListInfosResult)  
查询数据库慢日志信息。
