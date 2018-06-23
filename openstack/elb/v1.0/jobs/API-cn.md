# Package antiddos
    import "github.com/huaweicloud/golangsdk/openstack/elb/v1.0/jobs"
**[概述](#概述)**  

**[目录](#目录)**  

**[API对应表](#API对应表)**  

**[开始](#开始)**  

## 概述
查询job的状态，可用于查询创建负载均衡器、删除负载均衡器等API的执行状态。

示例代码, 查询job的状态，可用于查询创建负载均衡器、删除负载均衡器等API的执行状态。

    
    tenantID := "57e98940a77f4bb988a21a7d0603a626"
     jobId := "8aace0c86326772c016326c9f5dd0f3f"
     result, err := jobs.Get(client, tenantID, jobId).Extract()
    
     if err != nil {
         panic(err)
     }
## 目录
**[func Get(*golangsdk.ServiceClient, string, string) (GetResult)](#func-get)**  
## API对应表
|类别|API|EndPoint|
|----|---|--------|
|elb|func Get(*golangsdk.ServiceClient, string, string) (GetResult)|GET /v1.0/{tenant_id}/jobs/{job_id}|
## 开始
## func Get
    func Get(*golangsdk.ServiceClient, string, string) (GetResult)  
查询job的状态，可用于查询创建负载均衡器、删除负载均衡器等API的执行状态。
