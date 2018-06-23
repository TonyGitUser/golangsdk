package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/backups"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

func HandleCreatePolicySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/6926d5168444416fa28646de8a67450fno01/backups/policy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "{}")
	})
}

var GetPolicyOutput = `
{"policy":{"starttime":"00:00:00","keepday":7}}
`

var GetPolicyResponse = backups.GetPolicyResponse{
	Policy: struct {
		Keepday   int    `json:"keepday,omitempty"`
		Starttime string `json:"starttime,omitempty"`
	}{
		Keepday:   7,
		Starttime: "00:00:00",
	},
}

func HandleGetPolicySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/6926d5168444416fa28646de8a67450fno01/backups/policy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetPolicyOutput)
	})
}

var BackupOutput = `
{
  "backup": {
    "created": "2018-05-20T14:47:09",
    "dataStore": {
      "type": "MySQL",
      "version": "5.6.39",
      "version_id": "9dae9226-5f8e-4b10-bd41-adae2e693e89"
    },
    "description": "My Backup",
    "id": "1e24c1bfe18647a890fc0aeb7e775616br01",
    "instance_id": "6926d5168444416fa28646de8a67450fno01",
    "locationRef": null,
    "name": "backup",
    "parent_id": null,
    "size": null,
    "status": "BUILDING",
    "updated": "2018-05-20T14:47:09",
    "backuptype": "1"
  },
  "extendparam": {
    "jobs": ["1a462af6-85db-4029-943d-544ceadd5119"]
  }
}
`

var BackupResponse = backups.BackupResponse{
	Backup: struct {
		Created   string `json:"created,omitempty"`
		DataStore struct {
			Type      string `json:"type,omitempty"`
			Version   string `json:"version,omitempty"`
			VersionId string `json:"version_id,omitempty"`
		} `json:"dataStore,omitempty"`
		Description string  `json:"description,omitempty"`
		Id          string  `json:"id,omitempty"`
		InstanceId  string  `json:"instance_id,omitempty"`
		LocationRef string  `json:"locationRef,omitempty"`
		Name        string  `json:"name,omitempty"`
		ParentId    string  `json:"parent_id,omitempty"`
		Size        float64 `json:"size,omitempty"`
		Status      string  `json:"status,omitempty"`
		Updated     string  `json:"updated,omitempty"`
		Backuptype  string  `json:"backuptype,omitempty"`
	}{
		Created: "2018-05-20T14:47:09",
		DataStore: struct {
			Type      string `json:"type,omitempty"`
			Version   string `json:"version,omitempty"`
			VersionId string `json:"version_id,omitempty"`
		}{
			Type:      "MySQL",
			Version:   "5.6.39",
			VersionId: "9dae9226-5f8e-4b10-bd41-adae2e693e89",
		},
		Description: "My Backup",
		Id:          "1e24c1bfe18647a890fc0aeb7e775616br01",
		InstanceId:  "6926d5168444416fa28646de8a67450fno01",
		LocationRef: "",
		Name:        "backup",
		ParentId:    "",
		Status:      "BUILDING",
		Updated:     "2018-05-20T14:47:09",
		Backuptype:  "1",
	},
	Extendparam: struct {
		// Indicates the task ID.
		Jobs []string `json:"jobs,omitempty"`
	}{
		Jobs: []string{"1a462af6-85db-4029-943d-544ceadd5119"},
	},
}

func HandleBackupSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/backups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, BackupOutput)
	})
}

var ListOutput = `
{
  "backups": [{
    "created": "2018-05-20T14:47:09",
    "dataStore": {
      "type": "MySQL",
      "version": "5.6.39",
      "version_id": "9dae9226-5f8e-4b10-bd41-adae2e693e89"
    },
    "description": "My Backup",
    "id": "1e24c1bfe18647a890fc0aeb7e775616br01",
    "instance_id": "6926d5168444416fa28646de8a67450fno01",
    "locationRef": "",
    "name": "backup",
    "parent_id": null,
    "size": 0.002175,
    "status": "COMPLETED",
    "updated": "2018-05-20T14:48:54",
    "backuptype": "1"
  }]
}
`

var ListResponse = backups.ListResponse{
	Backups: []struct {
		Id          string  `json:"id,omitempty"`
		Name        string  `json:"name,omitempty"`
		Description string  `json:"description,omitempty"`
		LocationRef string  `json:"locationRef,omitempty"`
		Created     string  `json:"created,omitempty"`
		Updated     string  `json:"updated,omitempty"`
		Size        float64 `json:"size,omitempty"`
		Status      string  `json:"status,omitempty"`
		Backuptype  string  `json:"backuptype,omitempty"`
		DataStore   struct {
			Type      string `json:"type,omitempty"`
			Version   string `json:"version,omitempty"`
			VersionId string `json:"version_id,omitempty"`
		} `json:"dataStore,omitempty"`
		InstanceId string `json:"instance_id,omitempty"`
		ParentId   string `json:"parent_id,omitempty"`
	}{
		{
			Created: "2018-05-20T14:47:09",
			DataStore: struct {
				Type      string `json:"type,omitempty"`
				Version   string `json:"version,omitempty"`
				VersionId string `json:"version_id,omitempty"`
			}{
				Type:      "MySQL",
				Version:   "5.6.39",
				VersionId: "9dae9226-5f8e-4b10-bd41-adae2e693e89",
			},
			Description: "My Backup",
			Id:          "1e24c1bfe18647a890fc0aeb7e775616br01",
			InstanceId:  "6926d5168444416fa28646de8a67450fno01",
			LocationRef: "",
			Size:        0.002175,
			Name:        "backup",
			ParentId:    "",
			Status:      "COMPLETED",
			Updated:     "2018-05-20T14:48:54",
			Backuptype:  "1",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/backups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var RestoreOutput = `
{"extendparam":{"jobs":[{"id":"6d020c8f-ab41-428a-b7c5-50abd2271e47"}]}}
`

var RestoreResponse = backups.RestoreResponse{
	Extendparam: struct {
		Jobs []struct {
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	}{
		Jobs: []struct {
			Id string `json:"id,omitempty"`
		}{
			{
				Id: "6d020c8f-ab41-428a-b7c5-50abd2271e47",
			},
		},
	},
}

func HandleRestoreSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/6926d5168444416fa28646de8a67450fno01/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RestoreOutput)
	})
}

