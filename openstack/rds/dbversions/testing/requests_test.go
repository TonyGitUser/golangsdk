package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/dbversions"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	datastoreName := "MySQL"
	actual, err := dbversions.List(client.ServiceClient(), versionId, projectId, datastoreName).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}
