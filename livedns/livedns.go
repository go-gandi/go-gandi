package livedns

import (
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the LiveDNS API client
func New(apikey string, sharingid string, debug bool, dryRun bool) *LiveDNSAPI {
	client := client.New(apikey, sharingid, debug, dryRun)
	client.SetEndpoint("livedns/")
	return &LiveDNSAPI{client: *client}
}

// NewFromClient returns an instance of the LiveDNS API client
func NewFromClient(g client.Gandi) *LiveDNSAPI {
	g.SetEndpoint("livedns/")
	return &LiveDNSAPI{client: g}
}
