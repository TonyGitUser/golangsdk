/*
This interface is used to query the status of tasks, such as creating or deleting a load balancer.

Sample Code, This interface is used to query the status of tasks, such as creating or deleting a load balancer.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
     jobId := "8aace0c86326772c016326c9f5dd0f3f"
     result, err := jobs.Get(client, tenantID, jobId).Extract()

     if err != nil {
         panic(err)
     }
*/
package jobs
