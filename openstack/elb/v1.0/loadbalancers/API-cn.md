# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/loadbalancers"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
弹性负载均衡（Elastic Load Balance ，以下简称ELB）通过监听规则将访问流量自动分发到后端多台弹性云服务器，扩展应用系统对外的服务能力，实现更高水平的应用程序容错性能。

示例代码, 创建负载均衡器。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := loadbalancers.Create(client, tenantID, loadbalancers.CreateOpts{
        Bandwidth:       1,
        SecurityGroupId: "",
        VpcId:           "773c3c42-d315-417b-9063-87091713148c",
        AdminStateUp:    1,
        VipSubnetId:     "",
        Type:            "External",
        Name:            "TestABC",
        Description:     "Show Load Balancer",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除负载均衡器。如果是外网负载均衡器，该接口会同时删除负载均衡器绑定的弹性IP。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
    loadbalancers.Delete(client, tenantID, loadBalancerId)
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除外网负载均衡器时保留弹性IP，如果需要删除弹性IP，请参考删除负载均衡器章节。

    
     tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ebe982b8afe04851bbc26ad4609003bf"
    loadbalancers.DeleteKeepEIP(client, tenantID, loadBalancerId)
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询指定负载均衡器的详细信息。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
    result, err := loadbalancers.Get(client, tenantID, loadBalancerId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询某租户下所有的负载均衡器的信息。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := loadbalancers.List(client, tenantID).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 修改负载均衡器，名称、描述、带宽和管理状态可修改。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
    result, err := loadbalancers.Update(client, tenantID, loadBalancerId, loadbalancers.UpdateOpts{
        Name:        "UpdatedLB",
        Description: "TEST LB",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func DeleteKeepEIP(*golangsdk.ServiceClient, string, string) (DeleteKeepEIPResult)](#func-deletekeepeip)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|elb|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /v1.0/{tenant_id}/elbaas/loadbalancers|
|elb|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /v1.0/{tenant_id}/elbaas/loadbalancers/{loadbalancer_id}|
|elb|func DeleteKeepEIP(*golangsdk.ServiceClient, string, string) (DeleteKeepEIPResult)|DELETE /v1.0/{tenant_id}/elbaas/loadbalancers/{loadbalancer_id}/keep-eip|
|elb|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1.0/{tenant_id}/elbaas/loadbalancers/{loadbalancer_id}|
|elb|func List(*golangsdk.ServiceClient, string) (ListResult)|GET /v1.0/{tenant_id}/elbaas/loadbalancers|
|elb|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /v1.0/{tenant_id}/elbaas/loadbalancers/{loadbalancer_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
创建负载均衡器。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除负载均衡器。如果是外网负载均衡器，该接口会同时删除负载均衡器绑定的弹性IP。
## func DeleteKeepEIP
    func DeleteKeepEIP(*golangsdk.ServiceClient, string, string) (DeleteKeepEIPResult)  
删除外网负载均衡器时保留弹性IP，如果需要删除弹性IP，请参考删除负载均衡器章节。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询指定负载均衡器的详细信息。
## func List
    func List(*golangsdk.ServiceClient, string) (ListResult)  
查询某租户下所有的负载均衡器的信息。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
修改负载均衡器，名称、描述、带宽和管理状态可修改。
