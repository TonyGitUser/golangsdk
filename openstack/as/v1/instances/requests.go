package instances

import (
	"github.com/huaweicloud/golangsdk"
)

type ActionOpts struct {

	// Specifies the ECS ID.
	InstancesId []string `json:"instances_id,omitempty"`

	// Specifies whether to delete ECS instances when they are removed from an AS group. The value can be no (default) or yes.This parameter takes effect only when the action is set to REMOVE.
	InstanceDelete string `json:"instance_delete,omitempty"`

	// Specifies an action to be performed on instances in batches. The options are as follows:ADD: adds instances to the AS group.REMOVE: removes instances from the AS group.PROTECT: enables instance protection.UNPROTECT: disables instance protection.
	Action string `json:"action,omitempty"`
}

type ActionOptsBuilder interface {
	ToActionMap() (map[string]interface{}, error)
}

func (opts ActionOpts) ToActionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Action(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts ActionOptsBuilder) (r ActionResult) {
	b, _ := opts.ToActionMap()

	_, r.Err = client.Post(ActionURL(client, tenantId, scalingGroupId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type DeleteOpts struct {

	// Specifies whether the instances are deleted when they are removed from the AS group. The value can be yes or no (default).
	InstanceDelete string `q:"instance_delete"`
}

type DeleteOptsBuilder interface {
	ToDeleteQuery() (string, error)
}

func (opts DeleteOpts) ToDeleteQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func Delete(client *golangsdk.ServiceClient, tenantId string, instanceId string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := DeleteURL(client, tenantId, instanceId)
	if opts != nil {
		query, _ := opts.ToDeleteQuery()
		url += query
	}

	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies the instance lifecycle status in the AS group.INSERVICE: The instance in the AS group is in use.PENDING: The instance is being added to the AS group.REMOVING: The instance is being removed from the AS group.PENDING_WAIT: The instance is waiting to be added to the AS group.REMOVING_WAIT: The instance is waiting to be removed from the AS group.
	LifeCycleState string `q:"life_cycle_state"`

	// Specifies the instance health status.INITIALIZING: The instance is initializing.NORMAL: The instance is normal.ERROR: The instance is abnormal.
	HealthStatus string `q:"health_status"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default is 20 and the maximum is 100.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, tenantId string, scalingGroupId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId, scalingGroupId)
	if opts != nil {
		query, _ := opts.ToListQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
