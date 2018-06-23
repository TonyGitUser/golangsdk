package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/policies"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := policies.Create(client.ServiceClient(), tenantId, policies.CreateOpts{
		ScalingPolicyName: "as-policy-7a75",
		ScalingPolicyAction: struct {
			Operation          string `json:"operation,omitempty"`
			InstanceNumber     int    `json:"instance_number,omitempty"`
			InstancePercentage int    `json:"instance_percentage,omitempty"`
		}{
			Operation:      "ADD",
			InstanceNumber: 1,
		},
		CoolDownTime: 900,
		ScheduledPolicy: struct {
			LaunchTime      string `json:"launch_time,omitempty"`
			RecurrenceType  string `json:"recurrence_type,omitempty"`
			RecurrenceValue string `json:"recurrence_value,omitempty"`
			StartTime       string `json:"start_time,omitempty"`
			EndTime         string `json:"end_time,omitempty"`
		}{
			LaunchTime:      "16:00",
			RecurrenceType:  "Daily",
			RecurrenceValue: "",
			StartTime:       "2015-12-14T03:34Z",
			EndTime:         "2019-12-27T03:34Z",
		},
		ScalingPolicyType: "RECURRENCE",
		ScalingGroupId:    "60dcec94-5d5b-4dbf-9f50-4ccd7d841432",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
	actual, err := policies.Update(client.ServiceClient(), tenantId, policyId, policies.UpdateOpts{
		ScalingPolicyType: "RECURRENCE",
		ScalingPolicyName: "policy_01",
		ScheduledPolicy: struct {
			LaunchTime      string `json:"launch_time,omitempty"`
			RecurrenceType  string `json:"recurrence_type,omitempty"`
			RecurrenceValue string `json:"recurrence_value,omitempty"`
			StartTime       string `json:"start_time,omitempty"`
			EndTime         string `json:"end_time,omitempty"`
		}{
			LaunchTime:      "16:00",
			RecurrenceType:  "Daily",
			RecurrenceValue: "",
			EndTime:         "2020-02-08T17:31Z",
			StartTime:       "2019-01-08T17:31Z",
		},
		CoolDownTime: 300,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
	actual, err := policies.Get(client.ServiceClient(), tenantId, policyId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	scalingGroupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := policies.List(client.ServiceClient(), tenantId, scalingGroupId, policies.ListOpts{
		Limit: 2,
	}).Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestAction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleActionSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
	policies.Action(client.ServiceClient(), tenantId, policyId, policies.ActionOpts{
		Action: "resume",
	})
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	policyId := "1a68e2a0-a19c-430e-a9c1-2a41b662ff57"
	policies.Delete(client.ServiceClient(), tenantId, policyId)
}
