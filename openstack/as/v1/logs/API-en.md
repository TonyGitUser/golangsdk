# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/logs"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


Sample Code, This interface is used to query scaling action logs. The results are filtered based on the conditions you input and displayed by page.

    
    tenantId  := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := logs.List(client, tenantId, groupId, logs.ListOpts{
        Limit: 2,
    }).Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|as|func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/scaling_activity_log/{scaling_group_id}|
## Content
## func List
    func List(*golangsdk.ServiceClient, string, string, ListOptsBuilder) (ListResult)  
This interface is used to query scaling action logs. The results are filtered based on the conditions you input and displayed by page.
