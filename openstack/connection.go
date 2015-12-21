package openstack

import (
	"github.com/rackspace/gophercloud"
	_openstack "github.com/rackspace/gophercloud/openstack"
)

type Connection struct {
	opts     gophercloud.AuthOptions
	provider *gophercloud.ProviderClient
	object   *gophercloud.ServiceClient
}

func (conn *Connection) connect() bool {
	var err error
	conn.opts, err = _openstack.AuthOptionsFromEnv()
	conn.provider, err = _openstack.AuthenticatedClient(conn.opts)
	if err != nil {
		return false
	}
	return true
}

func (conn *Connection) flavors() bool {
	var err error
	epo := gophercloud.EndpointOpts{Region: "NCW"}
	conn.object, err = _openstack.NewObjectStorageV1(conn.provider, epo)
	if err != nil {
		return false
	}
	return true
}
