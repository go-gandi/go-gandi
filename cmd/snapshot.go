package main

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
				Action:      aList,
				Description: "List all snapshots for a zone",
			},
			actionDescription{
				Action:      aCreate,
				Description: "Create a snapshot",
			},
			actionDescription{
				Action:      aGet,
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
	jsonPrint(g.ListSnapshots((*args)[0]))
}

func createSnapshot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone for which to create a snapshot",
		})
		return
	}
	jsonPrint(g.CreateSnapshot((*args)[0]))
}

func getSnapshot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone for which to get a snapshot",
			"UUID of the snapshot",
		})
		return
	}
	jsonPrint(g.GetSnapshot((*args)[0], (*args)[1]))
}
