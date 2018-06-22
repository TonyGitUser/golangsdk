package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/subnets"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := subnets.Create(client.ServiceClient(), tenantID, subnets.CreateOpts{
		Subnet: subnets.Subnet{
			Name:         "subnet",
			Cidr:         "192.168.20.0/24",
			GatewayIp:    "192.168.20.1",
			DhcpEnable:   true,
			PrimaryDns:   "114.114.114.114",
			SecondaryDns: "114.114.115.115",
			DnsList: []string{
				"114.114.114.114",
				"114.114.115.115",
			},
			AvailabilityZone: "cn-north-1a",
			VpcId:            "ea3b0efe-0d6a-4288-8b16-753504a994b9",
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
	actual, err := subnets.Update(client.ServiceClient(), tenantID, "ea3b0efe-0d6a-4288-8b16-753504a994b9", "c9aba52d-ec14-40cb-930f-c8153e93c2db", subnets.UpdateOpts{
		Subnet: subnets.Subnet{
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
	actual, err := subnets.Get(client.ServiceClient(), tenantID, "c9aba52d-ec14-40cb-930f-c8153e93c2db").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := subnets.List(client.ServiceClient(), tenantID, subnets.ListOpts{
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
	subnets.Delete(client.ServiceClient(), tenantID, "ea3b0efe-0d6a-4288-8b16-753504a994b9", "c9aba52d-ec14-40cb-930f-c8153e93c2db")
}
