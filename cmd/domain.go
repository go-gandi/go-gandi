package main

func domain() {
	switch *action {
	case aList:
		jsonPrint(g.ListDomains())
	case aAdd:
		addDomainToZone()
	case aGet:
		getDomain()
	case aSign:
		signDomain()
	case aKeys:
		getDomainKeys()
	case aDelKey:
		deleteDomainKey()
	case aNS:
		getDomainNS()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      aList,
				Description: "List all domains",
			},
			actionDescription{
				Action:      aAdd,
				Description: "Add a domains in a zone",
			},
			actionDescription{
				Action:      aGet,
				Description: "Get a domain",
			},
			actionDescription{
				Action:      aSign,
				Description: "Ask the Gandi server to sign a domain",
			},
			actionDescription{
				Action:      aKeys,
				Description: "Return the signing keys created for domain",
			},
			actionDescription{
				Action:      aDelKey,
				Description: "Delete a signing key",
			},
			actionDescription{
				Action:      aNS,
				Description: "Get nameservers for a domain",
			},
		})
	}
}

func addDomainToZone() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"FQDN of the domain to be added",
			"UUID of the zone where to add the domain",
		})
		return
	}
	jsonPrint(g.AddDomainToZone((*args)[0], (*args)[1]))
}

func getDomain() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain to be returned",
		})
		return
	}
	jsonPrint(g.GetDomain((*args)[0]))
}

func changeAssociatedZone() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"FQDN of the domain to be added",
			"UUID of the zone where to move the domain",
		})
		return
	}
	jsonPrint(g.ChangeAssociatedZone((*args)[0], (*args)[1]))
}

func signDomain() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain to be signed",
		})
		return
	}
	jsonPrint(g.SignDomain((*args)[0]))
}

func getDomainKeys() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain for which to return keys",
		})
		return
	}
	jsonPrint(g.GetDomainKeys((*args)[0]))
}

func deleteDomainKey() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"FQDN of the domain for which to delete a key",
			"UUID of the key to delete",
		})
		return
	}
	noPrint(g.DeleteDomainKey((*args)[0], (*args)[1]))
}

func getDomainNS() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain for which to return the nameservers",
		})
		return
	}
	jsonPrint(g.GetDomainNS((*args)[0]))
}
