package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/configures"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{"scaling_configuration_id":"55f8df1b-47e2-4cf9-844a-acde7469012a"}
`

var CreateResponse = configures.CreateResponse{
	ScalingConfigurationId: "55f8df1b-47e2-4cf9-844a-acde7469012a",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_configuration", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var GetOutput = `
{
  "scaling_configuration": {
    "tenant": "57e98940a77f4bb988a21a7d0603a626",
    "scaling_configuration_id": "f109bb4f-09f0-4dac-9115-6b429d548750",
    "scaling_configuration_name": "as-config-sh2u",
    "instance_config": {
      "instance_name": null,
      "instance_id": null,
      "flavorRef": "s2.small.1",
      "imageRef": "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
      "disk": [{
        "size": 40,
        "volume_type": "SATA",
        "disk_type": "SYS",
        "metadata": null,
        "dedicated_storage_id": null,
        "data_disk_image_id": null,
        "snapshot_id": null
      }],
      "key_name": null,
      "adminPass": "***",
      "personality": null,
      "public_ip": null,
      "user_data": "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
      "metadata": {
        "admin_pass": "***"
      }
    },
    "create_time": "2018-05-09T16:12:57Z"
  }
}
`
var GetResponse = configures.GetResponse{
	ScalingConfiguration: configures.ScalingConfiguration{
		Tenant:                   "57e98940a77f4bb988a21a7d0603a626",
		ScalingConfigurationId:   "f109bb4f-09f0-4dac-9115-6b429d548750",
		ScalingConfigurationName: "as-config-sh2u",
		InstanceConfig: configures.InstanceConfig{
			InstanceName: "",
			InstanceId:   "",
			FlavorRef:    "s2.small.1",
			ImageRef:     "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
			Disk: []configures.Disk{
				{
					Size:               40,
					VolumeType:         "SATA",
					DiskType:           "SYS",
					DedicatedStorageId: "",
					DataDiskImageId:    "",
					SnapshotId:         "",
				},
			},
			KeyName:     "",
			AdminPass:   "***",
			Personality: nil,
			PublicIp:    configures.PublicIp{},
			UserData:    "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
			Metadata: map[string]interface{}{
				"admin_pass": "***",
			},
		},
		CreateTime: "2018-05-09T16:12:57Z",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_configuration/f109bb4f-09f0-4dac-9115-6b429d548750", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "total_number": 3,
  "start_number": 0,
  "limit": 2,
  "scaling_configurations": [{
    "tenant": "57e98940a77f4bb988a21a7d0603a626",
    "scaling_configuration_id": "55f8df1b-47e2-4cf9-844a-acde7469012a",
    "scaling_configuration_name": "as-config-test",
    "instance_config": {
      "instance_name": null,
      "instance_id": null,
      "flavorRef": "s2.small.1",
      "imageRef": "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
      "disk": [{
        "size": 40,
        "volume_type": "SATA",
        "disk_type": "SYS",
        "metadata": null,
        "dedicated_storage_id": null,
        "data_disk_image_id": null,
        "snapshot_id": null
      }],
      "key_name": null,
      "adminPass": "***",
      "personality": null,
      "public_ip": null,
      "user_data": "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
      "metadata": {}
    },
    "create_time": "2018-05-10T16:26:53Z"
  }, {
    "tenant": "57e98940a77f4bb988a21a7d0603a626",
    "scaling_configuration_id": "34147aea-e7bb-4290-8248-794b26f355a6",
    "scaling_configuration_name": "as-config-test",
    "instance_config": {
      "instance_name": null,
      "instance_id": null,
      "flavorRef": "s2.small.1",
      "imageRef": "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
      "disk": [{
        "size": 40,
        "volume_type": "SATA",
        "disk_type": "SYS",
        "metadata": null,
        "dedicated_storage_id": null,
        "data_disk_image_id": null,
        "snapshot_id": null
      }],
      "key_name": null,
      "adminPass": "***",
      "personality": null,
      "public_ip": null,
      "user_data": "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
      "metadata": {}
    },
    "create_time": "2018-05-10T16:12:33Z"
  }]
}
`

var ListResponse = configures.ListResponse{
	TotalNumber: 3,
	StartNumber: 0,
	Limit:       2,
	ScalingConfigurations: []configures.ScalingConfiguration{
		{
			Tenant:                   "57e98940a77f4bb988a21a7d0603a626",
			ScalingConfigurationId:   "55f8df1b-47e2-4cf9-844a-acde7469012a",
			ScalingConfigurationName: "as-config-test",
			InstanceConfig: configures.InstanceConfig{
				InstanceName: "",
				InstanceId:   "",
				FlavorRef:    "s2.small.1",
				ImageRef:     "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
				Disk: []configures.Disk{
					{
						Size:               40,
						VolumeType:         "SATA",
						DiskType:           "SYS",
						DedicatedStorageId: "",
						DataDiskImageId:    "",
						SnapshotId:         "",
					},
				},
				KeyName:     "",
				AdminPass:   "***",
				Personality: nil,
				PublicIp:    configures.PublicIp{},
				UserData:    "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
				Metadata:    map[string]interface{}{},
			},
			CreateTime: "2018-05-10T16:26:53Z",
		},
		{
			Tenant:                   "57e98940a77f4bb988a21a7d0603a626",
			ScalingConfigurationId:   "34147aea-e7bb-4290-8248-794b26f355a6",
			ScalingConfigurationName: "as-config-test",
			InstanceConfig: configures.InstanceConfig{
				InstanceName: "",
				InstanceId:   "",
				FlavorRef:    "s2.small.1",
				ImageRef:     "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
				Disk: []configures.Disk{
					{
						Size:               40,
						VolumeType:         "SATA",
						DiskType:           "SYS",
						DedicatedStorageId: "",
						DataDiskImageId:    "",
						SnapshotId:         "",
					},
				},
				KeyName:     "",
				AdminPass:   "***",
				Personality: nil,
				PublicIp:    configures.PublicIp{},
				UserData:    "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
				Metadata:    map[string]interface{}{},
			},
			CreateTime: "2018-05-10T16:12:33Z",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_configuration", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_configuration/34147aea-e7bb-4290-8248-794b26f355a6", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}

func HandleDeleteWithBatchSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_configurations", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
