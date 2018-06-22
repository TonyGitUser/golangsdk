package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/vpcs"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := vpcs.Create(client.ServiceClient(), tenantID, vpcs.CreateOpts{
		Vpc: vpcs.VPC{
			Name: "ABC",
			Cidr: "192.168.0.0/16",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := vpcs.Update(client.ServiceClient(), tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0", vpcs.UpdateOpts{
		Vpc: vpcs.VPC{
			Name: "ABC-back",
			Cidr: "192.168.0.0/24",
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
	actual, err := vpcs.Get(client.ServiceClient(), tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := vpcs.List(client.ServiceClient(), tenantID, vpcs.ListOpts{
		Limit: 2,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	vpcs.Delete(client.ServiceClient(), tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0")
}
