package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/securitygroups"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := securitygroups.Create(client.ServiceClient(), securitygroups.CreateOpts{
		SecurityGroup: securitygroups.CreateSecurityGroup{
			Name:        "EricSG",
			Description: "Test SecurityGroup",
			TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := securitygroups.Update(client.ServiceClient(), "a988649d-cfbf-4c2a-bd91-3b84d2403747", securitygroups.UpdateOpts{
		SecurityGroup: securitygroups.UpdateSecurityGroup{
			Name:        "EricSG",
			Description: "Modified SecurityGroup",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := securitygroups.Get(client.ServiceClient(), "a988649d-cfbf-4c2a-bd91-3b84d2403747").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, client.ServiceClient().Endpoint)

	allPage, err := securitygroups.List(client.ServiceClient(), securitygroups.ListOpts{
		Limit: 2,
	}).AllPages()

	actual, err := securitygroups.ExtractList(allPage.(securitygroups.ListPage))
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	securitygroups.Delete(client.ServiceClient(), "7af80d49-0a43-462d-aed8-a1e12ac91af6")
}
