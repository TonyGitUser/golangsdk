package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/floatingips"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "floatingip": {
    "id": "f9cd3b9c-cef7-439c-8963-971bec12c292",
    "status": "DOWN",
    "router_id": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": null,
    "floating_ip_address": "49.4.64.51",
    "port_id": null
  }
}
`

var CreateResponse = floatingips.CreateResponse{
	Floatingip: floatingips.ListFloatingIP{
		Id:                "f9cd3b9c-cef7-439c-8963-971bec12c292",
		Status:            "DOWN",
		RouterId:          "",
		TenantId:          "57e98940a77f4bb988a21a7d0603a626",
		FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
		FixedIpAddress:    "",
		FloatingIpAddress: "49.4.64.51",
		PortId:            "",
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/floatingips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "floatingip": {
    "id": "f9cd3b9c-cef7-439c-8963-971bec12c292",
    "status": "DOWN",
    "router_id": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": null,
    "floating_ip_address": "49.4.64.51",
    "port_id": null
  }
}
`

var UpdateResponse = floatingips.UpdateResponse{
	Floatingip: floatingips.ListFloatingIP{
		Id:                "f9cd3b9c-cef7-439c-8963-971bec12c292",
		Status:            "DOWN",
		RouterId:          "",
		TenantId:          "57e98940a77f4bb988a21a7d0603a626",
		FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
		FixedIpAddress:    "",
		FloatingIpAddress: "49.4.64.51",
		PortId:            "",
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/floatingips/f9cd3b9c-cef7-439c-8963-971bec12c292", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "floatingip": {
    "id": "f9cd3b9c-cef7-439c-8963-971bec12c292",
    "status": "DOWN",
    "router_id": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": null,
    "floating_ip_address": "49.4.64.51",
    "port_id": null
  }
}
`

var GetResponse = floatingips.GetResponse{
	Floatingip: floatingips.ListFloatingIP{
		Id:                "f9cd3b9c-cef7-439c-8963-971bec12c292",
		Status:            "DOWN",
		RouterId:          "",
		TenantId:          "57e98940a77f4bb988a21a7d0603a626",
		FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
		FixedIpAddress:    "",
		FloatingIpAddress: "49.4.64.51",
		PortId:            "",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/floatingips/f9cd3b9c-cef7-439c-8963-971bec12c292", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var List1Output = `
{
  "floatingips": [{
    "id": "3faa05bd-d878-44e2-a363-f6672a9761d3",
    "status": "DOWN",
    "router_id": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": null,
    "floating_ip_address": "49.4.22.32",
    "port_id": null
  }, {
    "id": "4d60bba4-0791-4e82-8262-9bdffaeb1d14",
    "status": "DOWN",
    "router_id": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": null,
    "floating_ip_address": "49.4.4.36",
    "port_id": null
  }],
  "floatingips_links": [{
    "href": "%s/v2.0/floatingips?limit=2&marker=4d60bba4-0791-4e82-8262-9bdffaeb1d14",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/floatingips?limit=2&marker=3faa05bd-d878-44e2-a363-f6672a9761d3&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List2Output = `
{
  "floatingips": [{
    "id": "82abaa86-8518-47db-8d63-ddf152824635",
    "status": "ACTIVE",
    "router_id": "773c3c42-d315-417b-9063-87091713148c",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": "192.168.1.215",
    "floating_ip_address": "117.78.46.129",
    "port_id": "12c60538-01bc-4d11-a05f-29fc2d1aabd2"
  }, {
    "id": "84a71976-a8c2-42e0-8826-7fc27b876e42",
    "status": "DOWN",
    "router_id": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": null,
    "floating_ip_address": "49.4.68.149",
    "port_id": null
  }],
  "floatingips_links": [{
    "href": "%s/v2.0/floatingips?limit=2&marker=84a71976-a8c2-42e0-8826-7fc27b876e42",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/floatingips?limit=2&marker=82abaa86-8518-47db-8d63-ddf152824635&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List3Output = `
{
  "floatingips": [{
    "id": "f9cd3b9c-cef7-439c-8963-971bec12c292",
    "status": "DOWN",
    "router_id": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "floating_network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
    "fixed_ip_address": null,
    "floating_ip_address": "49.4.64.51",
    "port_id": null
  }],
  "floatingips_links": [{
    "href": "https://None/v2.0/floatingips?limit=2&marker=f9cd3b9c-cef7-439c-8963-971bec12c292&page_reverse=True",
    "rel": "previous"
  }]
}
`
var ListResponse = floatingips.ListResponse{
	Floatingips: []floatingips.ListFloatingIP{
		{
			Id:                "3faa05bd-d878-44e2-a363-f6672a9761d3",
			Status:            "DOWN",
			RouterId:          "",
			TenantId:          "57e98940a77f4bb988a21a7d0603a626",
			FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			FixedIpAddress:    "",
			FloatingIpAddress: "49.4.22.32",
			PortId:            "",
		},
		{
			Id:                "4d60bba4-0791-4e82-8262-9bdffaeb1d14",
			Status:            "DOWN",
			RouterId:          "",
			TenantId:          "57e98940a77f4bb988a21a7d0603a626",
			FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			FixedIpAddress:    "",
			FloatingIpAddress: "49.4.4.36",
			PortId:            "",
		},
		{
			Id:                "82abaa86-8518-47db-8d63-ddf152824635",
			Status:            "ACTIVE",
			RouterId:          "773c3c42-d315-417b-9063-87091713148c",
			TenantId:          "57e98940a77f4bb988a21a7d0603a626",
			FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			FixedIpAddress:    "192.168.1.215",
			FloatingIpAddress: "117.78.46.129",
			PortId:            "12c60538-01bc-4d11-a05f-29fc2d1aabd2",
		},
		{
			Id:                "84a71976-a8c2-42e0-8826-7fc27b876e42",
			Status:            "DOWN",
			RouterId:          "",
			TenantId:          "57e98940a77f4bb988a21a7d0603a626",
			FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			FixedIpAddress:    "",
			FloatingIpAddress: "49.4.68.149",
			PortId:            "",
		},
		{
			Id:                "f9cd3b9c-cef7-439c-8963-971bec12c292",
			Status:            "DOWN",
			RouterId:          "",
			TenantId:          "57e98940a77f4bb988a21a7d0603a626",
			FloatingNetworkId: "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			FixedIpAddress:    "",
			FloatingIpAddress: "49.4.64.51",
			PortId:            "",
		},
	},
}

func HandleListSuccessfully(t *testing.T, endPoint string) {
	th.Mux.HandleFunc("/v2.0/floatingips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		switch r.RequestURI {
		case "/v2.0/floatingips?limit=2":
			fmt.Fprintf(w, List1Output, endPoint)
		case "/v2.0/floatingips?limit=2&marker=4d60bba4-0791-4e82-8262-9bdffaeb1d14":
			fmt.Fprintf(w, List2Output, endPoint)
		case "/v2.0/floatingips?limit=2&marker=84a71976-a8c2-42e0-8826-7fc27b876e42":
			fmt.Fprintf(w, List3Output)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/floatingips/f9cd3b9c-cef7-439c-8963-971bec12c292", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
