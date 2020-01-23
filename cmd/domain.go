package main

func domainType() {
	switch *resourceType {
	case "list":
		domainList()
	case "get":
		domainPrint()
	case "nameservers":
		nameserversList()
	}
}

func domainList() {
	jsonPrint(d.ListDomains())
}

func domainPrint() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain to get info for",
		})
		return
	}
	jsonPrint(d.GetDomain((*args)[0]))
}

func nameserversList() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain for which to return the nameservers",
		})
		return
	}
	jsonPrint(d.GetNameServers((*args)[0]))
}
