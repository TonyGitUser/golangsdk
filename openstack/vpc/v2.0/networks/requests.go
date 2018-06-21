package networks

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type CreateNetwork struct {

	// Default Value:None. Specifies the network name. The name cannot be?admin_external_net.
	Name string `json:"name,omitempty"`

	// Default Value:FALSE. Specifies whether the network is an external network. This is an extended attribute. This attribute is for administrators only. Tenants cannot configure or update this attribute and can only query it.
	Routerexternal bool `json:"router:external,omitempty"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Default Value:FALSE. Specifies whether the network can be shared by different tenants. This attribute is for administrators only. Tenants cannot configure or update this attribute and can only query it.
	Shared bool `json:"shared,omitempty"`

	// Specifies the physical network used by this network. This is an extended attribute. This attribute is available only to administrators.
	ProviderphysicalNetwork string `json:"provider:physical_network,omitempty"`

	// Specifies the network type. Only the VXLAN and GENEVE networks are supported. This is an extended attribute. This attribute is available only to administrators. Only GENEVE tenants can perform operations on this attribute.
	ProvidernetworkType string `json:"provider:network_type,omitempty"`

	// Specifies the network segment ID. The value is a VLAN ID for a VLAN network and is a VNI for a VXLAN network. This is an extended attribute.
	ProvidersegmentationId int `json:"provider:segmentation_id,omitempty"`

	// Default Value:TRUE. Specifies whether the security option is enabled for the port. If the option is not enabled, the security group and DHCP snooping settings of all VMs in the network do not take effect.
	PortSecurityEnabled bool `json:"port_security_enabled,omitempty"`
}

type UpdateNetwork struct {

	// Default Value:None. Specifies the network name. The name cannot be?admin_external_net.
	Name string `json:"name,omitempty"`

	// Default Value:FALSE. Specifies whether the network is an external network. This is an extended attribute. This attribute is for administrators only. Tenants cannot configure or update this attribute and can only query it.
	Routerexternal bool `json:"router:external,omitempty"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Default Value:FALSE. Specifies whether the network can be shared by different tenants. This attribute is for administrators only. Tenants cannot configure or update this attribute and can only query it.
	Shared bool `json:"shared,omitempty"`

	// Default Value:TRUE. Specifies whether the security option is enabled for the port. If the option is not enabled, the security group and DHCP snooping settings of all VMs in the network do not take effect.
	PortSecurityEnabled bool `json:"port_security_enabled,omitempty"`
}

type ListNetwork struct {

	// Default Value:ACTIVE. Specifies the network status. The value can be?ACTIVE,?BUILD,?DOWN, or?ERROR.
	Status string `json:"status,omitempty"`

	// Default Value:Empty list. Specifies IDs of the subnets associated with this network. The IDs are in a list. Only one subnet can be associated with each network.
	Subnets []string `json:"subnets,omitempty"`

	// Default Value:None. Specifies the network name. The name cannot be?admin_external_net.
	Name string `json:"name,omitempty"`

	// Default Value:FALSE. Specifies whether the network is an external network. This is an extended attribute. This attribute is for administrators only. Tenants cannot configure or update this attribute and can only query it.
	Routerexternal bool `json:"router:external,omitempty"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `json:"tenant_id,omitempty"`

	// Default Value:FALSE. Specifies whether the network can be shared by different tenants. This attribute is for administrators only. Tenants cannot configure or update this attribute and can only query it.
	Shared bool `json:"shared,omitempty"`

	// Default Value:Automatically generated. Specifies the network ID.
	Id string `json:"id,omitempty"`

	// Specifies the physical network used by this network. This is an extended attribute. This attribute is available only to administrators.
	ProviderphysicalNetwork string `json:"provider:physical_network,omitempty"`

	// Specifies the network type. Only the VXLAN and GENEVE networks are supported. This is an extended attribute. This attribute is available only to administrators. Only GENEVE tenants can perform operations on this attribute.
	ProvidernetworkType string `json:"provider:network_type,omitempty"`

	// Specifies the network segment ID. The value is a VLAN ID for a VLAN network and is a VNI for a VXLAN network. This is an extended attribute.
	ProvidersegmentationId int `json:"provider:segmentation_id,omitempty"`

	// Specifies the availability zones available to this network. The current version does not support cross-availability-zone network scheduling. An empty list is returned.
	AvailabilityZoneHints []string `json:"availability_zone_hints,omitempty"`

	// Specifies the availability zone of this network. An empty list is returned.
	AvailabilityZones []string `json:"availability_zones,omitempty"`

	// Default Value:TRUE. Specifies whether the security option is enabled for the port. If the option is not enabled, the security group and DHCP snooping settings of all VMs in the network do not take effect.
	PortSecurityEnabled bool `json:"port_security_enabled,omitempty"`
}

type CreateOpts struct {

	// Specifies the network list.
	Network CreateNetwork `json:"network,"`
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

func Delete(client *golangsdk.ServiceClient, networkId string) (r DeleteResult) {
	url := DeleteURL(client, networkId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, networkId string) (r GetResult) {
	url := GetURL(client, networkId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Default Value:ACTIVE. Specifies the network status. The value can be?ACTIVE,?BUILD,?DOWN, or?ERROR.
	Status string `q:"status"`

	// Default Value:None. Specifies the network name. The name cannot be?admin_external_net.
	Name string `q:"name"`

	// Default Value:TRUE. Specifies the administrative status. The value can only be?true.
	AdminStateUp bool `q:"admin_state_up"`

	// Specifies the tenant ID. Only the administrator can specify the tenant ID of other tenants.
	TenantId string `q:"tenant_id"`

	// Default Value:FALSE. Specifies whether the network can be shared by different tenants. This attribute is for administrators only. Tenants cannot configure or update this attribute and can only query it.
	Shared bool `q:"shared"`

	// Default Value:Automatically generated. Specifies the network ID.
	Id string `q:"id"`

	// Specifies the network type. Only the VXLAN and GENEVE networks are supported. This is an extended attribute. This attribute is available only to administrators. Only GENEVE tenants can perform operations on this attribute.
	ProvidernetworkType string `q:"provider:network_type"`

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

// Queries all networks accessible to the tenant submitting the request. A maximum of 2000 records can be returned for each query operation.
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
				postpagination.LinkedPageBase{PageResult: r, LinkPath: []string{"networks_links"}},
			}
			return p
		})
}

type UpdateOpts struct {

	// Specifies the network list.
	Network UpdateNetwork `json:"network,"`
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
