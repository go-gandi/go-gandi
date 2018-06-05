package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kingpin"
	gandi "github.com/tiramiseb/go-gandi-livedns"
)

const (
	// Actions names
	aAdd         = "add"
	aAttach      = "attach"
	aCreate      = "create"
	aDelete      = "delete"
	aDelKey      = "delkey"
	aDelSlave    = "delslave"
	aDetach      = "detach"
	aDomains     = "domains"
	aGet         = "get"
	aKeys        = "keys"
	aList        = "list"
	aGetBIND     = "bind"
	aGetPowerDNS = "powerdns"
	aGetNSD      = "nsd"
	aGetKnot     = "knot"
	aNS          = "ns"
	aSign        = "sign"
	aSlave       = "slave"
	aSlaves      = "slaves"
	aText        = "text"
	aUpdate      = "update"
)

var (
	resourceType = kingpin.Arg("type", "Resource type (zone, record, snapshot, domain or axfr)").Required().String()
	action       = kingpin.Arg("action", "Action (valid actions depend on the type - if you provide an erroneous action, a list of allowed actions will be displayed)").Required().String()
	args         = kingpin.Arg("args", "Arguments to the action (valid arguments depend on the action)").Strings()
	apiKey       = kingpin.Flag("key", "The Gandi LiveDNS API key (may be stored in the GANDI_KEY environment variable)").OverrideDefaultFromEnvar("GANDI_KEY").Short('k').String()
	sharing_id   = kingpin.Flag("sharing_id", "The Gandi LiveDNS sharing_id (may be stored in the GANDI_SHARING_ID environment variable)").OverrideDefaultFromEnvar("GANDI_SHARING_ID").Short('i').String()
	g            *gandi.Gandi
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
	g = gandi.New(*apiKey, *sharing_id)
	switch *resourceType {
	case "zone":
		zone()
	case "record":
		zoneRecord()
	case "snapshot":
		snapshot()
	case "domain":
		domain()
	case "axfr":
		axfr()
	default:
		kingpin.Usage()
	}
}

func jsonPrint(data interface{}, err error) {
	if err != nil {
		fmt.Printf("{\"error\": \"%s\"}\n", err)
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
		fmt.Printf("{\"error\": \"%s\"}\n", err)
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
