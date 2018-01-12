package main

import "fmt"

func axfr() {
	switch *action {
	case aList:
		jsonPrint(g.ListTsigs())
	case aListBind:
		listTsigBind()
	case aListNsd:
		listTsigNsd()
	case aListPowerdns:
		listTsigPowerdns()
	case aListKnot:
		listTsigKnot()
	case aCreate:
		jsonPrint(g.CreateTsig())
	case aAdd:
		addTsigToDomain()
	case aSlave:
		addSlaveToDomain()
	case aSlaves:
		listSlavesInDomain()
	case aDelSlave:
		delSlaveFromDomain()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      aList,
				Description: "List all tsigs",
			},
			actionDescription{
				Action:      aListBind,
				Description: "List BIND config example for tsig (in text, no JSON output)",
			},
			actionDescription{
				Action:      aListNsd,
				Description: "List NSD config example for tsig (in text, no JSON output)",
			},
			actionDescription{
				Action:      aListBind,
				Description: "List PowerDNS config example for tsig (in text, no JSON output)",
			},
			actionDescription{
				Action:      aListBind,
				Description: "List KNOT config example for tsig (in text, no JSON output)",
			},
			actionDescription{
				Action:      aCreate,
				Description: "Create a tsig (cannot be deleted)",
			},
			actionDescription{
				Action:      aAdd,
				Description: "Add a tsig to a domain",
			},
			actionDescription{
				Action:      aSlave,
				Description: "Add a slave to a domain",
			},
			actionDescription{
				Action:      aSlaves,
				Description: "List slaves in a domain",
			},
			actionDescription{
				Action:      aDelSlave,
				Description: "Remove a slave from a domain",
			},
		})
	}
}

func listTsigBind() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	// api.gandi.net tsig config examples only return text no JSON
	if r, e := g.ListTsigBind((*args)[0]); e != nil {
		jsonPrint(r, e)
	} else {
		fmt.Println(r)
	}

}

func listTsigNsd() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	// api.gandi.net tsig config examples only return text no JSON
	if r, e := g.ListTsigNsd((*args)[0]); e != nil {
		jsonPrint(r, e)
	} else {
		fmt.Println(r)
	}
}

func listTsigPowerdns() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	// api.gandi.net tsig config examples only return text no JSON
	if r, e := g.ListTsigPowerdns((*args)[0]); e != nil {
		jsonPrint(r, e)
	} else {
		fmt.Println(r)
	}
}

func listTsigKnot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	// api.gandi.net tsig config examples only return text no JSON
	if r, e := g.ListTsigKnot((*args)[0]); e != nil {
		jsonPrint(r, e)
	} else {
		fmt.Println(r)
	}
}

func addTsigToDomain() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"FQDN of the domain where to add the tsig",
			"UUID of the tsig to add",
		})
		return
	}
	noPrint(g.AddTsigToDomain((*args)[0], (*args)[1]))
}

func addSlaveToDomain() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"FQDN of the domain where to add the slave",
			"IP address of the slave to add",
		})
		return
	}
	noPrint(g.AddSlaveToDomain((*args)[0], (*args)[1]))
}

func listSlavesInDomain() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain where list slaves",
		})
		return
	}
	jsonPrint(g.ListSlavesInDomain((*args)[0]))
}

func delSlaveFromDomain() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"FQDN of the domain where to remove the slave",
			"IP address of the slave to remove",
		})
		return
	}
	noPrint(g.DelSlaveFromDomain((*args)[0], (*args)[1]))
}
