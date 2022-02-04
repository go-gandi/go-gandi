package livedns_test

import (
	"testing"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/livedns"
	"gopkg.in/h2non/gock.v1"
)

func TestSignDomain(t *testing.T) {
	defer gock.Off()
	expectedUUID := "29f246af-fe71-4d73-8693-ac2d0e49b23b"

	gock.Observe(gock.DumpRequest)
	gock.New("https://api.gandi.net/v5/").
		Post("livedns/domains/example.com/keys").
		JSON(map[string]interface{}{
			"flags": 257,
		}).
		Reply(201).
		SetHeader(
			"location",
			"https://api.gandi.net/v5/livedns/domains/example.com/keys/"+expectedUUID).
		JSON(map[string]string{
			"message": "Domain Key Created",
		})

	liveDNS := livedns.New(config.Config{})
	response, err := liveDNS.SignDomain("example.com")
	if err != nil {
		t.Fatal(err)
	}
	if response.UUID != expectedUUID {
		t.Fatalf("UUID should be '%s' (while it is %s)",
			expectedUUID, response.UUID)
	}
}
