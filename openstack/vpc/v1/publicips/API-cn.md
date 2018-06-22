# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/publicips"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
公网IP地址为可以直接访问因特网的IP地址。私有IP地址为公有云内局域网络所有的IP地址，私有IP地址禁止出现在Internet中。 弹性IP是基于互联网上的静态IP地址，将弹性IP地址和子网中关联的弹性云服务器绑定和解绑，可以实现VPC中的弹性云服务器通过固定的公网IP地址与互联网互通。

示例代码, 申请弹性IP。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := publicips.Create(client, tenantID, publicips.CreateOpts{
        Publicip: struct {
            Type      string `json:"type,"`
            IpAddress string `json:"ip_address,omitempty"`
        }{
            Type: "5_bgp",
        },
        Bandwidth: struct {
                Name string `json:"name,omitempty"`
                Size int `json:"size,omitempty"`
                Id string `json:"id,omitempty"`
                ShareType string `json:"share_type,"`
                ChargeMode string `json:"charge_mode,omitempty"`
        }{
            Name: "bandwidth-d62f",
            Size:1,
            ShareType: "WHOLE",
            ChargeMode: "traffic"},
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 更新弹性IP，将弹性IP跟一个网卡绑定或者解绑定。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := publicips.Update(client, tenantID, "84a71976-a8c2-42e0-8826-7fc27b876e42", publicips.UpdateOpts{
       Publicip: struct {
           PortId string `json:"port_id,"`
       } {
           PortId: "be44f0d9-f8c6-485a-971e-83042dd20d6c",
       },
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
示例代码, 查询弹性IP。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := publicips.Get(client, tenantID, "84a71976-a8c2-42e0-8826-7fc27b876e42").Extract()
    
示例代码, 查询弹性IP列表。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := publicips.List(client, tenantID, publicips.ListOpts{
       Limit: 2,
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
示例代码, 删除弹性IP。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := publicips.Delete(client, tenantID, "2412e868-f93a-4dfc-b171-5096baa27403")
    
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/publicips|
|vpc|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/publicips/{publicip_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/publicips/{publicip_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/publicips|
|vpc|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/{tenant_id}/publicips/{publicip_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
申请弹性IP。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除弹性IP。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询弹性IP。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
查询弹性IP列表。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
更新弹性IP，将弹性IP跟一个网卡绑定或者解绑定。
