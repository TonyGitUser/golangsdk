package listeners

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOpts struct {

	// Specifies the load balancer name.The value is a string of 1 to 64 characters that consist of Chinese characters, letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the listener.The value is a string of 0 to 128 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`

	// Specifies the ID of the load balancer to which the listener belongs.
	LoadbalancerId string `json:"loadbalancer_id,omitempty"`

	// Specifies the listening protocol used for layer 4 or 7.A listener using UDP is not allowed for a private network load balancer.The value can be HTTP, TCP, HTTPS, SSL, or UDP.
	Protocol string `json:"protocol,omitempty"`

	// Specifies the listening port.The value contains 1 to 65,535 characters.
	Port int `json:"port,omitempty"`

	// Specifies the backend protocol.If the value of protocol is UDP, the parameter value can only be UDP.If the value of protocol is SSL, the parameter value can only be TCP.The value can be HTTP, TCP, or UDP.
	BackendProtocol string `json:"backend_protocol,omitempty"`

	// The value contains 1 to 65,535 characters.The value ranges from 1 to 65535.
	BackendPort int `json:"backend_port,omitempty"`

	// Specifies the load balancing algorithm.The value can be roundrobin, leastconn, or source.
	LbAlgorithm string `json:"lb_algorithm,omitempty"`

	// Specifies whether to enable the sticky session feature.The value can be true or false. The feature is enabled when the value is true.If the value of protocol is SSL, the sticky session is not supported and the parameter is invalid.If the value of protocol is HTTP, HTTPS, or TCP and that of lb_algorithm is not roundrobin, the parameter value can only be false.
	SessionSticky bool `json:"session_sticky,omitempty"`

	// Specifies the cookie processing method. The value is insert.insert indicates that the cookie is inserted by the load balancer. This parameter is valid when protocol is set to HTTP and session_sticky to true. The default value is insert. This parameter is invalid when protocol is set to TCP, SSL, or UDP, which means the parameter is empty.
	StickySessionType string `json:"sticky_session_type,omitempty"`

	// Specifies the cookie timeout duration (minutes). This parameter is valid when protocol is set to HTTP, session_sticky to true, and sticky_session_type to insert. This parameter is invalid when protocol is set to TCP, UDP, or SSL.The value ranges from 1 to 1440.
	CookieTimeout int `json:"cookie_timeout,omitempty"`

	// Specifies the TCP timeout period (minutes). This parameter is valid when protocol is set to TCP.The value ranges from 1 to 1440.
	TcpTimeout int `json:"tcp_timeout,omitempty"`

	// Specifies whether to maintain the TCP connection to the backend ECS after the ECS is deleted. This parameter is valid when protocol is set to TCP.The value can be true or false.
	TcpDraining bool `json:"tcp_draining,omitempty"`

	// Specifies the timeout duration (minutes) for the TCP connection to the backend ECS after the ECS is deleted.This parameter is valid when protocol is set to TCP and tcp_draining to true.The value ranges from 0 to 60.
	TcpDrainingTimeout int `json:"tcp_draining_timeout,omitempty"`

	// Specifies the certificate ID. This parameter is mandatory when protocol is set to HTTPS or SSL.The value can be obtained from the details about the SSL certificate.
	CertificateId string `json:"certificate_id,omitempty"`

	// Specifies the SSL certificate ID list if the value of protocol is HTTPS.This parameter is mandatory in SNI scenarios.This parameter is valid only when the load balancer is a public network load balancer.
	Certificates []string `json:"certificates,omitempty"`

	// Specifies the UDP session timeout duration (minutes). This parameter is valid when protocol is set to UDP.The value ranges from 1 to 1440.
	UdpTimeout int `json:"udp_timeout,omitempty"`

	// Specifies the SSL protocol standard supported by a listener.This parameter is used for enabling specified encryption protocols and valid only when the value of protocol is set to HTTPS or SSL.The value is TLSv1.2 or TLSv1.2 TLSv1.1 TLSv1. The default value is TLSv1.2.
	SslProtocols string `json:"ssl_protocols,omitempty"`

	// Specifies the cipher suite of an encryption protocol. This parameter is valid only when the value of protocol is set to HTTPS or SSL.The value is Default, Extended, or Strict.The value of Default is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256.The value of Extended is ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:DHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:AES128-SHA:AES256-SHA:DHE-DSS-AES128-SHA:CAMELLIA128-SHA:EDH-RSA-DES-CBC3-SHA:DES-CBC3-SHA:ECDHE-RSA-RC4-SHA:RC4-SHA:DHE-RSA-AES256-SHA:DHE-DSS-AES256-SHA:DHE-RSA-CAMELLIA256-SHA:DHE-DSS-CAMELLIA256-SHA:CAMELLIA256-SHA:EDH-DSS-DES-CBC3-SHA:DHE-RSA-CAMELLIA128-SHA:DHE-DSS-CAMELLIA128-SHA.The value of Strict is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256.The default value is Default. The value can only be set to Extended if the value of ssl_protocols is set to TLSv1.2 TLSv1.1 TLSv1.
	SslCiphers string `json:"ssl_ciphers,omitempty"`
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, tenantId string, opts CreateOptsBuilder) (r CreateResult) {
	b, _ := opts.ToCreateMap()

	_, r.Err = client.Post(CreateURL(client, tenantId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, tenantId string, listenerId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, listenerId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, listenerId string) (r GetResult) {
	url := GetURL(client, tenantId, listenerId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies the ID of the load balancer to which the listener belongs.
	LoadbalancerId string `q:"loadbalancer_id"`

	// Specifies the SSL protocol standard supported by a listener.This parameter is used for enabling specified encryption protocols and valid only when the value of protocol is set to HTTPS or SSL.The value is TLSv1.2 or TLSv1.2 TLSv1.1 TLSv1. The default value is TLSv1.2.
	SslProtocols string `q:"ssl_protocols"`

	// Specifies the cipher suite of an encryption protocol. This parameter is valid only when the value of protocol is set to HTTPS or SSL.The value is Default, Extended, or Strict.The value of Default is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256.The value of Extended is ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:DHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:AES128-SHA:AES256-SHA:DHE-DSS-AES128-SHA:CAMELLIA128-SHA:EDH-RSA-DES-CBC3-SHA:DES-CBC3-SHA:ECDHE-RSA-RC4-SHA:RC4-SHA:DHE-RSA-AES256-SHA:DHE-DSS-AES256-SHA:DHE-RSA-CAMELLIA256-SHA:DHE-DSS-CAMELLIA256-SHA:CAMELLIA256-SHA:EDH-DSS-DES-CBC3-SHA:DHE-RSA-CAMELLIA128-SHA:DHE-DSS-CAMELLIA128-SHA.The value of Strict is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256.The default value is Default. The value can only be set to Extended if the value of ssl_protocols is set to TLSv1.2 TLSv1.1 TLSv1.
	SslCiphers string `q:"ssl_ciphers"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, tenantId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId)
	if opts != nil {
		query, _ := opts.ToListQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the load balancer name.The value is a string of 1 to 64 characters that consist of Chinese characters, letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the listener.The value is a string of 0 to 128 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`

	// Specifies the listening port.The value contains 1 to 65,535 characters.
	Port int `json:"port,omitempty"`

	// The value contains 1 to 65,535 characters.The value ranges from 1 to 65535.
	BackendPort int `json:"backend_port,omitempty"`

	// Specifies the load balancing algorithm.The value can be roundrobin, leastconn, or source.
	LbAlgorithm string `json:"lb_algorithm,omitempty"`

	// Specifies the TCP timeout period (minutes). This parameter is valid when protocol is set to TCP.The value ranges from 1 to 1440.
	TcpTimeout int `json:"tcp_timeout,omitempty"`

	// Specifies whether to maintain the TCP connection to the backend ECS after the ECS is deleted. This parameter is valid when protocol is set to TCP.The value can be true or false.
	TcpDraining bool `json:"tcp_draining,omitempty"`

	// Specifies the timeout duration (minutes) for the TCP connection to the backend ECS after the ECS is deleted.This parameter is valid when protocol is set to TCP and tcp_draining to true.The value ranges from 0 to 60.
	TcpDrainingTimeout int `json:"tcp_draining_timeout,omitempty"`

	// Specifies the UDP session timeout duration (minutes). This parameter is valid when protocol is set to UDP.The value ranges from 1 to 1440.
	UdpTimeout int `json:"udp_timeout,omitempty"`

	// Specifies the SSL protocol standard supported by a listener.This parameter is used for enabling specified encryption protocols and valid only when the value of protocol is set to HTTPS or SSL.The value is TLSv1.2 or TLSv1.2 TLSv1.1 TLSv1. The default value is TLSv1.2.
	SslProtocols string `json:"ssl_protocols,omitempty"`

	// Specifies the cipher suite of an encryption protocol. This parameter is valid only when the value of protocol is set to HTTPS or SSL.The value is Default, Extended, or Strict.The value of Default is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256.The value of Extended is ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:DHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:AES128-SHA:AES256-SHA:DHE-DSS-AES128-SHA:CAMELLIA128-SHA:EDH-RSA-DES-CBC3-SHA:DES-CBC3-SHA:ECDHE-RSA-RC4-SHA:RC4-SHA:DHE-RSA-AES256-SHA:DHE-DSS-AES256-SHA:DHE-RSA-CAMELLIA256-SHA:DHE-DSS-CAMELLIA256-SHA:CAMELLIA256-SHA:EDH-DSS-DES-CBC3-SHA:DHE-RSA-CAMELLIA128-SHA:DHE-DSS-CAMELLIA128-SHA.The value of Strict is ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256.The default value is Default. The value can only be set to Extended if the value of ssl_protocols is set to TLSv1.2 TLSv1.1 TLSv1.
	SslCiphers string `json:"ssl_ciphers,omitempty"`

	// Specifies the certificate ID. This parameter is mandatory when protocol is set to HTTPS or SSL.The value can be obtained from the details about the SSL certificate.
	CertificateId string `json:"certificate_id,omitempty"`

	// Specifies the SSL certificate ID list if the value of protocol is HTTPS.This parameter is mandatory in SNI scenarios.This parameter is valid only when the load balancer is a public network load balancer.
	Certificates []string `json:"certificates,omitempty"`
}

type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *golangsdk.ServiceClient, tenantId string, listenerId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, listenerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
