package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v2.0/routers"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "router": {
    "id": "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
    "name": "EricTestPort",
    "status": "ACTIVE",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "admin_state_up": true,
    "external_gateway_info": {
      "network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
      "enable_snat": false
    },
    "routes": []
  }
}
`

var CreateResponse = routers.CreateResponse{
	Router: routers.ListRouter{
		Id:           "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
		Name:         "EricTestPort",
		Status:       "ACTIVE",
		TenantId:     "57e98940a77f4bb988a21a7d0603a626",
		AdminStateUp: true,
		ExternalGatewayInfo: routers.ExternalGatewayInfo{
			NetworkId:  "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			EnableSnat: false,
		},
		Routes: []routers.Route{},
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/routers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "router": {
    "id": "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
    "name": "ModifiedRouter",
    "status": "ACTIVE",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "admin_state_up": true,
    "external_gateway_info": {
      "network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
      "enable_snat": false
    },
    "routes": []
  }
}
`

var UpdateResponse = routers.UpdateResponse{
	Router: routers.ListRouter{
		Id:           "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
		Name:         "ModifiedRouter",
		Status:       "ACTIVE",
		TenantId:     "57e98940a77f4bb988a21a7d0603a626",
		AdminStateUp: true,
		ExternalGatewayInfo: routers.ExternalGatewayInfo{
			NetworkId:  "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			EnableSnat: false,
		},
		Routes: []routers.Route{},
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/routers/a4d322dd-821e-45ca-b1eb-3ccdbaef407d", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "router": {
    "id": "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
    "name": "EricTestPort",
    "status": "ACTIVE",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "admin_state_up": true,
    "external_gateway_info": {
      "network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
      "enable_snat": false
    },
    "routes": []
  }
}
`

var GetResponse = routers.GetResponse{
	Router: routers.ListRouter{
		Id:           "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
		Name:         "EricTestPort",
		Status:       "ACTIVE",
		TenantId:     "57e98940a77f4bb988a21a7d0603a626",
		AdminStateUp: true,
		ExternalGatewayInfo: routers.ExternalGatewayInfo{
			NetworkId:  "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
			EnableSnat: false,
		},
		Routes: []routers.Route{},
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/routers/a4d322dd-821e-45ca-b1eb-3ccdbaef407d", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var List1Output = `
{
  "routers": [{
    "id": "773c3c42-d315-417b-9063-87091713148c",
    "name": "vpc-c8cb",
    "status": "ACTIVE",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "admin_state_up": true,
    "external_gateway_info": {
      "network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
      "enable_snat": false
    },
    "routes": []
  }, {
    "id": "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
    "name": "EricTestPort",
    "status": "ACTIVE",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "admin_state_up": true,
    "external_gateway_info": {
      "network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
      "enable_snat": false
    },
    "routes": []
  }],
  "routers_links": [{
    "href": "%s/v2.0/routers?limit=2&distributed=true&marker=a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/routers?limit=2&distributed=true&marker=773c3c42-d315-417b-9063-87091713148c&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List2Output = `
{
  "routers": [{
    "id": "e23caa95-2599-4aa8-a2db-be3444450e78",
    "name": "EricTestVPC",
    "status": "ACTIVE",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "admin_state_up": true,
    "external_gateway_info": {
      "network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
      "enable_snat": false
    },
    "routes": []
  }, {
    "id": "ea3b0efe-0d6a-4288-8b16-753504a994b9",
    "name": "ABC",
    "status": "ACTIVE",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "admin_state_up": true,
    "external_gateway_info": {
      "network_id": "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
      "enable_snat": false
    },
    "routes": []
  }],
  "routers_links": [{
    "href": "%s/v2.0/routers?limit=2&distributed=true&marker=ea3b0efe-0d6a-4288-8b16-753504a994b9",
    "rel": "next"
  }, {
    "href": "https://None/v2.0/routers?limit=2&distributed=true&marker=e23caa95-2599-4aa8-a2db-be3444450e78&page_reverse=True",
    "rel": "previous"
  }]
}
`

var List3Output = `
{
  "routers": [],
  "routers_links": [{
    "href": "https://None/v2.0/routers?limit=2&distributed=true&page_reverse=True",
    "rel": "previous"
  }]
}
`
var ListResponse = routers.ListResponse{
	Routers: []routers.ListRouter{
		{
			Id:           "773c3c42-d315-417b-9063-87091713148c",
			Name:         "vpc-c8cb",
			Status:       "ACTIVE",
			TenantId:     "57e98940a77f4bb988a21a7d0603a626",
			AdminStateUp: true,
			ExternalGatewayInfo: routers.ExternalGatewayInfo{
				NetworkId:  "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
				EnableSnat: false,
			},
			Routes: []routers.Route{},
		},
		{
			Id:           "a4d322dd-821e-45ca-b1eb-3ccdbaef407d",
			Name:         "EricTestPort",
			Status:       "ACTIVE",
			TenantId:     "57e98940a77f4bb988a21a7d0603a626",
			AdminStateUp: true,
			ExternalGatewayInfo: routers.ExternalGatewayInfo{
				NetworkId:  "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
				EnableSnat: false,
			},
			Routes: []routers.Route{},
		},
		{
			Id:           "e23caa95-2599-4aa8-a2db-be3444450e78",
			Name:         "EricTestVPC",
			Status:       "ACTIVE",
			TenantId:     "57e98940a77f4bb988a21a7d0603a626",
			AdminStateUp: true,
			ExternalGatewayInfo: routers.ExternalGatewayInfo{
				NetworkId:  "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
				EnableSnat: false,
			},
			Routes: []routers.Route{},
		},
		{
			Id:           "ea3b0efe-0d6a-4288-8b16-753504a994b9",
			Name:         "ABC",
			Status:       "ACTIVE",
			TenantId:     "57e98940a77f4bb988a21a7d0603a626",
			AdminStateUp: true,
			ExternalGatewayInfo: routers.ExternalGatewayInfo{
				NetworkId:  "0a2228f2-7f8a-45f1-8e09-9039e1d09975",
				EnableSnat: false,
			},
			Routes: []routers.Route{},
		},
	},
}

func HandleListSuccessfully(t *testing.T, endPoint string) {
	th.Mux.HandleFunc("/v2.0/routers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		switch r.RequestURI {
		case "/v2.0/routers?limit=2":
			fmt.Fprintf(w, List1Output, endPoint)
		case "/v2.0/routers?limit=2&distributed=true&marker=a4d322dd-821e-45ca-b1eb-3ccdbaef407d":
			fmt.Fprintf(w, List2Output, endPoint)
		case "/v2.0/routers?limit=2&distributed=true&marker=ea3b0efe-0d6a-4288-8b16-753504a994b9":
			fmt.Fprintf(w, List3Output)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/routers/a4d322dd-821e-45ca-b1eb-3ccdbaef407d", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
