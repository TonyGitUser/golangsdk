# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/privateips"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
公网IP地址为可以直接访问因特网的IP地址。私有IP地址为公有云内局域网络所有的IP地址，私有IP地址禁止出现在Internet中。 弹性IP是基于互联网上的静态IP地址，将弹性IP地址和子网中关联的弹性云服务器绑定和解绑，可以实现VPC中的弹性云服务器通过固定的公网IP地址与互联网互通。

示例代码, 申请私有IP。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := privateips.Create(client, tenantID, privateips.CreateOpts{
       Privateips: [] privateips.CreatePrivateIp{
           {
               SubnetId:"5ae24488-454f-499c-86c4-c0355704005d",
               IpAddress: "192.168.0.12",
           },
       },
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
示例代码, 指定IP的ID查询私有IP。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := privateips.Get(client, tenantID, "ea274524-f1cc-4078-8e67-c002be25c413").Extract()
    
    if err != nil {
      panic(err)
    }
    
示例代码, 查询指定子网下的私有IP列表。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    subnetID := "5ae24488-454f-499c-86c4-c0355704005d"
    result, err := privateips.List(client, tenantID, subnetID, privateips.ListOpts{
      Limit: 2,
    }).Extract()
    
    if err != nil {
      panic(err)
    }
    
示例代码, 删除私有IP。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := privateips.Delete(client, tenantID, "ea274524-f1cc-4078-8e67-c002be25c413")
    
    if err != nil {
      panic(err)
    }
    
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/privateips|
|vpc|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/privateips/{privateip_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/privateips/{privateip_id}|
|vpc|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/subnets/{subnet_id}/privateips|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
申请私有IP。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除私有IP。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
指定IP的ID查询私有IP。
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
查询指定子网下的私有IP列表。
