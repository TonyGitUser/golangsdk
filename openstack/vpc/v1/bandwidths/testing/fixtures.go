package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/vpc/v1/bandwidths"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var UpdateOutput = `
{
  "bandwidth": {
    "id": "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
    "name": "bandwidth-ABCD",
    "size": 1,
    "share_type": "PER",
    "publicip_info": [{
      "publicip_id": "3faa05bd-d878-44e2-a363-f6672a9761d3",
      "publicip_address": "49.4.22.32",
      "publicip_type": "5_bgp"
    }],
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "bandwidth_type": "bgp",
    "charge_mode": "bandwidth",
    "billing_info": "CS1803212335SQ1UD:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626"
  }
}
`

var UpdateResponse = bandwidths.UpdateResponse{
	Bandwidth: bandwidths.BandWidth{
		Id:        "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
		Name:      "bandwidth-ABCD",
		Size:      1,
		ShareType: "PER",
		PublicipInfo: []struct {
			PublicipId      string `json:"publicip_id,"`
			PublicipAddress string `json:"publicip_address,"`
			PublicipType    string `json:"publicip_type,"`
		}{
			{
				PublicipId:      "3faa05bd-d878-44e2-a363-f6672a9761d3",
				PublicipAddress: "49.4.22.32",
				PublicipType:    "5_bgp",
			},
		},
		TenantId:      "57e98940a77f4bb988a21a7d0603a626",
		BandwidthType: "bgp",
		ChargeMode:    "bandwidth",
		BillingInfo:   "CS1803212335SQ1UD:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626",
	},
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/bandwidths/3c43e46e-4af1-45b8-a84d-ee6d04488d2a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "bandwidth": {
    "id": "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
    "name": "bandwidth-6f78",
    "size": 1,
    "share_type": "PER",
    "publicip_info": [{
      "publicip_id": "3faa05bd-d878-44e2-a363-f6672a9761d3",
      "publicip_address": "49.4.22.32",
      "publicip_type": "5_bgp"
    }],
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "bandwidth_type": "bgp",
    "charge_mode": "bandwidth",
    "billing_info": "CS1803212335SQ1UD:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626"
  }
}
`

var GetResponse = bandwidths.GetResponse{
	Bandwidth: bandwidths.BandWidth{
		Id:        "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
		Name:      "bandwidth-6f78",
		Size:      1,
		ShareType: "PER",
		PublicipInfo: []struct {
			PublicipId      string `json:"publicip_id,"`
			PublicipAddress string `json:"publicip_address,"`
			PublicipType    string `json:"publicip_type,"`
		}{
			{
				PublicipId:      "3faa05bd-d878-44e2-a363-f6672a9761d3",
				PublicipAddress: "49.4.22.32",
				PublicipType:    "5_bgp",
			},
		},
		TenantId:      "57e98940a77f4bb988a21a7d0603a626",
		BandwidthType: "bgp",
		ChargeMode:    "bandwidth",
		BillingInfo:   "CS1803212335SQ1UD:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626",
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/bandwidths/3c43e46e-4af1-45b8-a84d-ee6d04488d2a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "bandwidths": [{
    "id": "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
    "name": "bandwidth-6f78",
    "size": 1,
    "share_type": "PER",
    "publicip_info": [{
      "publicip_id": "3faa05bd-d878-44e2-a363-f6672a9761d3",
      "publicip_address": "49.4.22.32",
      "publicip_type": "5_bgp"
    }],
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "bandwidth_type": "bgp",
    "charge_mode": "bandwidth",
    "billing_info": "CS1803212335SQ1UD:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626"
  }, {
    "id": "99cb8fb5-6f23-47c3-84a6-6ac4b7729c73",
    "name": "bandwidth-b540",
    "size": 1,
    "share_type": "PER",
    "publicip_info": [{
      "publicip_id": "4d60bba4-0791-4e82-8262-9bdffaeb1d14",
      "publicip_address": "49.4.4.36",
      "publicip_type": "5_bgp"
    }],
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "bandwidth_type": "bgp",
    "charge_mode": "bandwidth",
    "billing_info": "CS1803212337UFH9F:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626"
  }]
}
`

var ListResponse = bandwidths.ListResponse{
	Bandwidths: []bandwidths.BandWidth{
		{
			Id:        "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
			Name:      "bandwidth-6f78",
			Size:      1,
			ShareType: "PER",
			PublicipInfo: []struct {
				PublicipId      string `json:"publicip_id,"`
				PublicipAddress string `json:"publicip_address,"`
				PublicipType    string `json:"publicip_type,"`
			}{
				{
					PublicipId:      "3faa05bd-d878-44e2-a363-f6672a9761d3",
					PublicipAddress: "49.4.22.32",
					PublicipType:    "5_bgp",
				},
			},
			TenantId:      "57e98940a77f4bb988a21a7d0603a626",
			BandwidthType: "bgp",
			ChargeMode:    "bandwidth",
			BillingInfo:   "CS1803212335SQ1UD:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626",
		},
		{
			Id:        "99cb8fb5-6f23-47c3-84a6-6ac4b7729c73",
			Name:      "bandwidth-b540",
			Size:      1,
			ShareType: "PER",
			PublicipInfo: []struct {
				PublicipId      string `json:"publicip_id,"`
				PublicipAddress string `json:"publicip_address,"`
				PublicipType    string `json:"publicip_type,"`
			}{
				{
					PublicipId:      "4d60bba4-0791-4e82-8262-9bdffaeb1d14",
					PublicipAddress: "49.4.4.36",
					PublicipType:    "5_bgp",
				},
			},
			TenantId:      "57e98940a77f4bb988a21a7d0603a626",
			BandwidthType: "bgp",
			ChargeMode:    "bandwidth",
			BillingInfo:   "CS1803212337UFH9F:60564b0c1f484feaad26cda1acc6d4d8:cn-north-1:57e98940a77f4bb988a21a7d0603a626",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/bandwidths", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
