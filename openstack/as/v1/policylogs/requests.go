package policylogs

import (
	"github.com/huaweicloud/golangsdk"
)

type ListOpts struct {

	// Specifies the log ID.
	LogId string `q:"log_id"`

	// Specifies the scaling resource type.AS group: SCALING_GROUP,Bandwidth: BANDWIDTH
	ScalingResourceType string `q:"scaling_resource_type"`

	// Specifies the scaling resource ID.
	ScalingResourceId string `q:"scaling_resource_id"`

	// Specifies the AS policy execution type.SCHEDULE: automatically triggered at a specified time point,RECURRENCE: automatically triggered at a specified time period,ALARM: alarm-triggered,MANUAL: manually triggered
	ExecuteType string `q:"execute_type"`

	// Specifies the start time for querying scaling action logs. The format of the start time is yyyy-MM-ddThh:mm:ssZ.
	StartTime string `q:"start_time"`

	// Specifies the end time for querying scaling action logs. The format of the end time is yyyy-MM-ddThh:mm:ssZ.
	EndTime string `q:"end_time"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default value is 20 and the maximum value is 100.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, tenantId string, scalingPolicyId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId, scalingPolicyId)
	if opts != nil {
		query, _ := opts.ToListQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
