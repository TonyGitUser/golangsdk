package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/subnets"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := subnets.Create(client.ServiceClient(), subnets.CreateOpts{
		Subnet: subnets.CreateSubnet{
			Name:      "subnetV2",
			IpVersion: 4,
			NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
			Cidr:      "192.168.20.0/24",
			GatewayIp: "192.168.20.1",
			AllocationPools: []subnets.AllocationPool{
				{
					Start: "192.168.20.2",
					End:   "192.168.20.254",
				},
			},
			EnableDhcp:     true,
			DnsNameservers: []string{"114.114.114.114", "114.114.115.115"},
			TenantId:       "57e98940a77f4bb988a21a7d0603a626",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := subnets.Update(client.ServiceClient(), "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28", subnets.UpdateOpts{
		Subnet: subnets.UpdateSubnet{
			Name:           "subnetV2",
			GatewayIp:      "192.168.20.1",
			EnableDhcp:     true,
			DnsNameservers: []string{"8.8.8.8", "4.4.4.4"},
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := subnets.Get(client.ServiceClient(), "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, client.ServiceClient().Endpoint)

	allPage, err := subnets.List(client.ServiceClient(), subnets.ListOpts{
		Limit: 2,
	}).AllPages()

	actual, err := subnets.ExtractList(allPage.(subnets.ListPage))
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	subnets.Delete(client.ServiceClient(), "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28")
}
