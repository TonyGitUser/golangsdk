package floatingips

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type CreateFloatingIP struct {

	// Default Value:None. Specifies the floating IP address.
	FloatingIpAddress string `json:"floating_ip_address,omitempty"`

	// Specifies the external network ID. You can only use fixed external network. You can use?GET /v2.0/networks?router:external=True?or?GET /v2.0/networks?name={floating_network}?or run the?neutron net-external-listcommand to obtain information about the external network.
	FloatingNetworkId string `json:"floating_network_id,omitempty"`

	// Default Value:None. Specifies the port ID.
	PortId string `json:"port_id,omitempty"`

	// Default Value:None. Specifies the private IP address of the associated port. Content entered by users will be ignored.
	FixedIpAddress string `json:"fixed_ip_address,omitempty"`

	// Default Value:ID of the authenticated tenant. Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`
}

type UpdateFloatingIP struct {

	// Default Value:None. Specifies the port ID.
	PortId string `json:"port_id,omitempty"`

	// Default Value:None. Specifies the private IP address of the associated port. Content entered by users will be ignored.
	FixedIpAddress string `json:"fixed_ip_address,omitempty"`
}

type ListFloatingIP struct {

	// Default Value:DOWN. Specifies the network status. The value can be?ACTIVE,?DOWN, or?ERROR.
	Status string `json:"status,omitempty"`

	// Default Value:Automatically generated. Specifies the floating IP address ID.
	Id string `json:"id,omitempty"`

	// Default Value:None. Specifies the floating IP address.
	FloatingIpAddress string `json:"floating_ip_address,omitempty"`

	// Specifies the external network ID. You can only use fixed external network. You can use?GET /v2.0/networks?router:external=True?or?GET /v2.0/networks?name={floating_network}?or run the?neutron net-external-listcommand to obtain information about the external network.
	FloatingNetworkId string `json:"floating_network_id,omitempty"`

	// Default Value:None. Specifies the ID of the belonged router.
	RouterId string `json:"router_id,omitempty"`

	// Default Value:None. Specifies the port ID.
	PortId string `json:"port_id,omitempty"`

	// Default Value:None. Specifies the private IP address of the associated port. Content entered by users will be ignored.
	FixedIpAddress string `json:"fixed_ip_address,omitempty"`

	// Default Value:ID of the authenticated tenant. Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`
}

type CreateOpts struct {

	// Specifies the floating IP address list.
	Floatingip CreateFloatingIP `json:"floatingip,"`
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

func Delete(client *golangsdk.ServiceClient, floatingipId string) (r DeleteResult) {
	url := DeleteURL(client, floatingipId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, floatingipId string) (r GetResult) {
	url := GetURL(client, floatingipId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// 默认值:自动生成。 浮动IP地址的id。
	Id string `q:"id"`

	// 默认值:空。 浮动IP地址。
	FloatingIpAddress string `q:"floating_ip_address"`

	// 外部网络的id。 只能使用固定的外网，外部网络的信息请通过GET /v2.0/networks?router:external=True或GET /v2.0/networks?name={floating_network}或neutron net-external-list方式查询。
	FloatingNetworkId string `q:"floating_network_id"`

	// 默认值:空。 所属路由器id。
	RouterId string `q:"router_id"`

	// 默认值:空。 端口id
	PortId string `q:"port_id"`

	// 默认值:空。 关联端口的私有IP地址。 忽略用户输入。
	FixedIpAddress string `q:"fixed_ip_address"`

	// 默认值:当前认证租户。 租户id，只有管理员用户才允许指定非本租户的tenant_id。
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

// Queries all floating IP addresses accessible to the tenant submitting the request.
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
				postpagination.LinkedPageBase{PageResult: r, LinkPath: []string{"floatingips_links"}},
			}
			return p
		})
}

type UpdateOpts struct {

	// Specifies the floating IP address list.
	Floatingip UpdateFloatingIP `json:"floatingip,"`
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

func Update(client *golangsdk.ServiceClient, floatingipId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, floatingipId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
