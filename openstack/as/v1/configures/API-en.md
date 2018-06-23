# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/configures"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to create an AS configuration. The AS configuration defines configurations of instances in the AS group. ECS instances are created based on the configurations if AS automatically adds instances to an AS group.

    
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
    
Sample Code, This interface is used to query details about an AS configuration.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    configureId := "f109bb4f-09f0-4dac-9115-6b429d548750"
    result, err := configures.Get(client, tenantId, configureId).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query AS configurations based on the conditions you input and display the results by page.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := configures.List(client, tenantId, configures.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a specified AS configuration. AS configurations used by AS groups cannot be deleted.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingConfigurationId := "34147aea-e7bb-4290-8248-794b26f355a6"
    result := configures.Delete(client, tenantId, scalingConfigurationId)
    
Sample Code, This interface is used to batch delete AS configurations. AS configurations used by AS groups cannot be deleted.

    
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
## Index
**[func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)](#func-delete)**  
**[func DeleteWithBatch(*golangsdk.ServiceClient, string, DeleteWithBatchOptsBuilder) (DeleteWithBatchResult)](#func-deletewithbatch)**  
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
**[func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|as|func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_configuration|
|as|func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_configuration/{scaling_configuration_id}|
|as|func DeleteWithBatch(*golangsdk.ServiceClient, string, DeleteWithBatchOptsBuilder) (DeleteWithBatchResult)|POST /autoscaling-api/v1/{tenant_id}/scaling_configurations|
|as|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_configuration/{scaling_configuration_id}|
|as|func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_configuration|
## Content
## func Create
    func Create(*golangsdk.ServiceClient, string, CreateOptsBuilder) (CreateResult)  
This interface is used to create an AS configuration. The AS configuration defines configurations of instances in the AS group. ECS instances are created based on the configurations if AS automatically adds instances to an AS group.
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string) (DeleteResult)  
This interface is used to delete a specified AS configuration. AS configurations used by AS groups cannot be deleted.
## func DeleteWithBatch
    func DeleteWithBatch(*golangsdk.ServiceClient, string, DeleteWithBatchOptsBuilder) (DeleteWithBatchResult)  
This interface is used to batch delete AS configurations. AS configurations used by AS groups cannot be deleted.
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
This interface is used to query details about an AS configuration.
## func List
    func List(*golangsdk.ServiceClient, string, ListOptsBuilder) (ListResult)  
This interface is used to query AS configurations based on the conditions you input and display the results by page.
