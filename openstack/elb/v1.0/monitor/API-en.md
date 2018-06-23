# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/monitor"
**[Overview](#overview)**  

**[Index](#index)**  

**[API Mapping](#api-mapping)**  

**[Content](#content)**  

## Overview
This interface is used to query all layer-4 and layer-7 traffic monitoring metrics.Only common tenants can query these monitoring metrics.

Sample Code, This interface is used to query all layer-4 and layer-7 traffic monitoring metrics.Only common tenants can query these monitoring metrics.

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := monitor.List(client, tenantID).Extract()
    
    if err != nil {
        panic(err)
    }
## Index
**[func List(*golangsdk.ServiceClient, string) (ListResult)](#func-list)**  
## API Mapping
|Catalog|API|EndPoint|
|----|---|--------|
|elb|func List(*golangsdk.ServiceClient, string) (ListResult)|GET /v1.0/{tenant_id}/elbaas/monitor|
## Content
## func List
    func List(*golangsdk.ServiceClient, string) (ListResult)  
This interface is used to query all layer-4 and layer-7 traffic monitoring metrics.Only common tenants can query these monitoring metrics.
