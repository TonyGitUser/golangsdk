package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/quotas"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
{
  "quotas": {
    "resources": [{
      "type": "vpc",
      "used": 3,
      "quota": 5,
      "min": 0
    }]
  }
}
`

var ListResponse = quotas.ListResponse{
	Quotas: quotas.Quota{
		[]quotas.Resource{
			{
				Type:  "vpc",
				Used:  3,
				Quota: 5,
				Min:   0,
			},
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
