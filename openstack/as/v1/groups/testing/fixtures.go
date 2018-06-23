package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/groups"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{"scaling_group_id":"60dcec94-5d5b-4dbf-9f50-4ccd7d841432"}
`

var CreateResponse = groups.CreateResponse{
	ScalingGroupId: "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{"scaling_group_id":"60dcec94-5d5b-4dbf-9f50-4ccd7d841432"}
`

var UpdateResponse = groups.UpdateResponse{
	ScalingGroupId: "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group/e4a97959-34d0-4c58-ab85-1dda4d0b9d11", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "scaling_group": {
    "scaling_group_name": "as-group-matt",
    "scaling_group_id": "e4a97959-34d0-4c58-ab85-1dda4d0b9d11",
    "scaling_group_status": "INSERVICE",
    "scaling_configuration_id": "f109bb4f-09f0-4dac-9115-6b429d548750",
    "scaling_configuration_name": "as-config-sh2u",
    "current_instance_number": 0,
    "desire_instance_number": 0,
    "min_instance_number": 0,
    "max_instance_number": 1,
    "cool_down_time": 900,
    "lb_listener_id": null,
    "lbaas_listeners": [],
    "networks": [{
      "id": "f6a0db7b-397c-4162-bc35-9a1f008b4373"
    }],
    "cloud_location_id": null,
    "available_zones": ["cn-north-1a"],
    "security_groups": [{
      "id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }],
    "create_time": "2018-05-09T16:12:57Z",
    "vpc_id": "773c3c42-d315-417b-9063-87091713148c",
    "detail": null,
    "health_periodic_audit_method": "NOVA_AUDIT",
    "health_periodic_audit_time": 5,
    "instance_terminate_policy": "OLD_CONFIG_OLD_INSTANCE",
    "is_scaling": false,
    "notifications": [],
    "delete_publicip": true
  }
}
`
var GetResponse = groups.GetResponse{
	ScalingGroup: groups.ScalingGroup{
		ScalingGroupName:         "as-group-matt",
		ScalingGroupId:           "e4a97959-34d0-4c58-ab85-1dda4d0b9d11",
		ScalingGroupStatus:       "INSERVICE",
		ScalingConfigurationId:   "f109bb4f-09f0-4dac-9115-6b429d548750",
		ScalingConfigurationName: "as-config-sh2u",
		CurrentInstanceNumber:    0,
		DesireInstanceNumber:     0,
		MinInstanceNumber:        0,
		MaxInstanceNumber:        1,
		CoolDownTime:             900,
		LbListenerId:             "",
		LbaasListeners:           []string{},
		Networks: []struct {
			// Specifies the network ID.
			Id string `json:"id,omitempty"`
		}{
			{Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373"},
		},
		CloudLocationId: "",
		AvailableZones:  []string{"cn-north-1a"},
		SecurityGroups: []struct {
			// Specifies the ID of the security group.
			Id string `json:"id,omitempty"`
		}{
			{Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"},
		},
		CreateTime:                "2018-05-09T16:12:57Z",
		VpcId:                     "773c3c42-d315-417b-9063-87091713148c",
		Detail:                    "",
		HealthPeriodicAuditMethod: "NOVA_AUDIT",
		HealthPeriodicAuditTime:   5,
		InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
		IsScaling:                 false,
		Notifications:             []string{},
		DeletePublicip:            true,
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group/e4a97959-34d0-4c58-ab85-1dda4d0b9d11", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "scaling_groups": [{
    "scaling_group_name": "as-group-Test",
    "scaling_group_id": "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
    "scaling_group_status": "PAUSED",
    "scaling_configuration_id": "f109bb4f-09f0-4dac-9115-6b429d548750",
    "scaling_configuration_name": "as-config-sh2u",
    "current_instance_number": 0,
    "desire_instance_number": 0,
    "min_instance_number": 0,
    "max_instance_number": 1,
    "cool_down_time": 900,
    "lb_listener_id": null,
    "lbaas_listeners": [],
    "networks": [{
      "id": "f6a0db7b-397c-4162-bc35-9a1f008b4373"
    }],
    "cloud_location_id": null,
    "available_zones": ["cn-north-1a"],
    "security_groups": [{
      "id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }],
    "create_time": "2018-05-10T17:22:54Z",
    "vpc_id": "773c3c42-d315-417b-9063-87091713148c",
    "detail": null,
    "health_periodic_audit_method": "NOVA_AUDIT",
    "health_periodic_audit_time": 5,
    "instance_terminate_policy": "OLD_CONFIG_OLD_INSTANCE",
    "is_scaling": false,
    "notifications": [],
    "delete_publicip": true
  }, {
    "scaling_group_name": "as-group-matt",
    "scaling_group_id": "e4a97959-34d0-4c58-ab85-1dda4d0b9d11",
    "scaling_group_status": "INSERVICE",
    "scaling_configuration_id": "f109bb4f-09f0-4dac-9115-6b429d548750",
    "scaling_configuration_name": "as-config-sh2u",
    "current_instance_number": 0,
    "desire_instance_number": 0,
    "min_instance_number": 0,
    "max_instance_number": 1,
    "cool_down_time": 900,
    "lb_listener_id": null,
    "lbaas_listeners": [],
    "networks": [{
      "id": "f6a0db7b-397c-4162-bc35-9a1f008b4373"
    }],
    "cloud_location_id": null,
    "available_zones": ["cn-north-1a"],
    "security_groups": [{
      "id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }],
    "create_time": "2018-05-09T16:12:57Z",
    "vpc_id": "773c3c42-d315-417b-9063-87091713148c",
    "detail": null,
    "health_periodic_audit_method": "NOVA_AUDIT",
    "health_periodic_audit_time": 5,
    "instance_terminate_policy": "OLD_CONFIG_OLD_INSTANCE",
    "is_scaling": false,
    "notifications": [],
    "delete_publicip": true
  }],
  "total_number": 2,
  "start_number": 0,
  "limit": 2
}
`

var ListResponse = groups.ListResponse{
	TotalNumber: 2,
	StartNumber: 0,
	Limit:       2,
	ScalingGroups: []groups.ScalingGroup{
		{
			ScalingGroupName:         "as-group-Test",
			ScalingGroupId:           "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
			ScalingGroupStatus:       "PAUSED",
			ScalingConfigurationId:   "f109bb4f-09f0-4dac-9115-6b429d548750",
			ScalingConfigurationName: "as-config-sh2u",
			CurrentInstanceNumber:    0,
			DesireInstanceNumber:     0,
			MinInstanceNumber:        0,
			MaxInstanceNumber:        1,
			CoolDownTime:             900,
			LbListenerId:             "",
			LbaasListeners:           []string{},
			Networks: []struct {
				// Specifies the network ID.
				Id string `json:"id,omitempty"`
			}{
				{Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373"},
			},
			CloudLocationId: "",
			AvailableZones:  []string{"cn-north-1a"},
			SecurityGroups: []struct {
				// Specifies the ID of the security group.
				Id string `json:"id,omitempty"`
			}{
				{Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"},
			},
			CreateTime:                "2018-05-10T17:22:54Z",
			VpcId:                     "773c3c42-d315-417b-9063-87091713148c",
			Detail:                    "",
			HealthPeriodicAuditMethod: "NOVA_AUDIT",
			HealthPeriodicAuditTime:   5,
			InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
			IsScaling:                 false,
			Notifications:             []string{},
			DeletePublicip:            true,
		},
		{
			ScalingGroupName:         "as-group-matt",
			ScalingGroupId:           "e4a97959-34d0-4c58-ab85-1dda4d0b9d11",
			ScalingGroupStatus:       "INSERVICE",
			ScalingConfigurationId:   "f109bb4f-09f0-4dac-9115-6b429d548750",
			ScalingConfigurationName: "as-config-sh2u",
			CurrentInstanceNumber:    0,
			DesireInstanceNumber:     0,
			MinInstanceNumber:        0,
			MaxInstanceNumber:        1,
			CoolDownTime:             900,
			LbListenerId:             "",
			LbaasListeners:           []string{},
			Networks: []struct {
				// Specifies the network ID.
				Id string `json:"id,omitempty"`
			}{
				{Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373"},
			},
			CloudLocationId: "",
			AvailableZones:  []string{"cn-north-1a"},
			SecurityGroups: []struct {
				// Specifies the ID of the security group.
				Id string `json:"id,omitempty"`
			}{
				{Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"},
			},
			CreateTime:                "2018-05-09T16:12:57Z",
			VpcId:                     "773c3c42-d315-417b-9063-87091713148c",
			Detail:                    "",
			HealthPeriodicAuditMethod: "NOVA_AUDIT",
			HealthPeriodicAuditTime:   5,
			InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
			IsScaling:                 false,
			Notifications:             []string{},
			DeletePublicip:            true,
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleEnableSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group/e4a97959-34d0-4c58-ab85-1dda4d0b9d11/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_group/e4a97959-34d0-4c58-ab85-1dda4d0b9d11", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
