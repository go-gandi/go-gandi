package main

func livednsSnapshot() {
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
				Description: "Create a livednsSnapshot",
			},
			actionDescription{
				Action:      aGet,
				Description: "Get a livednsSnapshot",
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
	jsonPrint(l.ListSnapshots((*args)[0]))
}

func createSnapshot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone for which to create a livednsSnapshot",
		})
		return
	}
	jsonPrint(l.CreateSnapshot((*args)[0]))
}

func getSnapshot() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone for which to get a livednsSnapshot",
			"UUID of the livednsSnapshot",
		})
		return
	}
	jsonPrint(l.GetSnapshot((*args)[0], (*args)[1]))
}
