package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/quotas"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := quotas.List(client.ServiceClient(), tenantID, quotas.ListOpts{
		Type: "vpc",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}
