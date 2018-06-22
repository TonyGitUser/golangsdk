package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/bandwidths"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := bandwidths.Update(client.ServiceClient(), tenantID, "3c43e46e-4af1-45b8-a84d-ee6d04488d2a", bandwidths.UpdateOpts{
		Bandwidth: struct {
			Name string `json:"name,"`
			Size int    `json:"size,"`
		}{
			Name: "bandwidth-ABCD",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := bandwidths.Get(client.ServiceClient(), tenantID, "3c43e46e-4af1-45b8-a84d-ee6d04488d2a").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := bandwidths.List(client.ServiceClient(), tenantID, bandwidths.ListOpts{
		Limit: 2,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}
