# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/bandwidths"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
带宽是指弹性云服务器通过弹性IP访问公网时使用的带宽。 可以通过带宽展示网络的使用情况，作为服务计费的依据。

示例代码, 更新带宽。

    
示例代码, 查询带宽。

    
示例代码, 查询带宽列表。

    
## 目录
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/bandwidths/{bandwidth_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/bandwidths|
|vpc|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/{tenant_id}/bandwidths/{bandwidth_id}|
## 开始
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询带宽。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
查询带宽列表。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
更新带宽。
