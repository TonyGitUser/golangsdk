package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/quotas"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := quotas.List(client.ServiceClient(), tenantID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestListWithInstances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListWithInstancesSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
	actual, err := quotas.ListWithInstances(client.ServiceClient(), tenantId, groupId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListWithInstancesResponse, actual)
}
