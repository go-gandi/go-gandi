package livedns

// ListTsigs lists all tsigs
func (g *LiveDNSAPI) ListTsigs() (tsigs []Tsig, err error) {
	_, err = g.client.Get("axfr/tsig", nil, &tsigs)
	return
}

// GetTsig lists more tsig details
func (g *LiveDNSAPI) GetTsig(uuid string) (tsig Tsig, err error) {
	_, err = g.client.Get("axfr/tsig/"+uuid, nil, &tsig)
	return
}

// GetTsigBIND shows a BIND nameserver config, and includes the nameservers available for zone transfers
func (g *LiveDNSAPI) GetTsigBIND(uuid string) ([]byte, error) {
	_, content, err := g.client.GetBytes("axfr/tsig/"+uuid+"/config/bind", nil)
	return content, err
}

// GetTsigPowerDNS shows a PowerDNS nameserver config, and includes the nameservers available for zone transfers
func (g *LiveDNSAPI) GetTsigPowerDNS(uuid string) ([]byte, error) {
	_, content, err := g.client.GetBytes("axfr/tsig/"+uuid+"/config/powerdns", nil)
	return content, err
}

// GetTsigNSD shows a NSD nameserver config, and includes the nameservers available for zone transfers
func (g *LiveDNSAPI) GetTsigNSD(uuid string) ([]byte, error) {
	_, content, err := g.client.GetBytes("axfr/tsig/"+uuid+"/config/nsd", nil)
	return content, err
}

// GetTsigKnot shows a Knot nameserver config, and includes the nameservers available for zone transfers
func (g *LiveDNSAPI) GetTsigKnot(uuid string) ([]byte, error) {
	_, content, err := g.client.GetBytes("axfr/tsig/"+uuid+"/config/knot", nil)
	return content, err
}

// CreateTsig creates a tsig
func (g *LiveDNSAPI) CreateTsig() (tsig Tsig, err error) {
	_, err = g.client.Post("axfr/tsig", nil, &tsig)
	return
}

// AddTsigToDomain adds a tsig to a domain
func (g *LiveDNSAPI) AddTsigToDomain(fqdn, uuid string) (err error) {
	_, err = g.client.Put("domains/"+fqdn+"/axfr/tsig/"+uuid, nil, nil)
	return
}

// AddSlaveToDomain adds a slave to a domain
func (g *LiveDNSAPI) AddSlaveToDomain(fqdn, host string) (err error) {
	_, err = g.client.Put("domains/"+fqdn+"/axfr/slaves/"+host, nil, nil)
	return
}

// ListSlavesInDomain lists slaves in a domain
func (g *LiveDNSAPI) ListSlavesInDomain(fqdn string) (slaves []string, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/axfr/slaves", nil, &slaves)
	return
}

// DelSlaveFromDomain removes a slave from a domain
func (g *LiveDNSAPI) DelSlaveFromDomain(fqdn, host string) (err error) {
	_, err = g.client.Delete("domains/"+fqdn+"/axfr/slaves/"+host, nil, nil)
	return
}
