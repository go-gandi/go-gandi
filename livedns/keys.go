package livedns

import (
	"github.com/go-gandi/go-gandi/types"
)

// GetTSIGKeys retrieves all the TSIG keys for the account
func (g *LiveDNSAPI) GetTSIGKeys() (response []TSIGKey, err error) {
	_, err = g.client.Get("axfr/tsig", nil, &response)
	return
}

// GetTSIGKey retrieves the specified TSIG key
func (g *LiveDNSAPI) GetTSIGKey(id string) (response TSIGKey, err error) {
	_, err = g.client.Get("axfr/tsig/"+id, nil, &response)
	return
}

// CreateTSIGKey creates a TSIG key
func (g *LiveDNSAPI) CreateTSIGKey(fqdn string) (response TSIGKey, err error) {
	_, err = g.client.Post("axfr/tsig", nil, &response)
	return
}

// GetDomainTSIGKeys retrieves the specified TSIG key
func (g *LiveDNSAPI) GetDomainTSIGKeys(fqdn string) (response []TSIGKey, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/axfr/tsig", nil, &response)
	return
}

// AssociateTSIGKeyWithDomain retrieves the specified TSIG key
func (g *LiveDNSAPI) AssociateTSIGKeyWithDomain(fqdn string, id string) (response types.StandardResponse, err error) {
	_, err = g.client.Put("domains/"+fqdn+"/axfr/tsig/"+id, nil, &response)
	return
}

// RemoveTSIGKeyFromDomain retrieves the specified TSIG key
func (g *LiveDNSAPI) RemoveTSIGKeyFromDomain(fqdn string, id string) (err error) {
	_, err = g.client.Delete("domains/"+fqdn+"/axfr/tsig/"+id, nil, nil)
	return
}

// SignDomain creates a DNSKEY and asks Gandi servers to automatically sign the domain
func (g *LiveDNSAPI) SignDomain(fqdn string) (response types.StandardResponse, err error) {
	f := SigningKey{Flags: 257}
	_, err = g.client.Post("domains/"+fqdn+"/keys", f, &response)
	return
}

// GetDomainKeys returns data about the signing keys created for a domain
func (g *LiveDNSAPI) GetDomainKeys(fqdn string) (keys []SigningKey, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/keys", nil, &keys)
	return
}

// GetDomainKey deletes a signing key from a domain
func (g *LiveDNSAPI) GetDomainKey(fqdn, uuid string) (key SigningKey, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/keys/"+uuid, nil, &key)
	return
}

// DeleteDomainKey deletes a signing key from a domain
func (g *LiveDNSAPI) DeleteDomainKey(fqdn, uuid string) (err error) {
	_, err = g.client.Delete("domains/"+fqdn+"/keys/"+uuid, nil, nil)
	return
}

// UpdateDomainKey updates a signing key for a domain (only the deleted status, actually...)
func (g *LiveDNSAPI) UpdateDomainKey(fqdn, uuid string, deleted bool) (err error) {
	_, err = g.client.Put("domains/"+fqdn+"/keys/"+uuid, SigningKey{Deleted: &deleted}, nil)
	return
}
