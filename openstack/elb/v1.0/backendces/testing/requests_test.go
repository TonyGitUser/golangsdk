package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/backendces"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	listenerId := "d5f3197c6bd8491ca1dfc905350d85ea"
	actual, err := backendces.Create(client.ServiceClient(), tenantId, listenerId, backendces.CreateOpts{
		Items: []struct {
			ServerId string `json:"server_id,omitempty"`
			Address  string `json:"address,omitempty"`
		}{
			{
				ServerId: "e1040c67-b130-41ee-9405-07c6c6c20208",
				Address:  "192.168.1.215",
			},
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	listenerId := "d5f3197c6bd8491ca1dfc905350d85ea"
	actual, err := backendces.List(client.ServiceClient(), tenantId, listenerId, backendces.ListOpts{
		Limit: "1",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	listenerId := "ca082827b61d4902bfaf32e8174e244a"
	backendces.Delete(client.ServiceClient(), tenantId, listenerId, backendces.DeleteOpts{
		RemoveMember: []struct {
			Id string `json:"id,omitempty"`
		}{
			{
				Id: "e1040c67-b130-41ee-9405-07c6c6c20208",
			},
		},
	})
}
