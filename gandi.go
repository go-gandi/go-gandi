package gandi

import (
	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/email"
	"github.com/go-gandi/go-gandi/livedns"
	"github.com/go-gandi/go-gandi/simplehosting"
)

// NewDomainClient returns a client to the Gandi Domains API
// It expects an API key, available from https://account.gandi.net/en/
func NewDomainClient(apikey string, config config.Config) *domain.Domain {
	return domain.New(apikey, config)
}

// NewEmailClient returns a client to the Gandi Email API
// It expects an API key, available from https://account.gandi.net/en/
func NewEmailClient(apikey string, config config.Config) *email.Email {
	return email.New(apikey, config)
}

// NewLiveDNSClient returns a client to the Gandi Domains API
// It expects an API key, available from https://account.gandi.net/en/
func NewLiveDNSClient(apikey string, config config.Config) *livedns.LiveDNS {
	return livedns.New(apikey, config)
}

// NewSimpleHostingClient returns a client to the Gandi Simple Hosting API
// It expects an API key, available from https://account.gandi.net/en/
func NewSimpleHostingClient(apikey string, config config.Config) *simplehosting.SimpleHosting {
	return simplehosting.New(apikey, config)
}
