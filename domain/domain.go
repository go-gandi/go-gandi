package domain

import (
	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Domain API client
func New(config config.Config) *Domain {
	client := client.New(config.APIKey, config.APIURL, config.SharingID, config.Debug, config.DryRun)
	client.SetEndpoint("domain/")
	return &Domain{client: *client}
}

// NewFromClient returns an instance of the Domain API client
func NewFromClient(g client.Gandi) *Domain {
	g.SetEndpoint("domain/")
	return &Domain{client: g}
}

// ListDomains requests the set of Domains
// It returns a slice of domains and any error encountered
func (g *Domain) ListDomains() (domains []ListResponse, err error) {
	_, err = g.client.Get("domains", nil, &domains)
	return
}

// GetDomain requests a single Domain
// It returns a Details object and any error encountered
func (g *Domain) GetDomain(domain string) (domainResponse Details, err error) {
	_, err = g.client.Get("domains/"+domain, nil, &domainResponse)
	return
}

// CreateDomain creates a single Domain
func (g *Domain) CreateDomain(req CreateRequest) (err error) {
	_, err = g.client.Post("domains", req, nil)
	return
}

// GetNameServers returns the configured nameservers for a domain
func (g *Domain) GetNameServers(domain string) (nameservers []string, err error) {
	_, err = g.client.Get("domains/"+domain+"/nameservers", nil, &nameservers)
	return
}

// UpdateNameServers sets the list of the nameservers for a domain
func (g *Domain) UpdateNameServers(domain string, ns []string) (err error) {
	_, err = g.client.Put("domains/"+domain+"/nameservers", Nameservers{ns}, nil)
	return
}

// GetContacts returns the contact objects for a domain
func (g *Domain) GetContacts(domain string) (contacts Contacts, err error) {
	_, err = g.client.Get("domains/"+domain+"/contacts", nil, &contacts)
	return
}

// SetContacts sets the contact objects for a domain
func (g *Domain) SetContacts(domain string, contacts Contacts) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/contacts", contacts, nil)
	return
}

// SetAutoRenew enables or disables auto renew on the given Domain
func (g *Domain) SetAutoRenew(domain string, autorenew bool) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/autorenew", AutoRenew{Enabled: &autorenew}, nil)
	return
}

func (g *Domain) ListDNSSECKeys(domain string) (keys []DNSSECKey, err error) {
	_, err = g.client.Get("domains/"+domain+"/dnskeys", nil, &keys)
	return
}

func (g *Domain) CreateDNSSECKey(domain string, key DNSSECKeyCreateRequest) (err error) {
	_, err = g.client.Post("domains/"+domain+"/dnskeys", key, nil)
	return
}

func (g *Domain) DeleteDNSSECKey(domain string, keyid string) (err error) {
	_, err = g.client.Delete("domains/"+domain+"/dnskeys/"+keyid, nil, nil)
	return
}

func (g *Domain) CreateGlueRecord(domain string, gluerecord GlueRecordCreateRequest) (err error) {
	_, err = g.client.Post("domains/"+domain+"/hosts", gluerecord, nil)
	return
}

func (g *Domain) ListGlueRecords(domain string) (gluerecords []GlueRecord, err error) {
	_, err = g.client.Get("domains/"+domain+"/hosts", nil, &gluerecords)
	return
}

func (g *Domain) UpdateGlueRecord(domain string, name string, ips []string) (err error) {
	_, err = g.client.Put("domains/"+domain+"/hosts/"+name, GlueRecordUpdateRequest{ips}, nil)
	return
}

func (g *Domain) DeleteGlueRecord(domain string, name string) (err error) {
	_, err = g.client.Delete("domains/"+domain+"/hosts/"+name, nil, nil)
	return
}

func (g *Domain) CreateWebRedirection(domain string, webredir WebRedirectionCreateRequest) (err error) {
	_, err = g.client.Post("domains/"+domain+"/webredirs", webredir, nil)
	return
}

func (g *Domain) ListWebRedirections(domain string) (webredirs []WebRedirection, err error) {
	_, err = g.client.Get("domains/"+domain+"/webredirs", nil, &webredirs)
	return
}

func (g *Domain) DeleteWebRedirection(domain string, host string) (err error) {
	_, err = g.client.Delete("domains/"+domain+"/webredirs/"+host, nil, nil)
	return
}
