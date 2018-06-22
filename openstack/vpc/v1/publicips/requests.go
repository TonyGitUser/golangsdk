package publicips

import (
	"github.com/huaweicloud/golangsdk"
)

type PublicIP struct {

	// Specifies the ID of the elastic IP address, which uniquely identifies the elastic IP address.
	Id string `json:"id,"`

	// Specifies the status of the elastic IP address.
	Status string `json:"status,"`

	// Specifies the type of the elastic IP address.
	Type string `json:"type,"`

	// Specifies the obtained elastic IP address.
	PublicIpAddress string `json:"public_ip_address,"`

	// Specifies the private IP address bound to the elastic IP address.
	PrivateIpAddress string `json:"private_ip_address,omitempty"`

	// Specifies the port ID.
	PortId string `json:"port_id,omitempty"`

	// Specifies the tenant ID of the operator.
	TenantId string `json:"tenant_id,"`

	// Specifies the time for applying for the elastic IP address.
	CreateTime string `json:"create_time,"`

	// Specifies the bandwidth ID of the elastic IP address.
	BandwidthId string `json:"bandwidth_id,"`

	// Specifies the bandwidth size.
	BandwidthSize int `json:"bandwidth_size,"`

	// Specifies whether the bandwidth is shared or exclusive.
	BandwidthShareType string `json:"bandwidth_share_type,"`

	// Specifies the bandwidth name.
	BandwidthName string `json:"bandwidth_name,"`
}

type CreateOpts struct {

	// Specifies the elastic IP address objects.
	Publicip struct {

		// Specifies the type of the elastic IP address. The value can the 5_telcom, 5_union, 5_bgp, or 5_sbgp. China Northeast: 5_telcom and 5_union China South: 5_sbgp China East: 5_sbgp China North: 5_bgp and 5_sbgp The value must be a type supported by the system. The value can be 5_telcom, 5_union, or 5_bgp.
		Type string `json:"type,"`

		// Specifies the elastic IP address to be obtained. The value must be a valid IP address in the available IP address segment.
		IpAddress string `json:"ip_address,omitempty"`
	} `json:"publicip,"`

	// Specifies the bandwidth objects.
	Bandwidth struct {

		// Specifies the bandwidth name. The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), and hyphens (-). This parameter is mandatory when share_type is set to PER and is optional when share_type is set to WHOLE with an ID specified.
		Name string `json:"name,omitempty"`

		// Specifies the bandwidth size. The value ranges from 1 Mbit/s to 300 Mbit/s. This parameter is mandatory when share_type is set to PER and is optional when share_type is set to WHOLE with an ID specified.
		Size int `json:"size,omitempty"`

		// Specifies the ID of the bandwidth. You can specify an earlier shared bandwidth when applying for an elastic IP address for the bandwidth whose type is set to WHOLE. The bandwidth whose type is set to WHOLE exclusively uses its own ID. The value can be the ID of the bandwidth whose type is set to WHOLE.
		Id string `json:"id,omitempty"`

		// Specifies whether the bandwidth is shared or exclusive. The value can be PER or WHOLE.
		ShareType string `json:"share_type,"`

		// Specifies the charging mode (by traffic or by bandwidth). The value can be bandwidth or traffic. If the value is an empty character string or no value is specified, default value bandwidth is used.
		ChargeMode string `json:"charge_mode,omitempty"`
	} `json:"bandwidth,"`
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

func Delete(client *golangsdk.ServiceClient, tenantId string, publicipId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, publicipId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, publicipId string) (r GetResult) {
	url := GetURL(client, tenantId, publicipId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies the resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the number of records returned on each page. The value ranges from 0 to intmax.
	Limit int `q:"limit"`
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

	// Specifies the elastic IP address objects.
	Publicip struct {

		// Specifies the port ID. Constraints: The value must be an existing port ID. If this parameter is not included or the parameter value is left blank, the elastic IP address is unbound. If the specified port ID does not exist or has been bound to an elastic IP address, an error message will be displayed.
		PortId string `json:"port_id,"`
	} `json:"publicip,"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, publicipId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, publicipId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
