package monitor

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response.Items)
	return &response, err
}

type ListResponse struct {
	Items []struct {

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
	}
}
