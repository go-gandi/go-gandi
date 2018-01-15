package gandi

// Tsig contains tsig data (no kidding!)
type Tsig struct {
	KeyName     string `json:"key_name, omitempty"`
	Secret      string `json:"secret,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	AxfrTsigURL string `json:"axfr_tsig_url,omitempty"`
}

// ListTsigs lists all tsigs
func (g *Gandi) ListTsigs() (tsigs []Tsig, err error) {
	_, err = g.askGandi(mGET, "axfr/tsig", nil, &tsigs)
	return
}

// By gandi API design, an axfr/tsig/uuid/config/[bind|knot|powerdns|nsd] request will not return a valid JSON,
// but rather the actual configuration example.
// Hence we have to treat the return as 'configexample string' instead of 'configexample []string'.

// ListTsigBind shows a BIND nameserver config, and includes the nameservers available for zone transfers
func (g *Gandi) ListTsigBind(uuid string) (configexample string, err error) {
	_, err = g.askGandi(mGET, "axfr/tsig/"+uuid+"/config/bind", nil, &configexample)
	return
}

// ListTsigPowerdns shows a PowerDNS nameserver config, and includes the nameservers available for zone transfers
func (g *Gandi) ListTsigPowerdns(uuid string) (configexample string, err error) {
	_, err = g.askGandi(mGET, "axfr/tsig/"+uuid+"/config/powerdns", nil, &configexample)
	return
}

// ListTsigNsd shows a NSD nameserver config, and includes the nameservers available for zone transfers
func (g *Gandi) ListTsigNsd(uuid string) (configexample string, err error) {
	_, err = g.askGandi(mGET, "axfr/tsig/"+uuid+"/config/nsd", nil, &configexample)
	return
}

// ListTsigKnot shows a KNOT nameserver config, and includes the nameservers available for zone transfers
func (g *Gandi) ListTsigKnot(uuid string) (configexample string, err error) {
	_, err = g.askGandi(mGET, "axfr/tsig/"+uuid+"/config/knot", nil, &configexample)
	return
}

// CreateTsig creates a tsig
func (g *Gandi) CreateTsig() (tsig Tsig, err error) {
	_, err = g.askGandi(mPOST, "axfr/tsig", nil, &tsig)
	return
}

// AddTsigToDomain adds a tsig to a domain
func (g *Gandi) AddTsigToDomain(fqdn, uuid string) (err error) {
	_, err = g.askGandi(mPUT, "domains/"+fqdn+"/axfr/tsig/"+uuid, nil, nil)
	return
}

// AddSlaveToDomain adds a slave to a domain
func (g *Gandi) AddSlaveToDomain(fqdn, host string) (err error) {
	_, err = g.askGandi(mPUT, "domains/"+fqdn+"/axfr/slaves/"+host, nil, nil)
	return
}

// ListSlavesInDomain lists slaves in a domain
func (g *Gandi) ListSlavesInDomain(fqdn string) (slaves []string, err error) {
	_, err = g.askGandi(mGET, "domains/"+fqdn+"/axfr/slaves", nil, &slaves)
	return
}

// DelSlaveFromDomain removes a slave from a domain
func (g *Gandi) DelSlaveFromDomain(fqdn, host string) (err error) {
	_, err = g.askGandi(mDELETE, "domains/"+fqdn+"/axfr/slaves/"+host, nil, nil)
	return
}
