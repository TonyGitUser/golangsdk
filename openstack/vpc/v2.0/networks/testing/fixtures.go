package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/networks"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "network": {
    "id": "1c6af92e-bd06-4350-8f51-5ec32167814f",
    "name": "NetWork Test",
    "status": "ACTIVE",
    "shared": false,
    "subnets": [],
    "availability_zone_hints": [],
    "availability_zones": ["cn-north-1a", "cn-north-1b"],
    "admin_state_up": true,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "provider:network_type": "vxlan",
    "router:external": false
  }
}
`

var CreateResponse = networks.CreateResponse{
	Network: networks.ListNetwork{
		Id:                    "1c6af92e-bd06-4350-8f51-5ec32167814f",
		Name:                  "NetWork Test",
		Status:                "ACTIVE",
		Shared:                false,
		Subnets:               []string{},
		AvailabilityZoneHints: []string{},
		AvailabilityZones:     []string{"cn-north-1a", "cn-north-1b"},
		AdminStateUp:          true,
		TenantId:              "57e98940a77f4bb988a21a7d0603a626",
		ProvidernetworkType:   "vxlan",
		Routerexternal:        false,
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "network": {
    "id": "1c6af92e-bd06-4350-8f51-5ec32167814f",
    "name": "NetWork Test",
    "status": "ACTIVE",
    "shared": false,
    "subnets": [],
    "availability_zone_hints": [],
    "availability_zones": ["cn-north-1a", "cn-north-1b"],
    "admin_state_up": true,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "provider:network_type": "vxlan",
    "router:external": false
  }
}
`

