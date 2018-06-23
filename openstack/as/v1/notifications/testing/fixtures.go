package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/notifications"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var EnableOutput = `
{
  "topic_urn": "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
  "topic_scene": ["SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL"],
  "topic_name": "Topic1"
}
`

var EnableResponse = notifications.EnableResponse{
	TopicUrn:   "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
	TopicScene: []string{"SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL"},
}

func HandleEnableSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_notification/60dcec94-5d5b-4dbf-9f50-4ccd7d841432", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, EnableOutput)
	})
}

var ListOutput = `
{
  "topics": [{
    "topic_urn": "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
    "topic_scene": ["SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL"],
    "topic_name": "Topic1"
  }]
}
`

var ListResponse = notifications.ListResponse{
	Topics: []struct {
		TopicUrn   string   `json:"topic_urn,omitempty"`
		TopicScene []string `json:"topic_scene,omitempty"`
		TopicName  string   `json:"topic_name,omitempty"`
	}{
		{
			TopicUrn:   "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
			TopicScene: []string{"SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL"},
			TopicName:  "Topic1",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_notification/60dcec94-5d5b-4dbf-9f50-4ccd7d841432", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_notification/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
