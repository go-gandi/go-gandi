package main

type liveDNSGetRecordsCmd struct {
	Name string `kong:"arg,optional,name='name',help='The name of the record to fetch'"`
	Type string `kong:"arg,optional,help='The type of the record to retrieve. You must specify the name too.'"`
}

func (d *liveDNSGetRecordsCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	if d.Name != "" && d.Type != "" {
		return jsonPrint(l.GetDomainRecordByNameAndType(fqdn, d.Name, d.Type))
	} else if d.Name != "" {
		return jsonPrint(l.GetDomainRecordsByName(fqdn, d.Name))
	} else {
		return jsonPrint(l.GetDomainRecords(fqdn))
	}
}

type liveDNSCreateRecordCmd struct {
	Name   string   `kong:"arg,help='The name of the record to fetch'"`
	Type   string   `kong:"arg,help='The type of the record to retrieve. You must specify the name too.'"`
	TTL    int      `kong:"arg,help='The TTL of the record'"`
	Values []string `kong:"arg,help='The values of the record'"`
}

func (d *liveDNSCreateRecordCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.CreateDomainRecord(fqdn, d.Name, d.Type, d.TTL, d.Values))
}

type liveDNSUpdateRecordCmd struct {
	Name   string   `kong:"arg,help='The name of the record to fetch'"`
	Type   string   `kong:"arg,help='The type of the record to retrieve. You must specify the name too.'"`
	TTL    int      `kong:"arg,help='The TTL of the record'"`
	Values []string `kong:"arg,help='The values of the record'"`
}

func (d *liveDNSUpdateRecordCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	return jsonPrint(l.UpdateDomainRecordByNameAndType(fqdn, d.Name, d.Type, d.TTL, d.Values))
}

type liveDNSDeleteRecordCmd struct {
	Name string `kong:"arg,optional,name='name',help='The name of the record to fetch'"`
	Type string `kong:"arg,optional,help='The type of the record to retrieve. You must specify the name too.'"`
}

func (d *liveDNSDeleteRecordCmd) Run(g *globals) error {
	fqdn := c.LiveDNS.Manage.Name.Name
	l := g.liveDNSHandle
	if d.Name != "" && d.Type != "" {
		return noPrint(l.DeleteDomainRecord(fqdn, d.Name, d.Type))
	} else if d.Name != "" {
		return noPrint(l.DeleteDomainRecordsByName(fqdn, d.Name))
	} else {
		return noPrint(l.DeleteAllDomainRecords(fqdn))
	}
}
