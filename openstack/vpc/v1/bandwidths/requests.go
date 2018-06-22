package bandwidths

import (
	"github.com/huaweicloud/golangsdk"
)

type BandWidth struct {

	// Specifies the bandwidth name. The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,"`

	// Specifies the bandwidth size. The value ranges from 1 Mbit/s to 300 Mbit/s.
	Size int `json:"size,"`

	// Specifies the bandwidth ID, which uniquely identifies the bandwidth.
	Id string `json:"id,"`

	// Specifies whether the bandwidth is shared or exclusive. The value can be PER or WHOLE.
	ShareType string `json:"share_type,"`

	// Specifies the elastic IP address of the bandwidth.  The bandwidth, whose type is set to WHOLE, supports up to 20 elastic IP addresses. The bandwidth, whose type is set to PER, supports only one elastic IP address.
	PublicipInfo []struct {

		// Specifies the tenant ID of the user.
		PublicipId string `json:"publicip_id,"`

		// Specifies the elastic IP address.
		PublicipAddress string `json:"publicip_address,"`

		// Specifies the elastic IP address type. The value can be 5_telcom, 5_union, or 5_bgp.
		PublicipType string `json:"publicip_type,"`
	} `json:"publicip_info,"`

	// Specifies the tenant ID of the user.
	TenantId string `json:"tenant_id,"`

	// Specifies the bandwidth type.
	BandwidthType string `json:"bandwidth_type,"`

	// Specifies the charging mode (by traffic or by bandwidth).
	ChargeMode string `json:"charge_mode,omitempty"`

	// Specifies the billing information.
	BillingInfo string `json:"billing_info,omitempty"`
}

func Get(client *golangsdk.ServiceClient, tenantId string, bandwidthId string) (r GetResult) {
	url := GetURL(client, tenantId, bandwidthId)
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

	// Specifies the bandwidth objects.
	Bandwidth struct {

		// Specifies the bandwidth name. At least one in parameter name or parameter size must be set. The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), and hyphens (-). If the value is left blank, the name of the bandwidth is not changed.
		Name string `json:"name,"`

		// Specifies the bandwidth size. Either parameter size or name must be set. The value ranges from 1 Mbit/s to 300 Mbit/s. If the parameter is not included, the bandwidth size is not changed.
		Size int `json:"size,"`
	} `json:"bandwidth,"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, bandwidthId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, bandwidthId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
