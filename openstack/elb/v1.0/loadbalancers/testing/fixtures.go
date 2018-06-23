package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/loadbalancers"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "uri": "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c6632674e0016326ad22c7088f",
  "job_id": "8aace0c6632674e0016326ad22c7088f"
}
`

var CreateResponse = loadbalancers.CreateResponse{
	Uri:   "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c6632674e0016326ad22c7088f",
	JobId: "8aace0c6632674e0016326ad22c7088f",
}

func HandleCreateSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "uri": "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c763267616016326b1d1dd0751",
  "job_id": "8aace0c763267616016326b1d1dd0751"
}
`

var UpdateResponse = loadbalancers.UpdateResponse{
	Uri:   "/v1/57e98940a77f4bb988a21a7d0603a626/jobs/8aace0c763267616016326b1d1dd0751",
	JobId: "8aace0c763267616016326b1d1dd0751",
}

func HandleUpdateSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/loadbalancers/ca082827b61d4902bfaf32e8174e244a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "vip_address": "49.4.9.216",
  "update_time": "2018-05-03 15:49:40",
  "create_time": "2018-05-03 15:44:43",
  "id": "ca082827b61d4902bfaf32e8174e244a",
  "status": "ACTIVE",
  "bandwidth": 1,
  "security_group_id": null,
  "vpc_id": "773c3c42-d315-417b-9063-87091713148c",
  "admin_state_up": 1,
  "vip_subnet_id": null,
  "type": "External",
  "name": "UpdatedLB",
  "description": "TEST LB"
}
`

var GetResponse = loadbalancers.GetResponse{
	VipAddress:      "49.4.9.216",
	UpdateTime:      "2018-05-03 15:49:40",
	CreateTime:      "2018-05-03 15:44:43",
	Id:              "ca082827b61d4902bfaf32e8174e244a",
	Status:          "ACTIVE",
	Bandwidth:       1,
	SecurityGroupId: "",
	VpcId:           "773c3c42-d315-417b-9063-87091713148c",
	AdminStateUp:    1,
	VipSubnetId:     "",
	Type:            "External",
	Name:            "UpdatedLB",
	Description:     "TEST LB",
}

func HandleGetSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/loadbalancers/ca082827b61d4902bfaf32e8174e244a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "loadbalancers": [{
    "vip_address": "49.4.9.241",
    "update_time": "2018-05-03 15:36:59",
    "create_time": "2018-05-03 15:36:59",
    "id": "b503ccbfd61f4a3fb7250dbfc85f0cd8",
    "status": "ACTIVE",
    "bandwidth": 1,
    "security_group_id": null,
    "vpc_id": "773c3c42-d315-417b-9063-87091713148c",
    "admin_state_up": 1,
    "vip_subnet_id": null,
    "type": "External",
    "name": "elb-630a",
    "description": ""
  }],
  "instance_num": "1"
}
`

var ListResponse = loadbalancers.ListResponse{
	Loadbalancers: []struct {
		VipAddress      string `json:"vip_address,omitempty"`
		UpdateTime      string `json:"update_time,omitempty"`
		CreateTime      string `json:"create_time,omitempty"`
		Id              string `json:"id,omitempty"`
		Status          string `json:"status,omitempty"`
		Bandwidth       int    `json:"bandwidth,omitempty"`
		VpcId           string `json:"vpc_id,omitempty"`
		AdminStateUp    int    `json:"admin_state_up,omitempty"`
		VipSubnetId     string `json:"vip_subnet_id,omitempty"`
		Type            string `json:"type,omitempty"`
		Name            string `json:"name,omitempty"`
		Description     string `json:"description,omitempty"`
		SecurityGroupId string `json:"security_group_id,omitempty"`
	}{
		{
			VipAddress:      "49.4.9.241",
			UpdateTime:      "2018-05-03 15:36:59",
			CreateTime:      "2018-05-03 15:36:59",
			Id:              "b503ccbfd61f4a3fb7250dbfc85f0cd8",
			Status:          "ACTIVE",
			Bandwidth:       1,
			SecurityGroupId: "",
			VpcId:           "773c3c42-d315-417b-9063-87091713148c",
			AdminStateUp:    1,
			VipSubnetId:     "",
			Type:            "External",
			Name:            "elb-630a",
			Description:     "",
		},
	},
	InstanceNum: "1",
}

func HandleListSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var DeleteOutput = ""

func HandleDeleteSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v1.0/57e98940a77f4bb988a21a7d0603a626/elbaas/loadbalancers/ca082827b61d4902bfaf32e8174e244a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DeleteOutput)
	})
}
