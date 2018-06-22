package privateips

import (
	"github.com/huaweicloud/golangsdk"
)

type PrivateIp struct {

	// 1、取值范围：ACTIVE ,DOWN 2、功能说明：私有IP的状态
	Status string `json:"status,omitempty"`

	// 功能说明：私有IP标识
	Id string `json:"id,omitempty"`

	// 功能说明：分配IP的子网标识
	SubnetId string `json:"subnet_id,omitempty"`

	// 功能说明：操作用户的租户ID
	TenantId string `json:"tenant_id,omitempty"`

	// 1、取值范围：network:dhcp，network:router_interface_distributed，compute:xxx(xxx对应具体的az名称，例如compute:aa-bb-cc表示是被aa-bb-cc上的虚拟机使用) 2、功能说明：私有IP的使用者，空表示未使用 3、约束：此处的取值范围只是本服务支持 的类型，其他类型未做标注
	DeviceOwner string `json:"device_owner,omitempty"`

	// 功能说明：申请到的私有IP
	IpAddress string `json:"ip_address,omitempty"`
}

type CreatePrivateIp struct {

	// 功能说明：分配IP的子网标识
	SubnetId string `json:"subnet_id,omitempty"`

	// 功能说明：申请到的私有IP
	IpAddress string `json:"ip_address,omitempty"`
}

type CreateOpts struct {

	// Specifies the private IP address list objects.
	Privateips []CreatePrivateIp `json:"privateips,"`
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

func Delete(client *golangsdk.ServiceClient, tenantId string, privateipId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, privateipId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, privateipId string) (r GetResult) {
	url := GetURL(client, tenantId, privateipId)
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

func List(client *golangsdk.ServiceClient, tenantId string, subnetId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId, subnetId)
	if opts != nil {
		query, _ := opts.ToListQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
