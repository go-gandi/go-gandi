package main

func axfr() {
	switch *action {
	case aList:
		jsonPrint(g.ListTsigs())
	case aCreate:
		jsonPrint(g.CreateTsig())
	case aAdd:
		addTsigToDomain()
	case aSlave:
		addSlaveToDomain()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      aList,
				Description: "List all tsigs",
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
		})
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
			"FQDN of the domain where to add the tsig",
			"IP address of the slave to add",
		})
		return
	}
	noPrint(g.AddSlaveToDomain((*args)[0], (*args)[1]))
}
