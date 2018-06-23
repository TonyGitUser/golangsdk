package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/logs"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListErrorsOutput = `
{
  "errorLogList": [{
    "datetime": "2016-08-30 09:55:39",
    "content": "[Warning] 'proxies_priv' entry '@ root@rds-bf83' ignored in --skip-name-resolve mode."
  }],
  "totalRecord": 1
}
`

var ListErrorsResponse = logs.ListErrorsResponse{
	ErrorLogList: []struct {
		Datetime string `json:"datetime,omitempty"`
		Content  string `json:"content,omitempty"`
	}{
		{
			Datetime: "2016-08-30 09:55:39",
			Content:  "[Warning] 'proxies_priv' entry '@ root@rds-bf83' ignored in --skip-name-resolve mode.",
		},
	},
	TotalRecord: 1,
}

func HandleListErrorsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/68e3010955d748099f62a0df726fe09b/instances/02cf1fd7-24ae-4fec-a621-329ec732e4f6/errorlog", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListErrorsOutput)
	})
}

var ListInfosOutput = `
{
  "slowLogList": [{
      "count": "409",
      "time": "1.29",
      "lockTime": " 0 ",
      "rowsSent": " 0 ",
      "rowsExamined": " 0 ",
      "database": " ",
      "users": " \trdsBackup@localhost  : (409) of query, (409) of all users",
      "querySample": "flush logs;"
    },
    {
      "count": "1",
      "time": "5.0",
      "lockTime": " 0 ",
      "rowsSent": " 1 ",
      "rowsExamined": " 0 ",
      "database": " ",
      "users": " \trdsAdmin@localhost  : (1) of query, (1) of all users",
      "querySample": "select sleep(5);"
    }
  ],
  "totalRecord": 2
}
`

var ListInfosResponse = logs.ListInfosResponse{
	SlowLogList: []struct {
		Count        string `json:"count,omitempty"`
		Time         string `json:"time,omitempty"`
		LockTime     string `json:"lockTime,omitempty"`
		RowsSent     string `json:"rowsSent,omitempty"`
		RowsExamined string `json:"rowsExamined,omitempty"`
		Database     string `json:"database,omitempty"`
		Users        string `json:"users,omitempty"`
		QuerySample  string `json:"querySample,omitempty"`
	}{
		{
			Count:        "409",
			Time:         "1.29",
			LockTime:     " 0 ",
			RowsSent:     " 0 ",
			RowsExamined: " 0 ",
			Database:     " ",
			Users:        " \trdsBackup@localhost  : (409) of query, (409) of all users",
			QuerySample:  "flush logs;",
		},
		{
			Count:        "1",
			Time:         "5.0",
			LockTime:     " 0 ",
			RowsSent:     " 1 ",
			RowsExamined: " 0 ",
			Database:     " ",
			Users:        " \trdsAdmin@localhost  : (1) of query, (1) of all users",
			QuerySample:  "select sleep(5);",
		},
	},
	TotalRecord: 2,
}

func HandleListInfosSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/68e3010955d748099f62a0df726fe09b/instances/02e47383-9222-4d29-bf5b-54b3013b0f71/slowlog", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListInfosOutput)
	})
}