var UpdateResponse = networks.UpdateResponse{
	Network: networks.ListNetwork{
		Id:                    "1c6af92e-bd06-4350-8f51-5ec32167814f",
		Name:                  "NetWork Test",
		Status:                "ACTIVE",
		Shared:                false,
		Subnets:               []string{},
		AvailabilityZoneHints: []string{},
		AvailabilityZones:     []string{"cn-north-1a", "cn-north-1b"},
		AdminStateUp:          true,
		TenantId:              "57e98940a77f4bb988a21a7d0603a626",
		ProvidernetworkType:   "vxlan",
		Routerexternal:        false,
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/networks/1c6af92e-bd06-4350-8f51-5ec32167814f", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "network": {
    "id": "1c6af92e-bd06-4350-8f51-5ec32167814f",
    "name": "NetWork Test",
    "status": "ACTIVE",
    "shared": false,
    "subnets": [],
    "availability_zone_hints": [],
    "availability_zones": ["cn-north-1a", "cn-north-1b"],
    "admin_state_up": true,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "provider:network_type": "vxlan",
    "router:external": false
  }
}
`

var GetResponse = networks.GetResponse{
	Network: networks.ListNetwork{
		Id:                    "1c6af92e-bd06-4350-8f51-5ec32167814f",
		Name:                  "NetWork Test",
		Status:                "ACTIVE",
		Shared:                false,
		Subnets:               []string{},
		AvailabilityZoneHints: []string{},
		AvailabilityZones:     []string{"cn-north-1a", "cn-north-1b"},
		AdminStateUp:          true,
		TenantId:              "57e98940a77f4bb988a21a7d0603a626",
		ProvidernetworkType:   "vxlan",
		Routerexternal:        false,
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/networks/1c6af92e-bd06-4350-8f51-5ec32167814f", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var List1Output = `
{
  "networks": [{
    "id": "1c6af92e-bd06-4350-8f51-5ec32167814f",
    "name": "NetWork Test",
    "status": "ACTIVE",
    "shared": false,
    "subnets": [],
    "availability_zone_hints": [],
    "availability_zones": ["cn-north-1a", "cn-north-1b"],
    "admin_state_up": true,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "provider:network_type": "vxlan",
    "router:external": false
  }, {
    "id": "5ae24488-454f-499c-86c4-c0355704005d",
    "name": "e23caa95-2599-4aa8-a2db-be3444450e78",
    "status": "ACTIVE",
    "shared": false,
    "subnets": ["7b4b101c-d5e2-4c52-9c6d-c6c7e1219919"],
    "availability_zone_hints": [],
    "availability_zones": ["cn-north-1a", "cn-north-1b"],
    "admin_state_up": true,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "provider:network_type": "vxlan",
    "router:external": false
  }],
  "networks_links": [{
    "href": "%s/v2.0/networks?limit=2&marker=5ae24488-454f-499c-86c4-c0355704005d",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/",
    "rel": "previous"
  }]
}
`

var List2Output = `
{
  "networks": [{
    "id": "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
    "name": "ea3b0efe-0d6a-4288-8b16-753504a994b9",
    "status": "ACTIVE",
    "shared": false,
    "subnets": ["4d1d4f45-6375-4821-822e-44c92d12a58c"],
    "availability_zone_hints": [],
    "availability_zones": ["cn-north-1a", "cn-north-1b"],
    "admin_state_up": true,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "provider:network_type": "vxlan",
    "router:external": false
  }, {
    "id": "f6a0db7b-397c-4162-bc35-9a1f008b4373",
    "name": "773c3c42-d315-417b-9063-87091713148c",
    "status": "ACTIVE",
    "shared": false,
    "subnets": ["c6e37ca7-b711-4969-b5c2-388acac38948"],
    "availability_zone_hints": [],
    "availability_zones": ["cn-north-1a", "cn-north-1b"],
    "admin_state_up": true,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "provider:network_type": "vxlan",
    "router:external": false
  }],
  "networks_links": [{
    "href": "%s/v2.0/networks?limit=2&marker=f6a0db7b-397c-4162-bc35-9a1f008b4373",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/",
    "rel": "previous"
  }]
}
`

var List3Output = `
{
  "networks": [],
  "networks_links": [{
    "href": "https://None/v2.0/",
    "rel": "previous"
  }]
}
`
var ListResponse = networks.ListResponse{
	Networks: []networks.ListNetwork{
		{
			Id:                    "1c6af92e-bd06-4350-8f51-5ec32167814f",
			Name:                  "NetWork Test",
			Status:                "ACTIVE",
			Shared:                false,
			Subnets:               []string{},
			AvailabilityZoneHints: []string{},
			AvailabilityZones:     []string{"cn-north-1a", "cn-north-1b"},
			AdminStateUp:          true,
			TenantId:              "57e98940a77f4bb988a21a7d0603a626",
			ProvidernetworkType:   "vxlan",
			Routerexternal:        false,
		},
		{
			Id:                    "5ae24488-454f-499c-86c4-c0355704005d",
			Name:                  "e23caa95-2599-4aa8-a2db-be3444450e78",
			Status:                "ACTIVE",
			Shared:                false,
			Subnets:               []string{"7b4b101c-d5e2-4c52-9c6d-c6c7e1219919"},
			AvailabilityZoneHints: []string{},
			AvailabilityZones:     []string{"cn-north-1a", "cn-north-1b"},
			AdminStateUp:          true,
			TenantId:              "57e98940a77f4bb988a21a7d0603a626",
			ProvidernetworkType:   "vxlan",
			Routerexternal:        false,
		},
		{
			Id:                    "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
			Name:                  "ea3b0efe-0d6a-4288-8b16-753504a994b9",
			Status:                "ACTIVE",
			Shared:                false,
			Subnets:               []string{"4d1d4f45-6375-4821-822e-44c92d12a58c"},
			AvailabilityZoneHints: []string{},
			AvailabilityZones:     []string{"cn-north-1a", "cn-north-1b"},
			AdminStateUp:          true,
			TenantId:              "57e98940a77f4bb988a21a7d0603a626",
			ProvidernetworkType:   "vxlan",
			Routerexternal:        false,
		},
		{
			Id:                    "f6a0db7b-397c-4162-bc35-9a1f008b4373",
			Name:                  "773c3c42-d315-417b-9063-87091713148c",
			Status:                "ACTIVE",
			Shared:                false,
			Subnets:               []string{"c6e37ca7-b711-4969-b5c2-388acac38948"},
			AvailabilityZoneHints: []string{},
			AvailabilityZones:     []string{"cn-north-1a", "cn-north-1b"},
			AdminStateUp:          true,
			TenantId:              "57e98940a77f4bb988a21a7d0603a626",
			ProvidernetworkType:   "vxlan",
			Routerexternal:        false,
		},
	},
}

func HandleListSuccessfully(t *testing.T, endPoint string) {
	th.Mux.HandleFunc("/v2.0/networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		switch r.RequestURI {
		case "/v2.0/networks?limit=2":
			fmt.Fprintf(w, List1Output, endPoint)
		case "/v2.0/networks?limit=2&marker=5ae24488-454f-499c-86c4-c0355704005d":
			fmt.Fprintf(w, List2Output, endPoint)
		case "/v2.0/networks?limit=2&marker=f6a0db7b-397c-4162-bc35-9a1f008b4373":
			fmt.Fprintf(w, List3Output)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/networks/1c6af92e-bd06-4350-8f51-5ec32167814f", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
