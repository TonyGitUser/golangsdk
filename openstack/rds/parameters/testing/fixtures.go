package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/parameters"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CraeteOutput = `
{"shouldRestart":"0","setParameteResult":"0"}
`

var CreateResponse = parameters.CreateResponse{
	ShouldRestart:     "0",
	SetParameteResult: "0",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/6926d5168444416fa28646de8a67450fno01/parameters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CraeteOutput)
	})
}

var ListOutput = `
{
  "configuration-parameters": [{
    "name": "autocommit",
    "type": "boolean",
    "restart_required": false,
    "datastore_version_id": "9dae9226-5f8e-4b10-bd41-adae2e693e89",
    "value_range": "ON|OFF",
    "description": "The autocommit mode. If set to ON, all changes to a table take effect immediately. If set to OFF, you must use COMMIT to accept a transaction or ROLLBACK to cancel it. "
  }, {
    "name": "automatic_sp_privileges",
    "type": "boolean",
    "restart_required": false,
    "datastore_version_id": "9dae9226-5f8e-4b10-bd41-adae2e693e89",
    "value_range": "ON|OFF",
    "description": "When this variable has a value of ON (the default), the server automatically grants the EXECUTE and ALTER ROUTINE privileges to the creator of a stored routine, if the user cannot already execute and alter or drop the routine."
  }, {
    "name": "auto_increment_increment",
    "min": "1",
    "max": "65535",
    "type": "integer",
    "restart_required": false,
    "datastore_version_id": "9dae9226-5f8e-4b10-bd41-adae2e693e89",
    "value_range": "1-65535",
    "description": "auto_increment_increment and auto_increment_offset are intended for use with master-to-master replication, and can be used to control the operation of AUTO_INCREMENT columns."
  }, {
    "name": "auto_increment_offset",
    "min": "1",
    "max": "65535",
    "type": "integer",
    "restart_required": false,
    "datastore_version_id": "9dae9226-5f8e-4b10-bd41-adae2e693e89",
    "value_range": "1-65535",
    "description": "auto_increment_increment and auto_increment_offset are intended for use with master-to-master replication, and can be used to control the operation of AUTO_INCREMENT columns. "
  }]
}
`

var ListResponse = parameters.ListResponse{
	ConfigurationParameters: []struct {
		Name               string `json:"name,omitempty"`
		Min                string `json:"min,omitempty"`
		Max                string `json:"max,omitempty"`
		Type               string `json:"type,omitempty"`
		ValueRange         string `json:"value_range,omitempty"`
		Description        string `json:"description,omitempty"`
		RestartRequired    bool   `json:"restart_required,omitempty"`
		DatastoreVersionId string `json:"datastore_version_id,omitempty"`
	}{
		{
			Name:               "autocommit",
			Type:               "boolean",
			RestartRequired:    false,
			DatastoreVersionId: "9dae9226-5f8e-4b10-bd41-adae2e693e89",
			ValueRange:         "ON|OFF",
			Description:        "The autocommit mode. If set to ON, all changes to a table take effect immediately. If set to OFF, you must use COMMIT to accept a transaction or ROLLBACK to cancel it. ",
		}, {
			Name:               "automatic_sp_privileges",
			Type:               "boolean",
			RestartRequired:    false,
			DatastoreVersionId: "9dae9226-5f8e-4b10-bd41-adae2e693e89",
			ValueRange:         "ON|OFF",
			Description:        "When this variable has a value of ON (the default), the server automatically grants the EXECUTE and ALTER ROUTINE privileges to the creator of a stored routine, if the user cannot already execute and alter or drop the routine.",
		}, {
			Name:               "auto_increment_increment",
			Min:                "1",
			Max:                "65535",
			Type:               "integer",
			RestartRequired:    false,
			DatastoreVersionId: "9dae9226-5f8e-4b10-bd41-adae2e693e89",
			ValueRange:         "1-65535",
			Description:        "auto_increment_increment and auto_increment_offset are intended for use with master-to-master replication, and can be used to control the operation of AUTO_INCREMENT columns.",
		}, {
			Name:               "auto_increment_offset",
			Min:                "1",
			Max:                "65535",
			Type:               "integer",
			RestartRequired:    false,
			DatastoreVersionId: "9dae9226-5f8e-4b10-bd41-adae2e693e89",
			ValueRange:         "1-65535",
			Description:        "auto_increment_increment and auto_increment_offset are intended for use with master-to-master replication, and can be used to control the operation of AUTO_INCREMENT columns. ",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/datastores/versions/9dae9226-5f8e-4b10-bd41-adae2e693e89/parameters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var GetOutput = `
{
  "name": "auto_increment_offset",
  "min": "1",
  "max": "65535",
  "type": "integer",
  "restart_required": false,
  "datastore_version_id": "9dae9226-5f8e-4b10-bd41-adae2e693e89",
  "value_range": "1-65535",
  "description": "auto_increment_increment and auto_increment_offset are intended for use with master-to-master replication, and can be used to control the operation of AUTO_INCREMENT columns. "
}
`

var GetResponse = parameters.GetResponse{
	Name:               "auto_increment_offset",
	Min:                "1",
	Max:                "65535",
	Type:               "integer",
	RestartRequired:    false,
	DatastoreVersionId: "9dae9226-5f8e-4b10-bd41-adae2e693e89",
	ValueRange:         "1-65535",
	Description:        "auto_increment_increment and auto_increment_offset are intended for use with master-to-master replication, and can be used to control the operation of AUTO_INCREMENT columns. ",
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/datastores/versions/9dae9226-5f8e-4b10-bd41-adae2e693e89/parameters/auto_increment_offset", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var RestoreOutput = `
{"shouldRestart":"1","setParameteResult":"0"}
`

var RestoreResponse = parameters.RestoreResponse{
	ShouldRestart:     "1",
	SetParameteResult: "0",
}

func HandleRestoreSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/6926d5168444416fa28646de8a67450fno01/parameters/default", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RestoreOutput)
	})
}
