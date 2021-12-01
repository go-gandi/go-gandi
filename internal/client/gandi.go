package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/types"
	"moul.io/http2curl"
)

// Gandi is the handle used to interact with the Gandi API
type Gandi struct {
	apikey    string
	endpoint  string
	sharingID string
	debug     bool
	dryRun    bool
}

// New instantiates a new Gandi client
func New(apikey string, apiurl string, sharingID string, debug bool, dryRun bool) *Gandi {
	if apiurl == "" {
		apiurl = config.APIURL
	}
	endpoint := apiurl + "/v5/"
	return &Gandi{apikey: apikey, endpoint: endpoint, sharingID: sharingID, debug: debug, dryRun: dryRun}
}

// SetEndpoint sets the URL to the endpoint. It takes a string defining the subpath under https://api.gandi.net/v5/
func (g *Gandi) SetEndpoint(endpoint string) {
	g.endpoint = g.endpoint + endpoint
}

// GetEndpoint gets the URL of the endpoint.
func (g *Gandi) GetEndpoint() string {
	return g.endpoint
}

// Get issues a GET request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Get(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodGet, path, params, recipient)
}

// Post issues a POST request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Post(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodPost, path, params, recipient)
}

// Patch issues a PATCH request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Patch(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodPatch, path, params, recipient)
}

// Delete issues a DELETE request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Delete(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodDelete, path, params, recipient)
}

// Put issues a PUT request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Put(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodPut, path, params, recipient)
}

func (g *Gandi) askGandi(method, path string, params, recipient interface{}) (http.Header, error) {
	header, body, err := g.doAskGandi(method, path, params, nil)
	if err != nil {
		return nil, err
	}
	if recipient == nil {
		return header, nil
	}

	return header, json.Unmarshal(body, &recipient)
}

// GetBytes issues a GET request but does not attempt to parse any response into JSON.
// It returns the response headers, a byteslice of the response, and any error
func (g *Gandi) GetBytes(path string, params interface{}) (http.Header, []byte, error) {
	headers := [][2]string{
		{"Accept", "text/plain"},
	}
	return g.doAskGandi(http.MethodGet, path, params, headers)
}

func (g *Gandi) doAskGandi(method, path string, p interface{}, extraHeaders [][2]string) (http.Header, []byte, error) {
	var (
		err error
		req *http.Request
	)
	params, err := json.Marshal(p)
	if err != nil {
		return nil, nil, fmt.Errorf("Fail to json.Marshal request params (error '%w')", err)
	}
	client := &http.Client{}
	suffix := ""
	if len(g.sharingID) != 0 {
		suffix += "?sharing_id=" + g.sharingID
	}
	if params != nil && string(params) != "null" {
		req, err = http.NewRequest(method, g.endpoint+path+suffix, bytes.NewReader(params))
	} else {
		req, err = http.NewRequest(method, g.endpoint+path+suffix, nil)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("Fail to create the request (error '%w')", err)
	}
	req.Header.Add("Authorization", "Apikey "+g.apikey)
	req.Header.Add("Content-Type", "application/json")
	if g.dryRun {
		req.Header.Add("Dry-Run", "1")
	}
	for _, header := range extraHeaders {
		req.Header.Add(header[0], header[1])
	}
	if g.debug {
		command, _ := http2curl.GetCurlCommand(req)
		log.Println("Request: ", command)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("Fail to do the request (error '%w')", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Fail to read the body (error '%w')", err)
	}
	if g.debug {
		var header bytes.Buffer
		for k, e := range resp.Header {
			header.WriteString(fmt.Sprintf("%s: %s ", k, e))
		}
		log.Printf("Response : [%s] %s", resp.Status, header.String())
		log.Printf("Response body: %s", string(body))
	}
	// Delete queries can return a 204 code. In this case, the
	// body is empty. See for instance:
	// https://api.gandi.net/docs/simplehosting/#delete-v5-simplehosting-instances-instance_id
	if resp.StatusCode == 204 {
		return resp.Header, []byte("{}"), err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var message types.StandardResponse

		if err = json.Unmarshal(body, &message); err != nil {
			return nil, nil, fmt.Errorf("Fail to decode the response body (error '%w')", err)
		}
		if message.Message != "" {
			err = fmt.Errorf("%d: %s", resp.StatusCode, message.Message)
		} else if len(message.Errors) > 0 {
			var errors []string
			for _, oneError := range message.Errors {
				errors = append(errors, fmt.Sprintf("%s: %s", oneError.Name, oneError.Description))
			}
			err = fmt.Errorf(strings.Join(errors, ", "))
		} else {
			err = fmt.Errorf("%d", resp.StatusCode)

		}
	}
	return resp.Header, body, err
}
