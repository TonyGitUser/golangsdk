/*


Sample Code, This interface is used to query an AS group notification list.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := notifications.List(client, tenantId, scalingGroupId).Extract()
    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a notification for a specified AS group.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    topicUrnId := "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1"
    result := notifications.Delete(client, tenantId, scalingGroupId, topicUrnId)

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to enable notification for an AS group. Each time this interface is called, the AS group adds a notification topic and scenario. Each AS group supports up to five topics. The notification topic is pre-configured by you in SMN. When the live network complies with the notification scenario that matches the notification topic, the AS group sends a notification to you.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := notifications.Enable(client, tenantId, scalingGroupId, notifications.EnableOpts{
        TopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
        TopicScene: [] string{
            "SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package notifications