var RestoreWithNewOutput = `
{
  "instance": {
    "id": "941295f7a35a472cac4f6ecfc3d56454no01",
    "status": "BUILD",
    "name": "trove-newinstance_node0",
    "created": "",
    "hostname": "",
    "type": "master",
    "region": "cn-north-1",
    "updated": "",
    "availabilityZone": "cn-north-1a",
    "vpc": "773c3c42-d315-417b-9063-87091713148c",
    "nics": {
      "subnetId": "f6a0db7b-397c-4162-bc35-9a1f008b4373"
    },
    "securityGroup": {
      "id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    },
    "flavor": {
      "id": "18207182-0c02-467f-ae05-9c1650df1717"
    },
    "volume": {
      "type": "COMMON",
      "size": 100
    },
    "dataStoreInfo": null,
    "dbPort": 8635,
    "extendparam": {
      "jobs": [{
        "id": "da136b51-0079-444e-9956-5916095fd400"
      }]
    },
    "backupStrategy": {
      "startTime": "02:00:00",
      "keepDays": 7
    }
  }
}
`

var RestoreWithNewResponse = backups.RestoreWithNewResponse{
	Instance: struct {
		Id               string `json:"id,omitempty"`
		Status           string `json:"status,omitempty"`
		Name             string `json:"name,omitempty"`
		Created          string `json:"created,omitempty"`
		Hostname         string `json:"hostname,omitempty"`
		Type             string `json:"type,omitempty"`
		Region           string `json:"region,omitempty"`
		Updated          string `json:"updated,omitempty"`
		AvailabilityZone string `json:"availabilityZone,omitempty"`
		Vpc              string `json:"vpc,omitempty"`
		Nics             struct {
			SubnetId string `json:"subnetId,omitempty"`
		} `json:"nics,omitempty"`
		SecurityGroup struct {
			Id string `json:"id,omitempty"`
		} `json:"securityGroup,omitempty"`
		Flavor struct {
			Id string `json:"id,omitempty"`
		} `json:"flavor,omitempty"`
		Volume struct {
			Type string `json:"type,omitempty"`
			Size int    `json:"size,omitempty"`
		} `json:"volume,omitempty"`
		DbPort        int    `json:"dbPort,omitempty"`
		DataStoreInfo string `json:"dataStoreInfo,omitempty"`
		Extendparam   struct {
			Jobs []struct {
				Id string `json:"id,omitempty"`
			} `json:"jobs,omitempty"`
		} `json:"extendparam,omitempty"`
		BackupStrategy struct {
			StartTime string `json:"startTime,omitempty"`
			KeepDays  int    `json:"keepDays,omitempty"`
		} `json:"backupStrategy,omitempty"`
	}{
		Id:               "941295f7a35a472cac4f6ecfc3d56454no01",
		Status:           "BUILD",
		Name:             "trove-newinstance_node0",
		Created:          "",
		Hostname:         "",
		Type:             "master",
		Region:           "cn-north-1",
		Updated:          "",
		AvailabilityZone: "cn-north-1a",
		Vpc:              "773c3c42-d315-417b-9063-87091713148c",
		Nics: struct {
			SubnetId string `json:"subnetId,omitempty"`
		}{
			SubnetId: "f6a0db7b-397c-4162-bc35-9a1f008b4373",
		},
		SecurityGroup: struct {
			Id string `json:"id,omitempty"`
		}{
			Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
		},
		Flavor: struct {
			Id string `json:"id,omitempty"`
		}{
			Id: "18207182-0c02-467f-ae05-9c1650df1717",
		},
		Volume: struct {
			Type string `json:"type,omitempty"`
			Size int    `json:"size,omitempty"`
		}{
			Type: "COMMON",
			Size: 100,
		},
		DataStoreInfo: "",
		DbPort:        8635,
		Extendparam: struct {
			Jobs []struct {
				Id string `json:"id,omitempty"`
			} `json:"jobs,omitempty"`
		}{
			Jobs: []struct {
				Id string `json:"id,omitempty"`
			}{
				{
					Id: "da136b51-0079-444e-9956-5916095fd400",
				},
			},
		},
		BackupStrategy: struct {
			StartTime string `json:"startTime,omitempty"`
			KeepDays  int    `json:"keepDays,omitempty"`
		}{
			StartTime: "02:00:00",
			KeepDays:  7,
		},
	},
}

func HandleRestoreWithNewSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RestoreWithNewOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/backups/1e24c1bfe18647a890fc0aeb7e775616br01", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
