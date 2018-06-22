package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/publicips"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	result, err := publicips.Create(client.ServiceClient(), tenantID, publicips.CreateOpts{
		Publicip: struct {
			Type      string `json:"type,"`
			IpAddress string `json:"ip_address,omitempty"`
		}{
			Type: "5_bgp",
		},
		Bandwidth: struct {
			Name       string `json:"name,omitempty"`
			Size       int    `json:"size,omitempty"`
			Id         string `json:"id,omitempty"`
			ShareType  string `json:"share_type,"`
			ChargeMode string `json:"charge_mode,omitempty"`
		}{
			Name:       "bandwidth-d62f",
			Size:       1,
			ShareType:  "WHOLE",
			ChargeMode: "traffic"},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, result)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := publicips.Update(client.ServiceClient(), tenantID, "84a71976-a8c2-42e0-8826-7fc27b876e42", publicips.UpdateOpts{
		Publicip: struct {
			PortId string `json:"port_id,"`
		}{
			PortId: "",
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
	actual, err := publicips.Get(client.ServiceClient(), tenantID, "84a71976-a8c2-42e0-8826-7fc27b876e42").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := publicips.List(client.ServiceClient(), tenantID, publicips.ListOpts{
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
	publicips.Delete(client.ServiceClient(), tenantID, "7ffddb5f-6731-43d8-9476-1444aaa40bc0")
}
