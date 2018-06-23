package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/healthcheck"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := healthcheck.Create(client.ServiceClient(), tenantId, healthcheck.CreateOpts{
		HealthcheckConnectPort: 80,
		HealthcheckInterval:    5,
		HealthcheckProtocol:    "HTTP",
		HealthcheckTimeout:     10,
		HealthyThreshold:       3,
		ListenerId:             "d5f3197c6bd8491ca1dfc905350d85ea",
		UnhealthyThreshold:     3,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	healthcheckId := "c74c17f02c5142b3befc973056f220d3"
	actual, err := healthcheck.Update(client.ServiceClient(), tenantId, healthcheckId, healthcheck.UpdateOpts{
		HealthcheckConnectPort: 81,
		HealthcheckInterval:    6,
		HealthcheckProtocol:    "HTTP",
		HealthcheckTimeout:     20,
		HealthyThreshold:       4,
		UnhealthyThreshold:     4,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	healthcheckId := "c74c17f02c5142b3befc973056f220d3"
	actual, err := healthcheck.Get(client.ServiceClient(), tenantId, healthcheckId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	healthcheckId := "c74c17f02c5142b3befc973056f220d3"
	healthcheck.Delete(client.ServiceClient(), healthcheckId)
}
