# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/rds//instances"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 获取实例列表。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := instances.List(client, versionId, projectId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 获取指定实例详细信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := instances.Get(client, versionId, projectId, instanceId).Extract()
    
    if err != nil {
        panic(err)
    }
    
    
示例代码, 获取指定数据库ID和区域下的所有实例规格信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := instances.ListFlavors(client, versionId, projectId, instances.ListFlavorsOpts{
        DbId:   "9dae9226-5f8e-4b10-bd41-adae2e693e89",
        Region: "cn-north-1",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 获取指定规格ID下的实例规格信息。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    favorId := "30a0baf6-d4c2-44d1-8569-68844512ae3a"
    result, err := instances.GetFlavor(client, versionId, projectId, favorId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 创建的RDS实例，可分为单实例、HA实例和只读副本。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := instances.Create(client, versionId, projectId, instances.CreateOpts{
        Instance: struct {
            Name string `json:"name,omitempty"`
            Datastore struct {
                Type    string `json:"type,omitempty"`
                Version string `json:"version,omitempty"`
            } `json:"datastore,omitempty"`
            FlavorRef string `json:"flavorRef,omitempty"`
            Volume struct {
                Type string `json:"type,omitempty"`
                Size int    `json:"size,omitempty"`
            } `json:"volume,omitempty"`
            Region           string `json:"region,omitempty"`
            AvailabilityZone string `json:"availabilityZone,omitempty"`
            Vpc              string `json:"vpc,omitempty"`
            Nics struct {
                SubnetId string `json:"subnetId,omitempty"`
            } `json:"nics,omitempty"`
            SecurityGroup struct {
                Id string `json:"id,omitempty"`
            } `json:"securityGroup,omitempty"`
            DbPort int `json:"dbPort,omitempty"`
            BackupStrategy struct {
                StartTime string `json:"startTime,omitempty"`
                KeepDays  int    `json:"keepDays,omitempty"`
            } `json:"backupStrategy,omitempty"`
            DbRtPd string `json:"dbRtPd,omitempty"`
            Ha struct {
                Enable          bool   `json:"enable,omitempty"`
                ReplicationMode string `json:"replicationMode,omitempty"`
            } `json:"ha,omitempty"`
        }{
            Name:             "MYSQL_TEST_CREATE",
            Region:           "cn-north-1",
            AvailabilityZone: "cn-north-1a",
            Vpc:              "773c3c42-d315-417b-9063-87091713148c",
            Datastore: struct {
                Type    string `json:"type,omitempty"`
                Version string `json:"version,omitempty"`
            }{
                Type:    "MySQL",
                Version: "5.6.30",
            },
            Nics: struct {
                SubnetId string `json:"subnetId,omitempty"`
            }{
                SubnetId: "f6a0db7b-397c-4162-bc35-9a1f008b4373",
            },
            SecurityGroup: struct {
                Id string `json:"id,omitempty"`
            }{
                Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
            },
            FlavorRef: "18207182-0c02-467f-ae05-9c1650df1717",
            Volume: struct {
                Type string `json:"type,omitempty"`
                Size int    `json:"size,omitempty"`
            }{
                Type: "COMMON",
                Size: 100,
            },
            DbPort: 8635,
            BackupStrategy: struct {
                StartTime string `json:"startTime,omitempty"`
                KeepDays  int    `json:"keepDays,omitempty"`
            }{
                StartTime: "19:40:00",
                KeepDays:  7,
            },
            DbRtPd: "@xA123456",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 调整实例的规格。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instancesId := "a0472c5aafac40688c82f8980f3ecd0ano01"
    result, err := instances.Update(client, versionId, projectId, instancesId, instances.UpdateOpts{
        Resize: struct {
            FlavorRef string `json:"flavorRef,omitempty"`
            Volume struct {
                Size int `json:"size,omitempty"`
            } `json:"volume,omitempty"`
        } {
            Volume: struct {
                Size int `json:"size,omitempty"`
            } {
                Size: 220,
            },
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 重启实例的数据库服务。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instancesId := "a0472c5aafac40688c82f8980f3ecd0ano01"
    result, err := instances.Reboot(client, versionId, projectId, instancesId, instances.RebootOpts{
        Restart: struct{}{},
    }).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 删除数据库实例。

    
    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instancesId := "a0472c5aafac40688c82f8980f3ecd0ano01"
    result, err := instances.Delete(client, versionId, projectId, instancesId, instances.DeleteOpts{}).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)](#func-create)**  
**[func Delete(*golangsdk.ServiceClient, string, string, string, DeleteOptsBuilder) (DeleteResult)](#func-delete)**  
**[func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)](#func-get)**  
**[func GetFlavor(*golangsdk.ServiceClient, string, string, string) (GetFlavorResult)](#func-getflavor)**  
**[func List(*golangsdk.ServiceClient, string, string) (ListResult)](#func-list)**  
**[func ListFlavors(*golangsdk.ServiceClient, string, string, ListFlavorsOptsBuilder) (ListFlavorsResult)](#func-listflavors)**  
**[func Reboot(*golangsdk.ServiceClient, string, string, string, RebootOptsBuilder) (RebootResult)](#func-reboot)**  
**[func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|rds|func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)|POST /rds/{versionId}/{project_id}/instances|
|rds|func Delete(*golangsdk.ServiceClient, string, string, string, DeleteOptsBuilder) (DeleteResult)|DELETE /rds/{versionId}/{project_id}/instances/{instanceId}|
|rds|func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)|GET /rds/{versionId}/{project_id}/instances/{instanceId}|
|rds|func GetFlavor(*golangsdk.ServiceClient, string, string, string) (GetFlavorResult)|GET /rds/{versionId}/{project_id}/flavors/{flavorId}|
|rds|func List(*golangsdk.ServiceClient, string, string) (ListResult)|GET /rds/{versionId}/{project_id}/instances|
|rds|func ListFlavors(*golangsdk.ServiceClient, string, string, ListFlavorsOptsBuilder) (ListFlavorsResult)|GET /rds/{versionId}/{project_id}/flavors|
|rds|func Reboot(*golangsdk.ServiceClient, string, string, string, RebootOptsBuilder) (RebootResult)|POST /rds/{versionId}/{project_id}/instances/{instanceId}/action|
|rds|func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)|POST /rds/{versionId}/{project_id}/instances/{instanceId}/action|
## 开始
## func Create
    func Create(*golangsdk.ServiceClient, string, string, CreateOptsBuilder) (CreateResult)  
创建的RDS实例，可分为单实例、HA实例和只读副本。
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, string, DeleteOptsBuilder) (DeleteResult)  
删除数据库实例。
## func Get
    func Get(*golangsdk.ServiceClient, string, string, string) (GetResult)  
获取指定实例详细信息。
## func GetFlavor
    func GetFlavor(*golangsdk.ServiceClient, string, string, string) (GetFlavorResult)  
获取指定规格ID下的实例规格信息。
## func List
    func List(*golangsdk.ServiceClient, string, string) (ListResult)  
获取实例列表。
## func ListFlavors
    func ListFlavors(*golangsdk.ServiceClient, string, string, ListFlavorsOptsBuilder) (ListFlavorsResult)  
获取指定数据库ID和区域下的所有实例规格信息。
## func Reboot
    func Reboot(*golangsdk.ServiceClient, string, string, string, RebootOptsBuilder) (RebootResult)  
重启实例的数据库服务。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)  
调整实例的规格。
