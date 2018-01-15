package main

func axfr() {
	switch *action {
	case aList:
		jsonPrint(g.ListTsigs())
  case aGet:
    getTsig()
	case aGetBIND:
		getTsigBIND()
	case aGetNSD:
		getTsigNSD()
	case aGetPowerDNS:
		getTsigPowerDNS()
	case aGetKnot:
		getTsigKnot()
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
				Action:      aGet,
				Description: "Get tsig details",
			},
			actionDescription{
				Action:      aGetBIND,
				Description: "Get BIND config example for tsig",
			},
			actionDescription{
				Action:      aGetNSD,
				Description: "Get NSD config example for tsig",
			},
			actionDescription{
				Action:      aGetPowerDNS,
				Description: "Get PowerDNS config example for tsig",
			},
			actionDescription{
				Action:      aGetKnot,
				Description: "Get Knot config example for tsig",
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

func getTsig() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the tsig to get details from",
		})
		return
	}
	jsonPrint(g.GetTsig((*args)[0]))
}

func getTsigBIND() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	textPrint(g.GetTsigBIND((*args)[0]))
}

func getTsigNSD() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	textPrint(g.GetTsigNSD((*args)[0]))
}

func getTsigPowerDNS() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	textPrint(g.GetTsigPowerDNS((*args)[0]))
}

func getTsigKnot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the TSIG key",
		})
		return
	}
	textPrint(g.GetTsigKnot((*args)[0]))
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
