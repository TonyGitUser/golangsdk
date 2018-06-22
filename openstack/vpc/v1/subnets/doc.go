/*
A subnet is a network that manages ECS network planes. It supports IP address management and DNS. The IP addresses of all ECSs in a subnet belong to the subnet. By default, ECSs in all subnets of the same VPC can communicate with one another, while ECSs in different VPCs cannot communicate with one another.

Sample Code, This interface is used to create a subnet.

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.Create(client, tenantID, subnets.CreateOpts{
       Subnet: subnets.Subnet {
           Name: "subnet",
           Cidr: "192.168.20.0/24",
           GatewayIp: "192.168.20.1",
           DhcpEnable: true,
           PrimaryDns: "114.114.114.114",
           SecondaryDns: "114.114.115.115",
           DnsList: [] string{
               "114.114.114.114",
               "114.114.115.115",
           },
           AvailabilityZone:"cn-north-1a",
           VpcId:"ea3b0efe-0d6a-4288-8b16-753504a994b9",
       },
    }).Extract()

    if err != nil {
       panic(err)
    }

Sample Code, This interface is used to update information about a subnet.

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.Update(client, tenantID, "ea3b0efe-0d6a-4288-8b16-753504a994b9","c9aba52d-ec14-40cb-930f-c8153e93c2db", subnets.UpdateOpts{
       Subnet: subnets.Subnet{
           Name: "ABC-back",
           Cidr: "192.168.0.0/24",
       },
    }).Extract()

    if err != nil {
       panic(err)
    }

Sample Code, This interface is used to query details about a subnet.

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.Get(client, tenantID, "c9aba52d-ec14-40cb-930f-c8153e93c2db").Extract()

Sample Code, This interface is used to query subnets using search criteria and to display the subnets in a list.

    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := subnets.List(client, tenantID, subnets.ListOpts{
      Limit: 2,
    }).Extract()

    if err != nil {
      panic(err)
    }

Sample Code, This interface is used to delete a subnet.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := subnets.Delete(client, tenantID, "ea3b0efe-0d6a-4288-8b16-753504a994b9","c9aba52d-ec14-40cb-930f-c8153e93c2db")
*/
package subnets
