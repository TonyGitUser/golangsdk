/*
This interface is used to query network resource quotas for the VPC service of a tenant. The network resources include VPCs, subnets, security groups, security group rules, elastic IP addresses, and VPNs.

Sample Code, This interface is used to create a security group.


    result, err := securitygroups.Create(client, securitygroups.CreateOpts{
        SecurityGroup: securitygroups.CreateSecurityGroup{
            Name:        "EricSG",
            Description: "Test SecurityGroup",
            TenantId: tenantID,
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query details about a security group.


    result, err := securitygroups.Get(client, "a988649d-cfbf-4c2a-bd91-3b84d2403747").Extract()

    if err != nil {
      panic(err)
    }

Sample Code, This interface is used to query security groups using search criteria and to display the security groups in a list.


    allPages, err := securitygroups.List(client, securitygroups.ListOpts{
        Limit: 2,
    }).AllPages()

    result, err := securitygroups.ExtractList(allPages.(securitygroups.ListPage))

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a security group.


    result := securitygroups.Delete(client, tenantID, "2465d913-1084-4a6a-91e7-2fd6f490ecb3")

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to update a security group.


    result, err := securitygroups.Update(client, "7af80d49-0a43-462d-aed8-a1e12ac91af6", securitygroups.UpdateOpts{
        SecurityGroup: securitygroups.UpdateSecurityGroup{
            Name:        "EricSG",
            Description: "Modified SecurityGroup",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package securitygroups
