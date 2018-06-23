package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/tags"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	resourceType := "scaling_group_tag"
	resourceId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	tags.Update(client.ServiceClient(), tenantId, resourceType, resourceId, tags.UpdateOpts{
		Tags: []struct {
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
	})
}

func TestListTenantTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListTenantTagsSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	resourceType := "scaling_group_tag"
	actual, err := tags.ListTenantTags(client.ServiceClient(), tenantId, resourceType).Extract()
	if err != nil {
		panic(err)
	}

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListTenantTagsResponse, actual)
}

func TestListResourceTags(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListResourceTagsSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	resourceType := "scaling_group_tag"
	resourceId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := tags.ListResourceTags(client.ServiceClient(), tenantId, resourceType, resourceId).Extract()
	if err != nil {
		panic(err)
	}

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResourceTagsResponse, actual)
}
