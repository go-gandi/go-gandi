package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kingpin"
)

const (
	create = "create"
	list   = "list"
)

var (
	resourceType = kingpin.Arg("type", "Resource type (zone, record, snapshot or domain)").Required().String()
	action       = kingpin.Arg("action", "Action (valid actions depend on the type - if you provide an erroneous action, a list of allowed actions will be displayed)").Required().String()
	args         = kingpin.Arg("args", "Arguments to the action (valid arguments depend on the action)").Strings()
	apiKey       = kingpin.Flag("key", "The Gandi LiveDNS API key (may be stored in the GANDI_KEY environment variable)").OverrideDefaultFromEnvar("GANDI_KEY").Short('k').String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
	switch *resourceType {
	case "zone":
		zone()
	case "record":
		zoneRecord()
	case "snapshot":
		snapshot()
	case "domain":
		domain()
	default:
		kingpin.Usage()
	}
}

func jsonPrint(data interface{}, err error) {
	if err != nil {
		fmt.Printf("{\"error\": \"%s\"\n", err)
		return
	}
	response, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(response))
}
func textPrint(data []byte, err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(data))
}
func noPrint(err error) {
	if err != nil {
		fmt.Printf("{\"error\": \"%s\"\n", err)
		return
	}
	fmt.Println("OK")
}

type actionDescription struct {
	Action      string
	Description string
}

func displayActionsList(actions []actionDescription) {
	fmt.Printf("Here are the actions accepted by type \"%s\":\n", *resourceType)
	for _, action := range actions {
		fmt.Printf("  %s: %s\n", action.Action, action.Description)
	}
}

func displayArgsList(arguments []string) {
	fmt.Printf("Here are the actions accepted by type \"%s\" and action \"%s\":\n", *resourceType, *action)
	for _, arg := range arguments {
		fmt.Printf("  * %s\n", arg)
	}
}
