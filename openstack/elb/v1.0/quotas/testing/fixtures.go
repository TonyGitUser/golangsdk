package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/quotas"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
{
  "quotas": {
    "resources": [{
      "max": 255,
      "type": "elb",
      "used": 2,
      "min": 1,
      "quota": 5
    }, {
      "max": 255,
      "type": "listener",
      "used": 0,
      "min": 1,
      "quota": 5
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
			Min   int    `json:"min,omitempty"`
		} `json:"resources,omitempty"`
	}{
		Resources: []struct {
			Type  string `json:"type,omitempty"`
			Used  int    `json:"used,omitempty"`
			Quota int    `json:"quota,omitempty"`
			Max   int    `json:"max,omitempty"`
			Min   int    `json:"min,omitempty"`
		}{
			{
				Max:   255,
				Type:  "elb",
				Used:  2,
				Min:   1,
				Quota: 5,
			},
			{
				Max:   255,
				Type:  "listener",
				Used:  0,
				Min:   1,
				Quota: 5,
			},
		},
	},
}

func HandleListSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
