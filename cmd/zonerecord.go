package main

import (
	"strconv"
	"strings"
)

func livednsDomainRecord() {
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
		jsonPrint(l.ListDomainRecords(uuidOrFQDN))
	} else {
		jsonPrint(l.ListDomainRecordsByName(uuidOrFQDN, (*args)[1]))
	}
}

func listRecordsAsText() {
	if len(*args) < 1 {
		displayArgsList([]string{
			"UUID of the zone containing the records to be listed",
		})
		return
	}
	textPrint(l.ListDomainRecordsAsText((*args)[0]))
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
	jsonPrint(l.CreateDomainRecord(uuidOrFQDN, (*args)[1], (*args)[2], ttl, (*args)[4:]))
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
	jsonPrint(l.GetDomainRecordByNameAndType(uuidOrFQDN, (*args)[1], (*args)[2]))
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
	jsonPrint(l.UpdateDomainRecordByNameAndType((*args)[0], (*args)[1], (*args)[2], ttl, values))
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
			noPrint(l.DeleteAllDomainRecords(uuidOrFQDN))
		} else if len(*args) < 3 {
			noPrint(l.DeleteDomainRecords(uuidOrFQDN, (*args)[1]))
		} else {
			noPrint(l.DeleteDomainRecord(uuidOrFQDN, (*args)[1], (*args)[2]))
		}
	} else {
		if len(*args) < 2 {
			noPrint(l.DeleteAllDomainRecords(uuidOrFQDN))
		} else if len(*args) < 3 {
			noPrint(l.DeleteDomainRecords(uuidOrFQDN, (*args)[1]))
		} else {
			noPrint(l.DeleteDomainRecord(uuidOrFQDN, (*args)[1], (*args)[2]))
		}
	}
}
