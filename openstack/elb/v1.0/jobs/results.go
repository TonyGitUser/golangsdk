package jobs

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

	// Specifies the task status.SUCCESS: The job was successfully executed.RUNNING: The job is in progress.FAIL: indicates that the task failed.INIT: The job is being initialized.
	Status string `json:"status,omitempty"`

	// Specifies the resource information or error information. The ELB resource ID is used as an example in the response example.
	Entities map[string]interface{} `json:"entities,omitempty"`

	// Specifies the unique ID assigned to the task for querying the execution status in Combined API.
	JobId string `json:"job_id,omitempty"`

	// Specifies the task type.
	JobType string `json:"job_type,omitempty"`

	// Specifies the time when the job started.
	BeginTime string `json:"begin_time,omitempty"`

	// Specifies the time when the job ended.
	EndTime string `json:"end_time,omitempty"`

	// Specifies the error code returned after the job execution fails.
	ErrorCode string `json:"error_code,omitempty"`

	// Specifies the cause of the task execution failure.
	FailReason string `json:"fail_reason,omitempty"`

	// Specifies the error message returned after an error occurs.
	Message string `json:"message,omitempty"`

	// Specifies the error code returned after an error occurs.For details about error codes, see Error Codes.https://support.huaweicloud.com/en-us/api-elb/en-us_topic_0020311892.html
	Code string `json:"code,omitempty"`

	// Specifies the execution information of a sub-task. When no sub-task exists, the value of this parameter is left blank. The structure of each sub-task is similar to that of the parent task.
	SubJobs string `json:"sub_jobs,omitempty"`
}
