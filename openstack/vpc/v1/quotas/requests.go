package quotas

import (
	"github.com/huaweicloud/golangsdk"
)

type Quota struct {

	// Specifies the resource list objects.
	Resources []Resource `json:"resources,"`
}

type Resource struct {

	// Specifies the resource type. The value can be vpc, subnet, securityGroup, securityGroupRule, publicIp, vpn, physicalConnect, virtualInterface, vpcPeer, loadbalancer, listener, firewall, or shareBandwidthIP.
	Type string `json:"type,"`

	// Specifies the number of created network resources. The value ranges from 0 to the value of quota.
	Used int `json:"used,"`

	// Specifies the maximum quota values for the resources. The quotas can be changed only in the FusionSphere OpenStack system. If it is left blank, -1 is displayed and the resources cannot be created. The default quotas for different resources are as follows:  VPC: 2 Subnet: 100 Security group: 100 Security group rule: 5000 Elastic IP address: 10  VPN: 5 Physical connection: 10 Virtual interface: 50 Load balancer: 10 Listener: 10 VPC peering connection: 50 Firewall: 200 IP address with shared bandwidth: 20 The value ranges from the default quota value to the maximum quota value.
	Quota int `json:"quota,"`

	// Specifies the minimum quota value allowed.
	Min int `json:"min,"`
}

type ListOpts struct {

	// Specifies the resource type.
	Type string `q:"type"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, tenantId string, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client, tenantId)
	if opts != nil {
		query, _ := opts.ToListQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202, 204},
	})
	return
}
