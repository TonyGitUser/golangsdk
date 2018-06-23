package loadbalancers

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOpts struct {

	// Specifies the load balancer name.The value is a string of 1 to 64 characters that consist of Chinese characters, letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the load balancer.The value is a string of 0 to 128 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`

	// Specifies the VPC ID.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies the bandwidth (Mbit/s). This parameter is mandatory when type is set to External.The value ranges from 1 to 300.
	Bandwidth int `json:"bandwidth,omitempty"`

	// Specifies the load balancer type.The value is Internal or External.
	Type string `json:"type,omitempty"`

	// Specifies the status of the load balancer.Optional values:0 or false: indicates that the load balancer is stopped. Only tenants are allowed to enter these two values.1 or true: indicates that the load balancer is running properly.2 or false: indicates that the load balancer is frozen. Only tenants are allowed to enter these two values.
	AdminStateUp int `json:"admin_state_up,omitempty"`

	// Specifies the subnet ID of backend ECSs. This parameter is mandatory when type is set to Internal.
	VipSubnetId string `json:"vip_subnet_id,omitempty"`

	// This field is valid when type is set to Internal. If type is set to Internal and an AZ is specified, the specified AZ must support private network load balancers. Otherwise, an error message is returned. For details, see Regions and Endpoints.https://developer.huaweicloud.com/en-us/endpoint
	Az string `json:"az,omitempty"`

	// Specifies the charging mode.Specifies the charging mode. It can be bandwidth or traffic. The default value is bandwidth.
	ChargeMode string `json:"charge_mode,omitempty"`

	// Specifies the elastic IP address type.The configured value must be supported by the system.The value can be 5_telcom, 5_union, or 5_bgp. The default value is 5_bgp.
	EipType string `json:"eip_type,omitempty"`

	// Specifies the security group ID.The value is a string of 1 to 200 characters that consists of uppercase and lowercase letters, digits, and hyphens (-).This parameter is mandatory when type is set to Internal.
	SecurityGroupId string `json:"security_group_id,omitempty"`

	// Specifies the IP address used by ELB for providing services. When type is set to External, the parameter value is the elastic IP address. When type is set to Internal, the parameter value is the private network IP address.You can select an existing elastic IP address and create a public network load balancer. When this parameter is configured, parameters bandwidth, charge_mode, and eip_type are invalid.
	VipAddress string `json:"vip_address,omitempty"`

	// Specifies the tenant ID.This parameter is mandatory when type is set to Internal.
	TenantId string `json:"tenantId,omitempty"`
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

func Delete(client *golangsdk.ServiceClient, tenantId string, loadbalancerId string) (r DeleteResult) {
	url := DeleteURL(client, tenantId, loadbalancerId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func DeleteKeepEIP(client *golangsdk.ServiceClient, tenantId string, loadbalancerId string) (r DeleteKeepEIPResult) {
	url := DeleteKeepEIPURL(client, tenantId, loadbalancerId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

func Get(client *golangsdk.ServiceClient, tenantId string, loadbalancerId string) (r GetResult) {
	url := GetURL(client, tenantId, loadbalancerId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

func List(client *golangsdk.ServiceClient, tenantId string) (r ListResult) {
	url := ListURL(client, tenantId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type UpdateOpts struct {

	// Specifies the load balancer name.The value is a string of 1 to 64 characters that consist of Chinese characters, letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the load balancer.The value is a string of 0 to 128 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`

	// Specifies the bandwidth (Mbit/s).The value ranges from 1 to 500.
	Bandwidth int `json:"bandwidth,omitempty"`

	// Specifies the status of the load balancer.Optional values:0 or false: indicates that the load balancer is stopped. Only tenants are allowed to enter these two values.1 or true: indicates that the load balancer is running properly.2 or false: indicates that the load balancer is frozen. Only tenants are allowed to enter these two values.
	AdminStateUp int `json:"admin_state_up,omitempty"`
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

func Update(client *golangsdk.ServiceClient, tenantId string, loadbalancerId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Put(UpdateURL(client, tenantId, loadbalancerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
