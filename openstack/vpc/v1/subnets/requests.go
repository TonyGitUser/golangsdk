package subnets

import (
	"github.com/huaweicloud/golangsdk"
)

type Subnet struct {

	// Specifies a resource ID in UUID format.
	Id string `json:"id,"`

	// Specifies the subnet name. The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,"`

	// Specifies the network segment on which the subnet resides. The value must be in CIDR format. The value must be within the CIDR block of the VPC. The subnet mask cannot be greater than 28.
	Cidr string `json:"cidr,"`

	// Specifies the gateway of the subnet. The value must be a valid IP address. The value must be an IP address in the subnet segment.
	GatewayIp string `json:"gateway_ip,"`

	// Specifies whether the DHCP function is enabled for the subnet. The value can be true or false. If this parameter is left blank, it is set to true by default.
	DhcpEnable bool `json:"dhcp_enable,omitempty"`

	// Specifies the IP address of DNS server 1 on the subnet. The value must be a valid IP address.
	PrimaryDns string `json:"primary_dns,omitempty"`

	// Specifies the IP address of DNS server 2 on the subnet. The value must be a valid IP address.
	SecondaryDns string `json:"secondary_dns,omitempty"`

	// Specifies the DNS server address list of a subnet. This field is required if you need to use more than two DNS servers. This parameter value is the superset of both DNS server address 1 and DNS server address 2.
	DnsList []string `json:"dnsList,omitempty"`

	// Identifies the availability zone (AZ) to which the subnet belongs. The value must be an existing AZ in the system.
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// Specifies the ID of the VPC to which the subnet belongs.
	VpcId string `json:"vpc_id,"`

	// Specifies the operations can be performed on security groups during subnet creation. This is a system default parameter. Users do not need to configure this parameter.
	PortSecurityEnable string `json:"port_security_enable,omitempty"`

	// Specifies the status of the subnet. The value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
	Status string `json:"status,"`

	// Specifies the network (Native OpenStack API) ID.
	NeutronNetworkId string `json:"neutron_network_id,"`

	// Specifies the subnet (Native OpenStack API) ID.
	NeutronSubnetId string `json:"neutron_subnet_id,"`
}

type CreateOpts struct {

	// Specifies the subnet objects.
	Subnet Subnet `json:"subnet,"`
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

func Delete(client *golangsdk.ServiceClient, tenantId string, vpcId string, subnetId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, vpcId, subnetId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, subnetId string) (r GetResult) {
	url := GetURL(client, tenantId, subnetId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies the resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page.
	Limit int `q:"limit"`

	// Specifies the VPC ID used as the query filter.
	VpcId string `q:"vpc_id"`
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

	// Specifies the subnet objects.
	Subnet Subnet `json:"subnet,"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, vpcId string, subnetId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, vpcId, subnetId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
