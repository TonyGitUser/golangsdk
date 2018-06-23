# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/quotas"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to query quotas of a tenant.

Sample Code, This interface is used to query quotas of a tenant.

## Index
**[func List(*golangsdk.ServiceClient, string) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|elb|func List(*golangsdk.ServiceClient, string) (ListResult)|GET /v1.0/{tenant_id}/elbaas/quotas|
## Content
## func List
    func List(*golangsdk.ServiceClient, string) (ListResult)  
This interface is used to query quotas of a tenant.
