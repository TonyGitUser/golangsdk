package vpcs

import (
	"github.com/huaweicloud/golangsdk"
)

type VPC struct {

	// Specifies a resource ID in UUID format.
	Id string `json:"id,"`

	// Specifies the name of the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Specifies the range of available subnets in the VPC. The value must be in CIDR format, for example, 192.168.0.0/16. The value ranges from 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to 192.168.255.0/24.
	Cidr string `json:"cidr,omitempty"`

	// Specifies the status of the VPC. The value can be CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.
	Status string `json:"status,"`
}

type CreateOpts struct {

	// Specifies the VPC objects.
	Vpc VPC `json:"vpc,"`
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

func Delete(client *golangsdk.ServiceClient, tenantId string, vpcId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, vpcId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, vpcId string) (r GetResult) {
	url := GetURL(client, tenantId, vpcId)
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

	// Specifies the VPC objects.
	Vpc VPC `json:"vpc,"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, vpcId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, vpcId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
