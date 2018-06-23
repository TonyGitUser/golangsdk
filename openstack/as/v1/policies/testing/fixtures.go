package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/policies"
)

var CreateOutput = `
{"scaling_policy_id":"f774d696-58be-4f52-9b6d-6906f90c072d"}
`

var CreateResponse = policies.CreateResponse{
	ScalingPolicyId: "f774d696-58be-4f52-9b6d-6906f90c072d",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_policy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{"scaling_policy_id":"1a68e2a0-a19c-430e-a9c1-2a41b662ff57"}
`

var UpdateResponse = policies.UpdateResponse{
	ScalingPolicyId: "1a68e2a0-a19c-430e-a9c1-2a41b662ff57",
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_policy/1a68e2a0-a19c-430e-a9c1-2a41b662ff57", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "scaling_policy": {
    "scaling_policy_id": "1a68e2a0-a19c-430e-a9c1-2a41b662ff57",
    "scaling_group_id": "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
    "scaling_policy_name": "as-policy-s4ph",
    "scaling_policy_type": "ALARM",
    "alarm_id": "al1526121470824LpAqvyjLR",
    "scheduled_policy": {},
    "cool_down_time": 900,
    "scaling_policy_action": {
      "operation": "ADD",
      "instance_number": 1
    },
    "policy_status": "INSERVICE",
    "create_time": "2018-05-12T10:37:51Z"
  }
}
`
var GetResponse = policies.GetResponse{
	ScalingPolicy: struct {
		ScalingGroupId    string `json:"scaling_group_id,omitempty"`
		ScalingPolicyName string `json:"scaling_policy_name,omitempty"`
		ScalingPolicyId   string `json:"scaling_policy_id,omitempty"`
		PolicyStatus      string `json:"policy_status,omitempty"`
		ScalingPolicyType string `json:"scaling_policy_type,omitempty"`
		AlarmId           string `json:"alarm_id,omitempty"`
		ScheduledPolicy   struct {
			LaunchTime      string `json:"launch_time,omitempty"`
			RecurrenceType  string `json:"recurrence_type,omitempty"`
			RecurrenceValue string `json:"recurrence_value,omitempty"`
			StartTime       string `json:"start_time,omitempty"`
			EndTime         string `json:"end_time,omitempty"`
		} `json:"scheduled_policy,omitempty"`
		ScalingPolicyAction struct {
			Operation          string `json:"operation,omitempty"`
			InstanceNumber     int    `json:"instance_number,omitempty"`
			InstancePercentage int    `json:"instance_percentage,omitempty"`
		} `json:"scaling_policy_action,omitempty"`
		CoolDownTime int    `json:"cool_down_time,omitempty"`
		CreateTime   string `json:"create_time,omitempty"`
	}{
		ScalingPolicyId:   "1a68e2a0-a19c-430e-a9c1-2a41b662ff57",
		ScalingGroupId:    "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
		ScalingPolicyName: "as-policy-s4ph",
		ScalingPolicyType: "ALARM",
		AlarmId:           "al1526121470824LpAqvyjLR",
		ScheduledPolicy: struct {
			LaunchTime      string `json:"launch_time,omitempty"`
			RecurrenceType  string `json:"recurrence_type,omitempty"`
			RecurrenceValue string `json:"recurrence_value,omitempty"`
			StartTime       string `json:"start_time,omitempty"`
			EndTime         string `json:"end_time,omitempty"`
		}{},
		CoolDownTime: 900,
		ScalingPolicyAction: struct {
			Operation          string `json:"operation,omitempty"`
			InstanceNumber     int    `json:"instance_number,omitempty"`
			InstancePercentage int    `json:"instance_percentage,omitempty"`
		}{
			Operation:      "ADD",
			InstanceNumber: 1,
		},
		PolicyStatus: "INSERVICE",
		CreateTime:   "2018-05-12T10:37:51Z",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_policy/1a68e2a0-a19c-430e-a9c1-2a41b662ff57", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "limit": 2,
  "total_number": 4,
  "start_number": 0,
  "scaling_policies": [{
    "scaling_policy_id": "1a68e2a0-a19c-430e-a9c1-2a41b662ff57",
    "scaling_group_id": "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
    "scaling_policy_name": "as-policy-s4ph",
    "scaling_policy_type": "ALARM",
    "alarm_id": "al1526121470824LpAqvyjLR",
    "scheduled_policy": {},
    "cool_down_time": 900,
    "scaling_policy_action": {
      "operation": "ADD",
      "instance_number": 1
    },
    "policy_status": "INSERVICE",
    "create_time": "2018-05-12T10:37:51Z"
  }, {
    "scaling_policy_id": "f774d696-58be-4f52-9b6d-6906f90c072d",
    "scaling_group_id": "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
    "scaling_policy_name": "as-policy-7a75",
    "scaling_policy_type": "RECURRENCE",
    "scheduled_policy": {
      "launch_time": "16:00",
      "recurrence_type": "Daily",
      "start_time": "2015-12-14T03:34Z",
      "end_time": "2019-12-27T03:34Z"
    },
    "cool_down_time": 900,
    "scaling_policy_action": {
      "operation": "ADD",
      "instance_number": 1
    },
    "policy_status": "INSERVICE",
    "create_time": "2018-05-13T15:05:02Z"
  }]
}
`

var ListResponse = policies.ListResponse{
	Limit:       2,
	TotalNumber: 4,
	StartNumber: 0,
	ScalingPolicies: []struct {
		ScalingGroupId    string `json:"scaling_group_id,omitempty"`
		ScalingPolicyName string `json:"scaling_policy_name,omitempty"`
		ScalingPolicyId   string `json:"scaling_policy_id,omitempty"`
		PolicyStatus      string `json:"policy_status,omitempty"`
		ScalingPolicyType string `json:"scaling_policy_type,omitempty"`
		AlarmId           string `json:"alarm_id,omitempty"`
		ScheduledPolicy   struct {
			LaunchTime      string `json:"launch_time,omitempty"`
			RecurrenceType  string `json:"recurrence_type,omitempty"`
			RecurrenceValue string `json:"recurrence_value,omitempty"`
			StartTime       string `json:"start_time,omitempty"`
			EndTime         string `json:"end_time,omitempty"`
		} `json:"scheduled_policy,omitempty"`
		ScalingPolicyAction struct {
			Operation          string `json:"operation,omitempty"`
			InstanceNumber     int    `json:"instance_number,omitempty"`
			InstancePercentage int    `json:"instance_percentage,omitempty"`
		} `json:"scaling_policy_action,omitempty"`
		CoolDownTime int    `json:"cool_down_time,omitempty"`
		CreateTime   string `json:"create_time,omitempty"`
	}{
		{
			ScalingPolicyId:   "1a68e2a0-a19c-430e-a9c1-2a41b662ff57",
			ScalingGroupId:    "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
			ScalingPolicyName: "as-policy-s4ph",
			ScalingPolicyType: "ALARM",
			AlarmId:           "al1526121470824LpAqvyjLR",
			ScheduledPolicy: struct {
				LaunchTime      string `json:"launch_time,omitempty"`
				RecurrenceType  string `json:"recurrence_type,omitempty"`
				RecurrenceValue string `json:"recurrence_value,omitempty"`
				StartTime       string `json:"start_time,omitempty"`
				EndTime         string `json:"end_time,omitempty"`
			}{},
			CoolDownTime: 900,
			ScalingPolicyAction: struct {
				Operation          string `json:"operation,omitempty"`
				InstanceNumber     int    `json:"instance_number,omitempty"`
				InstancePercentage int    `json:"instance_percentage,omitempty"`
			}{
				Operation:      "ADD",
				InstanceNumber: 1,
			},
			PolicyStatus: "INSERVICE",
			CreateTime:   "2018-05-12T10:37:51Z",
		}, {
			ScalingPolicyId:   "f774d696-58be-4f52-9b6d-6906f90c072d",
			ScalingGroupId:    "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
			ScalingPolicyName: "as-policy-7a75",
			ScalingPolicyType: "RECURRENCE",
			ScheduledPolicy: struct {
				LaunchTime      string `json:"launch_time,omitempty"`
				RecurrenceType  string `json:"recurrence_type,omitempty"`
				RecurrenceValue string `json:"recurrence_value,omitempty"`
				StartTime       string `json:"start_time,omitempty"`
				EndTime         string `json:"end_time,omitempty"`
			}{
				LaunchTime:     "16:00",
				RecurrenceType: "Daily",
				StartTime:      "2015-12-14T03:34Z",
				EndTime:        "2019-12-27T03:34Z",
			},
			CoolDownTime: 900,
			ScalingPolicyAction: struct {
				Operation          string `json:"operation,omitempty"`
				InstanceNumber     int    `json:"instance_number,omitempty"`
				InstancePercentage int    `json:"instance_percentage,omitempty"`
			}{
				Operation:      "ADD",
				InstanceNumber: 1,
			},
			PolicyStatus: "INSERVICE",
			CreateTime:   "2018-05-13T15:05:02Z",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_policy/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/list", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleActionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_policy/1a68e2a0-a19c-430e-a9c1-2a41b662ff57/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_policy/1a68e2a0-a19c-430e-a9c1-2a41b662ff57", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
