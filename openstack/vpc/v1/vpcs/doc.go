/*
The Virtual Private Cloud (VPC) service enables you to provision logically isolated, configurable, and manageable virtual networks for Elastic Cloud Servers (ECSs), improving the security of resources in the cloud system and simplifying network deployment.

Sample Code, This interface is used to create a VPC.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.Create(client, tenantID, vpcs.CreateOpts{
        Vpc: vpcs.VPC{
            Name: "ABC",
            Cidr: "192.168.0.0/16",
        },
    }).Extract()

Sample Code, This interface is used to update information about a VPC.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.Update(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0", vpcs.UpdateOpts{
        Vpc: vpcs.VPC{
            Name: "ABC-back",
            Cidr: "192.168.0.0/24",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query details about a VPC.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
        result, err := vpcs.Get(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0").Extract()

Sample Code, This interface is used to query VPCs using search criteria and to display the VPCs in a list.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := vpcs.List(client, tenantID, vpcs.ListOpts{
       Limit: 2,
    }).Extract()

    if err != nil {
       panic(err)
    }

Sample Code, This interface is used to delete a VPC.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := vpcs.Delete(client, tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0")
*/
package vpcs
