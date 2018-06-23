package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/as/v1/configures"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := configures.Create(client.ServiceClient(), tenantId, configures.CreateOpts{
		ScalingConfigurationName: "as-config-test",
		InstanceConfig: configures.CreateInstanceConfig{
			InstanceName: "",
			InstanceId:   "",
			FlavorRef:    "s2.small.1",
			ImageRef:     "e3c10a27-07fc-4144-b5a5-e11d1003a5fe",
			Disk: []configures.Disk{
				{
					Size:               40,
					VolumeType:         "SATA",
					DiskType:           "SYS",
					DedicatedStorageId: "",
					DataDiskImageId:    "",
					SnapshotId:         "",
				},
			},
			KeyName:     "",
			Personality: nil,
			UserData:    "IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPTDNuMkIkVXh6TG9pYklmeUtQY0hKNHphRnRnYWZsTGdpaXdGMDZoL3lPY2t4Q3RnYmsyWXBncW8uU2Jjd0pvdWdEaUhkVk5vUThyOHhvemV0eHJOYjh5c3FCYTEnIHwgY2hwYXNzd2QgLWU7",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := configures.Get(client.ServiceClient(), tenantID, "f109bb4f-09f0-4dac-9115-6b429d548750").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	tenantID := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := configures.List(client.ServiceClient(), tenantID, configures.ListOpts{
		Limit: 2,
	}).Extract()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	scalingConfigurationId := "34147aea-e7bb-4290-8248-794b26f355a6"
	configures.Delete(client.ServiceClient(), tenantId, scalingConfigurationId)
}

func TestDeleteWithBatch(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteWithBatchSuccessfully(t)

	tenantId := "57e98940a77f4bb988a21a7d0603a626"
	configures.DeleteWithBatch(client.ServiceClient(), tenantId, configures.DeleteWithBatchOpts{
		ScalingConfigurationId: []string{
			"483b6f84-d0d0-451f-9993-adb28e09a721",
			"e24c26dc-8a26-4e9c-85bf-4ac1ddc090bf",
		},
	})
}
