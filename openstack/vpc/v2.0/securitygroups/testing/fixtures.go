package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/securitygroups"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "security_group": {
    "id": "a988649d-cfbf-4c2a-bd91-3b84d2403747",
    "name": "EricSG",
    "description": "Test SecurityGroup",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "security_group_rules": [{
      "id": "0650c837-2ee0-478f-a184-d7ebc6cad353",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv4",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "a988649d-cfbf-4c2a-bd91-3b84d2403747"
    }, {
      "id": "d4866ba2-c085-4683-9dee-5932eb771106",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv6",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "a988649d-cfbf-4c2a-bd91-3b84d2403747"
    }]
  }
}
`

var CreateResponse = securitygroups.CreateResponse{
	SecurityGroup: securitygroups.ListSecurityGroup{
		Id:          "a988649d-cfbf-4c2a-bd91-3b84d2403747",
		Name:        "EricSG",
		Description: "Test SecurityGroup",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		SecurityGroupRules: []securitygroups.ListSecurityGroupRule{
			{
				Id:              "0650c837-2ee0-478f-a184-d7ebc6cad353",
				Direction:       "egress",
				Protocol:        "",
				Ethertype:       "IPv4",
				Description:     "",
				RemoteGroupId:   "",
				RemoteIpPrefix:  "",
				TenantId:        "57e98940a77f4bb988a21a7d0603a626",
				PortRangeMax:    0,
				PortRangeMin:    0,
				SecurityGroupId: "a988649d-cfbf-4c2a-bd91-3b84d2403747",
			},
			{
				Id:              "d4866ba2-c085-4683-9dee-5932eb771106",
				Direction:       "egress",
				Protocol:        "",
				Ethertype:       "IPv6",
				Description:     "",
				RemoteGroupId:   "",
				RemoteIpPrefix:  "",
				TenantId:        "57e98940a77f4bb988a21a7d0603a626",
				PortRangeMax:    0,
				PortRangeMin:    0,
				SecurityGroupId: "a988649d-cfbf-4c2a-bd91-3b84d2403747",
			},
		},
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "security_group": {
    "id": "7af80d49-0a43-462d-aed8-a1e12ac91af6",
    "name": "EricSG",
    "description": "Modified SecurityGroup",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "security_group_rules": [{
      "id": "11dad133-7a0d-414e-83e5-00b8836254da",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv6",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7af80d49-0a43-462d-aed8-a1e12ac91af6"
    }, {
      "id": "a6c51263-3134-4a47-83ba-bc434abb388c",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv4",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7af80d49-0a43-462d-aed8-a1e12ac91af6"
    }]
  }
}
`
var UpdateResponse = securitygroups.UpdateResponse{
	SecurityGroup: securitygroups.ListSecurityGroup{
		Id:          "7af80d49-0a43-462d-aed8-a1e12ac91af6",
		Name:        "EricSG",
		Description: "Modified SecurityGroup",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		SecurityGroupRules: []securitygroups.ListSecurityGroupRule{
			{
				Id:              "11dad133-7a0d-414e-83e5-00b8836254da",
				Direction:       "egress",
				Protocol:        "",
				Ethertype:       "IPv6",
				Description:     "",
				RemoteGroupId:   "",
				RemoteIpPrefix:  "",
				TenantId:        "57e98940a77f4bb988a21a7d0603a626",
				PortRangeMax:    0,
				PortRangeMin:    0,
				SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
			}, {
				Id:              "a6c51263-3134-4a47-83ba-bc434abb388c",
				Direction:       "egress",
				Protocol:        "",
				Ethertype:       "IPv4",
				Description:     "",
				RemoteGroupId:   "",
				RemoteIpPrefix:  "",
				TenantId:        "57e98940a77f4bb988a21a7d0603a626",
				PortRangeMax:    0,
				PortRangeMin:    0,
				SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
			},
		},
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/security-groups/a988649d-cfbf-4c2a-bd91-3b84d2403747", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "security_group": {
    "id": "a988649d-cfbf-4c2a-bd91-3b84d2403747",
    "name": "EricSG",
    "description": "Test SecurityGroup",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "security_group_rules": [{
      "id": "0650c837-2ee0-478f-a184-d7ebc6cad353",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv4",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "a988649d-cfbf-4c2a-bd91-3b84d2403747"
    }, {
      "id": "d4866ba2-c085-4683-9dee-5932eb771106",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv6",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "a988649d-cfbf-4c2a-bd91-3b84d2403747"
    }]
  }
}
`

var GetResponse = securitygroups.GetResponse{
	SecurityGroup: securitygroups.ListSecurityGroup{
		Id:          "a988649d-cfbf-4c2a-bd91-3b84d2403747",
		Name:        "EricSG",
		Description: "Test SecurityGroup",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		SecurityGroupRules: []securitygroups.ListSecurityGroupRule{
			{
				Id:              "0650c837-2ee0-478f-a184-d7ebc6cad353",
				Direction:       "egress",
				Protocol:        "",
				Ethertype:       "IPv4",
				Description:     "",
				RemoteGroupId:   "",
				RemoteIpPrefix:  "",
				TenantId:        "57e98940a77f4bb988a21a7d0603a626",
				PortRangeMax:    0,
				PortRangeMin:    0,
				SecurityGroupId: "a988649d-cfbf-4c2a-bd91-3b84d2403747",
			},
			{
				Id:              "d4866ba2-c085-4683-9dee-5932eb771106",
				Direction:       "egress",
				Protocol:        "",
				Ethertype:       "IPv6",
				Description:     "",
				RemoteGroupId:   "",
				RemoteIpPrefix:  "",
				TenantId:        "57e98940a77f4bb988a21a7d0603a626",
				PortRangeMax:    0,
				PortRangeMin:    0,
				SecurityGroupId: "a988649d-cfbf-4c2a-bd91-3b84d2403747",
			},
		},
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/security-groups/a988649d-cfbf-4c2a-bd91-3b84d2403747", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var List1Output = `
{
  "security_groups": [{
    "id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
    "name": "default",
    "description": "default",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "security_group_rules": [{
      "id": "2aafd879-c90f-49e0-87ee-0e4f0721c40c",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv4",
      "description": null,
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "id": "bf04230c-208a-4864-b0ff-876c21b33d0c",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv6",
      "description": null,
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "id": "ee07e0d4-b4ec-4c68-a754-ff01652d4e11",
      "direction": "ingress",
      "protocol": null,
      "ethertype": "IPv6",
      "description": null,
      "remote_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "id": "164bcf26-c9c0-417a-abc5-c129ad0c1854",
      "direction": "ingress",
      "protocol": "tcp",
      "ethertype": "IPv4",
      "description": null,
      "remote_group_id": null,
      "remote_ip_prefix": "0.0.0.0/0",
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": 3389,
      "port_range_min": 3389,
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "id": "ff0eb13a-9ea7-432c-9032-9316a298d43a",
      "direction": "ingress",
      "protocol": "tcp",
      "ethertype": "IPv4",
      "description": null,
      "remote_group_id": null,
      "remote_ip_prefix": "0.0.0.0/0",
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": 22,
      "port_range_min": 22,
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "id": "d534b6e7-d001-4c36-b34d-ae70aa48f988",
      "direction": "ingress",
      "protocol": null,
      "ethertype": "IPv4",
      "description": null,
      "remote_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }]
  }, {
    "id": "7af80d49-0a43-462d-aed8-a1e12ac91af6",
    "name": "EricSG",
    "description": "Modified SecurityGroup",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "security_group_rules": [{
      "id": "11dad133-7a0d-414e-83e5-00b8836254da",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv6",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7af80d49-0a43-462d-aed8-a1e12ac91af6"
    }, {
      "id": "a6c51263-3134-4a47-83ba-bc434abb388c",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv4",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "7af80d49-0a43-462d-aed8-a1e12ac91af6"
    }]
  }],
  "security_groups_links": [{
    "href": "%s/v2.0/security-groups?limit=2&marker=7af80d49-0a43-462d-aed8-a1e12ac91af6",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/security-groups?limit=2&marker=7844d4b4-d78f-45dc-9465-2b4d1bca83a5&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List2Output = `
{
  "security_groups": [{
    "id": "a988649d-cfbf-4c2a-bd91-3b84d2403747",
    "name": "EricSG",
    "description": "Test SecurityGroup",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "security_group_rules": [{
      "id": "0650c837-2ee0-478f-a184-d7ebc6cad353",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv4",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "a988649d-cfbf-4c2a-bd91-3b84d2403747"
    }, {
      "id": "d4866ba2-c085-4683-9dee-5932eb771106",
      "direction": "egress",
      "protocol": null,
      "ethertype": "IPv6",
      "description": "",
      "remote_group_id": null,
      "remote_ip_prefix": null,
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "port_range_max": null,
      "port_range_min": null,
      "security_group_id": "a988649d-cfbf-4c2a-bd91-3b84d2403747"
    }]
  }],
  "security_groups_links": [{
    "href": "https://None/v2.0/security-groups?limit=2&marker=a988649d-cfbf-4c2a-bd91-3b84d2403747&page_reverse=True",
    "rel": "previous"
  }]
}
`

var ListResponse = securitygroups.ListResponse{
	SecurityGroups: []securitygroups.ListSecurityGroup{
		{
			Id:          "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			Name:        "default",
			Description: "default",
			TenantId:    "57e98940a77f4bb988a21a7d0603a626",
			SecurityGroupRules: []securitygroups.ListSecurityGroupRule{
				{
					Id:              "2aafd879-c90f-49e0-87ee-0e4f0721c40c",
					Direction:       "egress",
					Protocol:        "",
					Ethertype:       "IPv4",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				}, {
					Id:              "bf04230c-208a-4864-b0ff-876c21b33d0c",
					Direction:       "egress",
					Protocol:        "",
					Ethertype:       "IPv6",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				}, {
					Id:              "ee07e0d4-b4ec-4c68-a754-ff01652d4e11",
					Direction:       "ingress",
					Protocol:        "",
					Ethertype:       "IPv6",
					Description:     "",
					RemoteGroupId:   "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				}, {
					Id:              "164bcf26-c9c0-417a-abc5-c129ad0c1854",
					Direction:       "ingress",
					Protocol:        "tcp",
					Ethertype:       "IPv4",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "0.0.0.0/0",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    3389,
					PortRangeMin:    3389,
					SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				}, {
					Id:              "ff0eb13a-9ea7-432c-9032-9316a298d43a",
					Direction:       "ingress",
					Protocol:        "tcp",
					Ethertype:       "IPv4",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "0.0.0.0/0",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    22,
					PortRangeMin:    22,
					SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				}, {
					Id:              "d534b6e7-d001-4c36-b34d-ae70aa48f988",
					Direction:       "ingress",
					Protocol:        "",
					Ethertype:       "IPv4",
					Description:     "",
					RemoteGroupId:   "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				},
			},
		},
		{
			Id:          "7af80d49-0a43-462d-aed8-a1e12ac91af6",
			Name:        "EricSG",
			Description: "Modified SecurityGroup",
			TenantId:    "57e98940a77f4bb988a21a7d0603a626",
			SecurityGroupRules: []securitygroups.ListSecurityGroupRule{
				{
					Id:              "11dad133-7a0d-414e-83e5-00b8836254da",
					Direction:       "egress",
					Protocol:        "",
					Ethertype:       "IPv6",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
				}, {
					Id:              "a6c51263-3134-4a47-83ba-bc434abb388c",
					Direction:       "egress",
					Protocol:        "",
					Ethertype:       "IPv4",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
				},
			},
		},
		{
			Id:          "a988649d-cfbf-4c2a-bd91-3b84d2403747",
			Name:        "EricSG",
			Description: "Test SecurityGroup",
			TenantId:    "57e98940a77f4bb988a21a7d0603a626",
			SecurityGroupRules: []securitygroups.ListSecurityGroupRule{
				{
					Id:              "0650c837-2ee0-478f-a184-d7ebc6cad353",
					Direction:       "egress",
					Protocol:        "",
					Ethertype:       "IPv4",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "a988649d-cfbf-4c2a-bd91-3b84d2403747",
				},
				{
					Id:              "d4866ba2-c085-4683-9dee-5932eb771106",
					Direction:       "egress",
					Protocol:        "",
					Ethertype:       "IPv6",
					Description:     "",
					RemoteGroupId:   "",
					RemoteIpPrefix:  "",
					TenantId:        "57e98940a77f4bb988a21a7d0603a626",
					PortRangeMax:    0,
					PortRangeMin:    0,
					SecurityGroupId: "a988649d-cfbf-4c2a-bd91-3b84d2403747",
				},
			},
		},
	},
}

func HandleListSuccessfully(t *testing.T, endPoint string) {
	th.Mux.HandleFunc("/v2.0/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		switch r.RequestURI {
		case "/v2.0/security-groups?limit=2":
			fmt.Fprintf(w, List1Output, endPoint)
		case "/v2.0/security-groups?limit=2&marker=7af80d49-0a43-462d-aed8-a1e12ac91af6":
			fmt.Fprintf(w, List2Output)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/security-groups/7af80d49-0a43-462d-aed8-a1e12ac91af6", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
