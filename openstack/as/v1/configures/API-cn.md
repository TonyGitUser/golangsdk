# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/configures"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 创建弹性伸缩配置。伸缩配置定义了用于创建弹性伸缩组中实例的配置。AutoScaling为某个伸缩组自动添加实例时，会根据配置创建实例。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := configures.Create(client, tenantId, configures.CreateOpts{
        ScalingConfigurationName: "as-config-test",
        InstanceConfig: configures.CreateInstanceConfig{
            InstanceName: "",
            InstanceId:   "",
            FlavorRef:    "s2.small.1",
            ImageRef:     "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
            Disk: []configures.Disk{
                {
                    Size:               40,
                    VolumeType:         "SATA",
                    DiskType:           "SYS",
                    DedicatedStorageId: "",
                    DataDiskImageId:    "",
                    SnapshotId:         "",
                },
            },
            KeyName:     "",
            Personality: nil,
            UserData:    "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 查询一个弹性伸缩配置的详细信息。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    configureId := "f109bb4f-09f0-4dac-9115-6b429d548750"
    result, err := configures.Get(client, tenantId, configureId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 根据输入条件，查询弹性伸缩配置列表，分页查询。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := configures.List(client, tenantId, configures.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除一个指定弹性伸缩配置。被伸缩组使用的配置不能被删除。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingConfigurationId := "34147aea-e7bb-4290-8248-794b26f355a6"
    result := configures.Delete(client, tenantId, scalingConfigurationId)
    
示例代码, 批量删除指定弹性伸缩配置。被伸缩组使用的配置不能被删除。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result := configures.DeleteWithBatch(client, tenantId, configures.DeleteWithBatchOpts{
        ScalingConfigurationId:[]string{
            "483b6f84-d0d0-451f-9993-adb28e09a721",
            "e24c26dc-8a26-4e9c-85bf-4ac1ddc090bf",
        },
    })
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func DeleteWithBatch(*golangsdk.ServiceClient, string, DeleteWithBatchOptsBuilder) (DeleteWithBatchResult)](#func-deletewithbatch)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_configuration|
|as|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_configuration/{scaling_configuration_id}|
|as|func DeleteWithBatch(*golangsdk.ServiceClient, string, DeleteWithBatchOptsBuilder) (DeleteWithBatchResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_configurations|
|as|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_configuration/{scaling_configuration_id}|
|as|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_configuration|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
创建弹性伸缩配置。伸缩配置定义了用于创建弹性伸缩组中实例的配置。AutoScaling为某个伸缩组自动添加实例时，会根据配置创建实例。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
删除一个指定弹性伸缩配置。被伸缩组使用的配置不能被删除。
## func DeleteWithBatch
    func DeleteWithBatch(*golangsdk.ServiceClient, string, DeleteWithBatchOptsBuilder) (DeleteWithBatchResult)  
批量删除指定弹性伸缩配置。被伸缩组使用的配置不能被删除。
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询一个弹性伸缩配置的详细信息。
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
根据输入条件，查询弹性伸缩配置列表，分页查询。
