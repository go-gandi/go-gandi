package main

import g "github.com/tiramiseb/go-gandi-livedns"

func zone() {
	switch *action {
	case "list":
		jsonPrint(g.ListZones(*apiKey))
	case "create":
		createZone()
	case "get":
		getZone()
	case "update":
		updateZone()
	case "delete":
		deleteZone()
	case "domains":
		getZoneDomains()
	case "attach":
		attachDomainToZone()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      "list",
				Description: "List all zones",
			},
		})
	}
}

func createZone() {
	if len(*args) == 0 {
		displayArgsList([]string{
			"Name of the zone to be created",
		})
		return
	}
	jsonPrint(g.CreateZone(*apiKey, (*args)[0]))
}

func getZone() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone to be returned",
		})
		return
	}
	jsonPrint(g.GetZone(*apiKey, (*args)[0]))
}

func updateZone() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"UUID of the zone to be updated",
			"New name of the zone",
		})
		return
	}
	jsonPrint(g.UpdateZone(*apiKey, (*args)[0], (*args)[1]))
}

func deleteZone() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone to be deleted",
		})
		return
	}
	noPrint(g.DeleteZone(*apiKey, (*args)[0]))
}

func getZoneDomains() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone to be checked",
		})
		return
	}
	jsonPrint(g.GetZoneDomains(*apiKey, (*args)[0]))
}

func attachDomainToZone() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"UUID of the zone to attach the domain to",
			"FQDN of the domain to be attached",
		})
		return
	}
	jsonPrint(g.AttachDomainToZone(*apiKey, (*args)[0], (*args)[1]))
}
