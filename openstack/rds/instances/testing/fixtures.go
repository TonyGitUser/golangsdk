package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/instances"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
{
  "instances": [{
    "id": "6926d5168444416fa28646de8a67450fno01",
    "status": "ACTIVE",
    "name": "MYSQL_TEST_node0",
    "created": "2018-04-16T14:36:17+0000",
    "hostname": "192.168.1.182",
    "type": "master",
    "region": "cn-north-1",
    "updated": "2018-04-16T14:45:25+0000",
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
    "dataStoreInfo": {
      "type": "MySQL",
      "version": "5.6.39"
    },
    "dbPort": 8635,
    "backupStrategy": {
      "startTime": "19:40:00",
      "keepDays": 7
    }
  }]
}
`

var ListResponse = instances.ListResponse{
	Instances: []struct {
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
		DataStoreInfo struct {
			Type    string `json:"type,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"dataStoreInfo,omitempty"`
		DbPort         int `json:"dbPort,omitempty"`
		BackupStrategy struct {
			StartTime string `json:"startTime,omitempty"`
			KeepDays  int    `json:"keepDays,omitempty"`
		} `json:"backupStrategy,omitempty"`
		SlaveId string `json:"slaveId,omitempty"`
		Ha      []struct {
			ReplicationMode string `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`
		ReplicaOf string `json:"replicaOf,omitempty"`
	}{
		{
			Id:               "6926d5168444416fa28646de8a67450fno01",
			Status:           "ACTIVE",
			Name:             "MYSQL_TEST_node0",
			Created:          "2018-04-16T14:36:17+0000",
			Hostname:         "192.168.1.182",
			Type:             "master",
			Region:           "cn-north-1",
			Updated:          "2018-04-16T14:45:25+0000",
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
			DataStoreInfo: struct {
				Type    string `json:"type,omitempty"`
				Version string `json:"version,omitempty"`
			}{
				Type:    "MySQL",
				Version: "5.6.39",
			},
			DbPort: 8635,
			BackupStrategy: struct {
				StartTime string `json:"startTime,omitempty"`
				KeepDays  int    `json:"keepDays,omitempty"`
			}{
				StartTime: "19:40:00",
				KeepDays:  7,
			},
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var GetOutput = `
{
  "instance": {
    "id": "6926d5168444416fa28646de8a67450fno01",
    "status": "ACTIVE",
    "name": "MYSQL_TEST_node0",
    "created": "2018-04-16T14:36:17+0000",
    "hostname": "192.168.1.182",
    "type": "master",
    "region": "cn-north-1",
    "updated": "2018-04-16T14:45:25+0000",
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
      "size": 100,
      "used": 0.2
    },
    "dataStoreInfo": {
      "type": "MySQL",
      "version": "5.6.39"
    },
    "dbPort": 8635,
    "backupStrategy": {
      "startTime": "19:40:00",
      "keepDays": 7
    }
  }
}
`

var GetResponse = instances.GetResponse{
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
		DataStoreInfo struct {
			Type    string `json:"type,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"dataStoreInfo,omitempty"`
		DbPort         int `json:"dbPort,omitempty"`
		BackupStrategy struct {
			StartTime string `json:"startTime,omitempty"`
			KeepDays  int    `json:"keepDays,omitempty"`
		} `json:"backupStrategy,omitempty"`
		SlaveId string `json:"slaveId,omitempty"`
		Ha      []struct {
			ReplicationMode string `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`
		ReplicaOf string `json:"replicaOf,omitempty"`
	}{
		Id:               "6926d5168444416fa28646de8a67450fno01",
		Status:           "ACTIVE",
		Name:             "MYSQL_TEST_node0",
		Created:          "2018-04-16T14:36:17+0000",
		Hostname:         "192.168.1.182",
		Type:             "master",
		Region:           "cn-north-1",
		Updated:          "2018-04-16T14:45:25+0000",
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
		DataStoreInfo: struct {
			Type    string `json:"type,omitempty"`
			Version string `json:"version,omitempty"`
		}{
			Type:    "MySQL",
			Version: "5.6.39",
		},
		DbPort: 8635,
		BackupStrategy: struct {
			StartTime string `json:"startTime,omitempty"`
			KeepDays  int    `json:"keepDays,omitempty"`
		}{
			StartTime: "19:40:00",
			KeepDays:  7,
		},
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/6926d5168444416fa28646de8a67450fno01", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListFlavorsOutput = `
{
  "flavors": [{
    "id": "a6ce31c1-d1f3-44c4-b872-8bc2ca639a64",
    "name": "OTC_MYCM_LARGE_RR",
    "ram": 8000,
    "specCode": "rds.dec.mysql.s1.large.rr"
  }, {
    "id": "1739acca-ceef-4479-afd3-d49f5c171352",
    "name": "OTC_MYCM_XLARGE_RR",
    "ram": 16000,
    "specCode": "rds.dec.mysql.s1.xlarge.rr"
  }, {
    "id": "f1a9db80-b2b4-4dc9-8eb0-c70cc3806db6",
    "name": "OTC_MYCM_2XLARGE_RR",
    "ram": 32000,
    "specCode": "rds.dec.mysql.s1.2xlarge.rr"
  }, {
    "id": "d3a3c623-9e73-4010-b964-11192c5254e2",
    "name": "OTC_MYCM_4XLARGE_RR",
    "ram": 64000,
    "specCode": "rds.dec.mysql.s1.4xlarge.rr"
  }, {
    "id": "4d83ec7f-2ff2-4353-a772-f4408b2e9956",
    "name": "OTC_MYCM_8XLARGE_RR",
    "ram": 128000,
    "specCode": "rds.dec.mysql.s1.8xlarge.rr"
  }, {
    "id": "ab27fcfc-d5ee-44c7-8695-f6527998b26d",
    "name": "OTC_MYCM_MEDIUM_RR",
    "ram": 4000,
    "specCode": "rds.dec.mysql.s1.medium.rr"
  }, {
    "id": "b6d65f18-bdf1-44ea-adad-bc46423aab46",
    "name": "OTC_MYCM_MEDIUM_HA",
    "ram": 4000,
    "specCode": "rds.dec.mysql.s1.medium.ha"
  }, {
    "id": "3af07c95-e315-47f0-9251-50721fe44d64",
    "name": "OTC_MYCM_MEDIUM",
    "ram": 4000,
    "specCode": "rds.dec.mysql.s1.medium"
  }, {
    "id": "8a7a6165-1f1f-462b-b28a-aead6b0a4ac9",
    "name": "OTC_MYCM_LARGE",
    "ram": 8000,
    "specCode": "rds.dec.mysql.s1.large"
  }, {
    "id": "20ac694f-5d9d-4a98-838e-ad696548bf5c",
    "name": "OTC_MYCM_XLARGE",
    "ram": 16000,
    "specCode": "rds.dec.mysql.s1.xlarge"
  }, {
    "id": "65708628-36bf-4191-bb8f-f99dafe53fb0",
    "name": "OTC_MYCM_2XLARGE",
    "ram": 32000,
    "specCode": "rds.dec.mysql.s1.2xlarge"
  }, {
    "id": "18207182-0c02-467f-ae05-9c1650df1717",
    "name": "OTC_MYCM_MEDIUM",
    "ram": 4000,
    "specCode": "rds.mysql.s1.medium"
  }, {
    "id": "f20f3a7c-9bc1-4f07-aa62-336b387d0bf2",
    "name": "OTC_MYCM_LARGE",
    "ram": 8000,
    "specCode": "rds.mysql.s1.large"
  }, {
    "id": "39c59aff-0791-475c-b135-2505c26f945f",
    "name": "OTC_MYCM_XLARGE",
    "ram": 16000,
    "specCode": "rds.mysql.s1.xlarge"
  }, {
    "id": "58010308-0922-4b79-92a6-1866bca062e9",
    "name": "OTC_MYCM_2XLARGE",
    "ram": 32000,
    "specCode": "rds.mysql.s1.2xlarge"
  }, {
    "id": "57e94375-8ff0-4981-983a-d5e481c8c567",
    "name": "OTC_MYCM_4XLARGE",
    "ram": 64000,
    "specCode": "rds.mysql.s1.4xlarge"
  }, {
    "id": "30a0baf6-d4c2-44d1-8569-68844512ae3a",
    "name": "OTC_MYCM_8XLARGE",
    "ram": 128000,
    "specCode": "rds.mysql.s1.8xlarge"
  }]
}
`

var ListFlavorsResponse = instances.ListFlavorsResponse{
	Flavors: []struct {
		Id       string `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Ram      int    `json:"ram,omitempty"`
		SpecCode string `json:"specCode,omitempty"`
	}{
		{
			Id:       "a6ce31c1-d1f3-44c4-b872-8bc2ca639a64",
			Name:     "OTC_MYCM_LARGE_RR",
			Ram:      8000,
			SpecCode: "rds.dec.mysql.s1.large.rr",
		}, {
			Id:       "1739acca-ceef-4479-afd3-d49f5c171352",
			Name:     "OTC_MYCM_XLARGE_RR",
			Ram:      16000,
			SpecCode: "rds.dec.mysql.s1.xlarge.rr",
		}, {
			Id:       "f1a9db80-b2b4-4dc9-8eb0-c70cc3806db6",
			Name:     "OTC_MYCM_2XLARGE_RR",
			Ram:      32000,
			SpecCode: "rds.dec.mysql.s1.2xlarge.rr",
		}, {
			Id:       "d3a3c623-9e73-4010-b964-11192c5254e2",
			Name:     "OTC_MYCM_4XLARGE_RR",
			Ram:      64000,
			SpecCode: "rds.dec.mysql.s1.4xlarge.rr",
		}, {
			Id:       "4d83ec7f-2ff2-4353-a772-f4408b2e9956",
			Name:     "OTC_MYCM_8XLARGE_RR",
			Ram:      128000,
			SpecCode: "rds.dec.mysql.s1.8xlarge.rr",
		}, {
			Id:       "ab27fcfc-d5ee-44c7-8695-f6527998b26d",
			Name:     "OTC_MYCM_MEDIUM_RR",
			Ram:      4000,
			SpecCode: "rds.dec.mysql.s1.medium.rr",
		}, {
			Id:       "b6d65f18-bdf1-44ea-adad-bc46423aab46",
			Name:     "OTC_MYCM_MEDIUM_HA",
			Ram:      4000,
			SpecCode: "rds.dec.mysql.s1.medium.ha",
		}, {
			Id:       "3af07c95-e315-47f0-9251-50721fe44d64",
			Name:     "OTC_MYCM_MEDIUM",
			Ram:      4000,
			SpecCode: "rds.dec.mysql.s1.medium",
		}, {
			Id:       "8a7a6165-1f1f-462b-b28a-aead6b0a4ac9",
			Name:     "OTC_MYCM_LARGE",
			Ram:      8000,
			SpecCode: "rds.dec.mysql.s1.large",
		}, {
			Id:       "20ac694f-5d9d-4a98-838e-ad696548bf5c",
			Name:     "OTC_MYCM_XLARGE",
			Ram:      16000,
			SpecCode: "rds.dec.mysql.s1.xlarge",
		}, {
			Id:       "65708628-36bf-4191-bb8f-f99dafe53fb0",
			Name:     "OTC_MYCM_2XLARGE",
			Ram:      32000,
			SpecCode: "rds.dec.mysql.s1.2xlarge",
		}, {
			Id:       "18207182-0c02-467f-ae05-9c1650df1717",
			Name:     "OTC_MYCM_MEDIUM",
			Ram:      4000,
			SpecCode: "rds.mysql.s1.medium",
		}, {
			Id:       "f20f3a7c-9bc1-4f07-aa62-336b387d0bf2",
			Name:     "OTC_MYCM_LARGE",
			Ram:      8000,
			SpecCode: "rds.mysql.s1.large",
		}, {
			Id:       "39c59aff-0791-475c-b135-2505c26f945f",
			Name:     "OTC_MYCM_XLARGE",
			Ram:      16000,
			SpecCode: "rds.mysql.s1.xlarge",
		}, {
			Id:       "58010308-0922-4b79-92a6-1866bca062e9",
			Name:     "OTC_MYCM_2XLARGE",
			Ram:      32000,
			SpecCode: "rds.mysql.s1.2xlarge",
		}, {
			Id:       "57e94375-8ff0-4981-983a-d5e481c8c567",
			Name:     "OTC_MYCM_4XLARGE",
			Ram:      64000,
			SpecCode: "rds.mysql.s1.4xlarge",
		}, {
			Id:       "30a0baf6-d4c2-44d1-8569-68844512ae3a",
			Name:     "OTC_MYCM_8XLARGE",
			Ram:      128000,
			SpecCode: "rds.mysql.s1.8xlarge",
		},
	},
}

func HandleListFlavorsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/flavors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListFlavorsOutput)
	})
}

