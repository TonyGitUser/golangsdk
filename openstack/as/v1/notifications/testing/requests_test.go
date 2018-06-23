package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/notifications"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestEnable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleEnableSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := notifications.Enable(client.ServiceClient(), tenantId, scalingGroupId, notifications.EnableOpts{
		TopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
		TopicScene: []string{
			"SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL",
		},
	}).Extract()

	if err != nil {
		panic(err)
	}
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &EnableResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := notifications.List(client.ServiceClient(), tenantId, scalingGroupId).Extract()
	if err != nil {
		panic(err)
	}

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	topicUrnId := "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1"
	notifications.Delete(client.ServiceClient(), tenantId, scalingGroupId, topicUrnId)
}
