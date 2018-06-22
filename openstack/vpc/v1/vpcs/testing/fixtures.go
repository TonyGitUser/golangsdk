package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/vpcs"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "vpc": {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC",
    "cidr": "192.168.0.0/16",
    "status": "CREATING",
    "routes": null
  }
}
`

var CreateResponse = vpcs.CreateResponse{
	Vpc: vpcs.VPC{
		Id:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
		Name:   "ABC",
		Cidr:   "192.168.0.0/16",
		Status: "CREATING",
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "vpc": {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC-back",
    "cidr": "192.168.0.0/24",
    "status": "OK",
    "routes": null
  }
}
`

var UpdateResponse = vpcs.UpdateResponse{
	Vpc: vpcs.VPC{
		Id:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
		Name:   "ABC-back",
		Cidr:   "192.168.0.0/24",
		Status: "OK",
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/vpcs/7ffddb5f-6731-43d8-9476-1444aaa40bc0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "vpc": {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC-back",
    "cidr": "192.168.0.0/24",
    "status": "OK",
    "routes": null
  }
}
`

var GetResponse = vpcs.GetResponse{
	Vpc: vpcs.VPC{
		Id:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
		Name:   "ABC-back",
		Cidr:   "192.168.0.0/24",
		Status: "OK",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/vpcs/7ffddb5f-6731-43d8-9476-1444aaa40bc0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "vpcs": [{
    "id": "773c3c42-d315-417b-9063-87091713148c",
    "name": "vpc-c8cb",
    "cidr": "192.168.0.0/16",
    "status": "OK",
    "routes": []
  }, {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC-back",
    "cidr": "192.168.0.0/24",
    "status": "OK",
    "routes": []
  }]
}
`

var ListResponse = vpcs.ListResponse{
	Vpcs: []vpcs.VPC{
		{
			Id:     "773c3c42-d315-417b-9063-87091713148c",
			Name:   "vpc-c8cb",
			Cidr:   "192.168.0.0/16",
			Status: "OK",
		},
		{
			Id:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
			Name:   "ABC-back",
			Cidr:   "192.168.0.0/24",
			Status: "OK",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/vpcs/7ffddb5f-6731-43d8-9476-1444aaa40bc0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
