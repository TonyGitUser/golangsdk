package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/healthcheck"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "healthcheck_interval": 5,
  "listener_id": "d5f3197c6bd8491ca1dfc905350d85ea",
  "id": "c74c17f02c5142b3befc973056f220d3",
  "healthcheck_protocol": "HTTP",
  "unhealthy_threshold": 3,
  "update_time": "2018-05-06 12:36:04",
  "create_time": "2018-05-06 12:36:04",
  "healthcheck_connect_port": 80,
  "healthcheck_timeout": 10,
  "healthy_threshold": 3
}
`

var CreateResponse = healthcheck.CreateResponse{
	HealthcheckInterval:    5,
	ListenerId:             "d5f3197c6bd8491ca1dfc905350d85ea",
	Id:                     "c74c17f02c5142b3befc973056f220d3",
	HealthcheckProtocol:    "HTTP",
	UnhealthyThreshold:     3,
	UpdateTime:             "2018-05-06 12:36:04",
	CreateTime:             "2018-05-06 12:36:04",
	HealthcheckConnectPort: 80,
	HealthcheckTimeout:     10,
	HealthyThreshold:       3,
}

func HandleCreateSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "healthcheck_interval": 6,
  "listener_id": "d5f3197c6bd8491ca1dfc905350d85ea",
  "id": "c74c17f02c5142b3befc973056f220d3",
  "healthcheck_protocol": "HTTP",
  "unhealthy_threshold": 4,
  "update_time": "2018-05-06 12:41:36",
  "create_time": "2018-05-06 12:36:04",
  "healthy_threshold": 4,
  "healthcheck_connect_port": 81,
  "healthcheck_timeout": 20,
  "healthcheck_uri": null
}
`

var UpdateResponse = healthcheck.UpdateResponse{
	HealthcheckInterval:    6,
	ListenerId:             "d5f3197c6bd8491ca1dfc905350d85ea",
	Id:                     "c74c17f02c5142b3befc973056f220d3",
	HealthcheckProtocol:    "HTTP",
	UnhealthyThreshold:     4,
	UpdateTime:             "2018-05-06 12:41:36",
	CreateTime:             "2018-05-06 12:36:04",
	HealthyThreshold:       4,
	HealthcheckConnectPort: 81,
	HealthcheckTimeout:     20,
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/healthcheck/c74c17f02c5142b3befc973056f220d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "healthcheck_interval": 6,
  "listener_id": "d5f3197c6bd8491ca1dfc905350d85ea",
  "id": "c74c17f02c5142b3befc973056f220d3",
  "healthcheck_protocol": "HTTP",
  "unhealthy_threshold": 4,
  "update_time": "2018-05-06 12:41:36",
  "create_time": "2018-05-06 12:36:04",
  "healthy_threshold": 4,
  "healthcheck_connect_port": 81,
  "healthcheck_timeout": 20,
  "healthcheck_uri": null
}
`

var GetResponse = healthcheck.GetResponse{
	HealthcheckInterval:    6,
	ListenerId:             "d5f3197c6bd8491ca1dfc905350d85ea",
	Id:                     "c74c17f02c5142b3befc973056f220d3",
	HealthcheckProtocol:    "HTTP",
	UnhealthyThreshold:     4,
	UpdateTime:             "2018-05-06 12:41:36",
	CreateTime:             "2018-05-06 12:36:04",
	HealthyThreshold:       4,
	HealthcheckConnectPort: 81,
	HealthcheckTimeout:     20,
}

func HandleGetSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/healthcheck/c74c17f02c5142b3befc973056f220d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var DeleteOutput = ""

func HandleDeleteSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/elbaas/healthcheck/c74c17f02c5142b3befc973056f220d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DeleteOutput)
	})
}
