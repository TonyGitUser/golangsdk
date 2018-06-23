/*


Sample Code, This interface is used to query tags of a specified resource type by tenant ID.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    result, err := tags.ListTenantTags(client, tenantId, resourceType).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query tags of a specified resource type by tenant ID and resource ID.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    resourceId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := tags.ListResourceTags(client, tenantId, resourceType, resourceId).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to update or delete a resource tag.Each AS group can have a maximum of 10 tags added to it.


    tenantId := "57e98940a77f4bb988a21a7d0603a626"
    resourceType := "scaling_group_tag"
    resourceId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
    result, err := tags.Update(client, tenantId, resourceType, resourceId, tags.UpdateOpts{
        Tags: [] struct {
            Key   string `json:"key,omitempty"`
            Value string `json:"value,omitempty"`
        }{
            {
                Key:   "ENV15",
                Value: "ENV15",
            },
            {
                Key:   "ENV151",
                Value: "ENV151",
            },
        },
        Action: "update",
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package tags
