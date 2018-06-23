package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/quotas"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
{
  "quotas": {
    "resources": [{
      "type": "scaling_Group",
      "used": 1,
      "quota": 10,
      "max": 50
    }, {
      "type": "scaling_Config",
      "used": 1,
      "quota": 100,
      "max": 200
    }, {
      "type": "scaling_Policy",
      "used": -1,
      "quota": 10,
      "max": 50
    }, {
      "type": "scaling_Instance",
      "used": -1,
      "quota": 300,
      "max": 1000
    }]
  }
}
`

var ListResponse = quotas.ListResponse{
	Quotas: struct {
		Resources []struct {
			Type  string `json:"type,omitempty"`
			Used  int    `json:"used,omitempty"`
			Quota int    `json:"quota,omitempty"`
			Max   int    `json:"max,omitempty"`
		} `json:"resources,omitempty"`
	}{
		Resources: []struct {
			Type  string `json:"type,omitempty"`
			Used  int    `json:"used,omitempty"`
			Quota int    `json:"quota,omitempty"`
			Max   int    `json:"max,omitempty"`
		}{
			{
				Type:  "scaling_Group",
				Used:  1,
				Quota: 10,
				Max:   50,
			}, {
				Type:  "scaling_Config",
				Used:  1,
				Quota: 100,
				Max:   200,
			}, {
				Type:  "scaling_Policy",
				Used:  -1,
				Quota: 10,
				Max:   50,
			}, {
				Type:  "scaling_Instance",
				Used:  -1,
				Quota: 300,
				Max:   1000,
			},
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var ListWithInstancesOutput = `
{
  "quotas": {
    "resources": [{
      "type": "scaling_Group",
      "used": 1,
      "quota": 10,
      "max": 50
    }, {
      "type": "scaling_Config",
      "used": 1,
      "quota": 100,
      "max": 200
    }, {
      "type": "scaling_Policy",
      "used": -1,
      "quota": 10,
      "max": 50
    }, {
      "type": "scaling_Instance",
      "used": -1,
      "quota": 300,
      "max": 1000
    }]
  }
}
`

var ListWithInstancesResponse = quotas.ListWithInstancesResponse{
	Quotas: struct {
		Resources []struct {
			Type  string `json:"type,omitempty"`
			Used  int    `json:"used,omitempty"`
			Quota int    `json:"quota,omitempty"`
			Max   int    `json:"max,omitempty"`
		} `json:"resources,omitempty"`
	}{
		Resources: []struct {
			Type  string `json:"type,omitempty"`
			Used  int    `json:"used,omitempty"`
			Quota int    `json:"quota,omitempty"`
			Max   int    `json:"max,omitempty"`
		}{
			{
				Type:  "scaling_Group",
				Used:  1,
				Quota: 10,
				Max:   50,
			}, {
				Type:  "scaling_Config",
				Used:  1,
				Quota: 100,
				Max:   200,
			}, {
				Type:  "scaling_Policy",
				Used:  -1,
				Quota: 10,
				Max:   50,
			}, {
				Type:  "scaling_Instance",
				Used:  -1,
				Quota: 300,
				Max:   1000,
			},
		},
	},
}

func HandleListWithInstancesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/quotas/e4a97959-34d0-4c58-ab85-1dda4d0b9d11", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListWithInstancesOutput)
	})
}
