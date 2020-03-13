package main

type domainCmd struct {
	List   domainListCmd   `kong:"cmd,help='List managed domains'"`
	Manage domainManageCmd `kong:"cmd,help='Manage a domain'"`
}

type domainListCmd struct{}

func (cmd *domainListCmd) Run(g *globals) error {
	d := g.domainHandle
	return jsonPrint(d.ListDomains())
}

type domainManageCmd struct {
	Name struct {
		Name        string                `kong:"arg"`
		Display     domainDisplayCmd      `kong:"cmd,help='Display the domain'"`
		NameServers domainDisplayNSCmd    `kong:"cmd,name='nameservers',help='Display the Name Servers for the domain'"`
		AutoRenew   domainSetAutoRenewCmd `kong:"cmd,name='autorenew',help='Enable or disable autorenew for the domain'"`
	} `kong:"arg"`
}

type domainDisplayCmd struct{}

func (cmd *domainDisplayCmd) Run(g *globals) error {
	fqdn := c.Domain.Manage.Name.Name
	d := g.domainHandle
	return jsonPrint(d.GetDomain(fqdn))
}

type domainDisplayNSCmd struct{}

func (cmd *domainDisplayNSCmd) Run(g *globals) error {
	fqdn := c.Domain.Manage.Name.Name
	d := g.domainHandle
	return jsonPrint(d.GetNameServers(fqdn))
}

type domainSetAutoRenewCmd struct {
	Enable bool `kong:"type='bool',help='Should the domain autorenew'"`
}

func (cmd *domainSetAutoRenewCmd) Run(g *globals) error {
	fqdn := c.Domain.Manage.Name.Name
	d := g.domainHandle
	return noPrint(d.SetAutoRenew(fqdn, cmd.Enable))
}
