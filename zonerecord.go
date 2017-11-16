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

// GetZoneRecordWithNameAndType gets the record with specific name and type in a zone
func GetZoneRecordWithNameAndType(key, uuid, name, recordtype string) (record ZoneRecord, err error) {
	err = askGandi(key, mGET, "zones/"+uuid+"/records/"+name+"/"+recordtype, nil, &record)
	return
}

// CreateZoneRecord creates a zone
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

type itemsPrefixForZoneRecords struct {
	Items []ZoneRecord `json:"items"`
}

// ChangeZoneRecords changes all zone records
func ChangeZoneRecords(key, uuid string, records []ZoneRecord) (response StandardResponse, err error) {
	prefixedRecords := itemsPrefixForZoneRecords{Items: records}
	err = askGandi(key, mPUT, "zones/"+uuid+"/records", prefixedRecords, &response)
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

// DeleteAllRecords deletes all records in a zone
func DeleteAllRecords(key, uuid string) (err error) {
	err = askGandi(key, mDELETE, "zones/"+uuid+"/records", nil, nil)
	return
}

// DeleteRecords deletes all records with the given name in a zone
func DeleteRecords(key, uuid, name string) (err error) {
	err = askGandi(key, mDELETE, "zones/"+uuid+"/records/"+name, nil, nil)
	return
}

// DeleteRecord deletes the record with the given name and the given type in a zone
func DeleteRecord(key, uuid, name, recordtype string) (err error) {
	err = askGandi(key, mDELETE, "zones/"+uuid+"/records/"+name+"/"+recordtype, nil, nil)
	return
}
