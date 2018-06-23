package backendces

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOpts struct {
	Items []struct {

		// Specifies the backend ECS ID.
		ServerId string `json:"server_id,omitempty"`

		// Specifies the private IP address of the backend ECS.
		Address string `json:"address,omitempty"`
	}
}

type CreateOptsBuilder interface {
	ToCreateMap() ([]map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() ([]map[string]interface{}, error) {
	returnItems := make([]map[string]interface{}, 0)
	for _, item := range opts.Items {
		b, err := golangsdk.BuildRequestBody(item, "")
		if err != nil {
			return nil, err
		}
		returnItems = append(returnItems, b)
	}
	return returnItems, nil
}

func Create(client *golangsdk.ServiceClient, tenantId string, listenerId string, opts CreateOptsBuilder) (r CreateResult) {
	b, _ := opts.ToCreateMap()

	_, r.Err = client.Post(CreateURL(client, tenantId, listenerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type DeleteOpts struct {

	// Lists the removed backend ECSs.
	RemoveMember []struct {

		// Specifies the backend ECS ID.
		Id string `json:"id,omitempty"`
	} `json:"removeMember,omitempty"`
}

type DeleteOptsBuilder interface {
	ToDeleteMap() (map[string]interface{}, error)
}

func (opts DeleteOpts) ToDeleteMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Delete(client *golangsdk.ServiceClient, tenantId string, listenerId string, opts DeleteOptsBuilder) (r DeleteResult) {
	b, _ := opts.ToDeleteMap()

	_, r.Err = client.Post(DeleteURL(client, tenantId, listenerId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}

type ListOpts struct {

	// Specifies the start resource ID of pagination query. If the parameter is left empty, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Specifies the quantity of records returned on each page.The value ranges from 0 to intmax.
	Limit string `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, tenantId string, listenerId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId, listenerId)
	if opts != nil {
		query, _ := opts.ToListQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
