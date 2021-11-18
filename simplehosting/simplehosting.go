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
