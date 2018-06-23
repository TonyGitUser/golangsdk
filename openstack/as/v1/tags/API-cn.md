# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/tags"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述


示例代码, 根据租户id查询指定资源类型的标签列表。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    result, err := tags.ListTenantTags(client, tenantId, resourceType).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 根据租户id和资源id查询指定资源类型的资源标签列表。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    resourceId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := tags.ListResourceTags(client, tenantId, resourceType, resourceId).Extract()
    
    if err != nil {
        panic(err)
    }
    
示例代码, 更新或删除指定资源的标签。每个伸缩组最多添加10个标签。

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    resourceId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := tags.Update(client, tenantId, resourceType, resourceId, tags.UpdateOpts{
        Tags: [] struct {
            Key   string `json:"key,omitempty"`
            Value string `json:"value,omitempty"`
        }{
            {
                Key:   "ENV15",
                Value: "ENV15",
            },
            {
                Key:   "ENV151",
                Value: "ENV151",
            },
        },
        Action: "update",
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## 目录
**[func ListResourceTags(*golangsdk.ServiceClient, string, string, string) (ListResourceTagsResult)](#func-listresourcetags)**  
**[func ListTenantTags(*golangsdk.ServiceClient, string, string) (ListTenantTagsResult)](#func-listtenanttags)**  
**[func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|as|func ListResourceTags(*golangsdk.ServiceClient, string, string, string) (ListResourceTagsResult)|GET /autoscaling-api/v1/{tenant_id}/{resource_type}/{resource_id}/tags|
|as|func ListTenantTags(*golangsdk.ServiceClient, string, string) (ListTenantTagsResult)|GET /autoscaling-api/v1/{tenant_id}/{resource_type}/tags|
|as|func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)|POST /autoscaling-api/v1/{tenant_id}/{resource_type}/{resource_id}/tags/action|
## 开始
## func ListResourceTags
    func ListResourceTags(*golangsdk.ServiceClient, string, string, string) (ListResourceTagsResult)  
根据租户id和资源id查询指定资源类型的资源标签列表。
## func ListTenantTags
    func ListTenantTags(*golangsdk.ServiceClient, string, string) (ListTenantTagsResult)  
根据租户id查询指定资源类型的标签列表。
## func Update
    func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)  
更新或删除指定资源的标签。每个伸缩组最多添加10个标签。
