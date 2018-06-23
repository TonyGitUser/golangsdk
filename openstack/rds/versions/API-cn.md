# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//versions"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 查询当前支持的API版本列表。

    
    result, err := versions.List(client).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询指定API版本信息。

    
    result, err := versions.Get(client, "v1").Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|rds|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /rds/{versionId}|
|rds|func List(*golangsdk.ServiceClient) (ListResult)|GET /rds/|
## 开始
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询指定API版本信息。
## func List
    func List(*golangsdk.ServiceClient) (ListResult)  
查询当前支持的API版本列表。
