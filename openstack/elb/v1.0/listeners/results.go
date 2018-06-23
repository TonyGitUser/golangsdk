package listeners

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {

	// Specifies the time when information about the listener was updated.
	UpdateTime string `json:"update_time,omitempty"`

	// Specifies the backend port.
	BackendPort int `json:"backend_port,omitempty"`

	// Specifies the listener ID.
	Id string `json:"id,omitempty"`

	// Specifies the backend protocol.
	BackendProtocol string `json:"backend_protocol,omitempty"`

	// Specifies the cookie processing method. insert indicates that the cookie is inserted by the load balancer. This parameter is valid when session_sticky is true. The default value is insert. This parameter is valid only when protocol is set to HTTP.
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

	// Specifies the listening port.
	Port int `json:"port,omitempty"`

	// Specifies the cookie timeout period (s). This parameter is valid when session_sticky is true and sticky_session_type is insert.The value ranges from 1 to 1440.
	CookieTimeout int `json:"cookie_timeout,omitempty"`

	// Specifies the status of the load balancer.Optional values:false: The load balancer is disabled.true: The load balancer is running properly.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies whether to enable the sticky session feature. The feature is enabled when the value is true.
	SessionSticky bool `json:"session_sticky,omitempty"`

	// Specifies the load balancing algorithm for the listener.
	LbAlgorithm string `json:"lb_algorithm,omitempty"`

	// Specifies the listener name.
	Name string `json:"name,omitempty"`

	// Specifies whether to maintain the TCP connection to the backend ECS after the ECS is deleted. This parameter is valid when protocol is set to TCP.The value can be true or false.
	TcpDraining bool `json:"tcp_draining,omitempty"`

	// Specifies whether to maintain the TCP connection to the backend ECS after the ECS is deleted.This parameter is valid when protocol is set to TCP and tcp_draining to true.The value ranges from 0 to 60.
	TcpDrainingTimeout int `json:"tcp_draining_timeout,omitempty"`

	// Specifies the SSL protocol standard supported by a listener.This parameter is used for enabling specified encryption protocols and displayed only when protocol is set to HTTPS or SSL.
	SslProtocols string `json:"ssl_protocols,omitempty"`

	// Specifies the cipher suite of an encryption protocol.This parameter is displayed only when protocol is set to HTTPS or SSL.
	SslCiphers string `json:"ssl_ciphers,omitempty"`

	// Specifies the default SSL certificate ID.This parameter is displayed only when protocol is set to HTTPS or SSL.
	CertificateId string `json:"certificate_id,omitempty"`

	// Specifies the SSL certificate ID list if the value of protocol is HTTPS.This parameter is mandatory in SNI scenarios.
	Certificates []string `json:"certificates,omitempty"`
}

type DeleteResult struct {
	commonResult
}

func (r DeleteResult) Extract() (*DeleteResponse, error) {
	var response DeleteResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteResponse struct {
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

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

	// Specifies the listening port.
	Port int `json:"port,omitempty"`

	// Specifies the cookie timeout duration (minutes).The value ranges from 1 to 1440. This parameter is valid when session_sticky is true and sticky_session_type is insert.
	CookieTimeout int `json:"cookie_timeout,omitempty"`

	// Specifies the status of the load balancer.Optional values:false: The load balancer is disabled.true: The load balancer is running properly.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the quantity of backend ECSs.
	MemberNumber int `json:"member_number,omitempty"`

	// Specifies the health check task ID.
	HealthcheckId string `json:"healthcheck_id,omitempty"`

	// Specifies whether to enable the sticky session feature. The feature is enabled when the value is true. This parameter is valid only when protocol is set to HTTP.
	SessionSticky bool `json:"session_sticky,omitempty"`

	// Specifies the load balancing algorithm for the listener.
	LbAlgorithm string `json:"lb_algorithm,omitempty"`

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

	// Specifies the SSL protocol standard supported by a listener.This parameter is used for enabling specified encryption protocols and valid only when the value of protocol is set to HTTPS or SSL.The value is TLSv1.2 or TLSv1.2 TLSv1.1 TLSv1. The default value is TLSv1.2.
	SslProtocols string `json:"ssl_protocols,omitempty"`

	// Specifies the cipher suite of an encryption protocol. This parameter is valid only when the value of protocol is set to HTTPS or SSL.The value is Default, Extended, or Strict.The value of Default is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256.The value of Extended is ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:DHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:AES128-SHA:AES256-SHA:DHE-DSS-AES128-SHA:CAMELLIA128-SHA:EDH-RSA-DES-CBC3-SHA:DES-CBC3-SHA:ECDHE-RSA-RC4-SHA:RC4-SHA:DHE-RSA-AES256-SHA:DHE-DSS-AES256-SHA:DHE-RSA-CAMELLIA256-SHA:DHE-DSS-CAMELLIA256-SHA:CAMELLIA256-SHA:EDH-DSS-DES-CBC3-SHA:DHE-RSA-CAMELLIA128-SHA:DHE-DSS-CAMELLIA128-SHA.The value of Strict is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256.The default value is Default. The value can only be set to Extended if the value of ssl_protocols is set to TLSv1.2 TLSv1.1 TLSv1.
	SslCiphers string `json:"ssl_ciphers,omitempty"`
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
	}
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*UpdateResponse, error) {
	var response UpdateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateResponse struct {

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

	// Specifies the listening port.
	Port int `json:"port,omitempty"`

	// Specifies the cookie timeout duration (minutes).The value ranges from 1 to 1440. This parameter is valid when session_sticky is true and sticky_session_type is insert.
	CookieTimeout int `json:"cookie_timeout,omitempty"`

	// Specifies the status of the load balancer.Optional values:false: The load balancer is disabled.true: The load balancer is running properly.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the health check task ID.
	HealthcheckId string `json:"healthcheck_id,omitempty"`

	// Specifies whether to enable the sticky session feature. The feature is enabled when the value is true. This parameter is valid only when protocol is set to HTTP.
	SessionSticky bool `json:"session_sticky,omitempty"`

	// Specifies the load balancing algorithm for the listener.
	LbAlgorithm string `json:"lb_algorithm,omitempty"`

	// Specifies the listener name.
	Name string `json:"name,omitempty"`

	// Specifies whether to maintain the TCP connection from the LVS node to the backend ECS after the ECS is removed. This parameter is valid when protocol is set to TCP.The value can be true or false.
	TcpDraining bool `json:"tcp_draining,omitempty"`

	// Specifies the timeout duration (minutes) for the TCP connection from the LVS node to the backend ECS after the ECS is removed.This parameter is valid when protocol is set to TCP and tcp_draining to true.The value ranges from 0 to 60.
	TcpDrainingTimeout int `json:"tcp_draining_timeout,omitempty"`

	// Specifies the ID of the SSL certificate for security authentication.This parameter is mandatory when protocol is set to HTTPS or SSL. Otherwise, null is displayed for this parameter.
	CertificateId string `json:"certificate_id,omitempty"`

	// Specifies the SSL certificate ID list if the value of protocol is HTTPS.This parameter is mandatory in SNI scenarios.
	Certificates []string `json:"certificates,omitempty"`
}
