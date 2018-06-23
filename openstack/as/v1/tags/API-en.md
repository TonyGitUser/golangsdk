# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/tags"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to query tags of a specified resource type by tenant ID.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    result, err := tags.ListTenantTags(client, tenantId, resourceType).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to query tags of a specified resource type by tenant ID and resource ID.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    resourceId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := tags.ListResourceTags(client, tenantId, resourceType, resourceId).Extract()
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to update or delete a resource tag.Each AS group can have a maximum of 10 tags added to it.

    
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
## Index
**[func ListResourceTags(*golangsdk.ServiceClient, string, string, string) (ListResourceTagsResult)](#func-listresourcetags)**  
**[func ListTenantTags(*golangsdk.ServiceClient, string, string) (ListTenantTagsResult)](#func-listtenanttags)**  
**[func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)](#func-update)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|as|func ListResourceTags(*golangsdk.ServiceClient, string, string, string) (ListResourceTagsResult)|GET /autoscaling-api/v1/{tenant_id}/{resource_type}/{resource_id}/tags|
|as|func ListTenantTags(*golangsdk.ServiceClient, string, string) (ListTenantTagsResult)|GET /autoscaling-api/v1/{tenant_id}/{resource_type}/tags|
|as|func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)|POST /autoscaling-api/v1/{tenant_id}/{resource_type}/{resource_id}/tags/action|
## Content
## func ListResourceTags
    func ListResourceTags(*golangsdk.ServiceClient, string, string, string) (ListResourceTagsResult)  
This interface is used to query tags of a specified resource type by tenant ID and resource ID.
## func ListTenantTags
    func ListTenantTags(*golangsdk.ServiceClient, string, string) (ListTenantTagsResult)  
This interface is used to query tags of a specified resource type by tenant ID.
## func Update
    func Update(*golangsdk.ServiceClient, string, string, string, UpdateOptsBuilder) (UpdateResult)  
This interface is used to update or delete a resource tag.Each AS group can have a maximum of 10 tags added to it.
