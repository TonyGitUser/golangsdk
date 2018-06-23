package backups

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type BackupResult struct {
	commonResult
}

func (r BackupResult) Extract() (*BackupResponse, error) {
	var response BackupResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type BackupResponse struct {

	// Specifies the manual backup information.
	Backup struct {

		// Indicates the creation time.
		Created string `json:"created,omitempty"`

		// Indicates the database version.
		DataStore struct {

			// Indicates the DB engine. Currently, MySQL, PostgreSQL, and Microsoft SQL Server are supported.
			Type string `json:"type,omitempty"`

			// Indicates the database version.
			Version string `json:"version,omitempty"`

			// Indicates the database version ID. Its value is unique.
			VersionId string `json:"version_id,omitempty"`
		} `json:"dataStore,omitempty"`

		// Indicates the backup description.
		Description string `json:"description,omitempty"`

		// Indicates the backup UUID.
		Id string `json:"id,omitempty"`

		// Indicates the DB instance ID.
		InstanceId string `json:"instance_id,omitempty"`

		// This is a reserved field and has no meanings.
		LocationRef string `json:"locationRef,omitempty"`

		// Indicates the backup file name.
		Name string `json:"name,omitempty"`

		// This is a reserved field and has no meanings.
		ParentId string `json:"parent_id,omitempty"`

		// Indicates the file size in GB. Here its value is .
		Size float64 `json:"size,omitempty"`

		// Indicates the backing up status with the value .
		Status string `json:"status,omitempty"`

		// Indicates the completion time with the returned value .
		Updated string `json:"updated,omitempty"`

		// The default value is , indicating a snapshot.
		Backuptype string `json:"backuptype,omitempty"`
	} `json:"backup,omitempty"`

	// Indicates the extended parameter.
	Extendparam struct {

		// Indicates the task ID.
		Jobs []string `json:"jobs,omitempty"`
	} `json:"extendparam,omitempty"`
}

type CreatePolicyResult struct {
	commonResult
}

func (r CreatePolicyResult) Extract() (*CreatePolicyResponse, error) {
	var response CreatePolicyResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreatePolicyResponse struct {
}

type DeleteResult struct {
	commonResult
}

func (r DeleteResult) Extract() (*DeleteResponse, error) {
	var response DeleteResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteResponse struct {
}

type GetPolicyResult struct {
	commonResult
}

func (r GetPolicyResult) Extract() (*GetPolicyResponse, error) {
	var response GetPolicyResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetPolicyResponse struct {

	// Indicates the backup policy object, including the backup retention period (days) and start time.
	Policy struct {

		// Indicates the number of days to retain the generated backup files.
		Keepday int `json:"keepday,omitempty"`

		// Indicates the backup start time that has been set. The backup task will be triggered within one hour after the backup start time.
		Starttime string `json:"starttime,omitempty"`
	} `json:"policy,omitempty"`
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*ListResponse, error) {
	var response ListResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListResponse struct {

	// Indicates the backup information.
	Backups []struct {

		// Indicates the backup UUID.
		Id string `json:"id,omitempty"`

		// Indicates the backup file name.
		Name string `json:"name,omitempty"`

		// Indicates the backup description.
		Description string `json:"description,omitempty"`

		// This is a reserved field and has no meanings.
		LocationRef string `json:"locationRef,omitempty"`

		// Indicates the snapshot creation time.
		Created string `json:"created,omitempty"`

		// Indicates the completion time.
		Updated string `json:"updated,omitempty"`

		// Indicates the file size in GB.
		Size float64 `json:"size,omitempty"`

		// Indicates the backup status.
		Status string `json:"status,omitempty"`

		// The default value is , indicating a snapshot.
		Backuptype string `json:"backuptype,omitempty"`

		// Indicates the database version.
		DataStore struct {

			// Indicates the DB engine. Currently, MySQL, PostgreSQL, and Microsoft SQL Server are supported.
			Type string `json:"type,omitempty"`

			// Indicates the database version.
			Version string `json:"version,omitempty"`

			// Indicates the database version ID. Its value is unique.
			VersionId string `json:"version_id,omitempty"`
		} `json:"dataStore,omitempty"`

		// Indicates the DB instance ID.
		InstanceId string `json:"instance_id,omitempty"`

		// This is a reserved field and has no meanings.
		ParentId string `json:"parent_id,omitempty"`
	} `json:"backups,omitempty"`
}

type RestoreResult struct {
	commonResult
}

func (r RestoreResult) Extract() (*RestoreResponse, error) {
	var response RestoreResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RestoreResponse struct {

	// Indicates the returned extendparam key-value pair.
	Extendparam struct {

		// Indicates the task ID.
		Jobs []struct {

			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	} `json:"extendparam,omitempty"`
}

type RestoreWithNewResult struct {
	commonResult
}

func (r RestoreWithNewResult) Extract() (*RestoreWithNewResponse, error) {
	var response RestoreWithNewResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RestoreWithNewResponse struct {

	// Indicates the snapshot information.
	Instance struct {

		// Indicates the DB instance ID.
		Id string `json:"id,omitempty"`

		// Indicates the DB instance status. The value is BUILD.
		Status string `json:"status,omitempty"`

		// Indicates the created DB instance name.
		Name string `json:"name,omitempty"`

		// Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset. For example, in the Beijing time zone, the time zone offset is shown as +0800.
		Created string `json:"created,omitempty"`

		// Indicates the instance connection address. It is a blank string.
		Hostname string `json:"hostname,omitempty"`

		// Indicates the DB instance type, which can be master or readreplica.
		Type string `json:"type,omitempty"`

		// Same as the request parameter.
		Region string `json:"region,omitempty"`

		// Indicates the updated time, which is the same as created.
		Updated string `json:"updated,omitempty"`

		// Same as the request parameter.
		AvailabilityZone string `json:"availabilityZone,omitempty"`

		// Same as the request parameter.
		Vpc string `json:"vpc,omitempty"`

		// Indicates the nics information.
		Nics struct {

			// Indicates the subnet ID.
			SubnetId string `json:"subnetId,omitempty"`
		} `json:"nics,omitempty"`

		// Indicates the security group information.
		SecurityGroup struct {

			// Indicates the security group ID.
			Id string `json:"id,omitempty"`
		} `json:"securityGroup,omitempty"`

		// Indicates the DB instance specifications.
		Flavor struct {

			// Indicates the specification ID.
			Id string `json:"id,omitempty"`
		} `json:"flavor,omitempty"`

		// Same as the request parameter.
		Volume struct {

			// Specifies the volume type.Valid value:It must be COMMON (SATA), HIGH (SAS), or ULTRAHIGH (SSD) and is case-sensitive.
			Type string `json:"type,omitempty"`

			// Specifies the volume size.Its value must be a multiple of 10 and the value range is from 40 GB to 2,000 GB.
			Size int `json:"size,omitempty"`
		} `json:"volume,omitempty"`

		// Same as the request parameter.
		DbPort int `json:"dbPort,omitempty"`

		// Indicates database information.
		DataStoreInfo string `json:"dataStoreInfo,omitempty"`

		// Indicates the returned extendparam key-value pair.
		Extendparam struct {

			// Indicates the returned  parameter information.
			Jobs []struct {

				// Indicates the task ID.
				Id string `json:"id,omitempty"`
			} `json:"jobs,omitempty"`
		} `json:"extendparam,omitempty"`

		// Indicates the backup policy.
		BackupStrategy struct {

			// Indicates the backup start time that has been set. The backup task will be triggered within one hour after the backup start time.
			StartTime string `json:"startTime,omitempty"`

			// Indicates the number of days to retain the generated backup files.
			KeepDays int `json:"keepDays,omitempty"`
		} `json:"backupStrategy,omitempty"`
	} `json:"instance,omitempty"`
}
