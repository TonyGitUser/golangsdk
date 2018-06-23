# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/notifications"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to query an AS group notification list.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := notifications.List(client, tenantId, scalingGroupId).Extract()
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to delete a notification for a specified AS group.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    topicUrnId := "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1"
    result := notifications.Delete(client, tenantId, scalingGroupId, topicUrnId)
    
    if err != nil {
        panic(err)
    }
    
Sample Code, This interface is used to enable notification for an AS group. Each time this interface is called, the AS group adds a notification topic and scenario. Each AS group supports up to five topics. The notification topic is pre-configured by you in SMN. When the live network complies with the notification scenario that matches the notification topic, the AS group sends a notification to you.

    
    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := notifications.Enable(client, tenantId, scalingGroupId, notifications.EnableOpts{
        TopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
        TopicScene: [] string{
            "SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL",
        },
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)](#func-delete)**  
**[func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)](#func-enable)**  
**[func List(*golangsdk.ServiceClient, string, string) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|as|func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)|DELETE /autoscaling-api/v1/{tenant_id}/scaling_notification/{scaling_group_id}/{topic_urn}|
|as|func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)|PUT /autoscaling-api/v1/{tenant_id}/scaling_notification/{scaling_group_id}|
|as|func List(*golangsdk.ServiceClient, string, string) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_notification/{scaling_group_id}|
## Content
## func Delete
    func Delete(*golangsdk.ServiceClient, string, string, string) (DeleteResult)  
This interface is used to delete a notification for a specified AS group.
## func Enable
    func Enable(*golangsdk.ServiceClient, string, string, EnableOptsBuilder) (EnableResult)  
This interface is used to enable notification for an AS group. Each time this interface is called, the AS group adds a notification topic and scenario. Each AS group supports up to five topics. The notification topic is pre-configured by you in SMN. When the live network complies with the notification scenario that matches the notification topic, the AS group sends a notification to you.
## func List
    func List(*golangsdk.ServiceClient, string, string) (ListResult)  
This interface is used to query an AS group notification list.
