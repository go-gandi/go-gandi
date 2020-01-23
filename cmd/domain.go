package main

func domain_type() {
	switch *resourceType {
	case "list":
		domain_list()
	case "get":
		domain_print()
	case "nameservers":
		nameservers_list()
	}
}

func domain_list() {
	jsonPrint(d.ListDomains())
}

func domain_print() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain to get info for",
		})
		return
	}
	jsonPrint(d.GetDomain((*args)[0]))
}

func nameservers_list() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain for which to return the nameservers",
		})
		return
	}
	jsonPrint(d.GetNameServers((*args)[0]))
}
