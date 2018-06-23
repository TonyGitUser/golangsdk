package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/groups"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := groups.Create(client.ServiceClient(), tenantId, groups.CreateOpts{
		ScalingGroupName:          "as-group-Test",
		ScalingConfigurationId:    "f109bb4f-09f0-4dac-9115-6b429d548750",
		DesireInstanceNumber:      0,
		MinInstanceNumber:         0,
		MaxInstanceNumber:         1,
		CoolDownTime:              900,
		LbListenerId:              "",
		VpcId:                     "773c3c42-d315-417b-9063-87091713148c",
		HealthPeriodicAuditMethod: "NOVA_AUDIT",
		HealthPeriodicAuditTime:   5,
		InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
		Notifications:             []string{},
		DeletePublicip:            true,
		AvailableZones:            []string{"cn-north-1a"},
		Networks: []struct {
			// Specifies the network ID.
			Id string `json:"id,omitempty"`
		}{
			{Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373"},
		},
		SecurityGroups: []struct {
			// Specifies the ID of the security group.
			Id string `json:"id,omitempty"`
		}{
			{Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"},
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	scalingGroupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
	actual, err := groups.Update(client.ServiceClient(), tenantId, scalingGroupId, groups.UpdateOpts{
		ScalingGroupName:       "as-group-TestModified",
		ScalingConfigurationId: "f109bb4f-09f0-4dac-9115-6b429d548750",
		DesireInstanceNumber:   0,
		MinInstanceNumber:      0,
		MaxInstanceNumber:      1,
		CoolDownTime:           900,
		LbListenerId:           "",
		Networks: []struct {
			// Specifies the network ID.
			Id string `json:"id,omitempty"`
		}{
			{Id: "f6a0db7b-397c-4162-bc35-9a1f008b4373"},
		},
		SecurityGroups: []struct {
			// Specifies the ID of the security group.
			Id string `json:"id,omitempty"`
		}{
			{Id: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"},
		},
		HealthPeriodicAuditMethod: "NOVA_AUDIT",
		HealthPeriodicAuditTime:   5,
		InstanceTerminatePolicy:   "OLD_CONFIG_OLD_INSTANCE",
		Notifications:             []string{},
		DeletePublicip:            true,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
	actual, err := groups.Get(client.ServiceClient(), tenantID, groupId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := groups.List(client.ServiceClient(), tenantID, groups.ListOpts{
		Limit: 2,
	}).Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestEnable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleEnableSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
	groups.Enable(client.ServiceClient(), tenantId, groupId, groups.EnableOpts{
		Action: "pause",
	})
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	groupId := "e4a97959-34d0-4c58-ab85-1dda4d0b9d11"
	groups.Delete(client.ServiceClient(), tenantId, groupId, groups.DeleteOpts{
		ForceDelete: "No",
	})
}
