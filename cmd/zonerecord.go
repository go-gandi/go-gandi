package main

import (
	"strconv"

	g "github.com/tiramiseb/go-gandi-livedns"
)

func zoneRecord() {
	switch *action {
	case "list":
		listRecords()
	case "text":
		listRecordsAsText()
	case "get":
		getRecord()
	case "create":
		createRecord()
	case "update":
		updateRecord()
	case "delete":
		deleteRecord()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      "list",
				Description: "List all records in a zone",
			},
			actionDescription{
				Action:      "text",
				Description: "List all records in a zone as text",
			},
			actionDescription{
				Action:      "get",
				Description: "Get a single record in a zone",
			},
			actionDescription{
				Action:      "create",
				Description: "Create a record",
			},
		})
	}
}

func listRecords() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone containing the records to be listed",
			"(optional) name of the records",
		})
		return
	}
	if len(*args) < 2 {
		jsonPrint(g.ListZoneRecords(*apiKey, (*args)[0]))
	} else {
		jsonPrint(g.ListZoneRecordsWithName(*apiKey, (*args)[0], (*args)[1]))
	}
}

func listRecordsAsText() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone containing the records to be listed",
		})
		return
	}
	textPrint(g.ListZoneRecordsAsText(*apiKey, (*args)[0]))
}

func createRecord() {
	if len(*args) < 5 {
		displayArgsList([]string{
			"UUID of the zone where to create a record",
			"Name of the record to be created",
			"Type of the record",
			"TTL of the record",
			"Values... (multiple values possible)",
		})
		return
	}
	ttl, err := strconv.Atoi((*args)[3])
	if err != nil {
		noPrint(err)
	}
	values := (*args)[4:]
	jsonPrint(g.CreateZoneRecord(*apiKey, (*args)[0], (*args)[1], (*args)[2], ttl, values))
}

func getRecord() {
	if len(*args) < 3 {
		displayArgsList([]string{
			"UUID of the zone containing the records to be listed",
			"Name of the record",
			"Type of the record",
		})
		return
	}
	jsonPrint(g.GetZoneRecordWithNameAndType(*apiKey, (*args)[0], (*args)[1], (*args)[2]))
}

func updateRecord() {
	if len(*args) < 5 {
		displayArgsList([]string{
			"UUID of the zone containing the records to be updated",
			"Name of the record",
			"Type of the record",
			"New TTL for the record",
			"New values... (multiple values possible)",
		})
		return
	}
	ttl, err := strconv.Atoi((*args)[3])
	if err != nil {
		noPrint(err)
	}
	values := (*args)[4:]
	jsonPrint(g.ChangeZoneRecordWithNameAndType(*apiKey, (*args)[0], (*args)[1], (*args)[2], ttl, values))
}

func deleteRecord() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone containing the record(s) to be deleted",
			"(optional) name of the record(s)",
			"(optional) type of the record",
		})
		return
	}
	if len(*args) < 2 {
		noPrint(g.DeleteAllRecords(*apiKey, (*args)[0]))
	} else if len(*args) < 3 {
		noPrint(g.DeleteRecords(*apiKey, (*args)[0], (*args)[1]))
	} else {
		noPrint(g.DeleteRecord(*apiKey, (*args)[0], (*args)[1], (*args)[2]))
	}
}
