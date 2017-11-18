package main

import g "github.com/tiramiseb/go-gandi-livedns"

func domain() {
	switch *action {
	case aList:
		jsonPrint(g.ListDomains(*apiKey))
	case aAdd:
		addDomainToZone()
	case aGet:
		getDomain()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      "list",
				Description: "List all domains",
			},
			actionDescription{
				Action:      "add",
				Description: "Add a domains in a zone",
			},
			/*
				actionDescription{
					Action:      "create",
					Description: "Create a domain",
				},
			*/
			actionDescription{
				Action:      "get",
				Description: "Get a domain",
			},
			/*
				actionDescription{
					Action:      "update",
					Description: "Update a domain",
				},
				actionDescription{
					Action:      "delete",
					Description: "Delete a domain",
				},
			*/
			/*
				actionDescription{
					Action:      "domains",
					Description: "List domains attached to a domain",
				},
				actionDescription{
					Action:      "attach",
					Description: "Attach a domain to a domain",
				},
			*/
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
	jsonPrint(g.AddDomainToZone(*apiKey, (*args)[0], (*args)[1]))
}

/*
func createDomain() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"Name of the domain to be created",
		})
		return
	}
	//jsonPrint(g.CreateDomain(*apiKey, (*args)[0]))
}
*/
func getDomain() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"FQDN of the domain to be returned",
		})
		return
	}
	jsonPrint(g.GetDomain(*apiKey, (*args)[0]))
}

func changeAssociatedZone() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"FQDN of the domain to be added",
			"UUID of the zone where to move the domain",
		})
		return
	}
	jsonPrint(g.ChangeAssociatedZone(*apiKey, (*args)[0], (*args)[1]))
}

/*
func updateDomain() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"UUID of the domain to be updated",
			"New name of the domain",
		})
		return
	}
	//jsonPrint(g.UpdateDomain(*apiKey, (*args)[0], (*args)[1]))
}

func deleteDomain() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the domain to be deleted",
		})
		return
	}
	//noPrint(g.DeleteDomain(*apiKey, (*args)[0]))
}

func getDomainDomains() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the domain to be checked",
		})
		return
	}
	jsonPrint(g.GetDomainDomains(*apiKey, (*args)[0]))
}

func attachDomainToDomain() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"UUID of the domain to attach the domain to",
			"FQDN of the domain to be attached",
		})
		return
	}
	jsonPrint(g.AttachDomainToDomain(*apiKey, (*args)[0], (*args)[1]))
}
*/
