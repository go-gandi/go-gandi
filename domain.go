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
