package subnets

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type CreateSubnet struct {

	// Default Value:Automatically generated. Specifies the subnet ID.
	Id string `json:"id,omitempty"`

	// Default Value:None. Specifies the subnet name.
	Name string `json:"name,omitempty"`

	// Specifies the Internet Protocol (IP) version. Only IPv4 is supported.
	IpVersion int `json:"ip_version,omitempty"`

	// Specifies the IPv6 addressing mode. This attribute is not supported.
	Ipv6AddressMode string `json:"ipv6_address_mode,omitempty"`

	// Specifies the IPv6 route broadcast mode. This attribute is not supported.
	Ipv6RaMode string `json:"ipv6_ra_mode,omitempty"`

	// Specifies the ID of the network to which the subnet belongs.
	NetworkId string `json:"network_id,omitempty"`

	// Specifies the CIDR format. Only the addresses in the 10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16 network segments are supported. In addition, the subnet mask cannot be greater than 28.
	Cidr string `json:"cidr,omitempty"`

	// Default Value:First IP address in a CIDR block. The gateway IP address cannot conflict with IP addresses configured for?allocation_pools. (If the parameter value is changed, this change does not take effect in the L3 plug-in delivered with FusionSphere OpenStack V100R006C10.)
	GatewayIp string `json:"gateway_ip,omitempty"`

	// Default Value:All IP addresses in a CIDR block excepting the gateway and broadcast addresses. Specifies the available IP address pool. For details about the?allocation_pool?object, see?Table 2. For example, [ { "start": "10.0.0.2", "end": "10.0.0.251"} ] The first and the last four IP addresses in each subnet are the ones reserved by the system. For example, in subnet?192.168.1.0/24, IP addresses 192.168.1.0, 192.168.1.252, 192.168.1.253, 192.168.1.254, and 192.168.1.255 are reserved by the system. By default, the IP addresses reserved by the system are not in the IP address pool specified by?allocation_pool. When updating an IP address pool, the?allocation_poolvalue can contain neither gateway nor broadcast IP addresses.
	AllocationPools []AllocationPool `json:"allocation_pools,omitempty"`

	// Default Value:Empty list. Specifies the DNS server address. For example, "dns_nameservers": ["8.8.8.8","8.8.4.4"].
	DnsNameservers []string `json:"dns_nameservers,omitempty"`

	// Default Value:Empty list. Specifies the static VM routes. For details, see the?host_route?object. Static routes are not supported, and entered information will be ignored.
	HostRoutes []HostRoute `json:"host_routes,omitempty"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Default Value:TRUE. Specifies whether to enable the DHCP function. Value?false?indicates that the DHCP function is not enabled. The value can only be?true.
	EnableDhcp bool `json:"enable_dhcp,omitempty"`
}

type UpdateSubnet struct {

	// Default Value:Automatically generated. Specifies the subnet ID.
	Id string `json:"id,omitempty"`

	// Default Value:None. Specifies the subnet name.
	Name string `json:"name,omitempty"`

	// Default Value:First IP address in a CIDR block. The gateway IP address cannot conflict with IP addresses configured for?allocation_pools. (If the parameter value is changed, this change does not take effect in the L3 plug-in delivered with FusionSphere OpenStack V100R006C10.)
	GatewayIp string `json:"gateway_ip,omitempty"`

	// Default Value:Empty list. Specifies the DNS server address. For example, "dns_nameservers": ["8.8.8.8","8.8.4.4"].
	DnsNameservers []string `json:"dns_nameservers,omitempty"`

	// Default Value:Empty list. Specifies the static VM routes. For details, see the?host_route?object. Static routes are not supported, and entered information will be ignored.
	HostRoutes []HostRoute `json:"host_routes,omitempty"`

	// Default Value:TRUE. Specifies whether to enable the DHCP function. Value?false?indicates that the DHCP function is not enabled. The value can only be?true.
	EnableDhcp bool `json:"enable_dhcp,omitempty"`
}

type ListSubnet struct {

	// Default Value:Automatically generated. Specifies the subnet ID.
	Id string `json:"id,omitempty"`

	// Default Value:None. Specifies the subnet name.
	Name string `json:"name,omitempty"`

	// Specifies the Internet Protocol (IP) version. Only IPv4 is supported.
	IpVersion int `json:"ip_version,omitempty"`

	// Specifies the IPv6 addressing mode. This attribute is not supported.
	Ipv6AddressMode string `json:"ipv6_address_mode,omitempty"`

	// Specifies the IPv6 route broadcast mode. This attribute is not supported.
	Ipv6RaMode string `json:"ipv6_ra_mode,omitempty"`

	// Specifies the ID of the network to which the subnet belongs.
	NetworkId string `json:"network_id,omitempty"`

	// Specifies the CIDR format. Only the addresses in the 10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16 network segments are supported. In addition, the subnet mask cannot be greater than 28.
	Cidr string `json:"cidr,omitempty"`

	// Default Value:First IP address in a CIDR block. The gateway IP address cannot conflict with IP addresses configured for?allocation_pools. (If the parameter value is changed, this change does not take effect in the L3 plug-in delivered with FusionSphere OpenStack V100R006C10.)
	GatewayIp string `json:"gateway_ip,omitempty"`

	// Default Value:All IP addresses in a CIDR block excepting the gateway and broadcast addresses. Specifies the available IP address pool. For details about the?allocation_pool?object, see?Table 2. For example, [ { "start": "10.0.0.2", "end": "10.0.0.251"} ] The first and the last four IP addresses in each subnet are the ones reserved by the system. For example, in subnet?192.168.1.0/24, IP addresses 192.168.1.0, 192.168.1.252, 192.168.1.253, 192.168.1.254, and 192.168.1.255 are reserved by the system. By default, the IP addresses reserved by the system are not in the IP address pool specified by?allocation_pool. When updating an IP address pool, the?allocation_poolvalue can contain neither gateway nor broadcast IP addresses.
	AllocationPools []AllocationPool `json:"allocation_pools,omitempty"`

	// Default Value:Empty list. Specifies the DNS server address. For example, "dns_nameservers": ["8.8.8.8","8.8.4.4"].
	DnsNameservers []string `json:"dns_nameservers,omitempty"`

	// Default Value:Empty list. Specifies the static VM routes. For details, see the?host_route?object. Static routes are not supported, and entered information will be ignored.
	HostRoutes []HostRoute `json:"host_routes,omitempty"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Default Value:TRUE. Specifies whether to enable the DHCP function. Value?false?indicates that the DHCP function is not enabled. The value can only be?true.
	EnableDhcp bool `json:"enable_dhcp,omitempty"`
}

