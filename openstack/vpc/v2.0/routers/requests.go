package routers

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type CreateRouter struct {

	// Default Value:None. Specifies the router name. The name can contain only digits, letters, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Specifies the external gateway information. This is an extended attribute. For details, see?Table 2.
	ExternalGatewayInfo ExternalGatewayInfo `json:"external_gateway_info,omitempty"`

	// Default Value:FALSE. Specifies the distributed deployment mode. Administrator permission required.
	Distributed bool `json:"distributed,omitempty"`

	// Default Value:FALSE. Specifies the HA deployment mode. Administrator permission required.
	Ha bool `json:"ha,omitempty"`
}

type UpdateRouter struct {

	// Default Value:None. Specifies the router name. The name can contain only digits, letters, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies information about the routes. This is an extended attribute. For details, see?Table 3.
	Routes []Route `json:"routes,omitempty"`

	// Default Value:FALSE. Specifies the distributed deployment mode. Administrator permission required.
	Distributed bool `json:"distributed,omitempty"`
}

type ListRouter struct {

	// Default Value:Automatically generated. Specifies the router ID.
	Id string `json:"id,omitempty"`

	// Default Value:None. Specifies the router name. The name can contain only digits, letters, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the router status. The value can be?ACTIVE,?DOWN, or?ERROR.
	Status string `json:"status,omitempty"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Specifies the external gateway information. This is an extended attribute. For details, see?Table 2.
	ExternalGatewayInfo ExternalGatewayInfo `json:"external_gateway_info,omitempty"`

	// Specifies information about the routes. This is an extended attribute. For details, see?Table 3.
	Routes []Route `json:"routes,omitempty"`

	// Default Value:FALSE. Specifies the distributed deployment mode. Administrator permission required.
	Distributed bool `json:"distributed,omitempty"`

	// Default Value:FALSE. Specifies the HA deployment mode. Administrator permission required.
	Ha bool `json:"ha,omitempty"`
}

type ExternalGatewayInfo struct {

	// Specifies the UUID of the external network. You can use GET /v2.0/networks?router:external=True or run the neutron net-external-list command to query information about the external network.
	NetworkId string `json:"network_id,omitempty"`

	// Specifies whether the SNAT function is enabled. The default value is false.
	EnableSnat bool `json:"enable_snat,omitempty"`
}

type Route struct {

	// Specifies the IP address segment. You can only configure the default route, and its value can only be 0.0.0.0/0.
	Destination string `json:"destination,omitempty"`

	// Specifies the next hop IP address. The IP address can only be one in the subnet associated with the router.
	Nexthop string `json:"nexthop,omitempty"`
}

type AddInterfaceOpts struct {

	// Specifies the subnet ID. Either subnet_id or port_id is used. Use the gateway IP address of the subnet to create a router interface.
	SubnetId string `json:"subnet_id,omitempty"`

	// Specifies the port ID. Either subnet_id or port_id is used. Use the port IP address to create a router interface.
	PortId string `json:"port_id,omitempty"`
}

type AddInterfaceOptsBuilder interface {
	ToAddInterfaceMap() (map[string]interface{}, error)
}

func (opts AddInterfaceOpts) ToAddInterfaceMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func AddInterface(client *golangsdk.ServiceClient, routerId string, opts AddInterfaceOptsBuilder) (r AddInterfaceResult) {
	b, _ := opts.ToAddInterfaceMap()

	_, r.Err = client.Put(AddInterfaceURL(client, routerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type CreateOpts struct {

	// Specifies the router list.
	Router CreateRouter `json:"router,"`
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

func Delete(client *golangsdk.ServiceClient, routerId string) (r DeleteResult) {
	url := DeleteURL(client, routerId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, routerId string) (r GetResult) {
	url := GetURL(client, routerId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Default Value:Automatically generated. Specifies the router ID.
	Id string `q:"id"`

	// Default Value:None. Specifies the router name. The name can contain only digits, letters, underscores (_), and hyphens (-).
	Name string `q:"name"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `q:"admin_state_up"`

	// Specifies the router status. The value can be?ACTIVE,?DOWN, or?ERROR.
	Status string `q:"status"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
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

// Queries all routers accessible to the tenant submitting the request.
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
				postpagination.LinkedPageBase{PageResult: r, LinkPath: []string{"routers_links"}},
			}
			return p
		})
}

type RemoveInterfaceOpts struct {

	// Specifies the subnet ID. Either subnet_id or port_id is used. Use the gateway IP address of the subnet to create a router interface.
	SubnetId string `json:"subnet_id,"`

	// Specifies the port ID. Either subnet_id or port_id is used. Use the port IP address to create a router interface.
	PortId string `json:"port_id,"`
}

type RemoveInterfaceOptsBuilder interface {
	ToRemoveInterfaceMap() (map[string]interface{}, error)
}

func (opts RemoveInterfaceOpts) ToRemoveInterfaceMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func RemoveInterface(client *golangsdk.ServiceClient, routerId string, opts RemoveInterfaceOptsBuilder) (r RemoveInterfaceResult) {
	b, _ := opts.ToRemoveInterfaceMap()

	_, r.Err = client.Put(RemoveInterfaceURL(client, routerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the router list.
	Router UpdateRouter `json:"router,"`
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

func Update(client *golangsdk.ServiceClient, routerId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, routerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
