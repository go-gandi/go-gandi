package main

import (
	"strconv"
	"strings"
)

func zoneRecord() {
	switch *action {
	case aList:
		listRecords()
	case aText:
		listRecordsAsText()
	case aGet:
		getRecord()
	case aCreate:
		createRecord()
	case aUpdate:
		updateRecord()
	case aDelete:
		deleteRecord()
	default:
		displayActionsList([]actionDescription{
			actionDescription{
				Action:      aList,
				Description: "List all records in a zone, by zone ID or by domain",
			},
			actionDescription{
				Action:      aText,
				Description: "List all records in a zone as text",
			},
			actionDescription{
				Action:      aGet,
				Description: "Get a single record in a zone, by zone ID or by domain",
			},
			actionDescription{
				Action:      aCreate,
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
			jsonPrint(g.ListDomainRecords(uuidOrFQDN))
		} else {
			jsonPrint(g.ListZoneRecords(uuidOrFQDN))
		}
	} else {
		if strings.Contains(uuidOrFQDN, ".") {
			jsonPrint(g.ListDomainRecordsWithName(uuidOrFQDN, (*args)[1]))
		} else {
			jsonPrint(g.ListZoneRecordsWithName(uuidOrFQDN, (*args)[1]))
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
	textPrint(g.ListZoneRecordsAsText((*args)[0]))
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
		jsonPrint(g.CreateDomainRecord(uuidOrFQDN, (*args)[1], (*args)[2], ttl, (*args)[4:]))
	} else {
		jsonPrint(g.CreateZoneRecord(uuidOrFQDN, (*args)[1], (*args)[2], ttl, (*args)[4:]))
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
		jsonPrint(g.GetDomainRecordWithNameAndType(uuidOrFQDN, (*args)[1], (*args)[2]))
	} else {
		jsonPrint(g.GetZoneRecordWithNameAndType(uuidOrFQDN, (*args)[1], (*args)[2]))
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
	jsonPrint(g.ChangeZoneRecordWithNameAndType((*args)[0], (*args)[1], (*args)[2], ttl, values))
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
			noPrint(g.DeleteAllDomainRecords(uuidOrFQDN))
		} else if len(*args) < 3 {
			noPrint(g.DeleteDomainRecords(uuidOrFQDN, (*args)[1]))
		} else {
			noPrint(g.DeleteDomainRecord(uuidOrFQDN, (*args)[1], (*args)[2]))
		}
	} else {
		if len(*args) < 2 {
			noPrint(g.DeleteAllZoneRecords(uuidOrFQDN))
		} else if len(*args) < 3 {
			noPrint(g.DeleteZoneRecords(uuidOrFQDN, (*args)[1]))
		} else {
			noPrint(g.DeleteZoneRecord(uuidOrFQDN, (*args)[1], (*args)[2]))
		}
	}
}
