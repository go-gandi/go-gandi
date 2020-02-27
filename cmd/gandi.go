package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/tiramiseb/go-gandi"
	"github.com/tiramiseb/go-gandi/domain"
	"github.com/tiramiseb/go-gandi/livedns"
)

type cli struct {
	globals
	LiveDNS   liveDNSCmd `kong:"cmd,name='livedns',help='Manage LiveDNS'"`
	Domain    domainCmd  `kong:"cmd,help='Manage Domains'"`
	Debug     bool       `kong:"short='d',help='Enable debug logging'"`
	DryRun    bool       `kong:"help='Enable dry run mode'"`
	APIKey    string     `kong:"env='GANDI_KEY',help='The Gandi LiveDNS API key (may be stored in the GANDI_KEY environment variable)'"`
	SharingID string     `kong:"short='i',env='GANDI_SHARING_ID',help='The Gandi LiveDNS sharingID (may be stored in the GANDI_SHARING_ID environment variable)'"`
}

type globals struct {
	liveDNSHandle *livedns.LiveDNS
	domainHandle  *domain.Domain
	Version       versionFlag `kong:"name='version',help='Print version information and quit'"`
}

var c cli

type versionFlag string

func (v versionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v versionFlag) IsBool() bool                         { return true }
func (v versionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

func main() {
	c = cli{
		globals: globals{
			Version: "0.0.1",
		},
	}
	ctx := kong.Parse(&c)
	g := gandi.Config{
		SharingID: c.SharingID,
		Debug:     c.Debug,
		DryRun:    c.DryRun,
	}
	c.globals.domainHandle = gandi.NewDomainClient(c.APIKey, g)
	c.globals.liveDNSHandle = gandi.NewLiveDNSClient(c.APIKey, g)
	err := ctx.Run(&c.globals)
	ctx.FatalIfErrorf(err)
}

func jsonPrint(data interface{}, err error) error {
	if err != nil {
		return fmt.Errorf("{\"error\": \"%w\"}\n", err)
	}
	response, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(response))
	return nil
}

func textPrint(data []byte, err error) error {
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

func noPrint(err error) error {
	if err != nil {
		return fmt.Errorf("{\"error\": \"%w\"}\n", err)
	}
	fmt.Println("OK")
	return nil
}
