package notifications

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type DeleteResult struct {
	commonResult
}

func (r DeleteResult) Extract() (*DeleteResponse, error) {
	var response DeleteResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteResponse struct {
}

type EnableResult struct {
	commonResult
}

func (r EnableResult) Extract() (*EnableResponse, error) {
	var response EnableResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type EnableResponse struct {

	// Specifies a unified topic in SMN.
	TopicUrn string `json:"topic_urn,omitempty"`

	// Specifies a notification scenario, which can be one of the following:
	TopicScene []string `json:"topic_scene,omitempty"`
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResponse struct {

	// Specifies the AS group notification list.
	Topics []struct {

		// Specifies a unified topic in SMN.
		TopicUrn string `json:"topic_urn,omitempty"`

		// Specifies a notification scenario, which can be one of the following:
		TopicScene []string `json:"topic_scene,omitempty"`

		// Specifies the topic name in SMN.
		TopicName string `json:"topic_name,omitempty"`
	} `json:"topics,omitempty"`
}
