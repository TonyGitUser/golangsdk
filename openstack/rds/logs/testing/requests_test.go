package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/logs"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListErrors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListErrorsSuccessfully(t)

	versionId := "v1"
	projectId := "68e3010955d748099f62a0df726fe09b"
	instanceId := "02cf1fd7-24ae-4fec-a621-329ec732e4f6"
	actual, err := logs.ListErrors(client.ServiceClient(), versionId, projectId, instanceId, logs.ListErrorsOpts{
		StartDate: "2018-01-29",
		EndDate:   "2018-08-29",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListErrorsResponse, actual)
}

func TestListInfos(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListInfosSuccessfully(t)

	versionId := "v1"
	projectId := "68e3010955d748099f62a0df726fe09b"
	instanceId := "02e47383-9222-4d29-bf5b-54b3013b0f71"
	actual, err := logs.ListInfos(client.ServiceClient(), versionId, projectId, instanceId, logs.ListInfosOpts{
		Top: 10,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListInfosResponse, actual)
}
