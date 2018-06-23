package quotas

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
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

	//
	Quotas struct {

		//
		Resources []struct {

			// Specifies the quota type.
			Type string `json:"type,omitempty"`

			// Specifies the used amount of the quota.When type is set to scaling_Policy or scaling_Instance, this parameter is reserved, and the system returns -1 as the parameter value. You can query the used quota amount of AS policies and AS instances in a specified AS group. For details, see Querying Quotas for AS Policies and AS Instances.
			Used int `json:"used,omitempty"`

			// Specifies the total amount of the quota.
			Quota int `json:"quota,omitempty"`

			// Specifies the quota upper limit.
			Max int `json:"max,omitempty"`
		} `json:"resources,omitempty"`
	} `json:"quotas,omitempty"`
}

type ListWithInstancesResult struct {
	commonResult
}

func (r ListWithInstancesResult) Extract() (*ListWithInstancesResponse, error) {
	var response ListWithInstancesResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListWithInstancesResponse struct {

	//
	Quotas struct {

		//
		Resources []struct {

			// Specifies the quota type.
			Type string `json:"type,omitempty"`

			// Specifies the used amount of the quota.When type is set to scaling_Policy or scaling_Instance, this parameter is reserved, and the system returns -1 as the parameter value. You can query the used quota amount of AS policies and AS instances in a specified AS group. For details, see Querying Quotas for AS Policies and AS Instances.
			Used int `json:"used,omitempty"`

			// Specifies the total amount of the quota.
			Quota int `json:"quota,omitempty"`

			// Specifies the quota upper limit.
			Max int `json:"max,omitempty"`
		} `json:"resources,omitempty"`
	} `json:"quotas,omitempty"`
}
