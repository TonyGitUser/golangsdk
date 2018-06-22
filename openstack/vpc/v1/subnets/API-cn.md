# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/subnets"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
子网是用来管理弹性云服务器网络平面的一个网络，可以提供IP地址管理、DNS服务，子网内的弹性云服务器IP地址都属于该子网。默认情况下，同一个VPC的所有子网内的弹性云服务器均可以进行通信，不同VPC的弹性云服务器不能进行通信。

示例代码, 创建子网。

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.Create(client, tenantID, subnets.CreateOpts{
       Subnet: subnets.Subnet {
           Name: "subnet",
           Cidr: "192.168.20.0/24",
           GatewayIp: "192.168.20.1",
           DhcpEnable: true,
           PrimaryDns: "114.114.114.114",
           SecondaryDns: "114.114.115.115",
           DnsList: [] string{
               "114.114.114.114",
               "114.114.115.115",
           },
           AvailabilityZone:"cn-north-1a",
           VpcId:"ea3b0efe-0d6a-4288-8b16-753504a994b9",
       },
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
示例代码, 更新子网。

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.Update(client, tenantID, "ea3b0efe-0d6a-4288-8b16-753504a994b9","c9aba52d-ec14-40cb-930f-c8153e93c2db", subnets.UpdateOpts{
       Subnet: subnets.Subnet{
           Name: "ABC-back",
           Cidr: "192.168.0.0/24",
       },
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
示例代码, 查询子网。

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.Get(client, tenantID, "c9aba52d-ec14-40cb-930f-c8153e93c2db").Extract()
    
示例代码, 查询子网列表。

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.List(client, tenantID, subnets.ListOpts{
      Limit: 2,
    }).Extract()
    
    if err != nil {
      panic(err)
    }
    
示例代码, 删除子网。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := subnets.Delete(client, tenantID, "ea3b0efe-0d6a-4288-8b16-753504a994b9","c9aba52d-ec14-40cb-930f-c8153e93c2db")
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/subnets|
|vpc|func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/vpcs/{vpc_id}/subnets/{subnet_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/subnets/{subnet_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/subnets|
|vpc|func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/{tenant_id}/vpcs/{vpc_id}/subnets/{subnet_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
创建子网。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)  
删除子网。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询子网。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
查询子网列表。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)  
更新子网。
