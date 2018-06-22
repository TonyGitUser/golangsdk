package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/ports"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := ports.Create(client.ServiceClient(), ports.CreateOpts{
		Port: ports.CreatePort{
			Name:      "EricTestPort",
			NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := ports.Update(client.ServiceClient(), "5e56a480-f337-4985-8ca4-98546cb4fdae", ports.UpdateOpts{
		Port: ports.UpdatePort{
			Name: "ModifiedPort",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := ports.Get(client.ServiceClient(), "5e56a480-f337-4985-8ca4-98546cb4fdae").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := ports.List(client.ServiceClient(), ports.ListOpts{
		Limit: 3,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	ports.Delete(client.ServiceClient(), "5e56a480-f337-4985-8ca4-98546cb4fdae")
}
