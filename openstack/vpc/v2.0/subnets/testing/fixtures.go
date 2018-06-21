package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/subnets"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "subnet": {
    "name": "subnetV2",
    "cidr": "192.168.20.0/24",
    "id": "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
    "enable_dhcp": true,
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "dns_nameservers": ["114.114.114.114", "114.114.115.115"],
    "allocation_pools": [{
      "start": "192.168.20.2",
      "end": "192.168.20.254"
    }],
    "host_routes": [],
    "ip_version": 4,
    "gateway_ip": "192.168.20.1"
  }
}
`

var CreateResponse = subnets.CreateResponse{
	Subnet: subnets.ListSubnet{
		Name:           "subnetV2",
		Cidr:           "192.168.20.0/24",
		Id:             "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
		EnableDhcp:     true,
		NetworkId:      "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:       "57e98940a77f4bb988a21a7d0603a626",
		DnsNameservers: []string{"114.114.114.114", "114.114.115.115"},
		AllocationPools: []subnets.AllocationPool{
			{
				Start: "192.168.20.2",
				End:   "192.168.20.254",
			},
		},
		HostRoutes: []subnets.HostRoute{},
		IpVersion:  4,
		GatewayIp:  "192.168.20.1",
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "subnet": {
    "name": "subnetV2",
    "cidr": "192.168.20.0/24",
    "id": "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
    "enable_dhcp": true,
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "dns_nameservers": ["8.8.8.8", "4.4.4.4"],
    "allocation_pools": [{
      "start": "192.168.20.2",
      "end": "192.168.20.254"
    }],
    "host_routes": [],
    "ip_version": 4,
    "gateway_ip": "192.168.20.1"
  }
}
`

