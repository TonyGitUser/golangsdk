package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/networks"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := networks.Create(client.ServiceClient(), networks.CreateOpts{
		Network: networks.CreateNetwork{
			Name:         "NetWork Test",
			Shared:       false,
			AdminStateUp: true,
			TenantId:     "57e98940a77f4bb988a21a7d0603a626",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := networks.Update(client.ServiceClient(), "1c6af92e-bd06-4350-8f51-5ec32167814f", networks.UpdateOpts{
		Network: networks.UpdateNetwork{
			Name:   "Test Net Works Updated",
			Shared: true,
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := networks.Get(client.ServiceClient(), "1c6af92e-bd06-4350-8f51-5ec32167814f").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, client.ServiceClient().Endpoint)

	allPage, err := networks.List(client.ServiceClient(), networks.ListOpts{
		Limit: 2,
	}).AllPages()

	actual, err := networks.ExtractList(allPage.(networks.ListPage))
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	networks.Delete(client.ServiceClient(), "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28")
}
