package gandi

// Zone represents a DNS Zone
type Zone struct {
	Retry           int    `json:"retry,omitempty"`
	UUID            string `json:"uuid,omitempty"`
	ZoneHref        string `json:"zone_href,omitempty"`
	Minimum         int    `json:"minimum,omitempty"`
	DomainsHref     string `json:"domains_href,omitempty"`
	Refresh         int    `json:"refresh,omitempty"`
	ZoneRecordsHref string `json:"zone_records_href,omitempty"`
	Expire          int    `json:"expire,omitempty"`
	SharingID       string `json:"sharing_id,omitempty"`
	Serial          int    `json:"serial,omitempty"`
	Email           string `json:"email,omitempty"`
	PrimaryNS       string `json:"primary_ns,omitempty"`
	Name            string `json:"name"`
}

// ListZones lists all zones
func ListZones(key string) (zones []Zone, err error) {
	err = askGandi(key, mGET, "zones", nil, &zones)
	return
}

// CreateZone creates a zone
func CreateZone(key, name string) (response StandardResponse, err error) {
	err = askGandi(key, mPOST, "zones", Zone{Name: name}, &response)
	return
}

// GetZone returns a zone
func GetZone(key, uuid string) (zone Zone, err error) {
	err = askGandi(key, mGET, "zones/"+uuid, nil, &zone)
	return
}

// UpdateZone updates a zone
func UpdateZone(key, uuid, name string) (response StandardResponse, err error) {
	err = askGandi(key, mPATCH, "zones/"+uuid, Zone{Name: name}, &response)
	return
}

// DeleteZone deletes a zone
func DeleteZone(key, uuid string) (err error) {
	err = askGandi(key, mDELETE, "zones/"+uuid, nil, nil)
	return
}

// GetZoneDomains returns domains attached to a zone
func GetZoneDomains(key, uuid string) (domains []Domain, err error) {
	err = askGandi(key, mGET, "zones/"+uuid+"/domains", nil, &domains)
	return
}

// AttachDomainToZone attaches a domain to a zone
func AttachDomainToZone(key, uuid, fqdn string) (response StandardResponse, err error) {
	err = askGandi(key, mPOST, "zones/"+uuid+"/domains/"+fqdn, nil, &response)
	return
}
