# Package antiddos
    import "github.com/huaweicloudsdk/golangsdk/openstack/vpc/v2.0/subnets"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
子网是用来管理弹性云服务器网络平面的一个网络，可以提供IP地址管理、DNS服务，子网内的弹性云服务器IP地址都属于该子网。默认情况下，同一个VPC的所有子网内的弹性云服务器均可以进行通信，不同VPC的弹性云服务器不能进行通信。

示例代码, 创建子网。

    
    result, err := subnets.Create(client, subnets.CreateOpts{
    Subnet: subnets.CreateSubnet{
        Name:      "subnetV2",
        IpVersion: 4,
        NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
        Cidr:      "192.168.20.0/24",
        GatewayIp: "192.168.20.1",
        AllocationPools: []subnets.AllocationPool{
            {
                Start: "192.168.20.2",
                End:   "192.168.20.254",
            },
        },
        EnableDhcp:     true,
        DnsNameservers: [] string{"114.114.114.114", "114.114.115.115"},
        TenantId:       "57e98940a77f4bb988a21a7d0603a626",
    },
    }).Extract()
    
    if err != nil {
    panic(err)
    }
    
示例代码, 更新子网。

    
    result, err := subnets.Update(client, "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28", subnets.UpdateOpts{
        Subnet: subnets.UpdateSubnet{
            Name:           "subnetV2",
            GatewayIp:      "192.168.20.1",
            EnableDhcp:     true,
            DnsNameservers: [] string{"8.8.8.8", "4.4.4.4"},
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询子网。

    result, err := subnets.Get(client, "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28").Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询子网列表。

    allPages, err := subnets.List(client, subnets.ListOpts {
        Limit: 2,
    }).AllPages()
    
    result, err := subnets.ExtractList(allPages.(subnets.ListPage))
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除子网。

    
    result := subnets.Delete(client, "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28")
    
    if err != nil {
      panic(err)
    }
    
## 目录
**[func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)|POST /v2.0/subnets|
|vpc|func Delete(*golangsdk.ServiceClient, string) (DeleteResult)|DELETE /v2.0/subnets/{subnet_id}|
|vpc|func Get(*golangsdk.ServiceClient, string) (GetResult)|GET /v2.0/subnets/{subnet_id}|
|vpc|func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)|GET /v2.0/subnets|
|vpc|func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)|PUT /v2.0/subnets/{subnet_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, CreateOptsBuilder) (CreateResult)  
创建子网。
## func Delete
    func Delete(*golangsdk.ServiceClient, string) (DeleteResult)  
删除子网。
## func Get
    func Get(*golangsdk.ServiceClient, string) (GetResult)  
查询子网。
## func List
    func List(*golangsdk.ServiceClient, ListOptsBuilder) (ListResult)  
查询子网列表。
## func Update
    func Update(*golangsdk.ServiceClient, string, UpdateOptsBuilder) (UpdateResult)  
更新子网。
