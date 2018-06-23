package dbversions

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

	// Indicates the list of database versions.
	DataStores []struct {

		// Indicates the database version ID. Its value is unique.
		Id string `json:"id,omitempty"`

		// Indicates the database version.
		Name string `json:"name,omitempty"`

		// Indicates the database ID.
		Datastore string `json:"datastore,omitempty"`

		// Indicates the database image ID.
		Image string `json:"image,omitempty"`

		// Indicates the database package version information.
		Packages string `json:"packages,omitempty"`

		// Indicates the current database version status. 0 indicates Non-activated, and 1 indicates Activated. The interface can only query information of versions that are activated.
		Active int `json:"active,omitempty"`
	} `json:"dataStores,omitempty"`
}
