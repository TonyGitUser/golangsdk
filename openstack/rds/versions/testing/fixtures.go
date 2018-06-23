package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/versions"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
{
  "versions": [{
    "id": "v1",
    "links": [],
    "status": "CURRENT",
    "updated": "2017-02-07T17:34:02Z"
  }, {
    "id": "v1.0",
    "links": [{
      "href": "",
      "rel": "self"
    }],
    "status": "CURRENT",
    "updated": "2017-03-23T17:34:02Z"
  }]
}
`

var ListResponse = versions.ListResponse{
	Versions: []versions.Version{
		{
			Id:      "v1",
			Links:   []versions.Link{},
			Status:  "CURRENT",
			Updated: "2017-02-07T17:34:02Z",
		},
		{
			Id: "v1.0",
			Links: []versions.Link{
				{
					Href: "",
					Rel:  "self",
				},
			},
			Status:  "CURRENT",
			Updated: "2017-03-23T17:34:02Z",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var GetOutput = `
{
  "version": {
    "id": "v1",
    "links": [],
    "status": "CURRENT",
    "updated": "2017-02-07T17:34:02Z"
  },
  "versions": {
    "id": "v1",
    "links": [],
    "status": "CURRENT",
    "updated": "2017-02-07T17:34:02Z"
  }
}
`

var GetResponse = versions.GetResponse{
	Versions: versions.Version{
		Id:      "v1",
		Links:   []versions.Link{},
		Status:  "CURRENT",
		Updated: "2017-02-07T17:34:02Z",
	},
	Version: versions.Version{
		Id:      "v1",
		Links:   []versions.Link{},
		Status:  "CURRENT",
		Updated: "2017-02-07T17:34:02Z",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}
