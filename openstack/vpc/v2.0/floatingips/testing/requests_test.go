package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/floatingips"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := floatingips.Create(client.ServiceClient(), floatingips.CreateOpts{
		Floatingip: floatingips.CreateFloatingIP{
			FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			TenantId:          "57e98940a77f4bb988a21a7d0603a626",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := floatingips.Update(client.ServiceClient(), "f9cd3b9c-cef7-439c-8963-971bec12c292", floatingips.UpdateOpts{
		Floatingip: floatingips.UpdateFloatingIP{
			FixedIpAddress: "49.4.64.52",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := floatingips.Get(client.ServiceClient(), "f9cd3b9c-cef7-439c-8963-971bec12c292").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, client.ServiceClient().Endpoint)

	allPage, err := floatingips.List(client.ServiceClient(), floatingips.ListOpts{
		Limit: 2,
	}).AllPages()

	actual, err := floatingips.ExtractList(allPage.(floatingips.ListPage))
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	floatingips.Delete(client.ServiceClient(), "f9cd3b9c-cef7-439c-8963-971bec12c292")
}
