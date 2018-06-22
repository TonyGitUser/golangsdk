package securitygroups

import (
	"github.com/huaweicloud/golangsdk"
)

type SecurityGroup struct {

	// Specifies the security group name.
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the security group.
	Description string `json:"description,omitempty"`

	// Specifies the security group ID, which uniquely identifies the security group.
	Id string `json:"id,omitempty"`

	// Specifies the resource ID of the VPC to which the security group belongs.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies the default security group rule, which ensures that hosts in the security group can communicate with one another.
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules,omitempty"`
}

type CreateSecurityGroup struct {

	// Specifies the security group name.
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the security group.
	Description string `json:"description,omitempty"`

	// Specifies the resource ID of the VPC to which the security group belongs.
	VpcId string `json:"vpc_id,omitempty"`
}

type SecurityGroupRule struct {

	// Specifies the security group rule ID.
	Id string `json:"id,omitempty"`

	// Specifies the security group ID.
	SecurityGroupId string `json:"security_group_id,omitempty"`

	// Specifies the direction of access control. The value can be?egress?or?ingress.
	Direction string `json:"direction,omitempty"`

	// Specifies the version of the Internet Protocol. The value can be?IPv4?or?IPv6.
	Ethertype string `json:"ethertype,omitempty"`

	// Specifies the protocol type. If the parameter is left blank, the security group supports all types of protocols. The value can be?icmp,?tcp, or?udp.
	Protocol string `json:"protocol,omitempty"`

	// Specifies the start port. The value ranges from 1 to 65,535. The value must be less than or equal to the value of?port_range_max. An empty value indicates all ports. If?protocol?is?icmp, the value range is determined by the ICMP-port range relationship table provided in Appendix A.2.
	PortRangeMin int `json:"port_range_min,omitempty"`

	// Specifies the end port. The value ranges from 1 to 65,535. The value must be greater than or equal to the value of?port_range_min. An empty value indicates all ports. If?protocol?is?icmp, the value range is determined by the ICMP-port range relationship table provided in Appendix A.2.
	PortRangeMax int `json:"port_range_max,omitempty"`

	// Specifies the remote IP address. If the access control direction is set to?egress, the parameter specifies the source IP address. If the access control direction is set to?ingress, the parameter specifies the destination IP address. The parameter is exclusive with parameter?remote_group_id. The value can be in the CIDR format or IP addresses.
	RemoteIpPrefix string `json:"remote_ip_prefix,omitempty"`

	// Specifies the ID of the peer security group. The value is exclusive with parameter?remote_ip_prefix.
	RemoteGroupId string `json:"remote_group_id,omitempty"`
}

type CreateOpts struct {

	// Specifies the security group objects.
	SecurityGroup CreateSecurityGroup `json:"security_group,"`
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

func Delete(client *golangsdk.ServiceClient, tenantId string, securityGroupId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, securityGroupId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, securityGroupId string) (r GetResult) {
	url := GetURL(client, tenantId, securityGroupId)
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
