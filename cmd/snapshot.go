package main

type liveDNSListSnapshotsCmd struct{}

func (d *liveDNSListSnapshotsCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.ListSnapshots(fqdn))
}

type liveDNSCreateSnapshotCmd struct{}

func (d *liveDNSCreateSnapshotCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.CreateSnapshot(fqdn))
}

type liveDNSGetSnapshotCmd struct {
	ID string `kong:"arg,help='The ID of the snapshot'"`
}

func (d *liveDNSGetSnapshotCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.GetSnapshot(fqdn, d.ID))
}
