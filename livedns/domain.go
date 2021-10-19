package livedns

import (
	"github.com/go-gandi/go-gandi/types"
)

// ListDomains lists all domains
func (g *LiveDNSAPI) ListDomains() (domains []Domain, err error) {
	_, err = g.client.Get("domains", nil, &domains)
	return
}

// CreateDomain adds a domain to a zone
func (g *LiveDNSAPI) CreateDomain(fqdn string, ttl int) (response types.StandardResponse, err error) {
	_, err = g.client.Post("domains", createDomainRequest{FQDN: fqdn, Zone: zone{TTL: ttl}}, &response)
	return
}

// GetDomain returns a domain
func (g *LiveDNSAPI) GetDomain(fqdn string) (domain Domain, err error) {
	_, err = g.client.Get("domains/"+fqdn, nil, &domain)
	return
}

// UpdateDomain changes the zone associated to a domain
func (g *LiveDNSAPI) UpdateDomain(fqdn string, details UpdateDomainRequest) (response types.StandardResponse, err error) {
	_, err = g.client.Patch("domains/"+fqdn, details, &response)
	return
}

// GetDomainNS returns the list of the nameservers for a domain
func (g *LiveDNSAPI) GetDomainNS(fqdn string) (ns []string, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/nameservers", nil, &ns)
	return
}
