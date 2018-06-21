/*
A subnet is a network that manages ECS network planes. It supports IP address management and DNS. The IP addresses of all ECSs in a subnet belong to the subnet. By default, ECSs in all subnets of the same VPC can communicate with one another, while ECSs in different VPCs cannot communicate with one another.

Sample Code, This interface is used to create a subnet.


    result, err := subnets.Create(client, subnets.CreateOpts{
    Subnet: subnets.CreateSubnet{
        Name:      "subnetV2",
        IpVersion: 4,
        NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
        Cidr:      "192.168.20.0/24",
        GatewayIp: "192.168.20.1",
        AllocationPools: []subnets.AllocationPool{
            {
                Start: "192.168.20.2",
                End:   "192.168.20.254",
            },
        },
        EnableDhcp:     true,
        DnsNameservers: [] string{"114.114.114.114", "114.114.115.115"},
        TenantId:       "57e98940a77f4bb988a21a7d0603a626",
    },
    }).Extract()

    if err != nil {
    panic(err)
    }

Sample Code, This interface is used to update information about a subnet.


    result, err := subnets.Update(client, "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28", subnets.UpdateOpts{
        Subnet: subnets.UpdateSubnet{
            Name:           "subnetV2",
            GatewayIp:      "192.168.20.1",
            EnableDhcp:     true,
            DnsNameservers: [] string{"8.8.8.8", "4.4.4.4"},
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query details about a subnet.

    result, err := subnets.Get(client, "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28").Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query subnets using search criteria and to display the subnets in a list.

    allPages, err := subnets.List(client, subnets.ListOpts {
        Limit: 2,
    }).AllPages()

    result, err := subnets.ExtractList(allPages.(subnets.ListPage))

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a subnet.


    result := subnets.Delete(client, "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28")

    if err != nil {
      panic(err)
    }

*/
package subnets
