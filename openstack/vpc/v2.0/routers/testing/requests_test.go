package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/routers"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := routers.Create(client.ServiceClient(), routers.CreateOpts{
		Router: routers.CreateRouter{
			Name:         "EricTestPort",
			AdminStateUp: true,
			ExternalGatewayInfo: routers.ExternalGatewayInfo{
				NetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			},
			TenantId: "57e98940a77f4bb988a21a7d0603a626",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := routers.Update(client.ServiceClient(), "a4d322dd-821e-45ca-b1eb-3ccdbaef407d", routers.UpdateOpts{
		Router: routers.UpdateRouter{
			Name: "ModifiedRouter",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := routers.Get(client.ServiceClient(), "a4d322dd-821e-45ca-b1eb-3ccdbaef407d").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, client.ServiceClient().Endpoint)

	allPage, err := routers.List(client.ServiceClient(), routers.ListOpts{
		Limit: 2,
	}).AllPages()

	actual, err := routers.ExtractList(allPage.(routers.ListPage))
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	routers.Delete(client.ServiceClient(), "a4d322dd-821e-45ca-b1eb-3ccdbaef407d")
}
