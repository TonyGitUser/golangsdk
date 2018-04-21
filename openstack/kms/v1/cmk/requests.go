package cmk

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/postpagination"
)

type CancelDeletionOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type CancelDeletionOptsBuilder interface {
	ToCancelDeletionMap() (map[string]interface{}, error)
}

func (opts CancelDeletionOpts) ToCancelDeletionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CancelDeletion(client *golangsdk.ServiceClient, opts CancelDeletionOptsBuilder) (r CancelDeletionResult) {
	b, err := opts.ToCancelDeletionMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CancelDeletionURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type CreateOpts struct {

	// Alias of a CMK (The alias's length ranges from 1 to 255 characters and matches the regular expression ^[a-zA-Z0-9:/_-]{1,255}$. In addition, it must be different from the alias of a Default Master Key created by another service.)
	KeyAlias string `json:"key_alias,"`

	// CMK description (The value ranges from 0 to 255 characters.)
	KeyDescription string `json:"key_description,omitempty"`

	// Origin of a CMK. The default value is kms. The following values are enumerated: kms indicates that the CMK material is generated by KMS. external indicates that the CMK material is imported.
	Origin string `json:"origin,omitempty"`

	// Region where a CMK resides.
	Realm string `json:"realm,omitempty"`

	// Key policy (This parameter is an extension field.)
	KeyPolicy string `json:"key_policy,omitempty"`

	// Purpose of a CMK (The default value is , that is, a key is used to encrypt and decrypt data.)
	KeyUsage string `json:"key_usage,omitempty"`

	// Type of a CMK
	KeyType string `json:"key_type,omitempty"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type DisableOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type DisableOptsBuilder interface {
	ToDisableMap() (map[string]interface{}, error)
}

func (opts DisableOpts) ToDisableMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Disable(client *golangsdk.ServiceClient, opts DisableOptsBuilder) (r DisableResult) {
	b, err := opts.ToDisableMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(DisableURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type EnableOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type EnableOptsBuilder interface {
	ToEnableMap() (map[string]interface{}, error)
}

func (opts EnableOpts) ToEnableMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Enable(client *golangsdk.ServiceClient, opts EnableOptsBuilder) (r EnableResult) {
	b, err := opts.ToEnableMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(EnableURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type GetOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type GetOptsBuilder interface {
	ToGetMap() (map[string]interface{}, error)
}

func (opts GetOpts) ToGetMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Get(client *golangsdk.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	b, err := opts.ToGetMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(GetURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type GrantOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f.
	KeyId string `json:"key_id,"`

	// 32-byte ID of a user to which permissions are granted that matches the regular expression ^[a-zA-Z0-9_-]{32}$ Example: 0d0466b00d0466b00d0466b00d0466b0.
	GranteePrincipal string `json:"grantee_principal,"`

	// Permissions that can be granted Values: create-datakey, create-datakey-without-plaintext, encrypt-datakey, decrypt-datakey, describe-key, create-grant, retire-grant, encrypt-data, decrypt-data create-grant cannot be the only value.
	Operations []string `json:"operations,"`

	// Name of a grant which can be 1 to 255 characters in length and matches the regular expression ^[a-zA-Z0-9:/_-]{1,255}$
	Name string `json:"name,omitempty"`

	// 32-byte ID of a user who can retire a grant that matches the regular expression ^[a-zA-Z0-9_-]{32}$ Example: 0d0466b00d0466b00d0466b00d0466b0
	RetiringPrincipal string `json:"retiring_principal,omitempty"`

	// 32-byte ID of a user who can retire a grant that matches the regular expression ^[a-zA-Z0-9_-]{32}$ Example: 0d0466b00d0466b00d0466b00d0466b0
	Sequence string `json:"sequence,omitempty"`

	// 64-byte ID of a grant
	GrantId string `json:"grant_id,"`
}

type GrantOptsBuilder interface {
	ToGrantMap() (map[string]interface{}, error)
}

func (opts GrantOpts) ToGrantMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Grant(client *golangsdk.ServiceClient, opts GrantOptsBuilder) (r GrantResult) {
	b, err := opts.ToGrantMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(GrantURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Instances(client *golangsdk.ServiceClient) (r InstancesResult) {
	url := InstancesURL(client)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {

	// This parameter specifies the number of entries returned. If the specified number is smaller than the actual number of existing entries,  will be returned for the response parameter , indicating that the query results will be displayed in separate pages. The value is within the range of the maximum number of CMKs, for example, .
	Limit int `json:"limit,omitempty"`

	// This parameter marks the starting location in a pagination query. If the  value is , you can send consecutive requests to obtain more record entries. The  value must be set to the  value in the response, for example, .
	Marker string `json:"marker,omitempty"`

	//  State of a CMK that matches the regular expression . The following values are enumerated:
	KeyState string `json:"key_state,omitempty"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type ListOptsBuilder interface {
	ToListMap() (map[string]interface{}, error)
}

func (opts ListOpts) ToListMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// This API allows you to query the list of all CMKs.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) postpagination.Pager {
	initialRequest, err := opts.ToListMap()
	if err != nil {
		return postpagination.Pager{Err: err}
	}

	return postpagination.NewPager(client, "POST", ListURL(client), initialRequest,
		func(r postpagination.PageResult) postpagination.Page {
			p := ListPage{
				postpagination.PostMarkerPageBase{PageResult: r, Request: initialRequest},
			}
			return p
		})
}

type ListGrantsOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// This parameter specifies the number of entries returned. If the specified number is smaller than the actual number of existing entries,  will be returned for the response parameter , indicating that the query results will be displayed in separate pages.
	Limit string `json:"limit,omitempty"`

	// This parameter marks the starting location in a pagination query.
	Marker string `json:"marker,omitempty"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type ListGrantsOptsBuilder interface {
	ToListGrantsMap() (map[string]interface{}, error)
}

func (opts ListGrantsOpts) ToListGrantsMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ListGrants(client *golangsdk.ServiceClient, opts ListGrantsOptsBuilder) (r ListGrantsResult) {
	b, err := opts.ToListGrantsMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(ListGrantsURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListRetirableOpts struct {

	// This parameter specifies the number of entries returned. If the specified number is smaller than the actual number of existing entries,  will be returned for the response parameter , indicating that the query results will be displayed in separate pages.
	Limit string `json:"limit,omitempty"`

	// This parameter marks the starting location in a pagination query.
	Marker string `json:"marker,omitempty"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type ListRetirableOptsBuilder interface {
	ToListRetirableMap() (map[string]interface{}, error)
}

func (opts ListRetirableOpts) ToListRetirableMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ListRetirable(client *golangsdk.ServiceClient, opts ListRetirableOptsBuilder) (r ListRetirableResult) {
	b, err := opts.ToListRetirableMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(ListRetirableURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Quotas(client *golangsdk.ServiceClient) (r QuotasResult) {
	url := QuotasURL(client)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type RetireOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// 64-byte ID of a grant that meets the regular expression
	GrantId string `json:"grant_id,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type RetireOptsBuilder interface {
	ToRetireMap() (map[string]interface{}, error)
}

func (opts RetireOpts) ToRetireMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Retire(client *golangsdk.ServiceClient, opts RetireOptsBuilder) (r RetireResult) {
	b, err := opts.ToRetireMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(RetireURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type RevokeOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// 64-byte ID of a grant that meets the regular expression
	GrantId string `json:"grant_id,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type RevokeOptsBuilder interface {
	ToRevokeMap() (map[string]interface{}, error)
}

func (opts RevokeOpts) ToRevokeMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Revoke(client *golangsdk.ServiceClient, opts RevokeOptsBuilder) (r RevokeResult) {
	b, err := opts.ToRevokeMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(RevokeURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ScheduleDeletionOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// Number of days after which a CMK is scheduled to be deleted (The value ranges from  to .)
	PendingDays string `json:"pending_days,"`

	// 36-byte serial number of a request message. Example: 919c82d4-8046-4722-9094-35c3c6524cff.
	Sequence string `json:"sequence,omitempty"`
}

type ScheduleDeletionOptsBuilder interface {
	ToScheduleDeletionMap() (map[string]interface{}, error)
}

func (opts ScheduleDeletionOpts) ToScheduleDeletionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ScheduleDeletion(client *golangsdk.ServiceClient, opts ScheduleDeletionOptsBuilder) (r ScheduleDeletionResult) {
	b, err := opts.ToScheduleDeletionMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(ScheduleDeletionURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type UpdateAliasOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// Alias of a CMK whose length is 1 to 255 characters and which matches the regular expression . Suffix of the alias cannot be .
	KeyAlias string `json:"key_alias,"`

	// 36-byte serial number of a request message, for example, 919c82d4-8046-4722-9094-35c3c6524cff
	Sequence string `json:"sequence,omitempty"`
}

type UpdateAliasOptsBuilder interface {
	ToUpdateAliasMap() (map[string]interface{}, error)
}

func (opts UpdateAliasOpts) ToUpdateAliasMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateAlias(client *golangsdk.ServiceClient, opts UpdateAliasOptsBuilder) (r UpdateAliasResult) {
	b, err := opts.ToUpdateAliasMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(UpdateAliasURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type UpdateDescriptionOpts struct {

	// 36-byte ID of a CMK that matches the regular expression ^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$ Example: 0d0466b0-e727-4d9c-b35d-f84bb474a37f
	KeyId string `json:"key_id,"`

	// CMK description (The value ranges from 0 to 255 characters.)
	KeyDescription string `json:"key_description,"`

	// 36-byte serial number of a request message, for example, 919c82d4-8046-4722-9094-35c3c6524cff
	Sequence string `json:"sequence,omitempty"`
}

type UpdateDescriptionOptsBuilder interface {
	ToUpdateDescriptionMap() (map[string]interface{}, error)
}

func (opts UpdateDescriptionOpts) ToUpdateDescriptionMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateDescription(client *golangsdk.ServiceClient, opts UpdateDescriptionOptsBuilder) (r UpdateDescriptionResult) {
	b, err := opts.ToUpdateDescriptionMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(UpdateDescriptionURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}