var GetFlavorOutput = `
{
  "flavor": {
    "id": "30a0baf6-d4c2-44d1-8569-68844512ae3a",
    "name": "OTC_MYCM_8XLARGE",
    "ram": 128000,
    "specCode": "rds.mysql.s1.8xlarge"
  }
}
`

var GetFlavorResponse = instances.GetFlavorResponse{
	Flavor: struct {
		Id       string `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Ram      int    `json:"ram,omitempty"`
		SpecCode string `json:"specCode,omitempty"`
	}{
		Id:       "30a0baf6-d4c2-44d1-8569-68844512ae3a",
		Name:     "OTC_MYCM_8XLARGE",
		Ram:      128000,
		SpecCode: "rds.mysql.s1.8xlarge",
	},
}

func HandleGetFlavorSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/flavors/30a0baf6-d4c2-44d1-8569-68844512ae3a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetFlavorOutput)
	})
}

var CreateOutput = `
{
  "instance": {
    "id": "a0472c5aafac40688c82f8980f3ecd0ano01",
    "status": "BUILD",
    "name": "MYSQL_TEST_CREATE_node0",
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
        "id": "a0dac722-227e-4d86-8f70-a69ba2e500b0"
      }]
    },
    "backupStrategy": {
      "startTime": "19:40:00",
      "keepDays": 7
    }
  }
}
`

var CreateResponse = instances.CreateResponse{
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
		SlaveId string `json:"slaveId,omitempty"`
		Ha      struct {
			ReplicationMode string `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`
		ReplicaOf string `json:"replica_of,omitempty"`
	}{
		Id:               "a0472c5aafac40688c82f8980f3ecd0ano01",
		Status:           "BUILD",
		Name:             "MYSQL_TEST_CREATE_node0",
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
					Id: "a0dac722-227e-4d86-8f70-a69ba2e500b0",
				},
			},
		},
		BackupStrategy: struct {
			StartTime string `json:"startTime,omitempty"`
			KeepDays  int    `json:"keepDays,omitempty"`
		}{
			StartTime: "19:40:00",
			KeepDays:  7,
		},
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{"extendparam":{"jobs":[{"id":"fccf5742-cef9-43bf-abc1-320630bbfc03"}]}}
`

