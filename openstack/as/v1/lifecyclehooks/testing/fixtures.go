package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/lifecyclehooks"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var CreateOutput = `
{
  "default_result": "ABANDON",
  "default_timeout": 3600,
  "notification_topic_urn": "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
  "lifecycle_hook_type": "INSTANCE_LAUNCHING",
  "notification_metadata": null,
  "lifecycle_hook_name": "test-hook1",
  "notification_topic_name": "Topic1",
  "create_time": null
}
`

var CreateResponse = lifecyclehooks.CreateResponse{
	DefaultResult:         "ABANDON",
	DefaultTimeout:        3600,
	NotificationTopicUrn:  "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
	LifecycleHookType:     "INSTANCE_LAUNCHING",
	NotificationMetadata:  "",
	LifecycleHookName:     "test-hook1",
	NotificationTopicName: "Topic1",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_lifecycle_hook/60dcec94-5d5b-4dbf-9f50-4ccd7d841432", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "default_result": "CONTINUE",
  "default_timeout": 3600,
  "notification_topic_urn": "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
  "lifecycle_hook_type": "INSTANCE_LAUNCHING",
  "notification_metadata": null,
  "lifecycle_hook_name": "test-hook1",
  "notification_topic_name": "Topic1",
  "create_time": "2018-05-12T08:44:04Z"
}
`

var UpdateResponse = lifecyclehooks.UpdateResponse{
	DefaultResult:         "CONTINUE",
	DefaultTimeout:        3600,
	NotificationTopicUrn:  "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
	LifecycleHookType:     "INSTANCE_LAUNCHING",
	NotificationMetadata:  "",
	LifecycleHookName:     "test-hook1",
	NotificationTopicName: "Topic1",
	CreateTime:            "2018-05-12T08:44:04Z",
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_lifecycle_hook/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/test-hook1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "default_result": "ABANDON",
  "default_timeout": 3600,
  "notification_topic_urn": "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
  "lifecycle_hook_type": "INSTANCE_LAUNCHING",
  "notification_metadata": null,
  "lifecycle_hook_name": "test-hook1",
  "notification_topic_name": "Topic1",
  "create_time": "2018-05-12T08:44:04Z"
}
`
var GetResponse = lifecyclehooks.GetResponse{
	DefaultResult:         "ABANDON",
	DefaultTimeout:        3600,
	NotificationTopicUrn:  "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
	LifecycleHookType:     "INSTANCE_LAUNCHING",
	NotificationMetadata:  "",
	LifecycleHookName:     "test-hook1",
	NotificationTopicName: "Topic1",
	CreateTime:            "2018-05-12T08:44:04Z",
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_lifecycle_hook/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/test-hook1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "lifecycle_hooks": [{
    "default_result": "ABANDON",
    "default_timeout": 3600,
    "notification_topic_urn": "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
    "lifecycle_hook_type": "INSTANCE_LAUNCHING",
    "notification_metadata": null,
    "lifecycle_hook_name": "test-hook1",
    "notification_topic_name": "Topic1",
    "create_time": "2018-05-12T08:44:04Z"
  }]
}
`

var ListResponse = lifecyclehooks.ListResponse{
	LifecycleHooks: []struct {
		LifecycleHookName     string `json:"lifecycle_hook_name,omitempty"`
		LifecycleHookType     string `json:"lifecycle_hook_type,omitempty"`
		DefaultResult         string `json:"default_result,omitempty"`
		DefaultTimeout        int    `json:"default_timeout,omitempty"`
		NotificationTopicUrn  string `json:"notification_topic_urn,omitempty"`
		NotificationTopicName string `json:"notification_topic_name,omitempty"`
		NotificationMetadata  string `json:"notification_metadata,omitempty"`
		CreateTime            string `json:"create_time,omitempty"`
	}{
		{
			DefaultResult:         "ABANDON",
			DefaultTimeout:        3600,
			NotificationTopicUrn:  "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
			LifecycleHookType:     "INSTANCE_LAUNCHING",
			NotificationMetadata:  "",
			LifecycleHookName:     "test-hook1",
			NotificationTopicName: "Topic1",
			CreateTime:            "2018-05-12T08:44:04Z",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_lifecycle_hook/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/list", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

var ListWithSuspensionOutput = `
{
  "instance_hanging_info": [{
    "lifecycle_action_key": "2c4aabb1-3ac3-4aba-9c6a-261147be6dfb",
    "instance_id": "6abadacf-87af-4225-b762-4a56853aec02",
    "lifecycle_hook_name": "test-hook1",
    "scaling_group_id": "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
    "lifecycle_hook_status": "HANGING",
    "timeout": "2018-05-12T10:16:30Z",
    "default_result": "ABANDON"
  }]
}
`

var ListWithSuspensionResponse = lifecyclehooks.ListWithSuspensionResponse{
	InstanceHangingInfo: []struct {
		LifecycleHookName   string `json:"lifecycle_hook_name,omitempty"`
		LifecycleActionKey  string `json:"lifecycle_action_key,omitempty"`
		InstanceId          string `json:"instance_id,omitempty"`
		ScalingGroupId      string `json:"scaling_group_id,omitempty"`
		LifecycleHookStatus string `json:"lifecycle_hook_status,omitempty"`
		Timeout             string `json:"timeout,omitempty"`
		DefaultResult       string `json:"default_result,omitempty"`
	}{
		{
			LifecycleActionKey:  "2c4aabb1-3ac3-4aba-9c6a-261147be6dfb",
			InstanceId:          "6abadacf-87af-4225-b762-4a56853aec02",
			LifecycleHookName:   "test-hook1",
			ScalingGroupId:      "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
			LifecycleHookStatus: "HANGING",
			Timeout:             "2018-05-12T10:16:30Z",
			DefaultResult:       "ABANDON",
		},
	},
}

func HandleListWithSuspensionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_instance_hook/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/list", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListWithSuspensionOutput)
	})
}

func HandleCallbackSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_instance_hook/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/callback", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/autoscaling-api/v1/57e98940a77f4bb988a21a7d0603a626/scaling_lifecycle_hook/60dcec94-5d5b-4dbf-9f50-4ccd7d841432/test-hook1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
