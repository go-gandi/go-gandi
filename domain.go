package gandi

// Domain represents a DNS domain
type Domain struct {
	FQDN              string `json:"fqdn"`
	DomainHref        string `json:"domain_href,omitempty"`
	DomainRecordsHref string `json:"domain_records_href,omitempty"`
}
