package livedns

import (
	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the LiveDNS API client
func New(apikey string, config config.Config) *LiveDNS {
	client := client.New(apikey, config.APIURL, config.SharingID, config.Debug, config.DryRun)
	client.SetEndpoint("livedns/")
	return &LiveDNS{client: *client}
}

// NewFromClient returns an instance of the LiveDNS API client
func NewFromClient(g client.Gandi) *LiveDNS {
	g.SetEndpoint("livedns/")
	return &LiveDNS{client: g}
}
