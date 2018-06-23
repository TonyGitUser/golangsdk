package parameters

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

	// Indicates whether the value needs to be restarted to take effect.  indicates the parameter modification takes effect after a restart.  indicates the parameter modification takes effect immediately.
	ShouldRestart string `json:"shouldRestart,omitempty"`

	// Indicates the parameter modification result.  indicates the modification has succeeded on the primary DB instance, and failed on the standby DB instance.  indicates the modification has succeeded on both DB instances.
	SetParameteResult string `json:"setParameteResult,omitempty"`
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

	// Indicates the parameter name.
	Name string `json:"name,omitempty"`

	// Indicates the minimum value of the parameter. Returned only when  is  or .
	Min string `json:"min,omitempty"`

	// Indicates the maximum value of the parameter. Returned only when  is  or .
	Max string `json:"max,omitempty"`

	// Indicates whether the instance needs to reboot for the parameter to take effect. The value is  or .
	RestartRequired bool `json:"restart_required,omitempty"`

	// Indicates the parameter type. The value can be , , , , or .
	Type string `json:"type,omitempty"`

	// Indicates the parameter value range and enumerated value. Returned only when  is  or . If  is , the parameter value can be  or .
	ValueRange string `json:"value_range,omitempty"`

	// Indicates the parameter description.
	Description string `json:"description,omitempty"`

	// Indicates the database version.
	DatastoreVersionId string `json:"datastore_version_id,omitempty"`
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

	// Indicates all the parameters that can be modified of the database version.
	ConfigurationParameters []struct {

		// Indicates the parameter name.
		Name string `json:"name,omitempty"`

		// Indicates the minimum value. Returned only when  is  or .
		Min string `json:"min,omitempty"`

		// Indicates the maximum value. Returned only when  is  or .
		Max string `json:"max,omitempty"`

		// Indicates the parameter type. The value can be , , , , or .
		Type string `json:"type,omitempty"`

		// Indicates the parameter value range and enumerated value. Returned only when  is  or . If  is , the parameter value can be  or .
		ValueRange string `json:"value_range,omitempty"`

		// Indicates the parameter description.
		Description string `json:"description,omitempty"`

		// Indicates whether the instance needs to reboot for the parameter to take effect. The value is  or .
		RestartRequired bool `json:"restart_required,omitempty"`

		// Indicates the database version ID.
		DatastoreVersionId string `json:"datastore_version_id,omitempty"`
	} `json:"configuration-parameters,omitempty"`
}

type RestoreResult struct {
	commonResult
}

func (r RestoreResult) Extract() (*RestoreResponse, error) {
	var response RestoreResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RestoreResponse struct {

	// Indicates whether the value needs to be restarted to take effect.  indicates the parameter modification takes effect after a restart.  indicates the parameter modification takes effect immediately.
	ShouldRestart string `json:"shouldRestart,omitempty"`

	// Indicates the parameter modification result.  indicates the modification has succeeded on the primary DB instance, and failed on the standby DB instance.  indicates the modification has succeeded on both DB instances.
	SetParameteResult string `json:"setParameteResult,omitempty"`
}
