package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/instances"

	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func HandleActionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group_instance/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}

//var UpdateOutput = `
//{"scaling_group_id":"60dcec94-5d5b-4dbf-9f50-4ccd7d841432"}
//`
//
//var UpdateResponse = groups.UpdateResponse{
//    ScalingGroupId: "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
//}
//
//func HandleUpdateSuccessfully(t *testing.T) {
//    th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group/e4a97959-34d0-4c58-ab85-1dda4d0b9d11", func(w http.ResponseWriter, r *http.Request) {
//        th.TestMethod(t, r, "PUT")
//        th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
//
//        w.Header().Add("Content-Type", "application/json")
//        fmt.Fprintf(w, UpdateOutput)
//    })
//}
//
//var GetOutput = `
//{
//  "scaling_group": {
//    "scaling_group_name": "as-group-matt",
//    "scaling_group_id": "e4a97959-34d0-4c58-ab85-1dda4d0b9d11",
//    "scaling_group_status": "INSERVICE",
//    "scaling_configuration_id": "f109bb4f-09f0-4dac-9115-6b429d548750",
//    "scaling_configuration_name": "as-config-sh2u",
//    "current_instance_number": 0,
//    "desire_instance_number": 0,
//    "min_instance_number": 0,
//    "max_instance_number": 1,
//    "cool_down_time": 900,
//    "lb_listener_id": null,
//    "lbaas_listeners": [],
//    "networks": [{
//      "id": "f6a0db7b-397c-4162-bc35-9a1f008b4373"
//    }],
//    "cloud_location_id": null,
//    "available_zones": ["cn-north-1a"],
//    "security_groups": [{
//      "id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
//    }],
//    "create_time": "2018-05-09T16:12:57Z",
//    "vpc_id": "773c3c42-d315-417b-9063-87091713148c",
//    "detail": null,
//    "health_periodic_audit_method": "NOVA_AUDIT",
//    "health_periodic_audit_time": 5,
//    "instance_terminate_policy": "OLD_CONFIG_OLD_INSTANCE",
//    "is_scaling": false,
//    "notifications": [],
//    "delete_publicip": true
//  }
//}
//`
//var GetResponse = groups.GetResponse{
//    ScalingGroup: groups.ScalingGroup{
//        ScalingGroupName:         "as-group-matt",
//        ScalingGroupId:           "e4a97959-34d0-4c58-ab85-1dda4d0b9d11",
//        ScalingGroupStatus:       "INSERVICE",
//        ScalingConfigurationId:   "f109bb4f-09f0-4dac-9115-6b429d548750",
//        ScalingConfigurationName: "as-config-sh2u",
//        CurrentInstanceNumber:    0,
//        DesireInstanceNumber:     0,
//        MinInstanceNumber:        0,
//        MaxInstanceNumber:        1,
//        CoolDownTime:             900,
//        LbListenerId:             "",
//        LbaasListeners:           []string{},
//        Networks: []struct {
//            // Specifies the network ID.
//            Id string `json:"id,omitempty"`
//        }{
//            {Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373",},
//        },
//        CloudLocationId: "",
//        AvailableZones:  []string{"cn-north-1a"},
//        SecurityGroups: []struct {
//            // Specifies the ID of the security group.
//            Id string `json:"id,omitempty"`
//        }{
//            {Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",},
//        },
//        CreateTime:                "2018-05-09T16:12:57Z",
//        VpcId:                     "773c3c42-d315-417b-9063-87091713148c",
//        Detail:                    "",
//        HealthPeriodicAuditMethod: "NOVA_AUDIT",
//        HealthPeriodicAuditTime:   5,
//        InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
//        IsScaling:                 false,
//        Notifications:             []string{},
//        DeletePublicip:            true,
//    },
//}
//
//func HandleGetSuccessfully(t *testing.T) {
//    th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group/e4a97959-34d0-4c58-ab85-1dda4d0b9d11", func(w http.ResponseWriter, r *http.Request) {
//        th.TestMethod(t, r, "GET")
//        th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
//
//        w.Header().Add("Content-Type", "application/json")
//        fmt.Fprintf(w, GetOutput)
//    })
//}
//
var ListOutput = `
{
  "limit": 2,
  "total_number": 1,
  "start_number": 0,
  "scaling_group_instances": [{
    "instance_id": "e1040c67-b130-41ee-9405-07c6c6c20208",
    "scaling_group_id": "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
    "scaling_group_name": "as-group-Test",
    "life_cycle_state": "INSERVICE",
    "health_status": "NORMAL",
    "scaling_configuration_name": null,
    "scaling_configuration_id": null,
    "create_time": "2018-05-12T08:17:54Z",
    "instance_name": "ecs-SDK",
    "protect_from_scaling_down": false
  }]
}
`

var ListResponse = instances.ListResponse{
	Limit:       2,
	TotalNumber: 1,
	StartNumber: 0,
	ScalingGroupInstances: []struct {
		InstanceId               string `json:"instance_id,omitempty"`
		InstanceName             string `json:"instance_name,omitempty"`
		ScalingGroupId           string `json:"scaling_group_id,omitempty"`
		ScalingGroupName         string `json:"scaling_group_name,omitempty"`
		LifeCycleState           string `json:"life_cycle_state,omitempty"`
		HealthStatus             string `json:"health_status,omitempty"`
		ScalingConfigurationName string `json:"scaling_configuration_name,omitempty"`
		ScalingConfigurationId   string `json:"scaling_configuration_id,omitempty"`
		CreateTime               string `json:"create_time,omitempty"`
		ProtectFromScalingDown   bool   `json:"protect_from_scaling_down,omitempty"`
	}{
		{
			InstanceId:               "e1040c67-b130-41ee-9405-07c6c6c20208",
			ScalingGroupId:           "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
			ScalingGroupName:         "as-group-Test",
			LifeCycleState:           "INSERVICE",
			HealthStatus:             "NORMAL",
			ScalingConfigurationName: "",
			ScalingConfigurationId:   "",
			CreateTime:               "2018-05-12T08:17:54Z",
			InstanceName:             "ecs-SDK",
			ProtectFromScalingDown:   false,
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group_instance/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/list", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group_instance/e1040c67-b130-41ee-9405-07c6c6c20208?instance_delete=no", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
