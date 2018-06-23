package backups

import (
	"github.com/huaweicloud/golangsdk"
)

type BackupOpts struct {

	// Specifies the snapshot object.
	Backup struct {

		// Specifies the backup description. It contains a maximum of 256 characters and cannot contain the following special characters: !<>=&"'
		Description string `json:"description,omitempty"`

		// Specifies a primary DB instance ID.
		Instance string `json:"instance,omitempty"`

		// Specifies the backup name. It must be 4 to 64 characters in length and start with a letter. It is case-insensitive and can contain only letters, digits, hyphens (-), and underscores (_).
		Name string `json:"name,omitempty"`
	} `json:"backup,"`
}

type BackupOptsBuilder interface {
	ToBackupMap() (map[string]interface{}, error)
}

func (opts BackupOpts) ToBackupMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Backup(client *golangsdk.ServiceClient, versionId string, projectId string, opts BackupOptsBuilder) (r BackupResult) {
	b, _ := opts.ToBackupMap()

	_, r.Err = client.Post(BackupURL(client, versionId, projectId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

type CreatePolicyOpts struct {

	// Specifies the backup policy object, including the backup retention period (days) and start time.
	Policy struct {

		// Specifies the number of days to retain the generated backup files. Its value range is from 0 to 35. If this parameter is not specified or set to 0, the automated backup policy is disabled.
		Keepday int `json:"keepday,"`

		// Indicates the backup start time that has been set. Valid value: The value cannot be empty. The format can be hh:mm:ss or hh:mm and must be valid. The current time is the UTC time.
		Starttime string `json:"starttime,"`
	} `json:"policy,"`
}

type CreatePolicyOptsBuilder interface {
	ToCreatePolicyMap() (map[string]interface{}, error)
}

func (opts CreatePolicyOpts) ToCreatePolicyMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreatePolicy(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts CreatePolicyOptsBuilder) (r CreatePolicyResult) {
	b, _ := opts.ToCreatePolicyMap()

	_, r.Err = client.Put(CreatePolicyURL(client, versionId, projectId, instanceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, versionId string, projectId string, backupId string) (r DeleteResult) {
	url := DeleteURL(client, versionId, projectId, backupId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
		MoreHeaders:  map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func GetPolicy(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) (r GetPolicyResult) {
	url := GetPolicyURL(client, versionId, projectId, instanceId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func List(client *golangsdk.ServiceClient, versionId string, projectId string) (r ListResult) {
	url := ListURL(client, versionId, projectId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

type RestoreOpts struct {

	// Specifies the restore information, including backupRef and restoreTime. At least one of them must be specified. If both of them are specified, only backupRef takes effect.
	Restore struct {

		// Specifies the full backup file ID.
		BackupRef string `json:"backupRef,omitempty"`

		// Specifies the time point to which the DB instance is restored.
		RestoreTime int64 `json:"restoreTime,omitempty"`
	} `json:"restore,"`
}

type RestoreOptsBuilder interface {
	ToRestoreMap() (map[string]interface{}, error)
}

func (opts RestoreOpts) ToRestoreMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Restore(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts RestoreOptsBuilder) (r RestoreResult) {
	b, _ := opts.ToRestoreMap()

	_, r.Err = client.Post(RestoreURL(client, versionId, projectId, instanceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

type RestoreWithNewOpts struct {

	// Specifies new DB instance details.
	Instance struct {

		// Specifies the DB instance name. The DB instance name of the same DB engine is unique for the same tenant. Valid value: The value must be 4 to 64 characters in length and start with a letter. It is case-insensitive and can contain only letters, digits, hyphens (-), and underscores (_).
		Name string `json:"name,"`

		// Specifies the specification ID (flavors.id in the response message in section Obtaining All DB Instance Specifications). Valid value: The value cannot be empty. The name length and whether the name complies with UUID regular expression rules are verified.
		FlavorRef string `json:"flavorRef,"`

		// Specifies the volume information. This parameter is mandatory for single or HA DB instance creation.
		Volume struct {

			// Specifies the volume size in GB. It cannot be less than the volume size of the original DB instance.
			Size int `json:"size,"`
		} `json:"volume,"`

		// Specifies the HA configuration parameter. It is mandatory for HA DB instance creation.
		Ha struct {

			// Specifies the HA configuration parameter. Valid value: The value is true or false. The value true indicates creating HA DB instances. The value false indicates creating a single DB instance.
			Enable bool `json:"enable,omitempty"`

			// Specifies the replication mode for the standby DB instance. The value cannot be empty. For MySQL, the value is async or semisync. For HWSQL, the value is async or semisync. For PostgreSQL, the value is async or sync. For Microsoft SQL Server, the value is sync.
			ReplicationMode bool `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`

		// Restores data to a new DB instance.
		RestorePoint struct {

			// Specifies the full backup file.
			BackupRef string `json:"backupRef,omitempty"`

			// Specifies the time point the DB instance is restored to. At least one of the backupRef and restoreTime parameters should be specified. If both parameters are specified, the DB instance is restored using the full backup file.
			RestoreTime int64 `json:"restoreTime,omitempty"`

			// Specifies the source DB instance ID.
			SourceInstanceId string `json:"sourceInstanceId,omitempty"`
		} `json:"restorePoint,"`
	} `json:"instance,"`
}

type RestoreWithNewOptsBuilder interface {
	ToRestoreWithNewMap() (map[string]interface{}, error)
}

func (opts RestoreWithNewOpts) ToRestoreWithNewMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func RestoreWithNew(client *golangsdk.ServiceClient, versionId string, projectId string, opts RestoreWithNewOptsBuilder) (r RestoreWithNewResult) {
	b, _ := opts.ToRestoreWithNewMap()

	_, r.Err = client.Post(RestoreWithNewURL(client, versionId, projectId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}
