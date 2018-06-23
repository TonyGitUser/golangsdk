package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/backups"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestCreatePolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreatePolicySuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	instanceId := "6926d5168444416fa28646de8a67450fno01"
	_, err := backups.CreatePolicy(client.ServiceClient(), versionId, projectId, instanceId, backups.CreatePolicyOpts{
		Policy: struct {
			Keepday   int    `json:"keepday,"`
			Starttime string `json:"starttime,"`
		}{
			Keepday:   7,
			Starttime: "00:00:00",
		},
	}).Extract()
	th.AssertNoErr(t, err)
}

func TestGetPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetPolicySuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	instanceId := "6926d5168444416fa28646de8a67450fno01"
	actual, err := backups.GetPolicy(client.ServiceClient(), versionId, projectId, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetPolicyResponse, actual)
}

func TestBackup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleBackupSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := backups.Backup(client.ServiceClient(), versionId, projectId, backups.BackupOpts{
		Backup: struct {
			Description string `json:"description,omitempty"`
			Instance    string `json:"instance,omitempty"`
			Name        string `json:"name,omitempty"`
		}{
			Description: "My Backup",
			Instance:    "6926d5168444416fa28646de8a67450fno01",
			Name:        "backup",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &BackupResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := backups.List(client.ServiceClient(), versionId, projectId).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestRestore(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRestoreSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	instanceId := "6926d5168444416fa28646de8a67450fno01"
	actual, err := backups.Restore(client.ServiceClient(), versionId, projectId, instanceId, backups.RestoreOpts{
		Restore: struct {
			BackupRef   string `json:"backupRef,omitempty"`
			RestoreTime int64  `json:"restoreTime,omitempty"`
		}{
			BackupRef: "1e24c1bfe18647a890fc0aeb7e775616br01",
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &RestoreResponse, actual)
}

func TestRestoreWithNew(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRestoreWithNewSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	actual, err := backups.RestoreWithNew(client.ServiceClient(), versionId, projectId, backups.RestoreWithNewOpts{
		Instance: struct {
			Name      string `json:"name,"`
			FlavorRef string `json:"flavorRef,"`
			Volume    struct {
				Size int `json:"size,"`
			} `json:"volume,"`
			Ha struct {
				Enable          bool `json:"enable,omitempty"`
				ReplicationMode bool `json:"replicationMode,omitempty"`
			} `json:"ha,omitempty"`
			RestorePoint struct {
				BackupRef        string `json:"backupRef,omitempty"`
				RestoreTime      int64  `json:"restoreTime,omitempty"`
				SourceInstanceId string `json:"sourceInstanceId,omitempty"`
			} `json:"restorePoint,"`
		}{

			Name:      "trove-newinstance",
			FlavorRef: "18207182-0c02-467f-ae05-9c1650df1717",
			Volume: struct {
				Size int `json:"size,"`
			}{
				Size: 100,
			},
			RestorePoint: struct {
				BackupRef        string `json:"backupRef,omitempty"`
				RestoreTime      int64  `json:"restoreTime,omitempty"`
				SourceInstanceId string `json:"sourceInstanceId,omitempty"`
			}{
				BackupRef:        "1e24c1bfe18647a890fc0aeb7e775616br01",
				SourceInstanceId: "6926d5168444416fa28646de8a67450fno01",
			},
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &RestoreWithNewResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	versionId := "v1"
	projectId := "57e98940a77f4bb988a21a7d0603a626"
	backupId := "1e24c1bfe18647a890fc0aeb7e775616br01"
	backups.Delete(client.ServiceClient(), versionId, projectId, backupId)
}
