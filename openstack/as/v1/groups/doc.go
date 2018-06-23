/*


Sample Code, This interface is used to query AS groups based on conditions you input. The results are displayed by page.

    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := groups.List(client, tenantId, groups.ListOpts{
        Limit: 2,
    }).Extract()
    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to modify information about a specified AS group.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    result, err := groups.Update(client, tenantId, scalingGroupId, groups.UpdateOpts{
        ScalingGroupName:       "as-group-TestModified",
        ScalingConfigurationId: "f109bb4f-09f0-4dac-9115-6b429d548750",
        DesireInstanceNumber:   0,
        MinInstanceNumber:      0,
        MaxInstanceNumber:      1,
        CoolDownTime:           900,
        LbListenerId:           "",
        Networks: []struct {
            // Specifies the network ID.
            Id string `json:"id,omitempty"`
        }{
            {Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373",},
        },
        SecurityGroups: []struct {
            // Specifies the ID of the security group.
            Id string `json:"id,omitempty"`
        }{
            {Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",},
        },
        HealthPeriodicAuditMethod: "NOVA_AUDIT",
        HealthPeriodicAuditTime:   5,
        InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
        Notifications:             []string{},
        DeletePublicip:            true,
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to create an AS group. An AS group is a set of ECS instances with the same application scenario configurations. An AS group defines the minimum and maximum numbers of ECS instances.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := groups.Create(client, tenantId, groups.CreateOpts{
        ScalingGroupName:       "as-group-Test",
        ScalingConfigurationId: "f109bb4f-09f0-4dac-9115-6b429d548750",
        DesireInstanceNumber:   0,
        MinInstanceNumber:      0,
        MaxInstanceNumber:      1,
        CoolDownTime:           900,
        LbListenerId:           "",
        Networks: []struct {
            // Specifies the network ID.
            Id string `json:"id,omitempty"`
        }{
            {Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373",},
        },
        AvailableZones: []string{"cn-north-1a"},
        SecurityGroups: []struct {
            // Specifies the ID of the security group.
            Id string `json:"id,omitempty"`
        }{
            {Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",},
        },
        VpcId:                     "773c3c42-d315-417b-9063-87091713148c",
        HealthPeriodicAuditMethod: "NOVA_AUDIT",
        HealthPeriodicAuditTime:   5,
        InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
        Notifications:             []string{},
        DeletePublicip:            true,
    }).Extract()

    if err != nil {
        panic(err)
    }


Sample Code, This interface is used to query details about a specified AS group.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    result, err := groups.Get(client, tenantId, groupId).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to enable or disable a specified AS group.If an AS group is disabled, AS actions will not be triggered based on the AS policy.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    result := groups.Enable(client, tenantId, groupId, groups.EnableOpts{
        Action: "resume",
    })

Sample Code, This interface is used to delete a specified AS group.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
    groups.Delete(client, tenantId, groupId, groups.DeleteOpts{
        ForceDelete: "no",
    })
*/
package groups
