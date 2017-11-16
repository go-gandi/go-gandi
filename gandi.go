package gandi

import (
	"bytes"
	"encoding/json"
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
	var (
		err              error
		marshalledParams []byte
		req              *http.Request
	)
	client := &http.Client{}
	if params != nil {
		marshalledParams, err = json.Marshal(params)
		if err != nil {
			return err
		}
		req, err = http.NewRequest(method, gandiEndpoint+path, bytes.NewReader(marshalledParams))
	} else {
		req, err = http.NewRequest(method, gandiEndpoint+path, nil)
	}
	if err != nil {
		return err
	}
	req.Header.Add("X-Api-Key", key)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(recipient)
	return nil
}

func askGandiToBytes(key, method, path string, params interface{}) ([]byte, error) {
	var (
		err              error
		marshalledParams []byte
		req              *http.Request
	)
	client := &http.Client{}
	if params != nil {
		marshalledParams, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, gandiEndpoint+path, bytes.NewReader(marshalledParams))
	} else {
		req, err = http.NewRequest(method, gandiEndpoint+path, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func askGandiFromBytes(key, method, path string, params []byte, recipient interface{}) error {
	var (
		err error
		req *http.Request
	)
	client := &http.Client{}
	if params != nil {
		req, err = http.NewRequest(method, gandiEndpoint+path, bytes.NewReader(params))
	} else {
		req, err = http.NewRequest(method, gandiEndpoint+path, nil)
	}
	if err != nil {
		return err
	}
	req.Header.Add("X-Api-Key", key)
	req.Header.Add("Content-Type", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(recipient)
	return nil
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
