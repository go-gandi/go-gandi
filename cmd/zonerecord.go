package main

import (
	"strconv"
	"strings"

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
				Description: "List all records in a zone, by zone ID or by domain",
			},
			actionDescription{
				Action:      "text",
				Description: "List all records in a zone as text",
			},
			actionDescription{
				Action:      "get",
				Description: "Get a single record in a zone, by zone ID or by domain",
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
			"UUID of the zone containing the records to be listed, or FQDN of an attached domain",
			"(optional) name of the records",
		})
		return
	}
	uuidOrFQDN := (*args)[0]
	if len(*args) < 2 {
		if strings.Contains(uuidOrFQDN, ".") {
			jsonPrint(g.ListDomainRecords(*apiKey, uuidOrFQDN))
		} else {
			jsonPrint(g.ListZoneRecords(*apiKey, uuidOrFQDN))
		}
	} else {
		if strings.Contains(uuidOrFQDN, ".") {
			jsonPrint(g.ListDomainRecordsWithName(*apiKey, uuidOrFQDN, (*args)[1]))
		} else {
			jsonPrint(g.ListZoneRecordsWithName(*apiKey, uuidOrFQDN, (*args)[1]))
		}
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
	uuidOrFQDN := (*args)[0]
	if strings.Contains(uuidOrFQDN, ".") {
		jsonPrint(g.CreateDomainRecord(*apiKey, uuidOrFQDN, (*args)[1], (*args)[2], ttl, (*args)[4:]))
	} else {
		jsonPrint(g.CreateZoneRecord(*apiKey, uuidOrFQDN, (*args)[1], (*args)[2], ttl, (*args)[4:]))
	}
}

func getRecord() {
	if len(*args) < 3 {
		displayArgsList([]string{
			"UUID of the zone containing the records to be listed, or FQDN of an attached domain",
			"Name of the record",
			"Type of the record",
		})
		return
	}
	uuidOrFQDN := (*args)[0]
	if strings.Contains(uuidOrFQDN, ".") {
		jsonPrint(g.GetDomainRecordWithNameAndType(*apiKey, uuidOrFQDN, (*args)[1], (*args)[2]))
	} else {
		jsonPrint(g.GetZoneRecordWithNameAndType(*apiKey, uuidOrFQDN, (*args)[1], (*args)[2]))
	}
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
	uuidOrFQDN := (*args)[0]
	if strings.Contains(uuidOrFQDN, ".") {
		if len(*args) < 2 {
			noPrint(g.DeleteAllDomainRecords(*apiKey, uuidOrFQDN))
		} else if len(*args) < 3 {
			noPrint(g.DeleteDomainRecords(*apiKey, uuidOrFQDN, (*args)[1]))
		} else {
			noPrint(g.DeleteDomainRecord(*apiKey, uuidOrFQDN, (*args)[1], (*args)[2]))
		}
	} else {
		if len(*args) < 2 {
			noPrint(g.DeleteAllZoneRecords(*apiKey, uuidOrFQDN))
		} else if len(*args) < 3 {
			noPrint(g.DeleteZoneRecords(*apiKey, uuidOrFQDN, (*args)[1]))
		} else {
			noPrint(g.DeleteZoneRecord(*apiKey, uuidOrFQDN, (*args)[1], (*args)[2]))
		}
	}
}
