package main

type liveDNSCmd struct {
	Create liveDNSCreateCmd `kong:"cmd,help='Enable LiveDNS for a Domain'"`
	List   liveDNSListCmd   `kong:"cmd,help='List LiveDNS Domains'"`
	Manage liveDNSManageCmd `kong:"cmd,help='Manage LiveDNS Domain'"`
}

type liveDNSCreateCmd struct {
	FQDN string `kong:"arg"`
	TTL  int    `kong:"arg,optional,default='60',help='The default TTL of the domain, default= 60'"`
}

func (d *liveDNSCreateCmd) Run(g *globals) error {
	l := g.liveDNSHandle
	return jsonPrint(l.CreateDomain(d.FQDN, d.TTL))
}

type liveDNSListCmd struct{}

func (c *liveDNSListCmd) Run(g *globals) error {
	l := g.liveDNSHandle
	return jsonPrint(l.ListDomains())
}

type liveDNSManageCmd struct {
	Name struct {
		Name    string              `kong:"arg"`
		Display liveDNSGetDomainCmd `kong:"cmd,help='Display the domain'"`
		Records struct {
			List   liveDNSGetRecordsCmd   `kong:"cmd,name='list',help='Display records for domain'"`
			Create liveDNSCreateRecordCmd `kong:"cmd,name='create',help='Create records for domain'"`
			Update liveDNSUpdateRecordCmd `kong:"cmd,name='update',help='Update records for domain'"`
			Delete liveDNSDeleteRecordCmd `kong:"cmd,name='delete',help='Delete records for domain'"`
		} `kong:"cmd"`
		Sign          liveDNSSignDomainCmd           `kong:"cmd,help='Sign the domain'"`
		Keys          liveDNSGetDomainKeysCmd        `kong:"cmd,help='Get the DNSSEC keys for the domain'"`
		DeleteKey     liveDNSDeleteDomainKeyCmd      `kong:"cmd,help='Delete a DNSSEC key for the domain'"`
		NameServers   liveDNSGetDomainNSCmd          `kong:"cmd,help='Get nameservers for the domain'"`
		Snapshot      liveDNSCreateSnapshotCmd       `kong:"cmd,help='Create a snapshot of the domain'"`
		GetSnapshot   liveDNSGetSnapshotCmd          `kong:"cmd,name='get-snapshot',help='Get a snapshot of the domain'"`
		ListSnapshots liveDNSListSnapshotsCmd        `kong:"cmd,name='list-snapshots',help='List snapshots of the domain'"`
		GetTsigs      liveDNSGetTSIGsCmd             `kong:"cmd,name='get-tsigs',help='Get TSIGs'"`
		AddTSIG       liveDNSAddTSIGToDomainCmd      `kong:"cmd,name='add-tsig',help='Add TSIG to domain'"`
		RemoveTSIG    liveDNSRemoveTSIGFromDomainCmd `kong:"cmd,name='remove-tsig',help='Remove TSIG from domain'"`
		GetAXFRs      liveDNSListAXFRSlavesCmd       `kong:"cmd,name='get-axfrs',help='Get AXFRs'"`
		AddAXFR       liveDNSAddAXFRSlaveCmd         `kong:"cmd,name='add-axfr',help='Add AXFR to domain'"`
		RemoveAXFR    liveDNSRemoveAXFRSlaveCmd      `kong:"cmd,name='remove-axfr',help='Remove AXFR from domain'"`
	} `kong:"arg"`
}

type liveDNSGetDomainCmd struct{}

func (d *liveDNSGetDomainCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.GetDomain(fqdn))
}

type liveDNSSignDomainCmd struct{}

func (d *liveDNSSignDomainCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.SignDomain(fqdn))
}

type liveDNSGetDomainKeysCmd struct{}

func (d *liveDNSGetDomainKeysCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.GetDomainKeys(fqdn))
}

type liveDNSDeleteDomainKeyCmd struct {
	ID string `kong:"arg,help='The ID of the key to delete'"`
}

func (d *liveDNSDeleteDomainKeyCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return noPrint(l.DeleteDomainKey(fqdn, d.ID))
}

type liveDNSGetDomainNSCmd struct{}

func (d *liveDNSGetDomainNSCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.GetDomainNS(fqdn))
}
