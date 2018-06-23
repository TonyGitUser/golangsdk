package instances

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOpts struct {

	// Specifies the detailed DB instance information.
	Instance struct {

		// Specifies the DB instance name. The DB instance name of the same DB engine is unique for the same tenant.Valid value:The value must be 4 to 64 characters in length and start with a letter. It is case-insensitive and can contain only letters, digits, hyphens (-), and underscores (_).
		Name string `json:"name,omitempty"`

		// Specifies database information.
		Datastore struct {

			// Specifies the DB engine. Currently, MySQL, HWSQL, PostgreSQL, and Microsoft SQL Server are supported.The value is MySQL, HWSQL, PostgreSQL, or SQLServer.
			Type string `json:"type,omitempty"`

			// Specifies the database version.MySQL databases support MySQL 5.6 and 5.7. Example value: 5.7.HWSQL databases support HWSQL 5.6. Example value: 5.6.PostgreSQL databases support PostgreSQL 9.5 and 9.6. Example value: 9.6.Microsoft SQL Server databases support Microsoft SQL Server 2014 SP2 SE, 2014 SP2 EE, 2016 SP1 SE, 2016 SP1 EE, and 2008 R2 SP3 SE. Example value: 2008_R2_SP3_SE
			Version string `json:"version,omitempty"`
		} `json:"datastore,omitempty"`

		// Specifies the specification ID (flavors.id in the response message in section Obtaining All DB Instance Specifications).Valid value:The value cannot be empty. The name length and whether the name complies with UUID regular expression rules are verified.
		FlavorRef string `json:"flavorRef,omitempty"`

		// Specifies the volume information.
		Volume struct {

			// Specifies the volume type.Valid value:It must be COMMON (SATA), HIGH (SAS), or ULTRAHIGH (SSD) and is case-sensitive.
			Type string `json:"type,omitempty"`

			// Specifies the volume size.Its value must be a multiple of 10 and the value range is from 40 GB to 2,000 GB.
			Size int `json:"size,omitempty"`
		} `json:"volume,omitempty"`

		// Specifies the region ID.Valid value:The value cannot be empty. For details about how to obtain this parameter value, see Regions and Endpoints.
		Region string `json:"region,omitempty"`

		// Specifies the ID of the AZ.Valid value:The value cannot be empty. For details about how to obtain this parameter value, see Regions and Endpoints.
		AvailabilityZone string `json:"availabilityZone,omitempty"`

		// Specifies the VPC ID. For details about how to obtain this parameter value, see section "Virtual Private Cloud" in the Virtual Private Cloud API Reference.Valid value:The value cannot be empty. The name length and whether the name complies with UUID regular expression rules are verified.
		Vpc string `json:"vpc,omitempty"`

		// Specifies the nics information. For details about how to obtain this parameter value, see section "Subnet" in the Virtual Private Cloud API Reference.
		Nics struct {

			// Specifies the subnet ID obtained from the VPC.Valid value:The value cannot be empty. The name length and whether the name complies with UUID regular expression rules are verified.
			SubnetId string `json:"subnetId,omitempty"`
		} `json:"nics,omitempty"`

		// Specifies the security group which the RDS DB instance belongs to.For details about how to obtain this parameter value, see section "Security Group" in the Virtual Private Cloud API Reference.
		SecurityGroup struct {

			// Valid value:The value cannot be empty. The name length and whether the name complies with UUID regular expression rules are verified.
			Id string `json:"id,omitempty"`
		} `json:"securityGroup,omitempty"`

		// Specifies the database port number.
		DbPort int `json:"dbPort,omitempty"`

		// Specifies the advanced backup policy.
		BackupStrategy struct {

			// Specifies the backup start time that has been set.Valid value:The value cannot be empty. It must use the hh:mm:ss format and must be valid. The current time is the UTC time.
			StartTime string `json:"startTime,omitempty"`

			// Specifies the number of days to retain the generated backup files.Its value range is from 0 to 35. If this parameter is not specified or set to 0, the automated backup policy is disabled.
			KeepDays int `json:"keepDays,omitempty"`
		} `json:"backupStrategy,omitempty"`

		// Specifies the password for user root of the database.Valid value:The value cannot be empty and should contain 8 to 32 characters, including uppercase and lowercase letters, digits, and the following special characters: ~!@#%^*-_=+?
		DbRtPd string `json:"dbRtPd,omitempty"`

		// Specifies the parameters configured on HA and is used when creating HA DB instances.
		Ha struct {

			// Specifies the HA configuration parameter.Valid value:The value is true or false. The value true indicates creating HA DB instances. The value false indicates creating a single DB instance.
			Enable bool `json:"enable,omitempty"`

			// Specifies the replication mode for the standby DB instance.The value cannot be empty.For MySQL, the value is async or semisync.For PostgreSQL, the value is async or sync.For Microsoft SQL Server, the value is sync.For HWSQL, the value is async or semisync.
			ReplicationMode string `json:"replicationMode,omitempty"`
		} `json:"ha,omitempty"`
	} `json:"instance,omitempty"`
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, versionId string, projectId string, opts CreateOptsBuilder) (r CreateResult) {
	b, _ := opts.ToCreateMap()

	_, r.Err = client.Post(CreateURL(client, versionId, projectId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

type DeleteOpts struct {

	// This parameter is obsolete. Its value can be left empty or set to an integer and has no impact on calling the interface.
	KeepLastManualBackup string `q:"keepLastManualBackup"`
}

type DeleteOptsBuilder interface {
	ToDeleteQuery() (string, error)
}

func (opts DeleteOpts) ToDeleteQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func Delete(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := DeleteURL(client, versionId, projectId, instanceId)
	if opts != nil {
		query, _ := opts.ToDeleteQuery()
		url += query
	}

	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200, 201, 202, 204},
		MoreHeaders:  map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func Get(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string) (r GetResult) {
	url := GetURL(client, versionId, projectId, instanceId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

func GetFlavor(client *golangsdk.ServiceClient, versionId string, projectId string, flavorId string) (r GetFlavorResult) {
	url := GetFlavorURL(client, versionId, projectId, flavorId)
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

type ListFlavorsOpts struct {

	// Specifies the database version ID (dataStores.id in the response message in Database Version Queries).Valid value: The value cannot be empty. The name length and whether the name complies with UUID regular expression rules are verified.
	DbId string `q:"dbId"`

	// Specifies the region where the DB instance exists.Valid value:The value cannot be empty. For details about how to obtain this parameter value, see Regions and Endpoints.
	Region string `q:"region"`
}

type ListFlavorsOptsBuilder interface {
	ToListFlavorsQuery() (string, error)
}

func (opts ListFlavorsOpts) ToListFlavorsQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func ListFlavors(client *golangsdk.ServiceClient, versionId string, projectId string, opts ListFlavorsOptsBuilder) (r ListFlavorsResult) {
	url := ListFlavorsURL(client, versionId, projectId)
	if opts != nil {
		query, _ := opts.ToListFlavorsQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

type RebootOpts struct {

	// Specifies an empty value.
	Restart struct {
	} `json:"restart,"`
}

type RebootOptsBuilder interface {
	ToRebootMap() (map[string]interface{}, error)
}

func (opts RebootOpts) ToRebootMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Reboot(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts RebootOptsBuilder) (r RebootResult) {
	b, _ := opts.ToRebootMap()

	_, r.Err = client.Post(RebootURL(client, versionId, projectId, instanceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}

type UpdateOpts struct {

	// Specifies the information about the returned parameter flavorRef.
	Resize struct {

		// Specifies the specification ID (flavors.id in the response message in Obtaining All DB Instance Specifications).
		FlavorRef string `json:"flavorRef,omitempty"`

		// Specifies the information about the request parameter size.
		Volume struct {

			// Specifies the disk size after scaling up. The value must a multiple of 10.Its value range is from 160 GB to 2,000 GB.
			Size int `json:"size,omitempty"`
		} `json:"volume,omitempty"`
	} `json:"resize,omitempty"`
}

type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *golangsdk.ServiceClient, versionId string, projectId string, instanceId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, _ := opts.ToUpdateMap()

	_, r.Err = client.Post(UpdateURL(client, versionId, projectId, instanceId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201, 202, 204},
		MoreHeaders: map[string]string{"X-Language": "en-us", "Content-Type": "application/json"},
	})
	return
}
