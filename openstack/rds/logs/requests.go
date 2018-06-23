package logs

import (
	"github.com/huaweicloud/golangsdk"
)

type ListErrorsOpts struct {

	// Specifies the start time following a specified time translation rule. For example, time "2016-08-29+06:35" is translated into "2016-08-29+06%!A(MISSING)35", where "%!A(MISSING)" is translated from ":" and other digits and characters remain unchanged.
	StartDate string `q:"startDate"`

	// Specifies the end time following a specified time translation rule. For example, time "2016-08-29+06:35" is translated into "2016-08-29+06%!A(MISSING)35", where "%!A(MISSING)" is translated from ":" and other digits and characters remain unchanged.
	EndDate string `q:"endDate"`

	// Specifies the current page number, such as 1, 2, 3, or 4. The parameter value is  by default if it is not specified.
	CurPage int `q:"curPage"`

	// Specifies the number of records on a page. Its value range is 1 to 100. The parameter value is  by default if it is not specified.
	PerPage int `q:"perPage"`
}

type ListErrorsOptsBuilder interface {
	ToListErrorsQuery() (string, error)
}

func (opts ListErrorsOpts) ToListErrorsQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func ListErrors(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts ListErrorsOptsBuilder) (r ListErrorsResult) {
	url := ListErrorsURL(client, versionId, projectId, instanceId)
	if opts != nil {
		query, _ := opts.ToListErrorsQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

type ListInfosOpts struct {

	// Specifies the statement type. Its value can be empty, INSERT, UPDATE, SELECT, DELETE, or CREATE. If the value is empty, all statement types exist.
	Sftype string `q:"sftype"`

	// Specifies how many records are returned. If this parameter is specified, the value range is from 1 to 50. If this parameter is not specified, the first 10 records are returned.
	Top int `q:"top"`
}

type ListInfosOptsBuilder interface {
	ToListInfosQuery() (string, error)
}

func (opts ListInfosOpts) ToListInfosQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func ListInfos(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts ListInfosOptsBuilder) (r ListInfosResult) {
	url := ListInfosURL(client, versionId, projectId, instanceId)
	if opts != nil {
		query, _ := opts.ToListInfosQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}
