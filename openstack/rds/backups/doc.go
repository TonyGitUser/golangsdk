/*


Sample Code, This interface is used to set the automated backup policy.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := backups.CreatePolicy(client, versionId, projectId, instanceId, backups.CreatePolicyOpts{
        Policy: struct {
            Keepday   int    `json:"keepday,"`
            Starttime string `json:"starttime,"`
        }{
            Keepday:   7,
            Starttime: "00:00:00",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to obtain the automated backup policy information.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := backups.GetPolicy(client, versionId, projectId, instanceId).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to create a manual backup.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := backups.Backup(client, versionId, projectId, backups.BackupOpts{
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

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to return backup information as a list.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := backups.List(client, versionId, projectId).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to restore data to the current DB instance.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    instanceId := "6926d5168444416fa28646de8a67450fno01"
    result, err := backups.Restore(client, versionId, projectId, instanceId, backups.RestoreOpts{
        Restore: struct {
            // Specifies the full backup file ID.
            BackupRef string `json:"backupRef,omitempty"`

            // Specifies the time point to which the DB instance is restored.
            RestoreTime int64 `json:"restoreTime,omitempty"`
        }{
            BackupRef: "1e24c1bfe18647a890fc0aeb7e775616br01",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to restore the specified DB instance data to a new DB instance.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    result, err := backups.RestoreWithNew(client, versionId, projectId, backups.RestoreWithNewOpts{
        Instance: struct {
            Name      string `json:"name,"`
            FlavorRef string `json:"flavorRef,"`
            Volume struct {
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
            } {
                Size: 100,
            },
            RestorePoint: struct {
                BackupRef        string `json:"backupRef,omitempty"`
                RestoreTime      int64  `json:"restoreTime,omitempty"`
                SourceInstanceId string `json:"sourceInstanceId,omitempty"`
            } {
                BackupRef:        "1e24c1bfe18647a890fc0aeb7e775616br01",
                SourceInstanceId: "6926d5168444416fa28646de8a67450fno01",
            },
        },
    }).Extract()

    if err != nil {
        panic(err)
    }


Sample Code, This interface is used to delete snapshot information.


    versionId := "v1"
    projectId := "57e98940a77f4bb988a21a7d0603a626"
    backupId := "1e24c1bfe18647a890fc0aeb7e775616br01"
    result, err := backups.Delete(client, versionId, projectId, backupId).Extract()

    if err != nil {
        panic(err)
    }
*/
package backups
