package gandi

// Snapshot represents a zone snapshot
type Snapshot struct {
	UUID        string       `json:"uuid,omitempty"`
	DateCreated string       `json:"date_created,omitempty"`
	ZoneUUID    string       `json:"zone_uuid,omitempty"`
	ZoneData    []ZoneRecord `json:"zone_data,omitempty"`
}

// ListSnapshots lists all zones
func ListSnapshots(key, uuid string) (snapshots []Snapshot, err error) {
	err = askGandi(key, mGET, "zones/"+uuid+"/snapshots", nil, &snapshots)
	return
}

// CreateSnapshot creates a zone
func CreateSnapshot(key, uuid string) (response StandardResponse, err error) {
	err = askGandi(key, mPOST, "zones/"+uuid+"/snapshots", nil, &response)
	return
}

// GetSnapshot returns a zone
func GetSnapshot(key, uuid, snapUUID string) (snapshot Snapshot, err error) {
	err = askGandi(key, mGET, "zones/"+uuid+"/snapshots/"+snapUUID, nil, &snapshot)
	return
}
