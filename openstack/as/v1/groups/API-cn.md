# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/groups"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 根据输入条件过滤查询弹性伸缩组。查询结果分页显示。

    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := groups.List(client, tenantId, groups.ListOpts{
        Limit: 2,
    }).Extract()
    if err != nil {
        panic(err)
    }
    
示例代码, 修改一个指定弹性伸缩组中的信息。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    result, err := groups.Update(client, tenantId, scalingGroupId, groups.UpdateOpts{
        ScalingGroupName:       "as-group-TestModified",
        ScalingConfigurationId: "f109bb4f-09f0-4dac-9115-6b429d548750",
        DesireInstanceNumber:   0,
        MinInstanceNumber:      0,
        MaxInstanceNumber:      1,
        CoolDownTime:           900,
        LbListenerId:           "",
        Networks: []struct {
            // Specifies the network ID.
            Id string `json:"id,omitempty"`
        }{
            {Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373",},
        },
        SecurityGroups: []struct {
            // Specifies the ID of the security group.
            Id string `json:"id,omitempty"`
        }{
            {Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",},
        },
        HealthPeriodicAuditMethod: "NOVA_AUDIT",
        HealthPeriodicAuditTime:   5,
        InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
        Notifications:             []string{},
        DeletePublicip:            true,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 创建弹性伸缩组。伸缩组是具有相同应用场景配置的实例集合。伸缩组内定义了实例最大数量和最小数量等信息。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := groups.Create(client, tenantId, groups.CreateOpts{
        ScalingGroupName:       "as-group-Test",
        ScalingConfigurationId: "f109bb4f-09f0-4dac-9115-6b429d548750",
        DesireInstanceNumber:   0,
        MinInstanceNumber:      0,
        MaxInstanceNumber:      1,
        CoolDownTime:           900,
        LbListenerId:           "",
        Networks: []struct {
            // Specifies the network ID.
            Id string `json:"id,omitempty"`
        }{
            {Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373",},
        },
        AvailableZones: []string{"cn-north-1a"},
        SecurityGroups: []struct {
            // Specifies the ID of the security group.
            Id string `json:"id,omitempty"`
        }{
            {Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",},
        },
        VpcId:                     "773c3c42-d315-417b-9063-87091713148c",
        HealthPeriodicAuditMethod: "NOVA_AUDIT",
        HealthPeriodicAuditTime:   5,
        InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
        Notifications:             []string{},
        DeletePublicip:            true,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
    
示例代码, 查询一个指定弹性伸缩组详情。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    result, err := groups.Get(client, tenantId, groupId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 启用或停止一个指定弹性伸缩组。停止一个弹性伸缩组后，就不会根据策略产生弹性伸缩活动

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    result := groups.Enable(client, tenantId, groupId, groups.EnableOpts{
        Action: "resume",
    })
    
示例代码, 删除一个指定弹性伸缩组。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    groups.Delete(client, tenantId, groupId, groups.DeleteOpts{
        ForceDelete: "no",
    })
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)](#func-delete)**  
**[func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)](#func-enable)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
**[func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_group|
|as|func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_group/{scaling_group_id}|
|as|func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_group/{scaling_group_id}/action|
|as|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_group/{scaling_group_id}|
|as|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_group|
|as|func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_group/{scaling_group_id}|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
创建弹性伸缩组。伸缩组是具有相同应用场景配置的实例集合。伸缩组内定义了实例最大数量和最小数量等信息。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, DeleteOptsBuilder) (DeleteResult)  
删除一个指定弹性伸缩组。
## func Enable
    func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)  
启用或停止一个指定弹性伸缩组。停止一个弹性伸缩组后，就不会根据策略产生弹性伸缩活动
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询一个指定弹性伸缩组详情。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
根据输入条件过滤查询弹性伸缩组。查询结果分页显示。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, UpdateOptsBuilder) (UpdateResult)  
修改一个指定弹性伸缩组中的信息。
