package securitygrouprules

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type CreateSecurityGroup struct {

	// Default Value:None. Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Default Value:None. Specifies the security group name.
	Name string `json:"name,omitempty"`

	// Default Value:None. Provides supplementary information about the security group.
	Description string `json:"description,omitempty"`
}

type UpdateSecurityGroup struct {

	// Default Value:None. Specifies the security group name.
	Name string `json:"name,omitempty"`

	// Default Value:None. Provides supplementary information about the security group.
	Description string `json:"description,omitempty"`
}

type ListSecurityGroup struct {

	// Default Value:None. Specifies the security group ID.
	Id string `json:"id,omitempty"`

	// Default Value:None. Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Default Value:None. Specifies the security group name.
	Name string `json:"name,omitempty"`

	// Default Value:None. Provides supplementary information about the security group.
	Description string `json:"description,omitempty"`

	// Default Value:None. Specifies the security group rule list. For details, see?Table 2.
	SecurityGroupRules []ListSecurityGroupRule `json:"security_group_rules,omitempty"`
}

type CreateSecurityGroupRule struct {

	// Default Value:None. Provides supplementary information about the security group rule.
	Description string `json:"description,omitempty"`

	// Default Value:None. Specifies the ID of the belonged security group.
	SecurityGroupId string `json:"security_group_id,omitempty"`

	// Default Value:None. Specifies the peer ID of the belonged security group.
	RemoteGroupId string `json:"remote_group_id,omitempty"`

	// Default Value:None. Specifies the direction of the traffic for which the security group rule takes effect.
	Direction string `json:"direction,omitempty"`

	// Default Value:None. Specifies the peer IP address segment.
	RemoteIpPrefix string `json:"remote_ip_prefix,omitempty"`

	// Default Value:None. Specifies the protocol type or the IP protocol number.
	Protocol string `json:"protocol,omitempty"`

	// Default Value:None. Specifies the maximum port number. When ICMP is used, the value is the ICMP code.
	PortRangeMax int `json:"port_range_max,omitempty"`

	// Default Value:None. Specifies the minimum port number. If the ICMP protocol is used, this parameter indicates the ICMP type. When the TCP or UDP protocol is used, both?port_range_max?and?port_range_min?must be specified, and the?port_range_max?value must be greater than the?port_range_minvalue. When the ICMP protocol is used, if you specify the ICMP code (port_range_max), you must also specify the ICMP type (port_range_min).
	PortRangeMin int `json:"port_range_min,omitempty"`

	// Default Value:IPv4. Specifies the network type. Only IPv4 is supported.
	Ethertype string `json:"ethertype,omitempty"`

	// Default Value:None. Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`
}

type UpdateSecurityGroupRule struct {

	// Default Value:None. Provides supplementary information about the security group rule.
	Description string `json:"description,omitempty"`
}

type ListSecurityGroupRule struct {

	// Default Value:None. Specifies the security group rule ID.
	Id string `json:"id,omitempty"`

	// Default Value:None. Provides supplementary information about the security group rule.
	Description string `json:"description,omitempty"`

	// Default Value:None. Specifies the ID of the belonged security group.
	SecurityGroupId string `json:"security_group_id,omitempty"`

	// Default Value:None. Specifies the peer ID of the belonged security group.
	RemoteGroupId string `json:"remote_group_id,omitempty"`

	// Default Value:None. Specifies the direction of the traffic for which the security group rule takes effect.
	Direction string `json:"direction,omitempty"`

	// Default Value:None. Specifies the peer IP address segment.
	RemoteIpPrefix string `json:"remote_ip_prefix,omitempty"`

	// Default Value:None. Specifies the protocol type or the IP protocol number.
	Protocol string `json:"protocol,omitempty"`

	// Default Value:None. Specifies the maximum port number. When ICMP is used, the value is the ICMP code.
	PortRangeMax int `json:"port_range_max,omitempty"`

	// Default Value:None. Specifies the minimum port number. If the ICMP protocol is used, this parameter indicates the ICMP type. When the TCP or UDP protocol is used, both?port_range_max?and?port_range_min?must be specified, and the?port_range_max?value must be greater than the?port_range_minvalue. When the ICMP protocol is used, if you specify the ICMP code (port_range_max), you must also specify the ICMP type (port_range_min).
	PortRangeMin int `json:"port_range_min,omitempty"`

	// Default Value:IPv4. Specifies the network type. Only IPv4 is supported.
	Ethertype string `json:"ethertype,omitempty"`

	// Default Value:None. Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`
}

type CreateOpts struct {

	// Specifies the security group rule.
	SecurityGroupRule CreateSecurityGroupRule `json:"security_group_rule,"`
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

func Delete(client *golangsdk.ServiceClient, securityGroupsRulesId string) (r DeleteResult) {
	url := DeleteURL(client, securityGroupsRulesId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, securityGroupsRulesId string) (r GetResult) {
	url := GetURL(client, securityGroupsRulesId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Default Value:None. Specifies the ID of the belonged security group.
	SecurityGroupId string `q:"security_group_id"`

	// Default Value:None. Specifies the peer ID of the belonged security group.
	RemoteGroupId string `q:"remote_group_id"`

	// Default Value:None. Specifies the direction of the traffic for which the security group rule takes effect.
	Direction string `q:"direction"`

	// Default Value:None. Specifies the peer IP address segment.
	RemoteIpPrefix string `q:"remote_ip_prefix"`

	// Default Value:None. Specifies the protocol type or the IP protocol number.
	Protocol string `q:"protocol"`

	// Default Value:None. Specifies the maximum port number. When ICMP is used, the value is the ICMP code.
	PortRangeMax int `q:"port_range_max"`

	// Default Value:None. Specifies the minimum port number. If the ICMP protocol is used, this parameter indicates the ICMP type. When the TCP or UDP protocol is used, both?port_range_max?and?port_range_min?must be specified, and the?port_range_max?value must be greater than the?port_range_minvalue. When the ICMP protocol is used, if you specify the ICMP code (port_range_max), you must also specify the ICMP type (port_range_min).
	PortRangeMin int `q:"port_range_min"`

	// Default Value:IPv4. Specifies the network type. Only IPv4 is supported.
	Ethertype string `q:"ethertype"`

	// Default Value:None. Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `q:"tenant_id"`

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

// This interface is used to query security group rules using search criteria and to display the security group rules in a list.
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
				postpagination.LinkedPageBase{PageResult: r, LinkPath: []string{"security_group_rules_links"}},
			}
			return p
		})
}
