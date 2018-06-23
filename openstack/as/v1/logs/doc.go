/*


Sample Code, This interface is used to query scaling action logs. The results are filtered based on the conditions you input and displayed by page.


    tenantId  := "57e98940a77f4bb988a21a7d0603a626"
    groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := logs.List(client, tenantId, groupId, logs.ListOpts{
        Limit: 2,
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package logs
