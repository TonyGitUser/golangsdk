package logs

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type ListErrorsResult struct {
	commonResult
}

func (r ListErrorsResult) Extract() (*ListErrorsResponse, error) {
	var response ListErrorsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListErrorsResponse struct {

	// Indicates detailed information.
	ErrorLogList []struct {

		// Indicates the date and time.
		Datetime string `json:"datetime,omitempty"`

		// Indicates the log content.
		Content string `json:"content,omitempty"`
	} `json:"errorLogList,omitempty"`

	// Indicates the total number of records.
	TotalRecord int `json:"totalRecord,omitempty"`
}

type ListInfosResult struct {
	commonResult
}

func (r ListInfosResult) Extract() (*ListInfosResponse, error) {
	var response ListInfosResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListInfosResponse struct {

	// Indicates detailed information.
	SlowLogList []struct {

		// Indicates the number of executions.
		Count string `json:"count,omitempty"`

		// Indicates the average execution duration.
		Time string `json:"time,omitempty"`

		// Indicates the average waiting time before locking.
		LockTime string `json:"lockTime,omitempty"`

		// Indicates the average number of rows contained in a result.
		RowsSent string `json:"rowsSent,omitempty"`

		// Indicates the average number of scanned rows.
		RowsExamined string `json:"rowsExamined,omitempty"`

		// Indicates the database which the slow query log belongs to.
		Database string `json:"database,omitempty"`

		// Indicates the account.
		Users string `json:"users,omitempty"`

		// Indicates the execution syntax.
		QuerySample string `json:"querySample,omitempty"`
	} `json:"slowLogList,omitempty"`

	// Indicates the total number of records.
	TotalRecord int `json:"totalRecord,omitempty"`
}
