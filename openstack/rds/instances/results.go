package instances

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {

	// Indicates the DB instance information.
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

		// Returned only when creating HA DB instances.
		SlaveId string `json:"slaveId,omitempty"`

		// Indicates the HA DB instance information. Returned only when creating HA DB instances.
		Ha struct {

			// Indicates the replication mode for the standby DB instance. Same as the request parameter.
			ReplicationMode string `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`

		// Same as the request parameter. Returned only when creating a read replica.
		ReplicaOf string `json:"replica_of,omitempty"`
	} `json:"instance,omitempty"`
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

	// Indicates the returned extendparam key-value pair.
	Extendparam struct {

		// Indicates the returned jobs parameter information.
		Jobs []struct {

			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	} `json:"extendparam,omitempty"`
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*GetResponse, error) {
	var response GetResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetResponse struct {

	// Indicates the DB instance ID.
	Instance struct {

		// Indicates the DB instance ID.
		Id string `json:"id,omitempty"`

		// Indicates the DB instance status.Value:If the value is BUILD, the instance is being created.If the value is ACTIVE, the instance is normal.If the value is FAILED, the instance is abnormal.If the value is MODIFYING, the instance is being scaled up.If the value is REBOOTING, the instance is being restarted.If the value is RESTORING, the instance is being restored.
		Status string `json:"status,omitempty"`

		// Indicates the created DB instance name.
		Name string `json:"name,omitempty"`

		// Indicates the creation time in the following format: yyyy-mm-ddThh:mm:ssZ.T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset. For example, in the Beijing time zone, the time zone offset is shown as +0800.
		Created string `json:"created,omitempty"`

		// Indicates the instance connection address. It is a blank string until an ECS is created.
		Hostname string `json:"hostname,omitempty"`

		// Indicates the DB instance type, which can be master, slave, or readreplica.
		Type string `json:"type,omitempty"`

		// Indicates the region where the DB instance exists.
		Region string `json:"region,omitempty"`

		// Indicates the updated time, which is the same as created in the format.
		Updated string `json:"updated,omitempty"`

		// Indicates the AZ ID.
		AvailabilityZone string `json:"availabilityZone,omitempty"`

		// Indicates the VPC ID.
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

		// Indicates the specification information.
		Flavor struct {

			// Indicates the specification ID.
			Id string `json:"id,omitempty"`
		} `json:"flavor,omitempty"`

		// Indicates the volume information.
		Volume struct {

			// Indicates the volume type.
			Type string `json:"type,omitempty"`

			// Indicates the volume size.
			Size int `json:"size,omitempty"`
		} `json:"volume,omitempty"`

		// Indicates the database information.
		DataStoreInfo struct {

			// Indicates the DB engine.
			Type string `json:"type,omitempty"`

			// Indicates the database version.
			Version string `json:"version,omitempty"`
		} `json:"dataStoreInfo,omitempty"`

		// Indicates the database port number.
		DbPort int `json:"dbPort,omitempty"`

		// Indicates the advanced backup policy.
		BackupStrategy struct {

			// Indicates the backup start time that has been set. The backup task will be triggered within one hour after the backup start time.The current time is the UTC time.
			StartTime string `json:"startTime,omitempty"`

			// Indicates the number of days to retain the generated backup files.Its value range is from 0 to 35. If this parameter is 0, the automated backup policy is not set.
			KeepDays int `json:"keepDays,omitempty"`
		} `json:"backupStrategy,omitempty"`

		// Returned only in HA DB instance creation.
		SlaveId string `json:"slaveId,omitempty"`

		// Indicates the HA DB instance information. Returned only when obtaining HA DB instances.
		Ha []struct {

			// Indicates the replication mode for the standby DB instance.For MySQL, the value is async or semisync.For HWSQL, the value is async or semisync.For PostgreSQL, the value is async or sync.For Microsoft SQL Server, the value is sync.
			ReplicationMode string `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`

		// Returned only when obtaining the read replica information.
		ReplicaOf string `json:"replicaOf,omitempty"`
	} `json:"instance,omitempty"`
}

type GetFlavorResult struct {
	commonResult
}

func (r GetFlavorResult) Extract() (*GetFlavorResponse, error) {
	var response GetFlavorResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type GetFlavorResponse struct {

	// Indicates the DB instance specifications information list.
	Flavor struct {

		// Indicates the specification ID. Its value is unique.
		Id string `json:"id,omitempty"`

		// Indicates the specification item name.
		Name string `json:"name,omitempty"`

		// Indicates the memory size in megabytes (MB).
		Ram int `json:"ram,omitempty"`

		// Indicates the resource specifications code.Use rds.mysql.m1.xlarge as an example.rds indicates RDS, mysql indicates the DB engine, and m1.xlarge indicates the performance specification (large-memory). The parameter containing rr indicates the read replica specifications. The parameter not containing rr indicates the single or HA DB instance specifications.
		SpecCode string `json:"specCode,omitempty"`
	} `json:"flavor,omitempty"`
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

	// Indicates the DB instance ID.
	Instances []struct {

		// Indicates the DB instance ID.
		Id string `json:"id,omitempty"`

		// Indicates the DB instance status.Value:If the value is BUILD, the instance is being created.If the value is ACTIVE, the instance is normal.If the value is FAILED, the instance is abnormal.If the value is MODIFYING, the instance is being scaled up.If the value is REBOOTING, the instance is being restarted.If the value is RESTORING, the instance is being restored.
		Status string `json:"status,omitempty"`

		// Indicates the created DB instance name.
		Name string `json:"name,omitempty"`

		// Indicates the creation time in the following format: yyyy-mm-ddThh:mm:ssZ.T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset. For example, in the Beijing time zone, the time zone offset is shown as +0800.
		Created string `json:"created,omitempty"`

		// Indicates the instance connection address. It is a blank string until an ECS is created.
		Hostname string `json:"hostname,omitempty"`

		// Indicates the DB instance type, which can be master, slave, or readreplica.
		Type string `json:"type,omitempty"`

		// Indicates the region where the DB instance exists.
		Region string `json:"region,omitempty"`

		// Indicates the updated time, which is the same as created in the format.
		Updated string `json:"updated,omitempty"`

		// Indicates the AZ ID.
		AvailabilityZone string `json:"availabilityZone,omitempty"`

		// Indicates the VPC ID.
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

		// Indicates the specification information.
		Flavor struct {

			// Indicates the specification ID.
			Id string `json:"id,omitempty"`
		} `json:"flavor,omitempty"`

		// Indicates the volume information.
		Volume struct {

			// Indicates the volume type.
			Type string `json:"type,omitempty"`

			// Indicates the volume size.
			Size int `json:"size,omitempty"`
		} `json:"volume,omitempty"`

		// Indicates the database information.
		DataStoreInfo struct {

			// Indicates the DB engine.
			Type string `json:"type,omitempty"`

			// Indicates the database version.
			Version string `json:"version,omitempty"`
		} `json:"dataStoreInfo,omitempty"`

		// Indicates the database port number.
		DbPort int `json:"dbPort,omitempty"`

		// Indicates the advanced backup policy.
		BackupStrategy struct {

			// Indicates the backup start time that has been set. The backup task will be triggered within one hour after the backup start time.The current time is the UTC time.
			StartTime string `json:"startTime,omitempty"`

			// Indicates the number of days to retain the generated backup files.Its value range is from 0 to 35. If this parameter is 0, the automated backup policy is not set.
			KeepDays int `json:"keepDays,omitempty"`
		} `json:"backupStrategy,omitempty"`

		// Returned only in HA DB instance creation.
		SlaveId string `json:"slaveId,omitempty"`

		// Indicates the HA DB instance information. Returned only when obtaining HA DB instances.
		Ha []struct {

			// Indicates the replication mode for the standby DB instance.For MySQL, the value is async or semisync.For HWSQL, the value is async or semisync.For PostgreSQL, the value is async or sync.For Microsoft SQL Server, the value is sync.
			ReplicationMode string `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`

		// Returned only when obtaining the read replica information.
		ReplicaOf string `json:"replicaOf,omitempty"`
	} `json:"instances,omitempty"`
}

type ListFlavorsResult struct {
	commonResult
}

func (r ListFlavorsResult) Extract() (*ListFlavorsResponse, error) {
	var response ListFlavorsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListFlavorsResponse struct {

	// Indicates the DB instance specifications information list.
	Flavors []struct {

		// Indicates the specification ID. Its value is unique.
		Id string `json:"id,omitempty"`

		// Indicates the specification item name.
		Name string `json:"name,omitempty"`

		// Indicates the memory size in megabytes (MB).
		Ram int `json:"ram,omitempty"`

		// Indicates the resource specifications code.Use rds.mysql.m1.xlarge as an example.rds indicates RDS, mysql indicates the DB engine, and m1.xlarge indicates the performance specification (large-memory). The parameter containing rr indicates the read replica specifications. The parameter not containing rr indicates the single or HA DB instance specifications.
		SpecCode string `json:"specCode,omitempty"`
	} `json:"flavors,omitempty"`
}

type RebootResult struct {
	commonResult
}

func (r RebootResult) Extract() (*RebootResponse, error) {
	var response RebootResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RebootResponse struct {

	// Indicates the returned extendparam key-value pair.
	Extendparam struct {

		// Indicates the returned jobs parameter information.
		Jobs []struct {

			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	} `json:"extendparam,omitempty"`
}

type UpdateResult struct {
	commonResult
}

func (r UpdateResult) Extract() (*UpdateResponse, error) {
	var response UpdateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type UpdateResponse struct {

	// Indicates the returned extendparam key-value pair.
	Extendparam struct {

		// Indicates the returned  parameter information.
		Jobs []struct {

			// Indicates the task ID.
			Id string `json:"id,omitempty"`
		} `json:"jobs,omitempty"`
	} `json:"extendparam,omitempty"`
}
