# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//dbversions"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 获取指定类型的数据库版本相关信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    datastoreName := "MySQL"
    result, err := dbversions.List(client, versionId, projectId, datastoreName).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func List(*golangsdk.ServiceClient, string, string, string) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|rds|func List(*golangsdk.ServiceClient, string, string, string) (ListResult)|GET /rds/{versionId}/{project_id}/datastores/{datastore_name}/versions|
## 开始
## func List
    func List(*golangsdk.ServiceClient, string, string, string) (ListResult)  
获取指定类型的数据库版本相关信息。
