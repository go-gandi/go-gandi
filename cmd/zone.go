package main

func zone() {
	switch *action {
	case aList:
		jsonPrint(g.ListZones())
	case aCreate:
		createZone()
	case aGet:
		getZone()
	case aUpdate:
		updateZone()
	case aDelete:
		deleteZone()
	case aDomains:
		getZoneDomains()
	case aAttach:
		attachDomainToZone()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      aList,
				Description: "List all zones",
			},
			actionDescription{
				Action:      aCreate,
				Description: "Create a zone",
			},
			actionDescription{
				Action:      aGet,
				Description: "Get a zone",
			},
			actionDescription{
				Action:      aUpdate,
				Description: "Update a zone",
			},
			actionDescription{
				Action:      aDelete,
				Description: "Delete a zone",
			},
			actionDescription{
				Action:      aDomains,
				Description: "List domains attached to a zone",
			},
			actionDescription{
				Action:      aAttach,
				Description: "Attach a domain to a zone",
			},
		})
	}
}

func createZone() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"Name of the zone to be created",
		})
		return
	}
	jsonPrint(g.CreateZone((*args)[0]))
}

func getZone() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone to be returned",
		})
		return
	}
	jsonPrint(g.GetZone((*args)[0]))
}

func updateZone() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"UUID of the zone to be updated",
			"New name of the zone",
		})
		return
	}
	jsonPrint(g.UpdateZone((*args)[0], (*args)[1]))
}

func deleteZone() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone to be deleted",
		})
		return
	}
	noPrint(g.DeleteZone((*args)[0]))
}

func getZoneDomains() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone to be checked",
		})
		return
	}
	jsonPrint(g.GetZoneDomains((*args)[0]))
}

func attachDomainToZone() {
	if len(*args) < 2 {
		displayArgsList([]string{
			"UUID of the zone to attach the domain to",
			"FQDN of the domain to be attached",
		})
		return
	}
	jsonPrint(g.AttachDomainToZone((*args)[0], (*args)[1]))
}
