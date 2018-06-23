package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/listeners"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := listeners.Create(client.ServiceClient(), tenantID, listeners.CreateOpts{
		Name:              "TestLS",
		Description:       "LS TESTING",
		LoadbalancerId:    "ebe982b8afe04851bbc26ad4609003bf",
		Protocol:          "HTTP",
		Port:              80,
		BackendProtocol:   "HTTP",
		BackendPort:       81,
		LbAlgorithm:       "roundrobin",
		SessionSticky:     true,
		StickySessionType: "insert",
		CookieTimeout:     1,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
	actual, err := listeners.Update(client.ServiceClient(), tenantID, listenerId, listeners.UpdateOpts{
		Name:        "ModifiedTestLS",
		Description: "LS TESTING",
		Port:        80,
		BackendPort: 81,
		LbAlgorithm: "roundrobin",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
	actual, err := listeners.Get(client.ServiceClient(), tenantID, listenerId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := listeners.List(client.ServiceClient(), tenantID, listeners.ListOpts{
		LoadbalancerId: "ebe982b8afe04851bbc26ad4609003bf",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	listenerId := "223a537e337e4f9b947b8b39eb1b1f6c"
	listeners.Delete(client.ServiceClient(), tenantID, listenerId)
}
