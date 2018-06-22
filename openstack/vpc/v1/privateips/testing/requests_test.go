package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/privateips"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	result, err := privateips.Create(client.ServiceClient(), tenantID, privateips.CreateOpts{
		Privateips: []privateips.CreatePrivateIp{
			{
				SubnetId:  "5ae24488-454f-499c-86c4-c0355704005d",
				IpAddress: "192.168.0.12",
			},
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, result)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := privateips.Get(client.ServiceClient(), tenantID, "ea274524-f1cc-4078-8e67-c002be25c413").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	subnetID := "5ae24488-454f-499c-86c4-c0355704005d"
	actual, err := privateips.List(client.ServiceClient(), tenantID, subnetID, privateips.ListOpts{
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
	privateips.Delete(client.ServiceClient(), tenantID, "ea274524-f1cc-4078-8e67-c002be25c413")
}