var UpdateResponse = instances.UpdateResponse{
	Extendparam: struct {
		// Indicates the returned  parameter information.
		Jobs []struct {
			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	}{
		Jobs: []struct {
			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		}{
			{
				Id: "fccf5742-cef9-43bf-abc1-320630bbfc03",
			},
		},
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/a0472c5aafac40688c82f8980f3ecd0ano01/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var RebootOutput = `
{"extendparam":{"jobs":[{"id":"2bf3387e9b714dc5978062424574"}]}}
`

var RebootResponse = instances.RebootResponse{
	Extendparam: struct {
		// Indicates the returned  parameter information.
		Jobs []struct {
			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	}{
		Jobs: []struct {
			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		}{
			{
				Id: "2bf3387e9b714dc5978062424574",
			},
		},
	},
}

func HandleRebootSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/a0472c5aafac40688c82f8980f3ecd0ano01/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RebootOutput)
	})
}

var DeleteOutput = `
{"extendparam":{"jobs":[{"id":"2bf3387e9b714dc5978062424574"}]}}
`

var DeleteResponse = instances.DeleteResponse{
	Extendparam: struct {
		// Indicates the returned  parameter information.
		Jobs []struct {
			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	}{
		Jobs: []struct {
			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		}{
			{
				Id: "2bf3387e9b714dc5978062424574",
			},
		},
	},
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/instances/a0472c5aafac40688c82f8980f3ecd0ano01", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DeleteOutput)
	})
}
