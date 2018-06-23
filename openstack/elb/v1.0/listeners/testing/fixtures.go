package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/listeners"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "update_time": "2018-05-06 04:31:26",
  "cookie_timeout": 1,
  "id": "223a537e337e4f9b947b8b39eb1b1f6c",
  "tcp_draining": false,
  "sticky_session_type": "insert",
  "certificates": [],
  "description": "LS TESTING",
  "loadbalancer_id": "ebe982b8afe04851bbc26ad4609003bf",
  "create_time": "2018-05-06 04:31:26",
  "status": "ACTIVE",
  "protocol": "HTTP",
  "port": 80,
  "backend_protocol": "HTTP",
  "lb_algorithm": "roundrobin",
  "session_sticky": true,
  "backend_port": 81,
  "name": "TestLS"
}
`

var CreateResponse = listeners.CreateResponse{
	UpdateTime:        "2018-05-06 04:31:26",
	CookieTimeout:     1,
	Id:                "223a537e337e4f9b947b8b39eb1b1f6c",
	TcpDraining:       false,
	StickySessionType: "insert",
	Certificates:      []string{},
	Description:       "LS TESTING",
	LoadbalancerId:    "ebe982b8afe04851bbc26ad4609003bf",
	CreateTime:        "2018-05-06 04:31:26",
	Status:            "ACTIVE",
	Protocol:          "HTTP",
	Port:              80,
	BackendProtocol:   "HTTP",
	LbAlgorithm:       "roundrobin",
	SessionSticky:     true,
	BackendPort:       81,
	Name:              "TestLS",
}

func HandleCreateSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "update_time": "2018-05-06 04:42:07",
  "backend_port": 81,
  "id": "223a537e337e4f9b947b8b39eb1b1f6c",
  "ssl_ciphers": null,
  "backend_protocol": "HTTP",
  "sticky_session_type": "insert",
  "certificate_id": null,
  "certificates": [],
  "admin_state_up": false,
  "description": "LS TESTING",
  "loadbalancer_id": "ebe982b8afe04851bbc26ad4609003bf",
  "status": "ACTIVE",
  "create_time": "2018-05-06 04:31:26",
  "cookie_timeout": 1,
  "ssl_protocols": null,
  "protocol": "HTTP",
  "port": 80,
  "udp_timeout": null,
  "tcp_draining": false,
  "tcp_timeout": null,
  "lb_algorithm": "roundrobin",
  "healthcheck_id": null,
  "session_sticky": true,
  "tcp_draining_timeout": null,
  "name": "ModifiedTestLS"
}
`

var UpdateResponse = listeners.UpdateResponse{
	UpdateTime:         "2018-05-06 04:42:07",
	BackendPort:        81,
	Id:                 "223a537e337e4f9b947b8b39eb1b1f6c",
	BackendProtocol:    "HTTP",
	StickySessionType:  "insert",
	CertificateId:      "",
	Certificates:       []string{},
	AdminStateUp:       false,
	Description:        "LS TESTING",
	LoadbalancerId:     "ebe982b8afe04851bbc26ad4609003bf",
	Status:             "ACTIVE",
	CreateTime:         "2018-05-06 04:31:26",
	CookieTimeout:      1,
	Protocol:           "HTTP",
	Port:               80,
	TcpDraining:        false,
	LbAlgorithm:        "roundrobin",
	HealthcheckId:      "",
	SessionSticky:      true,
	TcpDrainingTimeout: 0,
	Name:               "ModifiedTestLS",
}

func HandleUpdateSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners/223a537e337e4f9b947b8b39eb1b1f6c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "update_time": "2018-05-06 04:42:07",
  "backend_port": 81,
  "id": "223a537e337e4f9b947b8b39eb1b1f6c",
  "ssl_ciphers": null,
  "backend_protocol": "HTTP",
  "sticky_session_type": "insert",
  "certificate_id": null,
  "certificates": [],
  "lb_algorithm": "roundrobin",
  "admin_state_up": false,
  "description": "LS TESTING",
  "loadbalancer_id": "ebe982b8afe04851bbc26ad4609003bf",
  "status": "ACTIVE",
  "create_time": "2018-05-06 04:31:26",
  "cookie_timeout": 1,
  "ssl_protocols": null,
  "protocol": "HTTP",
  "port": 80,
  "udp_timeout": null,
  "tcp_draining": false,
  "tcp_timeout": null,
  "member_number": 0,
  "healthcheck_id": null,
  "session_sticky": true,
  "tcp_draining_timeout": null,
  "name": "ModifiedTestLS"
}
`

var GetResponse = listeners.GetResponse{
	UpdateTime:        "2018-05-06 04:42:07",
	BackendPort:       81,
	Id:                "223a537e337e4f9b947b8b39eb1b1f6c",
	BackendProtocol:   "HTTP",
	StickySessionType: "insert",
	CertificateId:     "",
	Certificates:      []string{},
	AdminStateUp:      false,
	Description:       "LS TESTING",
	LoadbalancerId:    "ebe982b8afe04851bbc26ad4609003bf",
	Status:            "ACTIVE",
	CreateTime:        "2018-05-06 04:31:26",
	CookieTimeout:     1,
	Protocol:          "HTTP",
	Port:              80,
	LbAlgorithm:       "roundrobin",
	HealthcheckId:     "",
	SessionSticky:     true,
	Name:              "ModifiedTestLS",
}

func HandleGetSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners/223a537e337e4f9b947b8b39eb1b1f6c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var ListOutput = `
[{
  "update_time": "2018-05-06 04:42:07",
  "backend_port": 81,
  "id": "223a537e337e4f9b947b8b39eb1b1f6c",
  "ssl_ciphers": null,
  "backend_protocol": "HTTP",
  "sticky_session_type": "insert",
  "certificate_id": null,
  "certificates": [],
  "lb_algorithm": "roundrobin",
  "admin_state_up": false,
  "description": "LS TESTING",
  "loadbalancer_id": "ebe982b8afe04851bbc26ad4609003bf",
  "status": "ACTIVE",
  "create_time": "2018-05-06 04:31:26",
  "cookie_timeout": 1,
  "ssl_protocols": null,
  "protocol": "HTTP",
  "port": 80,
  "udp_timeout": null,
  "tcp_draining": false,
  "tcp_timeout": null,
  "member_number": 0,
  "healthcheck_id": null,
  "session_sticky": true,
  "tcp_draining_timeout": null,
  "name": "ModifiedTestLS"
}]
`

var ListResponse = listeners.ListResponse{
	Items: []struct {
		// Specifies the time when information about the listener was updated.
		UpdateTime string `json:"update_time,omitempty"`

		// Specifies the backend port.
		BackendPort int `json:"backend_port,omitempty"`

		// Specifies the listener ID in UUID format.
		Id string `json:"id,omitempty"`

		// Specifies the backend protocol.
		BackendProtocol string `json:"backend_protocol,omitempty"`

		// Specifies the cookie processing method. insert indicates that the cookie is inserted by the load balancer. This parameter is valid when protocol is set to HTTP and session_sticky to true. The default value is insert. This parameter is invalid when protocol is set to TCP, which means the parameter is empty.
		StickySessionType string `json:"sticky_session_type,omitempty"`

		// Provides supplementary information about the listener.
		Description string `json:"description,omitempty"`

		// Specifies the ID of the load balancer to which the listener belongs.
		LoadbalancerId string `json:"loadbalancer_id,omitempty"`

		// Specifies the time when the listener was created.
		CreateTime string `json:"create_time,omitempty"`

		// Specifies the listener status. The value can be ACTIVE, PENDING_CREATE, or ERROR.
		Status string `json:"status,omitempty"`

		// Specifies the listening protocol used for layer 4 or 7.
		Protocol string `json:"protocol,omitempty"`

		// Specifies the load balancing algorithm for the listener.
		LbAlgorithm string `json:"lb_algorithm,omitempty"`

		// Specifies the status of the load balancer.Optional values:false: The load balancer is disabled.true: The load balancer is running properly.
		AdminStateUp bool `json:"admin_state_up,omitempty"`

		// Specifies the cookie timeout duration (minutes).The value ranges from 1 to 1440. This parameter is valid when session_sticky is true and sticky_session_type is insert.
		CookieTimeout int `json:"cookie_timeout,omitempty"`

		// Specifies the quantity of backend ECSs.
		MemberNumber int `json:"member_number,omitempty"`

		// Specifies the health check task ID.
		HealthcheckId string `json:"healthcheck_id,omitempty"`

		// Specifies whether to enable the sticky session feature. The feature is enabled when the value is true. This parameter is valid only when protocol is set to HTTP.
		SessionSticky bool `json:"session_sticky,omitempty"`

		// Specifies the listening port.
		Port int `json:"port,omitempty"`

		// Specifies the listener name.
		Name string `json:"name,omitempty"`

		// Specifies the ID of the SSL certificate for security authentication.This parameter is mandatory when protocol is set to HTTPS or SSL. Otherwise, null is displayed for this parameter.
		CertificateId string `json:"certificate_id,omitempty"`

		// Specifies the SSL certificate ID list if the value of protocol is HTTPS.This parameter is mandatory in SNI scenarios.
		Certificates []string `json:"certificates,omitempty"`

		// Specifies the timeout duration for a TCP request session.
		TcpTimeout int `json:"tcp_timeout,omitempty"`

		// Specifies the timeout duration for a UDP request session.
		UdpTimeout int `json:"udp_timeout,omitempty"`
	}{
		{
			UpdateTime:        "2018-05-06 04:42:07",
			BackendPort:       81,
			Id:                "223a537e337e4f9b947b8b39eb1b1f6c",
			BackendProtocol:   "HTTP",
			StickySessionType: "insert",
			CertificateId:     "",
			Certificates:      []string{},
			AdminStateUp:      false,
			Description:       "LS TESTING",
			LoadbalancerId:    "ebe982b8afe04851bbc26ad4609003bf",
			Status:            "ACTIVE",
			CreateTime:        "2018-05-06 04:31:26",
			CookieTimeout:     1,
			Protocol:          "HTTP",
			Port:              80,
			LbAlgorithm:       "roundrobin",
			HealthcheckId:     "",
			SessionSticky:     true,
			Name:              "ModifiedTestLS",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var DeleteOutput = ""

func HandleDeleteSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/listeners/223a537e337e4f9b947b8b39eb1b1f6c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DeleteOutput)
	})
}
