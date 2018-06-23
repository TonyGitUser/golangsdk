package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/backendces"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "uri": "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c763267616016334e3f759605a",
  "job_id": "8aace0c763267616016334e3f759605a"
}
`

var CreateResponse = backendces.CreateResponse{
	Uri:   "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c763267616016334e3f759605a",
	JobId: "8aace0c763267616016334e3f759605a",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners/d5f3197c6bd8491ca1dfc905350d85ea/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var ListOutput = `
[{
  "server_address": "192.168.1.215",
  "id": "7ef7f0ebc3964d8f9edf556a1bb943d3",
  "status": "ACTIVE",
  "address": "100.70.38.178",
  "listeners": [{
    "id": "d5f3197c6bd8491ca1dfc905350d85ea"
  }],
  "update_time": "2018-05-06 09:59:14",
  "create_time": "2018-05-06 09:59:14",
  "server_name": null,
  "server_id": "e1040c67-b130-41ee-9405-07c6c6c20208",
  "health_status": "UNAVAILABLE"
}]
`

var ListResponse = backendces.ListResponse{
	Items: []struct {
		ServerAddress string `json:"server_address,omitempty"`
		Id            string `json:"id,omitempty"`
		Address       string `json:"address,omitempty"`
		Status        string `json:"status,omitempty"`
		HealthStatus  string `json:"health_status,omitempty"`
		UpdateTime    string `json:"update_time,omitempty"`
		CreateTime    string `json:"create_time,omitempty"`
		ServerName    string `json:"server_name,omitempty"`
		ServerId      string `json:"server_id,omitempty"`
		Listeners     []struct {
			Id string `json:"id,omitempty"`
		} `json:"listeners,omitempty"`
	}{
		{
			ServerAddress: "192.168.1.215",
			Id:            "7ef7f0ebc3964d8f9edf556a1bb943d3",
			Status:        "ACTIVE",
			Address:       "100.70.38.178",
			Listeners: []struct {
				Id string `json:"id,omitempty"`
			}{
				{
					Id: "d5f3197c6bd8491ca1dfc905350d85ea",
				},
			},
			UpdateTime:   "2018-05-06 09:59:14",
			CreateTime:   "2018-05-06 09:59:14",
			ServerName:   "",
			ServerId:     "e1040c67-b130-41ee-9405-07c6c6c20208",
			HealthStatus: "UNAVAILABLE",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners/d5f3197c6bd8491ca1dfc905350d85ea/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var DeleteOutput = `
{
  "uri": "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c763267616016334e43fa26060",
  "job_id": "8aace0c763267616016334e43fa26060"
}
`

var DeleteResponse = backendces.DeleteResponse{
	Uri:   "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c763267616016334e43fa26060",
	JobId: "8aace0c763267616016334e43fa26060",
}

func HandleDeleteSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners/ca082827b61d4902bfaf32e8174e244a/members/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DeleteOutput)
	})
}
