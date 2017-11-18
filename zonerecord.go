package gandi

// ZoneRecord represents a DNS Record
type ZoneRecord struct {
	RrsetType   string   `json:"rrset_type,omitempty"`
	RrsetTTL    int      `json:"rrset_ttl,omitempty"`
	RrsetName   string   `json:"rrset_name,omitempty"`
	RrsetHref   string   `json:"rrset_href,omitempty"`
	RrsetValues []string `json:"rrset_values,omitempty"`
}

// ListZoneRecords lists all records in a zone
func ListZoneRecords(key, uuid string) (records []ZoneRecord, err error) {
	err = askGandi(key, mGET, "zones/"+uuid+"/records", nil, &records)
	return
}

// ListDomainRecords lists all records in the zone associated with a domain
func ListDomainRecords(key, fqdn string) (records []ZoneRecord, err error) {
	err = askGandi(key, mGET, "domains/"+fqdn+"/records", nil, &records)
	return
}

// ListZoneRecordsAsText lists all records in a zone and returns them as a text file
// ... and by text, I mean a slice of bytes
func ListZoneRecordsAsText(key, uuid string) ([]byte, error) {
	return askGandiToBytes(key, mGET, "zones/"+uuid+"/records", nil)
}

// ListZoneRecordsWithName lists all records with a specific name in a zone
func ListZoneRecordsWithName(key, uuid, name string) (records []ZoneRecord, err error) {
	err = askGandi(key, mGET, "zones/"+uuid+"/records/"+name, nil, &records)
	return
}

// ListDomainRecordsWithName lists all records with a specific name in a zone
func ListDomainRecordsWithName(key, fqdn, name string) (records []ZoneRecord, err error) {
	err = askGandi(key, mGET, "domains/"+fqdn+"/records/"+name, nil, &records)
	return
}

// GetZoneRecordWithNameAndType gets the record with specific name and type in a zone
func GetZoneRecordWithNameAndType(key, uuid, name, recordtype string) (record ZoneRecord, err error) {
	err = askGandi(key, mGET, "zones/"+uuid+"/records/"+name+"/"+recordtype, nil, &record)
	return
}

// GetDomainRecordWithNameAndType gets the record with specific name and type in the zone attached to the domain
func GetDomainRecordWithNameAndType(key, fqdn, name, recordtype string) (record ZoneRecord, err error) {
	err = askGandi(key, mGET, "domains/"+fqdn+"/records/"+name+"/"+recordtype, nil, &record)
	return
}

// CreateZoneRecord creates a record in a zone
func CreateZoneRecord(key, uuid, name, recordtype string, ttl int, values []string) (response StandardResponse, err error) {
	err = askGandi(key, mPOST, "zones/"+uuid+"/records",
		ZoneRecord{
			RrsetType:   recordtype,
			RrsetTTL:    ttl,
			RrsetName:   name,
			RrsetValues: values,
		},
		&response)
	return
}

// CreateDomainRecord creates a record in the zone attached to a domain
func CreateDomainRecord(key, fqdn, name, recordtype string, ttl int, values []string) (response StandardResponse, err error) {
	err = askGandi(key, mPOST, "domains/"+fqdn+"/records",
		ZoneRecord{
			RrsetType:   recordtype,
			RrsetTTL:    ttl,
			RrsetName:   name,
			RrsetValues: values,
		},
		&response)
	return
}

type itemsPrefixForZoneRecords struct {
	Items []ZoneRecord `json:"items"`
}

// ChangeZoneRecords changes all records in a zone
func ChangeZoneRecords(key, uuid string, records []ZoneRecord) (response StandardResponse, err error) {
	prefixedRecords := itemsPrefixForZoneRecords{Items: records}
	err = askGandi(key, mPUT, "zones/"+uuid+"/records", prefixedRecords, &response)
	return
}

