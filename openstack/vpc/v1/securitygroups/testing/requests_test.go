package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/securitygroups"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := securitygroups.Create(client.ServiceClient(), tenantID, securitygroups.CreateOpts{
		SecurityGroup: securitygroups.CreateSecurityGroup{
			Name:        "EricSG",
			Description: "Test SecurityGroup",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := securitygroups.Get(client.ServiceClient(), tenantID, "f7616338-fa30-42b8-bf6b-754c0701aab8").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := securitygroups.List(client.ServiceClient(), tenantID, securitygroups.ListOpts{
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
	securitygroups.Delete(client.ServiceClient(), tenantID, "2465d913-1084-4a6a-91e7-2fd6f490ecb3")
}
