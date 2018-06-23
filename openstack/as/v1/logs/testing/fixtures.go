package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/logs"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
{
  "limit": 2,
  "scaling_activity_log": [{
    "id": "c24dec12-9fb3-4337-89b6-f14b9d914481",
    "instance_value": 0,
    "desire_value": 1,
    "scaling_value": 0,
    "start_time": "2018-05-12T09:14:51Z",
    "end_time": null,
    "instance_added_list": "",
    "instance_deleted_list": null,
    "instance_removed_list": "",
    "status": "DOING",
    "description": "{\"reason\":[{\"change_reason\":\"DIFF\",\"old_value\":0,\"change_time\":\"2018-05-12T09:14:51Z\",\"new_value\":1}]}"
  }, {
    "id": "ae352da9-333f-47cc-880d-90704c6bb04e",
    "instance_value": 0,
    "desire_value": 1,
    "scaling_value": 0,
    "start_time": "2018-05-12T09:13:39Z",
    "end_time": "2018-05-12T09:14:40Z",
    "instance_added_list": "",
    "instance_deleted_list": null,
    "instance_removed_list": "",
    "status": "FAIL",
    "description": "{\"reason\":[{\"change_reason\":\"MANNUAL_ADD\",\"old_value\":0,\"change_time\":\"2018-05-12T09:13:39Z\",\"new_value\":1}],\"scaling_num\":0}"
  }],
  "total_number": 4,
  "start_number": 0
}
`

var ListResponse = logs.ListResponse{
	Limit:       2,
	TotalNumber: 4,
	StartNumber: 0,
	ScalingActivityLog: []struct {
		Status              string `json:"status,omitempty"`
		StartTime           string `json:"start_time,omitempty"`
		EndTime             string `json:"end_time,omitempty"`
		Id                  string `json:"id,omitempty"`
		InstanceRemovedList string `json:"instance_removed_list,omitempty"`
		InstanceDeletedList string `json:"instance_deleted_list,omitempty"`
		InstanceAddedList   string `json:"instance_added_list,omitempty"`
		ScalingValue        int    `json:"scaling_value,omitempty"`
		Description         string `json:"description,omitempty"`
		InstanceValue       int    `json:"instance_value,omitempty"`
		DesireValue         int    `json:"desire_value,omitempty"`
	}{
		{
			Id:                  "c24dec12-9fb3-4337-89b6-f14b9d914481",
			InstanceValue:       0,
			DesireValue:         1,
			ScalingValue:        0,
			StartTime:           "2018-05-12T09:14:51Z",
			EndTime:             "",
			InstanceAddedList:   "",
			InstanceDeletedList: "",
			InstanceRemovedList: "",
			Status:              "DOING",
			Description:         "{\"reason\":[{\"change_reason\":\"DIFF\",\"old_value\":0,\"change_time\":\"2018-05-12T09:14:51Z\",\"new_value\":1}]}",
		},
		{
			Id:                  "ae352da9-333f-47cc-880d-90704c6bb04e",
			InstanceValue:       0,
			DesireValue:         1,
			ScalingValue:        0,
			StartTime:           "2018-05-12T09:13:39Z",
			EndTime:             "2018-05-12T09:14:40Z",
			InstanceAddedList:   "",
			InstanceDeletedList: "",
			InstanceRemovedList: "",
			Status:              "FAIL",
			Description:         "{\"reason\":[{\"change_reason\":\"MANNUAL_ADD\",\"old_value\":0,\"change_time\":\"2018-05-12T09:13:39Z\",\"new_value\":1}],\"scaling_num\":0}",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_activity_log/60dcec94-5d5b-4dbf-9f50-4ccd7d841432", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
