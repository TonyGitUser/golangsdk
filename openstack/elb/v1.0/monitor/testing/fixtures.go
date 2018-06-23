package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/monitor"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
[{
  "cps": 0,
  "create_time": "2018-05-06 13:25:22",
  "inact_conn": 0,
  "in_Bps": 0,
  "out_pps": 0,
  "loadbalancer_ip": "49.4.9.141",
  "in_pps": 0,
  "loadbalancer_name": "UpdatedLB",
  "loadbalancer_id": "ebe982b8afe04851bbc26ad4609003bf",
  "out_Bps": 0,
  "act_conn": 0,
  "ncps": 0
}]
`

var ListResponse = monitor.ListResponse{
	Items: []struct {
		// Specifies the number of active connections.
		ActConn int `json:"act_conn,omitempty"`

		// Specifies the number of concurrent connections per second.
		Cps int `json:"cps,omitempty"`

		// Specifies the monitoring period.
		CreateTime string `json:"create_time,omitempty"`

		// Specifies the inbound bandwidth per second. The unit is byte/second.
		InBps int `json:"in_Bps,omitempty"`

		// Specifies the number of incoming packets per second.
		InPps int `json:"in_pps,omitempty"`

		// Specifies the number of inactive connections.
		InactConn int `json:"inact_conn,omitempty"`

		// Specifies the load balancer ID.
		LoadbalancerId string `json:"loadbalancer_id,omitempty"`

		// Specifies the ELB IP address.
		LoadbalancerIp string `json:"loadbalancer_ip,omitempty"`

		// Specifies the load balancer name.
		LoadbalancerName string `json:"loadbalancer_name,omitempty"`

		// Specifies the number of new connections per second.
		Ncps int `json:"ncps,omitempty"`

		// Specifies the outbound bandwidth per second. The unit is byte/second.
		OutBps int `json:"out_Bps,omitempty"`

		// Specifies the number of outgoing packets per second.
		OutPps int `json:"out_pps,omitempty"`
	}{
		{
			Cps:              0,
			CreateTime:       "2018-05-06 13:25:22",
			InactConn:        0,
			InBps:            0,
			OutPps:           0,
			LoadbalancerIp:   "49.4.9.141",
			InPps:            0,
			LoadbalancerName: "UpdatedLB",
			LoadbalancerId:   "ebe982b8afe04851bbc26ad4609003bf",
			OutBps:           0,
			ActConn:          0,
			Ncps:             0,
		},
	},
}

func HandleListSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/monitor", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
