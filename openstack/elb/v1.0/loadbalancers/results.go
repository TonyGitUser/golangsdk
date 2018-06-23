package loadbalancers

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {

	// Specifies the URI of the task for creating a load balancer. It is returned by Combined API.
	Uri string `json:"uri,omitempty"`

	// Specifies the unique ID assigned to the task for creating a load balancer in Combined API.
	JobId string `json:"job_id,omitempty"`
}

type DeleteResult struct {
	commonResult
}

func (r DeleteResult) Extract() (*DeleteResponse, error) {
	var response DeleteResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteResponse struct {

	// Specifies the URI of the task for deleting a load balancer. It is returned by Combined API.
	Uri string `json:"uri,omitempty"`

	// Specifies the unique ID assigned to the task for deleting a load balancer in Combined API.
	JobId string `json:"job_id,omitempty"`
}

type DeleteKeepEIPResult struct {
	commonResult
}

func (r DeleteKeepEIPResult) Extract() (*DeleteKeepEIPResponse, error) {
	var response DeleteKeepEIPResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteKeepEIPResponse struct {

	// Specifies the task URI returned by Combined API after the task for deleting a load balancer is delivered.
	Uri string `json:"uri,omitempty"`

	// Specifies the unique ID assigned to the task for deleting a load balancer in Combined API.
	JobId string `json:"job_id,omitempty"`
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

	// Specifies the IP address used by ELB for providing services.
	VipAddress string `json:"vip_address,omitempty"`

	// Specifies the time when information about the load balancer was updated.
	UpdateTime string `json:"update_time,omitempty"`

	// Specifies the time when the load balancer was created.
	CreateTime string `json:"create_time,omitempty"`

	// Specifies the load balancer ID.
	Id string `json:"id,omitempty"`

	// Specifies the status of the load balancer.The value can be ACTIVE, PENDING_CREATE, or ERROR.
	Status string `json:"status,omitempty"`

	// Specifies the bandwidth.
	Bandwidth int `json:"bandwidth,omitempty"`

	// Specifies the VPC ID.
	VpcId string `json:"vpc_id,omitempty"`

	// Specifies the status of the load balancer.Optional values:0: The load balancer is disabled.1: The load balancer is running properly.2: The load balancer is frozen.
	AdminStateUp int `json:"admin_state_up,omitempty"`

	// Specifies the ID of the private network that the load balancer accesses. This parameter is valid when type is Internal.
	VipSubnetId string `json:"vip_subnet_id,omitempty"`

	// Specifies the load balancer type. The value is Internal or External.
	Type string `json:"type,omitempty"`

	// Specifies the load balancer name.
	Name string `json:"name,omitempty"`

	// Provides supplementary information about the load balancer.
	Description string `json:"description,omitempty"`

	// Specifies the security group ID.null is displayed for this parameter when type is set to External.
	SecurityGroupId string `json:"security_group_id,omitempty"`
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResponse struct {

	// Lists the load balancers.
	Loadbalancers []struct {

		// Specifies the IP address used by ELB for providing services.
		VipAddress string `json:"vip_address,omitempty"`

		// Specifies the time when information about the load balancer was updated.
		UpdateTime string `json:"update_time,omitempty"`

		// Specifies the time when the load balancer was created.
		CreateTime string `json:"create_time,omitempty"`

		// Specifies the load balancer ID.
		Id string `json:"id,omitempty"`

		// Specifies the status of the load balancer.The value can be ACTIVE, PENDING_CREATE, or ERROR.
		Status string `json:"status,omitempty"`

		// Specifies the bandwidth.
		Bandwidth int `json:"bandwidth,omitempty"`

		// Specifies the VPC ID.
		VpcId string `json:"vpc_id,omitempty"`

		// Specifies the status of the load balancer.Optional values:0: The load balancer is disabled.1: The load balancer is running properly.2: The load balancer is frozen.
		AdminStateUp int `json:"admin_state_up,omitempty"`

		// Specifies the ID of the private network that the load balancer accesses. This parameter is valid when type is Internal.
		VipSubnetId string `json:"vip_subnet_id,omitempty"`

		// Specifies the load balancer type. The value is Internal or External.
		Type string `json:"type,omitempty"`

		// Specifies the load balancer name.
		Name string `json:"name,omitempty"`

		// Provides supplementary information about the load balancer.
		Description string `json:"description,omitempty"`

		// Specifies the security group ID.null is displayed for this parameter when type is set to External.
		SecurityGroupId string `json:"security_group_id,omitempty"`
	} `json:"loadbalancers,omitempty"`

	// Specifies the quantity of load balancers.
	InstanceNum string `json:"instance_num,omitempty"`
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*UpdateResponse, error) {
	var response UpdateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateResponse struct {

	// Specifies the URI of the task for modifying a load balancer. It is returned by Combined API.
	Uri string `json:"uri,omitempty"`

	// Specifies the unique ID assigned to the task for modifying a load balancer in Combined API.
	JobId string `json:"job_id,omitempty"`
}
