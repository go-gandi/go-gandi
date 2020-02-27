package main

type liveDNSGetTSIGsCmd struct {
	Format string `kong:"arg,optional,enum='bind,nsd,powerdns,knot,none',default='none',help='The format of the TSIG config (bind, nsd, powerdns, knot)'"`
}

func (d *liveDNSGetTSIGsCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	switch d.Format {
	case "bind":
		return jsonPrint(l.GetTsigBIND(fqdn))
	case "nsd":
		return jsonPrint(l.GetTsigNSD(fqdn))
	case "powerdns":
		return jsonPrint(l.GetTsigPowerDNS(fqdn))
	case "knot":
		return jsonPrint(l.GetTsigKnot(fqdn))
	default:
		return jsonPrint(l.GetTSIGKey(fqdn))
	}
}

type liveDNSAddTSIGToDomainCmd struct {
	UUID string `kong:"arg,help='The UUID of the TSIG to add'"`
}

func (d *liveDNSAddTSIGToDomainCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return noPrint(l.AddTsigToDomain(fqdn, d.UUID))
}

type liveDNSRemoveTSIGFromDomainCmd struct {
	UUID string `kong:"arg,help='The UUID of the TSIG to add'"`
}

func (d *liveDNSRemoveTSIGFromDomainCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return noPrint(l.RemoveTSIGKeyFromDomain(fqdn, d.UUID))
}

type liveDNSAddAXFRSlaveCmd struct {
	IP string `kong:"arg,help='The IP of the AXFR requestor'"`
}

func (d *liveDNSAddAXFRSlaveCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return noPrint(l.AddSlaveToDomain(fqdn, d.IP))
}

type liveDNSListAXFRSlavesCmd struct{}

func (d *liveDNSListAXFRSlavesCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.ListSlavesInDomain(fqdn))
}

type liveDNSRemoveAXFRSlaveCmd struct {
	IP string `kong:"arg,help='The IP of the AXFR requestor'"`
}

func (d *liveDNSRemoveAXFRSlaveCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return noPrint(l.DelSlaveFromDomain(fqdn, d.IP))
}
