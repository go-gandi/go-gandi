package simplehosting

import (
	"fmt"
	"strings"

	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Simple Hosting API client
func New(apikey string, sharingid string, debug bool, dryRun bool) *SimpleHosting {
	client := client.New(apikey, sharingid, debug, dryRun)
	client.SetEndpoint("simplehosting/")
	return &SimpleHosting{client: *client}
}

// ListInstances requests the list of SimpleHosting instances
func (g *SimpleHosting) ListInstances() (simplehostings []ListInstancesResponse, err error) {
	_, err = g.client.Get("instances", nil, &simplehostings)
	return
}

// GetInstance requests a single Instance
func (g *SimpleHosting) GetInstance(instanceId string) (simplehostingResponse Instance, err error) {
	_, err = g.client.Get("instances/"+instanceId, nil, &simplehostingResponse)
	return
}

// CreateInstance creates a SimpleHosting instance
func (g *SimpleHosting) CreateInstance(req CreateInstanceRequest) (instanceId string, err error) {
	header, err := g.client.Post("instances", req, nil)
	if err != nil {
		return "", err
	}
	// We have to extract the instance ID from the
	// Content-Location response header.
	location := header.Get("Content-Location")
	endpoint := g.client.GetEndpoint() + "instances/"
	if strings.HasPrefix(location, endpoint) {
		return strings.TrimPrefix(location, endpoint), nil
	} else {
		return "", fmt.Errorf("Could not extract the instance ID from '%s'", location)
	}
}

// CreateInstance deletes a SimpleHosting instance
func (g *SimpleHosting) DeleteInstance(instanceId string) (response ErrorResponse, err error) {
	_, err = g.client.Delete("instances/"+instanceId, nil, &response)
	return
}

// ListVhosts lists vhosts of a Simple Hosting instance
func (g *SimpleHosting) ListVhosts(instanceId string) (response []Vhost, err error) {
	_, err = g.client.Get("instances/"+instanceId+"/vhosts", nil, &response)
	return
}

// ListVhosts creates a vhost for a Simple Hosting instance
func (g *SimpleHosting) CreateVhost(instanceId string, req CreateVhostRequest) (response Vhost, err error) {
	_, err = g.client.Post("instances/"+instanceId+"/vhosts", req, &response)
	if err != nil {
		return Vhost{}, err
	}
	return
}

// ListVhosts deletes vhosts of a Simple Hosting instance
func (g *SimpleHosting) DeleteVhost(instanceId string, fqdn string) (response ErrorResponse, err error) {
	_, err = g.client.Delete("instances/"+instanceId+"/vhosts/"+fqdn, nil, &response)
	return
}
