package gandi

// Domain represents a DNS domain
type Domain struct {
	FQDN              string `json:"fqdn,omitempty"`
	DomainHref        string `json:"domain_href,omitempty"`
	DomainKeysHref    string `json:"domain_keys_href,omitempty"`
	DomainRecordsHref string `json:"domain_records_href,omitempty"`
	ZoneUUID          string `json:"zone_uuid,omitempty"`
	ZoneHref          string `json:"zone_href,omitempty"`
	ZoneRecordsHref   string `json:"zone_records_href,omitempty"`
}

// SigningKey holds data about a DNSSEC signing key
type SigningKey struct {
	Status        string `json:"status,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Algorithm     int    `json:"algorithm,omitempty"`
	Deleted       *bool  `json:"deleted"`
	AlgorithmName string `json:"algorithm_name,omitempty"`
	FQDN          string `json:"fqdn,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	DS            string `json:"ds,omitempty"`
	KeyHref       string `json:"key_href,omitempty"`
}

// ListDomains lists all domains
func ListDomains(key string) (domains []Domain, err error) {
	err = askGandi(key, mGET, "domains", nil, &domains)
	return
}

// AddDomainToZone adds a domain to a zone
// It is equivalent to AttachDomainToZone, the only difference is the entry point in the LiveDNS API.
func AddDomainToZone(key, fqdn, uuid string) (response StandardResponse, err error) {
	err = askGandi(key, mPOST, "domains", Domain{FQDN: fqdn, ZoneUUID: uuid}, &response)
	return
}

// GetDomain returns a domain
func GetDomain(key, fqdn string) (domain Domain, err error) {
	err = askGandi(key, mGET, "domains/"+fqdn, nil, &domain)
	return
}

// ChangeAssociatedZone changes the zone associated to a domain
func ChangeAssociatedZone(key, fqdn, uuid string) (response StandardResponse, err error) {
	err = askGandi(key, mPATCH, "domains/"+fqdn, Domain{ZoneUUID: uuid}, &response)
	return
}

// SignDomain creates a DNSKEY and asks Gandi servers to automatically sign the domain
func SignDomain(key, fqdn string) (response StandardResponse, err error) {
	f := SigningKey{Flags: 257}
	err = askGandi(key, mPOST, "domains/"+fqdn+"/keys", f, &response)
	return
}

// GetDomainKeys returns data about the signing keys created for a domain
func GetDomainKeys(key, fqdn string) (keys []SigningKey, err error) {
	err = askGandi(key, mGET, "domains/"+fqdn+"/keys", nil, &keys)
	return
}

// DeleteDomainKey deletes a signing key from a domain
func DeleteDomainKey(key, fqdn, uuid string) (err error) {
	err = askGandi(key, mDELETE, "domains/"+fqdn+"/keys/"+uuid, nil, nil)
	return
}
