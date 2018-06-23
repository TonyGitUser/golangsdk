package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/parameters"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	instanceId := "6926d5168444416fa28646de8a67450fno01"
	actual, err := parameters.Create(client.ServiceClient(), versionId, projectId, instanceId, parameters.CreateOpts{
		Values: map[string]interface{}{
			"connect_timeout": 17,
			"sync_binlog":     1,
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	datastoreId := "9dae9226-5f8e-4b10-bd41-adae2e693e89"
	actual, err := parameters.List(client.ServiceClient(), versionId, projectId, datastoreId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	datastoreId := "9dae9226-5f8e-4b10-bd41-adae2e693e89"
	parameterName := "auto_increment_offset"
	actual, err := parameters.Get(client.ServiceClient(), versionId, projectId, datastoreId, parameterName).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestRestore(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRestoreSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	instanceId := "6926d5168444416fa28646de8a67450fno01"
	actual, err := parameters.Restore(client.ServiceClient(), versionId, projectId, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &RestoreResponse, actual)
}
