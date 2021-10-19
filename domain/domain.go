package domain

import (
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Domain API client
func New(apikey string, sharingid string, debug bool, dryRun bool) *DomainAPI {
	client := client.New(apikey, sharingid, debug, dryRun)
	client.SetEndpoint("domain/")
	return &DomainAPI{client: *client}
}

// NewFromClient returns an instance of the Domain API client
func NewFromClient(g client.Gandi) *DomainAPI {
	g.SetEndpoint("domain/")
	return &DomainAPI{client: g}
}

// ListDomains requests the set of Domains
// It returns a slice of domains and any error encountered
func (g *DomainAPI) ListDomains() (domains []ListResponse, err error) {
	_, err = g.client.Get("domains", nil, &domains)
	return
}

// GetDomain requests a single Domain
// It returns a Details object and any error encountered
func (g *DomainAPI) GetDomain(domain string) (domainResponse Details, err error) {
	_, err = g.client.Get("domains/"+domain, nil, &domainResponse)
	return
}

// CreateDomain creates a single Domain
func (g *DomainAPI) CreateDomain(req CreateRequest) (err error) {
	_, err = g.client.Post("domains", req, nil)
	return
}

// GetNameServers returns the configured nameservers for a domain
func (g *DomainAPI) GetNameServers(domain string) (nameservers []string, err error) {
	_, err = g.client.Get("domains/"+domain+"/nameservers", nil, &nameservers)
	return
}

// UpdateNameServers sets the list of the nameservers for a domain
func (g *DomainAPI) UpdateNameServers(domain string, ns []string) (err error) {
	_, err = g.client.Put("domains/"+domain+"/nameservers", Nameservers{ns}, nil)
	return
}

// GetContacts returns the contact objects for a domain
func (g *DomainAPI) GetContacts(domain string) (contacts Contacts, err error) {
	_, err = g.client.Get("domains/"+domain+"/contacts", nil, &contacts)
	return
}

// SetContacts sets the contact objects for a domain
func (g *DomainAPI) SetContacts(domain string, contacts Contacts) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/contacts", contacts, nil)
	return
}

// SetAutoRenew enables or disables auto renew on the given Domain
func (g *DomainAPI) SetAutoRenew(domain string, autorenew bool) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/autorenew", AutoRenew{Enabled: &autorenew}, nil)
	return
}
