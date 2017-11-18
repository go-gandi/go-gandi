package gandi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	gandiEndpoint = "https://dns.api.gandi.net/api/v5/"
	// HTTP Methods
	mPATCH  = http.MethodPatch
	mGET    = http.MethodGet
	mPOST   = http.MethodPost
	mDELETE = http.MethodDelete
	mPUT    = http.MethodPut
)

func askGandi(key, method, path string, params, recipient interface{}) error {
	marshalledParams, err := json.Marshal(params)
	if err != nil {
		return err
	}
	resp, err := doAskGandi(key, method, path, marshalledParams, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(recipient)
	return nil
}

func askGandiToBytes(key, method, path string, params interface{}) ([]byte, error) {
	headers := [][2]string{
		[2]string{"Accept", "text/plain"},
	}
	marshalledParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := doAskGandi(key, method, path, marshalledParams, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func askGandiFromBytes(key, method, path string, params []byte, recipient interface{}) error {
	resp, err := doAskGandi(key, method, path, params, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(recipient)
	return nil
}

func doAskGandi(key, method, path string, params []byte, extraHeaders [][2]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, gandiEndpoint+path, bytes.NewReader(params))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", key)
	req.Header.Add("Content-Type", "application/json")
	for _, header := range extraHeaders {
		req.Header.Add(header[0], header[1])
	}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		defer resp.Body.Close()
		var message StandardResponse
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&message)
		if message.Message != "" {
			err = fmt.Errorf("%d: %s", resp.StatusCode, message.Message)
		} else {
			err = fmt.Errorf("%d", resp.StatusCode)

		}
	}
	return resp, err
}

// StandardResponse is a standard response
type StandardResponse struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	UUID    string          `json:"uuid,omitempty"`
	Object  string          `json:"object,omitempty"`
	Cause   string          `json:"cause,omitempty"`
	Status  string          `json:"status,omitempty"`
	Errors  []StandardError `json:"errors,omitempty"`
}

// StandardError is embedded in a standard error
type StandardError struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
