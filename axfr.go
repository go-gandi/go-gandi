package gandi

// Tsig contains tsig data (no kidding!)
type Tsig struct {
	KeyName     string `json:"key_name, omitempty"`
	Secret      string `json:"secret,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	AxfrTsigURL string `json:"axfr_tsig_url,omitempty"`
}

// ListTsigs lists all tsigs
func ListTsigs(key string) (tsigs []Tsig, err error) {
	err = askGandi(key, mGET, "axfr/tsig", nil, &tsigs)
	return
}

// CreateTsig creates a tsig
func CreateTsig(key string) (tsig Tsig, err error) {
	err = askGandi(key, mPOST, "axfr/tsig", nil, &tsig)
	return
}

// AddTsigToDomain adds a tsig to a domain
func AddTsigToDomain(key, fqdn, uuid string) (err error) {
	err = askGandi(key, mPUT, "domains/"+fqdn+"/axfr/tsig/"+uuid, nil, nil)
	return
}

// AddSlaveToDomain adds a slave to a domain
func AddSlaveToDomain(key, fqdn, host string) (err error) {
	err = askGandi(key, mPUT, "domains/"+fqdn+"/axfr/slaves/"+host, nil, nil)
	return
}
