package main

import g "github.com/tiramiseb/go-gandi-livedns"

func snapshot() {
	switch *action {
	case aList:
		listSnapshots()
	case aCreate:
		createSnapshot()
	case aGet:
		getSnapshot()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      "list",
				Description: "List all snapshots for a zone",
			},
			actionDescription{
				Action:      "create",
				Description: "Create a snapshot",
			},
			actionDescription{
				Action:      "get",
				Description: "Get a snapshot",
			},
		})
	}
}

func listSnapshots() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone for which to list snapshots",
		})
		return
	}
	jsonPrint(g.ListSnapshots(*apiKey, (*args)[0]))
}

func createSnapshot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone for which to create a snapshot",
		})
		return
	}
	jsonPrint(g.CreateSnapshot(*apiKey, (*args)[0]))
}

func getSnapshot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone for which to get a snapshot",
			"UUID of the snapshot",
		})
		return
	}
	jsonPrint(g.GetSnapshot(*apiKey, (*args)[0], (*args)[1]))
}
