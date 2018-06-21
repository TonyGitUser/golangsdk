package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/securitygrouprules"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "security_group_rule": {
    "remote_group_id": null,
    "direction": "egress",
    "protocol": "tcp",
    "description": "Test SecurityGroup",
    "ethertype": "IPv4",
    "remote_ip_prefix": "10.10.0.0/24",
    "port_range_max": null,
    "security_group_id": "7af80d49-0a43-462d-aed8-a1e12ac91af6",
    "port_range_min": null,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "id": "26243298-ae79-46a3-bad9-34395762e033"
  }
}
`

var CreateResponse = securitygrouprules.CreateResponse{
	SecurityGroupRule: securitygrouprules.ListSecurityGroupRule{
		RemoteGroupId:   "",
		Direction:       "egress",
		Protocol:        "tcp",
		Description:     "Test SecurityGroup",
		Ethertype:       "IPv4",
		RemoteIpPrefix:  "10.10.0.0/24",
		PortRangeMax:    0,
		SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
		PortRangeMin:    0,
		TenantId:        "57e98940a77f4bb988a21a7d0603a626",
		Id:              "26243298-ae79-46a3-bad9-34395762e033",
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/security-group-rules", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var GetOutput = `
{"security_group_rule": {"remote_group_id": null, "direction": "egress", "protocol": "tcp", "description": "Test SecurityGroup", "ethertype": "IPv4", "remote_ip_prefix": "10.10.0.0/24", "port_range_max": null, "security_group_id": "7af80d49-0a43-462d-aed8-a1e12ac91af6", "port_range_min": null, "tenant_id": "57e98940a77f4bb988a21a7d0603a626", "id": "26243298-ae79-46a3-bad9-34395762e033"}}
`

var GetResponse = securitygrouprules.GetResponse{
	SecurityGroupRule: securitygrouprules.ListSecurityGroupRule{
		RemoteGroupId:   "",
		Direction:       "egress",
		Protocol:        "tcp",
		Description:     "Test SecurityGroup",
		Ethertype:       "IPv4",
		RemoteIpPrefix:  "10.10.0.0/24",
		PortRangeMax:    0,
		SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
		PortRangeMin:    0,
		TenantId:        "57e98940a77f4bb988a21a7d0603a626",
		Id:              "26243298-ae79-46a3-bad9-34395762e033",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/security-group-rules/26243298-ae79-46a3-bad9-34395762e033", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var List1Output = `
{
  "security_group_rules_links": [{
    "href": "%s/v2.0/security-group-rules?limit=2&protocol=tcp&marker=ff0eb13a-9ea7-432c-9032-9316a298d43a",
    "rel": "next"
  }, {
    "href": "https://network.region.cnnorth1.hwclouds.com/v2.0/security-group-rules?limit=2&protocol=tcp&marker=164bcf26-c9c0-417a-abc5-c129ad0c1854&page_reverse=True",
    "rel": "previous"
  }],
  "security_group_rules": [{
    "remote_group_id": null,
    "direction": "ingress",
    "protocol": "tcp",
    "description": null,
    "ethertype": "IPv4",
    "remote_ip_prefix": "0.0.0.0/0",
    "port_range_max": 3389,
    "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
    "port_range_min": 3389,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "id": "164bcf26-c9c0-417a-abc5-c129ad0c1854"
  }, {
    "remote_group_id": null,
    "direction": "ingress",
    "protocol": "tcp",
    "description": null,
    "ethertype": "IPv4",
    "remote_ip_prefix": "0.0.0.0/0",
    "port_range_max": 22,
    "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
    "port_range_min": 22,
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "id": "ff0eb13a-9ea7-432c-9032-9316a298d43a"
  }]
}
`

var List2Output = `
{
  "security_group_rules_links": [{
    "href": "https://network.region.cnnorth1.hwclouds.com/v2.0/security-group-rules?limit=2&protocol=tcp&page_reverse=True",
    "rel": "previous"
  }],
  "security_group_rules": []
}
`

var ListResponse = securitygrouprules.ListResponse{
	SecurityGroupRules: []securitygrouprules.ListSecurityGroupRule{
		{
			RemoteGroupId:   "",
			Direction:       "ingress",
			Protocol:        "tcp",
			Description:     "",
			Ethertype:       "IPv4",
			RemoteIpPrefix:  "0.0.0.0/0",
			PortRangeMax:    3389,
			SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			PortRangeMin:    3389,
			TenantId:        "57e98940a77f4bb988a21a7d0603a626",
			Id:              "164bcf26-c9c0-417a-abc5-c129ad0c1854",
		},
		{
			RemoteGroupId:   "",
			Direction:       "ingress",
			Protocol:        "tcp",
			Description:     "",
			Ethertype:       "IPv4",
			RemoteIpPrefix:  "0.0.0.0/0",
			PortRangeMax:    22,
			SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			PortRangeMin:    22,
			TenantId:        "57e98940a77f4bb988a21a7d0603a626",
			Id:              "ff0eb13a-9ea7-432c-9032-9316a298d43a",
		},
	},
}

func HandleListSuccessfully(t *testing.T, endPoint string) {
	th.Mux.HandleFunc("/v2.0/security-group-rules", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		switch r.RequestURI {
		case "/v2.0/security-group-rules?limit=2&protocol=tcp":
			fmt.Fprintf(w, List1Output, endPoint)
		case "/v2.0/security-group-rules?limit=2&protocol=tcp&marker=ff0eb13a-9ea7-432c-9032-9316a298d43a":
			fmt.Fprintf(w, List2Output)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/security-groups/26243298-ae79-46a3-bad9-34395762e033", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
