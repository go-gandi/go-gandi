package simplehosting_test

import (
	"testing"

	"github.com/go-gandi/go-gandi/simplehosting"
	"gopkg.in/h2non/gock.v1"
)

func TestDeleteInstance(t *testing.T) {
	defer gock.Off()
	instanceId := "23739138-4850-11ec-8973-00163ec4cb00"

	gock.New("https://api.gandi.net/v5/").
		Delete("/instances/" + instanceId).
		Reply(204)

	simpleHosting := simplehosting.New("", "", true, false)
	response, err := simpleHosting.DeleteInstance(instanceId)
	if err != nil {
		t.Fatal(err)
	}
	expected := simplehosting.ErrorResponse{}
	if response != expected {
		t.Fatalf("Response should be '%#v' (while it is %#v)", expected, response)
	}
}

// TestCreateInstance tests the instance ID is correctly returned.
func TestCreateInstance(t *testing.T) {
	defer gock.Off()
	instanceName := "new-instance"
	expectedInstanceId := "23739138-4850-11ec-8973-00163ec4cb00"

	gock.New("https://api.gandi.net/v5/").
		Post("instances").
		JSON(map[string]interface{}{
			"name":     instanceName,
			"location": "FR",
			"type": map[string]interface{}{
				"database": map[string]string{
					"name":    "mysql",
					"version": "",
				},
				"language": map[string]string{
					"name":    "php",
					"version": "",
				},
			},
			"size": "",
		}).
		Reply(202).
		SetHeader(
			"Content-Location",
			"https://api.gandi.net/v5/simplehosting/"+"instances/"+expectedInstanceId).
		JSON(map[string]string{
			"message": "Instance is being created",
		})

	simpleHosting := simplehosting.New("", "", true, false)
	instanceId, err := simpleHosting.CreateInstance(
		simplehosting.CreateInstanceRequest{
			Name:     instanceName,
			Location: "FR",
			Type: &simplehosting.InstanceType{
				Database: &simplehosting.Database{
					Name: "mysql",
				},
				Language: &simplehosting.Language{
					Name: "php",
				},
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}
	if instanceId != expectedInstanceId {
		t.Fatalf("InstanceId should be '%s' (while it is %s)",
			expectedInstanceId, instanceId)
	}
}
