package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/logs"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := logs.List(client.ServiceClient(), tenantId, groupId, logs.ListOpts{
		Limit: 2,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}
