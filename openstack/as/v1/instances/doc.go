/*


Sample Code, This interface is used to remove instances from an AS group or add instances to an AS group in batches, or enable or disable instance protection on the instances in an AS group in batches.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := instances.Action(client, tenantId, groupId, instances.ActionOpts{
        InstancesId:     []string{
            "e1040c67-b130-41ee-9405-07c6c6c20208",
        },
        InstanceDelete: "no",
        Action:         "ADD",
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query information about instances in an AS group. The results are filtered based on the conditions you input and displayed by page.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := instances.List(client, tenantId, groupId, instances.ListOpts{
        Limit: 2,
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to remove instances from a specified AS group.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "e1040c67-b130-41ee-9405-07c6c6c20208"
    result := instances.Delete(client, tenantId, instanceId, instances.DeleteOpts{
       InstanceDelete: "no",
    })

    if err != nil {
       panic(err)
    }
*/
package instances
