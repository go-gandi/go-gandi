package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-gandi/go-gandi/types"
	"io"
	"io/ioutil"
	"log"
	"moul.io/http2curl"
	"net/http"
	"strings"
)

const (
	gandiEndpoint = "https://api.gandi.net/v5/"
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
func New(apikey string, sharingID string, debug bool, dryRun bool) *Gandi {
	return &Gandi{apikey: apikey, endpoint: gandiEndpoint, sharingID: sharingID, debug: debug, dryRun: dryRun}
}

// SetEndpoint sets the URL to the endpoint. It takes a string defining the subpath under https://api.gandi.net/v5/
func (g *Gandi) SetEndpoint(endpoint string) {
	g.endpoint = gandiEndpoint + endpoint
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
	resp, err := g.doAskGandi(method, path, params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if recipient == nil {
		return resp.Header, nil
	}

	decoder := json.NewDecoder(resp.Body)

	return resp.Header, decoder.Decode(recipient)
}

// GetBytes issues a GET request but does not attempt to parse any response into JSON.
// It returns the response headers, a byteslice of the response, and any error
func (g *Gandi) GetBytes(path string, params interface{}) (http.Header, []byte, error) {
	headers := [][2]string{
		{"Accept", "text/plain"},
	}
	resp, err := g.doAskGandi(http.MethodGet, path, params, headers)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return resp.Header, content, err
}

func (g *Gandi) doAskGandi(method, path string, p interface{}, extraHeaders [][2]string) (*http.Response, error) {
	var (
		err error
		req *http.Request
	)
	params, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("Fail to json.Marshal request params (error '%s')", err)
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
		return nil, fmt.Errorf("Fail to create the request (error '%s')", err)
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
		return resp, fmt.Errorf("Fail to do the request (error '%s')", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf("Fail to read the body (error '%s')", err)
	}
	if g.debug {
		var header bytes.Buffer
		for k, e := range resp.Header {
			header.WriteString(fmt.Sprintf("%s: %s ", k, e))
		}
		log.Println(fmt.Sprintf("Response : [%s] %s", resp.Status, header.String()))
		log.Println(fmt.Sprintf("Response body: %s", string(body)))
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var message types.StandardResponse

		if err = json.Unmarshal(body, &message); err != nil {
			return resp, fmt.Errorf("Fail to decode the response body (error '%s')", err)
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
	return resp, err
}
