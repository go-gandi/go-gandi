Gandi LiveDNS Go library
========================

[![GoDoc](https://godoc.org/github.com/tiramiseb/go-gandi-livedns?status.svg)](https://godoc.org/github.com/tiramiseb/go-gandi-livedns)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/tiramiseb/go-gandi-livedns/master/LICENSE)

This library interacts with [Gandi's LiveDNS API](http://doc.livedns.gandi.net/), to manage domains hosted on Gandi. This library returns some data as HTTP headers, please note those are ignored by this library.

**Gandi is currently (as of Nov. 2017) migrating on a new platform, this library is for the NEW platform.**

A simple CLI is also shipped with this library. It returns responses to the requests in JSON.

Example
-------

This example mimics the steps of [the official LiveDNS documentation example](http://doc.livedns.gandi.net/#quick-example).

First (step 1), [get your API key](https://account.gandi.net/) from the "Security" section in new Account admin panel to be able to make authenticated requests to the API.

```go
import "github.com/tiramiseb/go-gandi-livedns"
apikey = "<the API key>"
// Step 2: create the zone
gandi.CreateZone(apikey, "example.com Zone")
// Step 2bis: fetch the zone UUID
var zoneUUID string
for _, zone = range gandi.ListZones() {
    if zone.Name == "example.com Zone" {
        zoneUUID = zone.UUID
        break
    }
}
// Step 3: create DNS records
gandi.CreateZoneRecord(apikey, zoneUUID, "www", "A", 3600, []string{"192.168.0.1"})
// Step 4: associate the domain
gandi.AttachDomainToZone(apikey, zoneUUID, "example.com")
// Step 5: change nameservers
// not implemented yet
// Step 6: setup automatic DNSSEC signing
gandi.SignDomain(apikey, "example.com")
// Getting the key href
gandi.GetDomainKeys(apikey, "example.com")
// Deleting the key
gandi.DeleteDomainKey(apikey, "example.com", "bb004a38-566b-4200-bd6e-830b48ea50cf")
// Recovering the key
// not implemented
// Step 7: adding extra security with a slave server
// Creating a TSIG key
tsig, _ := gandi.CreateTsig(apikey)
// Adding the TSIG key for AXFRs
gandi.AddTsigToDomain(apikey, "example.com", tsig.UUID)
// Adding two slaves servers to the domain
for _, host = range []string{"198.51.100.1", "2001:DB8::1"} {
    gandi.AddSlaveToDomain(apikey, "example.com", host)
}
// Getting sample configurations
// not implemented yet
```

Compiling the CLI
-----------------

```
cd cmd
go build -o gandi
```
