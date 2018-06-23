package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/lifecyclehooks"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := lifecyclehooks.Create(client.ServiceClient(), tenantId, groupId, lifecyclehooks.CreateOpts{
		LifecycleHookName:    "test-hook1",
		DefaultResult:        "ABANDON",
		DefaultTimeout:       3600,
		NotificationTopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
		LifecycleHookType:    "INSTANCE_LAUNCHING",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	lifecyclehookName := "test-hook1"
	actual, err := lifecyclehooks.Update(client.ServiceClient(), tenantId, groupId, lifecyclehookName, lifecyclehooks.UpdateOpts{
		DefaultResult:        "CONTINUE",
		DefaultTimeout:       1800,
		NotificationTopicUrn: "urn:smn:cn-north-1:57e98940a77f4bb988a21a7d0603a626:Topic1",
		LifecycleHookType:    "INSTANCE_LAUNCHING",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	lifecycleHookName := "test-hook1"
	actual, err := lifecyclehooks.Get(client.ServiceClient(), tenantId, groupId, lifecycleHookName).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := lifecyclehooks.List(client.ServiceClient(), tenantId, groupId).Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestListWithSuspension(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListWithSuspensionSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	actual, err := lifecyclehooks.ListWithSuspension(client.ServiceClient(), tenantId, groupId, lifecyclehooks.ListWithSuspensionOpts{
		InstanceId: "6abadacf-87af-4225-b762-4a56853aec02",
	}).Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListWithSuspensionResponse, actual)
}

func TestCallback(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCallbackSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	lifecyclehooks.CallBack(client.ServiceClient(), tenantId, groupId, lifecyclehooks.CallBackOpts{
		LifecycleActionResult: "ABANDON",
		InstanceId:            "e1040c67-b130-41ee-9405-07c6c6c20208",
		LifecycleHookName:     "test-hook1",
	})
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "60dcec94-5d5b-4dbf-9f50-4ccd7d841432"
	lifecycleHookName := "test-hook1"
	lifecyclehooks.Delete(client.ServiceClient(), tenantId, groupId, lifecycleHookName)
}
