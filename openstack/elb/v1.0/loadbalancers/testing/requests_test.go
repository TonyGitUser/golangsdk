package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/elb/v1.0/loadbalancers"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := loadbalancers.Create(client.ServiceClient(), tenantID, loadbalancers.CreateOpts{
		Bandwidth:       1,
		SecurityGroupId: "",
		VpcId:           "773c3c42-d315-417b-9063-87091713148c",
		AdminStateUp:    1,
		VipSubnetId:     "",
		Type:            "External",
		Name:            "TestABC",
		Description:     "Show Load Balancer",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
	actual, err := loadbalancers.Update(client.ServiceClient(), tenantID, loadBalancerId, loadbalancers.UpdateOpts{
		Name:        "UpdatedLB",
		Description: "TEST LB",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
	actual, err := loadbalancers.Get(client.ServiceClient(), tenantID, loadBalancerId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := loadbalancers.List(client.ServiceClient(), tenantID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	loadBalancerId := "ca082827b61d4902bfaf32e8174e244a"
	loadbalancers.Delete(client.ServiceClient(), tenantID, loadBalancerId)
}
