package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/tags"
)

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group_tag/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/tags/action ", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}

var ListTenantTagsOutput = `
{
  "tags": [{
    "key": "ENV15",
    "values": ["ENV15"]
  }, {
    "key": "ENV151",
    "values": ["ENV151"]
  }]
}
`

var ListTenantTagsResponse = tags.ListTenantTagsResponse{
	Tags: []struct {
		Key    string   `json:"key,omitempty"`
		Values []string `json:"values,omitempty"`
	}{
		{
			Key:    "ENV15",
			Values: []string{"ENV15"},
		}, {
			Key:    "ENV151",
			Values: []string{"ENV151"},
		},
	},
}

func HandleListTenantTagsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group_tag/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListTenantTagsOutput)
	})
}

var ListResourceTagsOutput = `
{
  "tags": [{
    "key": "ENV15",
    "value": "ENV15"
  }, {
    "key": "ENV151",
    "value": "ENV151"
  }]
}
`

var ListResourceTagsResponse = tags.ListResourceTagsResponse{
	Tags: []struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	}{
		{
			Key:   "ENV15",
			Value: "ENV15",
		}, {
			Key:   "ENV151",
			Value: "ENV151",
		},
	},
}

func HandleListResourceTagsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group_tag/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListResourceTagsOutput)
	})
}