var UpdateResponse = subnets.UpdateResponse{
	Subnet: subnets.ListSubnet{
		Name:           "subnetV2",
		Cidr:           "192.168.20.0/24",
		Id:             "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
		EnableDhcp:     true,
		NetworkId:      "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:       "57e98940a77f4bb988a21a7d0603a626",
		DnsNameservers: []string{"8.8.8.8", "4.4.4.4"},
		AllocationPools: []subnets.AllocationPool{
			{
				Start: "192.168.20.2",
				End:   "192.168.20.254",
			},
		},
		HostRoutes: []subnets.HostRoute{},
		IpVersion:  4,
		GatewayIp:  "192.168.20.1",
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/subnets/ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "subnet": {
    "name": "subnetV2",
    "cidr": "192.168.20.0/24",
    "id": "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
    "enable_dhcp": true,
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "dns_nameservers": ["114.114.114.114", "114.114.115.115"],
    "allocation_pools": [{
      "start": "192.168.20.2",
      "end": "192.168.20.254"
    }],
    "host_routes": [],
    "ip_version": 4,
    "gateway_ip": "192.168.20.1"
  }
}
`

var GetResponse = subnets.GetResponse{
	Subnet: subnets.ListSubnet{
		Name:           "subnetV2",
		Cidr:           "192.168.20.0/24",
		Id:             "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
		EnableDhcp:     true,
		NetworkId:      "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:       "57e98940a77f4bb988a21a7d0603a626",
		DnsNameservers: []string{"114.114.114.114", "114.114.115.115"},
		AllocationPools: []subnets.AllocationPool{
			{
				Start: "192.168.20.2",
				End:   "192.168.20.254",
			},
		},
		HostRoutes: []subnets.HostRoute{},
		IpVersion:  4,
		GatewayIp:  "192.168.20.1",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/subnets/ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var List1Output = `
{
  "subnets": [{
    "name": "subnet-3d95",
    "cidr": "192.168.0.0/24",
    "id": "4d1d4f45-6375-4821-822e-44c92d12a58c",
    "enable_dhcp": true,
    "network_id": "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "dns_nameservers": ["100.125.1.250", "114.114.114.114"],
    "allocation_pools": [{
      "start": "192.168.0.2",
      "end": "192.168.0.251"
    }],
    "host_routes": [],
    "ip_version": 4,
    "gateway_ip": "192.168.0.1"
  }, {
    "name": "EricTestSubNet",
    "cidr": "192.168.0.0/24",
    "id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
    "enable_dhcp": true,
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "dns_nameservers": ["100.125.1.250", "114.114.114.114"],
    "allocation_pools": [{
      "start": "192.168.0.2",
      "end": "192.168.0.251"
    }],
    "host_routes": [],
    "ip_version": 4,
    "gateway_ip": "192.168.0.1"
  }],
  "subnets_links": [{
    "href": "%s/v2.0/subnets?limit=2&marker=7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/subnets?limit=2&marker=4d1d4f45-6375-4821-822e-44c92d12a58c&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List2Output = `
{
  "subnets": [{
    "name": "subnet-c8cb",
    "cidr": "192.168.1.0/24",
    "id": "c6e37ca7-b711-4969-b5c2-388acac38948",
    "enable_dhcp": true,
    "network_id": "f6a0db7b-397c-4162-bc35-9a1f008b4373",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "dns_nameservers": ["114.114.114.114", "114.114.115.115"],
    "allocation_pools": [{
      "start": "192.168.1.2",
      "end": "192.168.1.251"
    }],
    "host_routes": [],
    "ip_version": 4,
    "gateway_ip": "192.168.1.1"
  }, {
    "name": "subnetV2",
    "cidr": "192.168.20.0/24",
    "id": "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
    "enable_dhcp": true,
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "dns_nameservers": ["8.8.8.8", "4.4.4.4"],
    "allocation_pools": [{
      "start": "192.168.20.2",
      "end": "192.168.20.254"
    }],
    "host_routes": [],
    "ip_version": 4,
    "gateway_ip": "192.168.20.1"
  }],
  "subnets_links": [{
    "href": "%s/v2.0/subnets?limit=2&marker=ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/subnets?limit=2&marker=c6e37ca7-b711-4969-b5c2-388acac38948&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List3Output = `
{"subnets":[],"subnets_links":[]}
`
var ListResponse = subnets.ListResponse{
	Subnets: []subnets.ListSubnet{
		{
			Name:           "subnet-3d95",
			Cidr:           "192.168.0.0/24",
			Id:             "4d1d4f45-6375-4821-822e-44c92d12a58c",
			EnableDhcp:     true,
			NetworkId:      "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
			TenantId:       "57e98940a77f4bb988a21a7d0603a626",
			DnsNameservers: []string{"100.125.1.250", "114.114.114.114"},
			AllocationPools: []subnets.AllocationPool{
				{
					Start: "192.168.0.2",
					End:   "192.168.0.251",
				},
			},
			HostRoutes: []subnets.HostRoute{},
			IpVersion:  4,
			GatewayIp:  "192.168.0.1",
		},
		{
			Name:           "EricTestSubNet",
			Cidr:           "192.168.0.0/24",
			Id:             "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
			EnableDhcp:     true,
			NetworkId:      "5ae24488-454f-499c-86c4-c0355704005d",
			TenantId:       "57e98940a77f4bb988a21a7d0603a626",
			DnsNameservers: []string{"100.125.1.250", "114.114.114.114"},
			AllocationPools: []subnets.AllocationPool{
				{
					Start: "192.168.0.2",
					End:   "192.168.0.251",
				},
			},
			HostRoutes: []subnets.HostRoute{},
			IpVersion:  4,
			GatewayIp:  "192.168.0.1",
		},
		{
			Name:           "subnet-c8cb",
			Cidr:           "192.168.1.0/24",
			Id:             "c6e37ca7-b711-4969-b5c2-388acac38948",
			EnableDhcp:     true,
			NetworkId:      "f6a0db7b-397c-4162-bc35-9a1f008b4373",
			TenantId:       "57e98940a77f4bb988a21a7d0603a626",
			DnsNameservers: []string{"114.114.114.114", "114.114.115.115"},
			AllocationPools: []subnets.AllocationPool{
				{
					Start: "192.168.1.2",
					End:   "192.168.1.251",
				},
			},
			HostRoutes: []subnets.HostRoute{},
			IpVersion:  4,
			GatewayIp:  "192.168.1.1",
		},
		{
			Name:           "subnetV2",
			Cidr:           "192.168.20.0/24",
			Id:             "ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28",
			EnableDhcp:     true,
			NetworkId:      "5ae24488-454f-499c-86c4-c0355704005d",
			TenantId:       "57e98940a77f4bb988a21a7d0603a626",
			DnsNameservers: []string{"8.8.8.8", "4.4.4.4"},
			AllocationPools: []subnets.AllocationPool{
				{
					Start: "192.168.20.2",
					End:   "192.168.20.254",
				},
			},
			HostRoutes: []subnets.HostRoute{},
			IpVersion:  4,
			GatewayIp:  "192.168.20.1",
		},
	},
}

func HandleListSuccessfully(t *testing.T, endPoint string) {
	th.Mux.HandleFunc("/v2.0/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		switch r.RequestURI {
		case "/v2.0/subnets?limit=2":
			fmt.Fprintf(w, List1Output, endPoint)
		case "/v2.0/subnets?limit=2&marker=7b4b101c-d5e2-4c52-9c6d-c6c7e1219919":
			fmt.Fprintf(w, List2Output, endPoint)
		case "/v2.0/subnets?limit=2&marker=ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28":
			fmt.Fprintf(w, List3Output)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/subnets/ffa56d29-0ad8-43b4-b8e8-dc74b6ef0c28 ", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
