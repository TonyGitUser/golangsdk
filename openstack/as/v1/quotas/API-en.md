# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/as/v1/quotas"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview


## Index
**[func List(*golangsdk.ServiceClient, string) (ListResult)](#func-list)**  
**[func ListWithInstances(*golangsdk.ServiceClient, string, string) (ListWithInstancesResult)](#func-listwithinstances)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|as|func List(*golangsdk.ServiceClient, string) (ListResult)|GET /autoscaling-api/v1/{tenant_id}/quotas|
|as|func ListWithInstances(*golangsdk.ServiceClient, string, string) (ListWithInstancesResult)|GET /autoscaling-api/v1/{tenant_id}/quotas/{scaling_group_id}|
## Content
## func List
    func List(*golangsdk.ServiceClient, string) (ListResult)  
This interface is used to query the total quota amount and used quota amount of AS groups and AS configurations for a specified tenant.In addition, this interface is also used to query the total amount of AS policies and AS instances of a specified tenant.
## func ListWithInstances
    func ListWithInstances(*golangsdk.ServiceClient, string, string) (ListWithInstancesResult)  
This interface is used to query the total quota amount and used quota amount of AS policies and AS instances of a specified AS group.