// ChangeDomainRecords changes all records in the zone attached to a domain
func ChangeDomainRecords(key, fqdn string, records []ZoneRecord) (response StandardResponse, err error) {
	prefixedRecords := itemsPrefixForZoneRecords{Items: records}
	err = askGandi(key, mPUT, "domains/"+fqdn+"/records", prefixedRecords, &response)
	return
}

// ChangeZoneRecordsAsText changes all zone records, taking them as text
// ... and by text, I mean a slice of bytes
func ChangeZoneRecordsAsText(key, uuid string, records []byte) (response StandardResponse, err error) {
	err = askGandiFromBytes(key, mPUT, "zones/"+uuid+"/records", records, &response)
	return
}

// ChangeZoneRecordsWithName changes all zone records with the given name
func ChangeZoneRecordsWithName(key, uuid, name string, records []ZoneRecord) (response StandardResponse, err error) {
	prefixedRecords := itemsPrefixForZoneRecords{Items: records}
	err = askGandi(key, mPUT, "zones/"+uuid+"/records/"+name, prefixedRecords, &response)
	return
}

// ChangeDomainRecordsWithName changes all records with the given name in the zone attached to the domain
func ChangeDomainRecordsWithName(key, fqdn, name string, records []ZoneRecord) (response StandardResponse, err error) {
	prefixedRecords := itemsPrefixForZoneRecords{Items: records}
	err = askGandi(key, mPUT, "domains/"+fqdn+"/records/"+name, prefixedRecords, &response)
	return
}

// ChangeZoneRecordWithNameAndType changes the zone record with the given name and the given type
func ChangeZoneRecordWithNameAndType(key, uuid, name, recordtype string, ttl int, values []string) (response StandardResponse, err error) {
	err = askGandi(key, mPUT, "zones/"+uuid+"/records/"+name+"/"+recordtype,
		ZoneRecord{
			RrsetTTL:    ttl,
			RrsetValues: values,
		},
		&response)
	return
}

// ChangeDomainRecordWithNameAndType changes the record with the given name and the given type in the zone attached to a domain
func ChangeDomainRecordWithNameAndType(key, fqdn, name, recordtype string, ttl int, values []string) (response StandardResponse, err error) {
	err = askGandi(key, mPUT, "domains/"+fqdn+"/records/"+name+"/"+recordtype,
		ZoneRecord{
			RrsetTTL:    ttl,
			RrsetValues: values,
		},
		&response)
	return
}

// DeleteAllZoneRecords deletes all records in a zone
func DeleteAllZoneRecords(key, uuid string) (err error) {
	err = askGandi(key, mDELETE, "zones/"+uuid+"/records", nil, nil)
	return
}

// DeleteAllDomainRecords deletes all records in the zone attached to a domain
func DeleteAllDomainRecords(key, fqdn string) (err error) {
	err = askGandi(key, mDELETE, "domains/"+fqdn+"/records", nil, nil)
	return
}

// DeleteZoneRecords deletes all records with the given name in a zone
func DeleteZoneRecords(key, uuid, name string) (err error) {
	err = askGandi(key, mDELETE, "zones/"+uuid+"/records/"+name, nil, nil)
	return
}

// DeleteDomainRecords deletes all records with the given name in the zone attached to a domain
func DeleteDomainRecords(key, fqdn, name string) (err error) {
	err = askGandi(key, mDELETE, "domains/"+fqdn+"/records/"+name, nil, nil)
	return
}

// DeleteZoneRecord deletes the record with the given name and the given type in a zone
func DeleteZoneRecord(key, uuid, name, recordtype string) (err error) {
	err = askGandi(key, mDELETE, "zones/"+uuid+"/records/"+name+"/"+recordtype, nil, nil)
	return
}

// DeleteDomainRecord deletes the record with the given name and the given type in the zone attached to a domain
func DeleteDomainRecord(key, fqdn, name, recordtype string) (err error) {
	err = askGandi(key, mDELETE, "domains/"+fqdn+"/records/"+name+"/"+recordtype, nil, nil)
	return
}