type AllocationPool struct {

	// Specifies the start IP address of a network pool.
	Start string `json:"start,omitempty"`

	// Specifies the end IP address of a network pool.
	End string `json:"end,omitempty"`
}

type HostRoute struct {

	// Specifies the destination subnet of a route.
	Destination string `json:"destination,omitempty"`

	// Specifies the next-hop IP address of a route.
	Nexthop string `json:"nexthop,omitempty"`
}

type CreateOpts struct {

	// Specifies the subnet objects.
	Subnet CreateSubnet `json:"subnet,"`
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

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, _ := opts.ToCreateMap()

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, subnetId string) (r DeleteResult) {
	url := DeleteURL(client, subnetId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, subnetId string) (r GetResult) {
	url := GetURL(client, subnetId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Default Value:None. Specifies the subnet name.
	Name string `q:"name"`

	// Specifies the Internet Protocol (IP) version. Only IPv4 is supported.
	IpVersion int `q:"ip_version"`

	// Specifies the ID of the network to which the subnet belongs.
	NetworkId string `q:"network_id"`

	// Specifies the CIDR format. Only the addresses in the 10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16 network segments are supported. In addition, the subnet mask cannot be greater than 28.
	Cidr string `q:"cidr"`

	// Default Value:First IP address in a CIDR block. The gateway IP address cannot conflict with IP addresses configured for?allocation_pools. (If the parameter value is changed, this change does not take effect in the L3 plug-in delivered with FusionSphere OpenStack V100R006C10.)
	GatewayIp string `q:"gateway_ip"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `q:"tenant_id"`

	// Default Value:TRUE. Specifies whether to enable the DHCP function. Value?false?indicates that the DHCP function is not enabled. The value can only be?true.
	EnableDhcp bool `q:"enable_dhcp"`

	// Specifies the resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`

	// Specifies the page direction. The value can be True or False.
	PageReverse bool `q:"page_reverse"`

	// The offset+1 record will be first displayed.
	Offset string `q:"offset"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// This interface is used to query subnets using search criteria and to display the subnets in a list.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) postpagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return postpagination.Pager{Err: err}
		}
		url += query
	}

	return postpagination.NewPager(client, "GET", url, nil,
		func(r postpagination.PageResult) postpagination.Page {
			p := ListPage{
				postpagination.LinkedPageBase{PageResult: r, LinkPath: []string{"subnets_links"}},
			}
			return p
		})
}

type UpdateOpts struct {

	// Specifies the subnet objects.
	Subnet UpdateSubnet `json:"subnet,"`
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

func Update(client *golangsdk.ServiceClient, subnetId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, subnetId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
