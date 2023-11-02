package client

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/go-gandi/go-gandi/types"
	"gopkg.in/h2non/gock.v1"
)

type element struct {
	Item string `json:"item"`
}

func TestAskGandiCollection(t *testing.T) {
	defer gock.Off()
	gock.New("https://api.gandi.net/v5/").
		Get("/domain/domains").
		Reply(200).
		SetHeader("link", "<https://api.gandi.net/v5/domain/domains?page=2&sort_by=fqdn>; rel=\"next\", <https://api.gandi.net/v5/domain/domains?sort_by=fqdn&page=2>; rel=\"last\"").
		JSON([]map[string]string{map[string]string{"item": "item1"}})

	gock.New("https://api.gandi.net/v5/").
		Get("/domain/domains").
		MatchParam("page", "2").
		MatchParam("sort_by", "fqdn").
		Reply(200).
		JSON([]map[string]string{map[string]string{"item": "item2"}})

	client := New("", "", "https://api.gandi.net", "", false, false, 1*time.Second)
	var elements []element
	_, rawMessages, err := client.askGandiCollection("GET", "domain/domains", nil)
	for _, rawMessage := range rawMessages {
		var element element
		err := json.Unmarshal(rawMessage, &element)
		if err != nil {
			t.Fatal(err)
		}
		elements = append(elements, element)
	}
	if err != nil {
		t.Fatal(err)
	}
	expected := []element{
		element{
			Item: "item1",
		},
		element{
			Item: "item2",
		},
	}
	if !reflect.DeepEqual(elements, expected) {
		t.Fatalf("Expected elements are '%#v' (actual %#v)", expected, elements)
	}

}

func TestAskGandiCollectionEmpty(t *testing.T) {
	defer gock.Off()
	gock.New("https://api.gandi.net/v5/").
		Get("/domain/domains").
		Reply(200).
		JSON([]map[string]string{})
	client := New("", "", "https://api.gandi.net", "", false, false, 1*time.Second)
	_, rawMessages, err := client.askGandiCollection("GET", "domain/domains", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(rawMessages) != 0 {
		t.Fatalf("Length of elements slice should be 0 (instead of %d)", len(rawMessages))
	}

}

func TestRequestError(t *testing.T) {
	defer gock.Off()
	gock.New("https://api.gandi.net/v5/").
		Get("/domain/domains").
		Reply(500).
		JSON(types.StandardResponse{})
	client := New("", "", "https://api.gandi.net", "", false, false, 1*time.Second)
	response := []map[string]string{}
	_, err := client.Get("domain/domains", nil, &response)

	var e *types.RequestError
	if errors.As(err, &e) {
		if e.StatusCode != 500 {
			t.Fatalf("Error StatusCode should be: %v)", e.StatusCode)
		}
	} else {
		t.Fatalf("Error type is not RequestError (actual: %v)", err)
	}
}

func TestRequestErrorNotJson(t *testing.T) {
	defer gock.Off()
	gock.New("https://api.gandi.net/v5/").
		Get("/domain/domains").
		Reply(400).
		AddHeader("Content-Type", "text/html").
		BodyString("<html><p>error</p></html>")
	client := New("", "https://api.gandi.net", "", false, false, 1*time.Second)
	response := []map[string]string{}
	_, err := client.Get("domain/domains", nil, &response)

	if err.Error() != "Response body is not json for status 400" {
		t.Fatalf("Invalid error for non Json response code")
	}
}
