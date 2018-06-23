package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/jobs"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var GetOutput = `
{
  "status": "SUCCESS",
  "entities": {
    "elb": {
      "id": "ebe982b8afe04851bbc26ad4609003bf",
      "name": "UpdatedLB",
      "description": "TEST LB",
      "type": "External",
      "status": "ACTIVE",
      "bandwidth": 1,
      "vip_address": "49.4.9.141",
      "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
      "admin_state_up": true,
      "vpc_id": "773c3c42-d315-417b-9063-87091713148c"
    }
  },
  "job_id": "8aace0c86326772c016326c9f5dd0f3f",
  "job_type": "updateELB",
  "begin_time": "2018-05-03T16:15:59.964Z",
  "end_time": "2018-05-03T16:16:04.587Z",
  "error_code": null,
  "fail_reason": null
}
`

var GetResponse = jobs.GetResponse{
	Status: "SUCCESS",
	Entities: map[string]interface{}{
		"elb": map[string]interface{}{
			"id":             "ebe982b8afe04851bbc26ad4609003bf",
			"name":           "UpdatedLB",
			"description":    "TEST LB",
			"type":           "External",
			"status":         "ACTIVE",
			"bandwidth":      float64(1),
			"vip_address":    "49.4.9.141",
			"tenant_id":      "57e98940a77f4bb988a21a7d0603a626",
			"admin_state_up": true,
			"vpc_id":         "773c3c42-d315-417b-9063-87091713148c",
		},
	},
	JobId:      "8aace0c86326772c016326c9f5dd0f3f",
	JobType:    "updateELB",
	BeginTime:  "2018-05-03T16:15:59.964Z",
	EndTime:    "2018-05-03T16:16:04.587Z",
	ErrorCode:  "",
	FailReason: "",
}

func HandleListSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c86326772c016326c9f5dd0f3f", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}
