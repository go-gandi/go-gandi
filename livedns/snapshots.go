package livedns

import (
	"github.com/go-gandi/go-gandi/types"
)

// ListSnapshots lists all snapshots for a domain
func (g *LiveDNS) ListSnapshots(fqdn string) (snapshots []Snapshot, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/snapshots", nil, &snapshots)
	return
}

// CreateSnapshot creates a snapshot for a domain
func (g *LiveDNS) CreateSnapshot(fqdn string) (response types.StandardResponse, err error) {
	_, err = g.client.Post("domains/"+fqdn+"/snapshots", nil, &response)
	return
}

// GetSnapshot returns a snapshot for a domain
func (g *LiveDNS) GetSnapshot(fqdn, snapUUID string) (snapshot Snapshot, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/snapshots/"+snapUUID, nil, &snapshot)
	return
}

// DeleteSnapshot deletes a snapshot for a domain
func (g *LiveDNS) DeleteSnapshot(fqdn, snapUUID string) (err error) {
	_, err = g.client.Delete("domains/"+fqdn+"/snapshots/"+snapUUID, nil, nil)
	return
}
