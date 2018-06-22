package ports

import (
	"github.com/huaweicloud/golangsdk"
)

type CreatePort struct {

	// Specifies the port name. The value can contain no more than 255 characters. This parameter is left blank by default.
	Name string `json:"name,omitempty"`

	// Specifies the ID of the network to which the port belongs. The network ID must be a real one in the network environment.
	NetworkId string `json:"network_id,omitempty"`

	// Specifies the administrative state of the port. The value can only be?true, and the default value is?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the port IP address. A port supports only one fixed IP address that cannot be changed.
	FixedIps []FixedIp `json:"fixed_ips,omitempty"`

	// Specifies the ID of the tenant. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Specifies the UUID of the security group, for example,? "security_groups": ["a0608cbf-d047-4f54-8b28-cd7b59853fff"]. This attribute is extended.
	SecurityGroups []string `json:"security_groups,omitempty"`

	// Specifies a set of zero or more extra DHCP option pairs. An option pair consists of an option value and name. This attribute is extended.
	ExtraDhcpOpts []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

type UpdatePort struct {

	// Specifies the port name. The value can contain no more than 255 characters. This parameter is left blank by default.
	Name string `json:"name,omitempty"`

	// Specifies the UUID of the security group, for example,? "security_groups": ["a0608cbf-d047-4f54-8b28-cd7b59853fff"]. This attribute is extended.
	SecurityGroups []string `json:"security_groups,omitempty"`

	// Specifies a set of zero or more extra DHCP option pairs. An option pair consists of an option value and name. This attribute is extended.
	ExtraDhcpOpts []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

type FixedIp struct {

	// 1、功能说明：所属子网ID 2、约束：不支持更新
	SubnetId string `json:"subnet_id,omitempty"`

	// 1、功能说明：端口IP地址 2、约束：不支持更新
	IpAddress string `json:"ip_address,omitempty"`
}

type ExtraDHCPOpt struct {

	// 功能说明：Option名称
	OptName string `json:"opt_name,omitempty"`

	// 功能说明：Option值
	OptValue string `json:"opt_value,omitempty"`
}

type Port struct {

	// Specifies the port ID, which uniquely identifies the port.
	Id string `json:"id,omitempty"`

	// Specifies the port name. The value can contain no more than 255 characters. This parameter is left blank by default.
	Name string `json:"name,omitempty"`

	// Specifies the ID of the network to which the port belongs. The network ID must be a real one in the network environment.
	NetworkId string `json:"network_id,omitempty"`

	// Specifies the administrative state of the port. The value can only be?true, and the default value is?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the port MAC address. The system automatically sets this parameter, and you are not allowed to configure the parameter value.
	MacAddress string `json:"mac_address,omitempty"`

	// Specifies the port IP address. A port supports only one fixed IP address that cannot be changed.
	FixedIps []FixedIp `json:"fixed_ips,omitempty"`

	// Specifies the ID of the device to which the port belongs. The system automatically sets this parameter, and you are not allowed to configure or change the parameter value.
	DeviceId string `json:"device_id,omitempty"`

	// Specifies the belonged device, which can be the DHCP server, router, load balancers, or Nova. The system automatically sets this parameter, and you are not allowed to configure or change the parameter value.
	DeviceOwner string `json:"device_owner,omitempty"`

	// Specifies the ID of the tenant. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Specifies the status of the port. The value can be?ACTIVE,?BUILD, or?DOWN.
	Status string `json:"status,omitempty"`

	// Specifies the UUID of the security group. This attribute is extended.
	SecurityGroups []string `json:"security_groups,omitempty"`

	// 1. Specifies a set of zero or more allowed address pairs. An address pair consists of an IP address and MAC address. This attribute is extended. For details, see parameter?allow_address_pair. 2. The IP address cannot be?0.0.0.0. 3. Configure an independent security group for the port if a large CIDR block (subnet mask less than 24) is configured for parameter?allowed_address_pairs.
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	// Specifies a set of zero or more extra DHCP option pairs. An option pair consists of an option value and name. This attribute is extended.
	ExtraDhcpOpts []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`

	// Specifies the interface type of the port. The value can be?ovs,?hw_veb, or others. This attribute is extended. This parameter is visible only to administrators.
	BindingvifType string `json:"binding:vif_type,omitempty"`

	// Specifies the host ID. This parameter is visible only to administrators.
	BindinghostId string `json:"binding:host_id,omitempty"`

	// Specifies the type of the bound vNIC. The value can be?normal?or?direct. Parameter?normal?indicates software switching. Parameter?direct?indicates SR-IOV PCIe passthrough, which is not supported.
	BindingvnicType string `json:"binding:vnic_type,omitempty"`
}

type AllowedAddressPair struct {

	// 功能说明：IP地址 约束：不支持0.0.0.0 如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。
	IpAddress string `json:"ip_address,omitempty"`

	// 功能说明：MAC地址
	MacAddress string `json:"mac_address,omitempty"`
}

type CreateOpts struct {

	// Specifies the port object.
	Port CreatePort `json:"port,"`
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

func Delete(client *golangsdk.ServiceClient, portId string) (r DeleteResult) {
	url := DeleteURL(client, portId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, portId string) (r GetResult) {
	url := GetURL(client, portId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies that the port ID is used as the filter.
	Id string `q:"id"`

	// Specifies that the port name is used as the filter.
	Name string `q:"name"`

	// Specifies that the administrative state is used as the filter.
	AdminStateUp bool `q:"admin_state_up"`

	// Specifies that the network ID is used as the filter.
	NetworkId string `q:"network_id"`

	// Specifies that the MAC address is used as the filter.
	MacAddress string `q:"mac_address"`

	// Specifies that the device ID is used as the filter.
	DeviceId string `q:"device_id"`

	// Specifies that the device owner is used as the filter.
	DeviceOwner string `q:"device_owner"`

	// Specifies that the status is used as the filter.
	Status string `q:"status"`

	// Specifies the resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client)
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

	// Specifies the port object.
	Port UpdatePort `json:"port,"`
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

func Update(client *golangsdk.ServiceClient, portId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, portId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
