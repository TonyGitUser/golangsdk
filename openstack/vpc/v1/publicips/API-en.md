# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/vpc/v1/publicips"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
A public IP address is an IP address that can be directly accessed over the Internet. Private IP addresses are all IP addresses on the local area network (LAN) of the public cloud and cannot exist on the Internet. An EIP is a static, public IP address. You can bind an EIP to and unbind an EIP from an ECS in your subnet. An EIP enables an ECS in your VPC to communicate with the Internet through a fixed public IP address.

Sample Code, This interface is used to apply for an elastic IP address.

    
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
    
Sample Code, This interface is used to bind an elastic IP address to a NIC or unbind an elastic IP address from a NIC.

    
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
    
Sample Code, This interface is used to query details about an elastic IP address.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := publicips.Get(client, tenantID, "84a71976-a8c2-42e0-8826-7fc27b876e42").Extract()
    
Sample Code, This interface is used to query elastic IP addresses using search criteria and to display the elastic IP addresses in a list.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := publicips.List(client, tenantID, publicips.ListOpts{
       Limit: 2,
    }).Extract()
    
    if err != nil {
       panic(err)
    }
    
Sample Code, This interface is used to delete an elastic IP address.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := publicips.Delete(client, tenantID, "2412e868-f93a-4dfc-b171-5096baa27403")
    
## Index
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|vpc|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1/{tenant_id}/publicips|
|vpc|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1/{tenant_id}/publicips/{publicip_id}|
|vpc|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1/{tenant_id}/publicips/{publicip_id}|
|vpc|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /v1/{tenant_id}/publicips|
|vpc|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1/{tenant_id}/publicips/{publicip_id}|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
This interface is used to apply for an elastic IP address.
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
This interface is used to delete an elastic IP address.
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
This interface is used to query details about an elastic IP address.
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
This interface is used to query elastic IP addresses using search criteria and to display the elastic IP addresses in a list.
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to bind an elastic IP address to a NIC or unbind an elastic IP address from a NIC.
