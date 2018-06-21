package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/ports"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "port": {
    "id": "d6d0258f-7bf2-4504-845b-ab6b8513a986",
    "name": "EricTestPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.35"
    }],
    "mac_address": "fa:16:3e:69:6d:8f",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }
}
`

var CreateResponse = ports.CreateResponse{
	Port: ports.ListPort{
		Id:           "d6d0258f-7bf2-4504-845b-ab6b8513a986",
		Name:         "EricTestPort",
		Status:       "DOWN",
		AdminStateUp: true,
		FixedIps: []ports.FixedIp{
			{
				SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
				IpAddress: "192.168.0.35",
			},
		},
		MacAddress:  "fa:16:3e:69:6d:8f",
		NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		DeviceId:    "",
		DeviceOwner: "",
		SecurityGroups: []string{
			"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
		},
		ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
		AllowedAddressPairs: []ports.AllowedAddressPair{},
		BindingvnicType:     "normal",
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "port": {
    "id": "d6d0258f-7bf2-4504-845b-ab6b8513a986",
    "name": "ModifiedPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.35"
    }],
    "mac_address": "fa:16:3e:69:6d:8f",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }
}
`

var UpdateResponse = ports.UpdateResponse{
	Port: ports.ListPort{
		Id:           "d6d0258f-7bf2-4504-845b-ab6b8513a986",
		Name:         "ModifiedPort",
		Status:       "DOWN",
		AdminStateUp: true,
		FixedIps: []ports.FixedIp{
			{
				SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
				IpAddress: "192.168.0.35",
			},
		},
		MacAddress:  "fa:16:3e:69:6d:8f",
		NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		DeviceId:    "",
		DeviceOwner: "",
		SecurityGroups: []string{
			"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
		},
		ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
		AllowedAddressPairs: []ports.AllowedAddressPair{},
		BindingvnicType:     "normal",
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/ports/d6d0258f-7bf2-4504-845b-ab6b8513a986", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "port": {
    "id": "d6d0258f-7bf2-4504-845b-ab6b8513a986",
    "name": "EricTestPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.35"
    }],
    "mac_address": "fa:16:3e:69:6d:8f",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }
}
`

var GetResponse = ports.GetResponse{
	Port: ports.ListPort{
		Id:           "d6d0258f-7bf2-4504-845b-ab6b8513a986",
		Name:         "EricTestPort",
		Status:       "DOWN",
		AdminStateUp: true,
		FixedIps: []ports.FixedIp{
			{
				SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
				IpAddress: "192.168.0.35",
			},
		},
		MacAddress:  "fa:16:3e:69:6d:8f",
		NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		DeviceId:    "",
		DeviceOwner: "",
		SecurityGroups: []string{
			"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
		},
		ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
		AllowedAddressPairs: []ports.AllowedAddressPair{},
		BindingvnicType:     "normal",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/ports/d6d0258f-7bf2-4504-845b-ab6b8513a986", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var List1Output = `
{
  "ports": [{
    "id": "52550ad3-a04f-420d-965d-00c30c040160",
    "name": "EricTestPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.117"
    }],
    "mac_address": "fa:16:3e:e4:32:fd",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }, {
    "id": "b02c7697-8659-4144-ba50-ac9d7a7e8adc",
    "name": "EricTestPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.148"
    }],
    "mac_address": "fa:16:3e:22:cb:c3",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }],
  "ports_links": [{
    "href": "%s/v2.0/ports?name=EricTestPort&limit=2&marker=b02c7697-8659-4144-ba50-ac9d7a7e8adc",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/ports?name=EricTestPort&limit=2&marker=52550ad3-a04f-420d-965d-00c30c040160&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List2Output = `
{
  "ports": [{
    "id": "d6d0258f-7bf2-4504-845b-ab6b8513a986",
    "name": "EricTestPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.35"
    }],
    "mac_address": "fa:16:3e:69:6d:8f",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }],
  "ports_links": [{
    "href": "https://None/v2.0/ports?name=EricTestPort&limit=2&marker=d6d0258f-7bf2-4504-845b-ab6b8513a986&page_reverse=True",
    "rel": "previous"
  }]
}
`

var ListResponse = ports.ListResponse{
	Ports: []ports.ListPort{
		{
			Id:           "52550ad3-a04f-420d-965d-00c30c040160",
			Name:         "EricTestPort",
			Status:       "DOWN",
			AdminStateUp: true,
			FixedIps: []ports.FixedIp{
				{
					SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
					IpAddress: "192.168.0.117",
				},
			},
			MacAddress:  "fa:16:3e:e4:32:fd",
			NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
			TenantId:    "57e98940a77f4bb988a21a7d0603a626",
			DeviceId:    "",
			DeviceOwner: "",
			SecurityGroups: []string{
				"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			},
			ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
			AllowedAddressPairs: []ports.AllowedAddressPair{},
			BindingvnicType:     "normal",
		},
		{
			Id:           "b02c7697-8659-4144-ba50-ac9d7a7e8adc",
			Name:         "EricTestPort",
			Status:       "DOWN",
			AdminStateUp: true,
			FixedIps: []ports.FixedIp{
				{
					SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
					IpAddress: "192.168.0.148",
				},
			},
			MacAddress:  "fa:16:3e:22:cb:c3",
			NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
			TenantId:    "57e98940a77f4bb988a21a7d0603a626",
			DeviceId:    "",
			DeviceOwner: "",
			SecurityGroups: []string{
				"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			},
			ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
			AllowedAddressPairs: []ports.AllowedAddressPair{},
			BindingvnicType:     "normal",
		},
		{
			Id:           "d6d0258f-7bf2-4504-845b-ab6b8513a986",
			Name:         "EricTestPort",
			Status:       "DOWN",
			AdminStateUp: true,
			FixedIps: []ports.FixedIp{
				{
					SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
					IpAddress: "192.168.0.35",
				},
			},
			MacAddress:  "fa:16:3e:69:6d:8f",
			NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
			TenantId:    "57e98940a77f4bb988a21a7d0603a626",
			DeviceId:    "",
			DeviceOwner: "",
			SecurityGroups: []string{
				"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			},
			ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
			AllowedAddressPairs: []ports.AllowedAddressPair{},
			BindingvnicType:     "normal",
		},
	},
}

func HandleListSuccessfully(t *testing.T, endPoint string) {
	th.Mux.HandleFunc("/v2.0/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		switch r.RequestURI {
		case "/v2.0/ports?limit=2&name=EricTestPort":
			fmt.Fprintf(w, List1Output, endPoint)
		case "/v2.0/ports?name=EricTestPort&limit=2&marker=b02c7697-8659-4144-ba50-ac9d7a7e8adc":
			fmt.Fprintf(w, List2Output)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/ports/d6d0258f-7bf2-4504-845b-ab6b8513a986", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
