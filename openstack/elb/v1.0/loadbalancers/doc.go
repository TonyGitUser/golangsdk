/*
Elastic Load Balance (ELB) is a service that automatically distributes access traffic to multiple Elastic Cloud Servers (ECSs) following configured listening rules to balance their workloads. It improves the fault tolerance and expands the service capabilities of your applications.

Sample Code, This interface is used to create a load balancer.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := loadbalancers.Create(client, tenantID, loadbalancers.CreateOpts{
        Bandwidth:       1,
        SecurityGroupId: "",
        VpcId:           "773c3c42-d315-417b-9063-87091713148c",
        AdminStateUp:    1,
        VipSubnetId:     "",
        Type:            "External",
        Name:            "TestABC",
        Description:     "Show Load Balancer",
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a load balancer. If the load balancer is a public network load balancer, this interface deletes the EIP bound to the load balancer.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
    loadbalancers.Delete(client, tenantID, loadBalancerId)

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a public network load balancer. The EIP bound to the load balancer will not be deleted. If you need to delete this IP address, refer to section Deleting a Load Balancer.


     tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ebe982b8afe04851bbc26ad4609003bf"
    loadbalancers.DeleteKeepEIP(client, tenantID, loadBalancerId)

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query details about a load balancer.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
    result, err := loadbalancers.Get(client, tenantID, loadBalancerId).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query load balancers using search criteria and to display the load balancers in a list.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := loadbalancers.List(client, tenantID).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to modify the name, description, bandwidth, and management status of a load balancer.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
    result, err := loadbalancers.Update(client, tenantID, loadBalancerId, loadbalancers.UpdateOpts{
        Name:        "UpdatedLB",
        Description: "TEST LB",
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package loadbalancers
