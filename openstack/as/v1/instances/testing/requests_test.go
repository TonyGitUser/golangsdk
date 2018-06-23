package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/instances"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestAction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleActionSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	instances.Action(client.ServiceClient(), tenantId, groupId, instances.ActionOpts{
		InstancesId: []string{
			"e1040c67-b130-41ee-9405-07c6c6c20208",
		},
		InstanceDelete: "no",
		Action:         "ADD",
	})
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := instances.List(client.ServiceClient(), tenantId, groupId, instances.ListOpts{
		Limit: 2,
	}).Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	instanceId := "e1040c67-b130-41ee-9405-07c6c6c20208"
	instances.Delete(client.ServiceClient(), tenantId, instanceId, instances.DeleteOpts{
		InstanceDelete: "no",
	})
}
