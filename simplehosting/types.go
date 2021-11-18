package simplehosting

import (
	"github.com/go-gandi/go-gandi/internal/client"
)

// SimpleHosting is the API client to the Gandi v5 Simple Hosting API
type SimpleHosting struct {
	client client.Gandi
}

// ListInstancesResponse is the response object returned by listing
// simplehosting instances
type ListInstancesResponse struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Size     string    `json:"size"`
	Status   string    `json:"status"`
	Database *Database `json:"database"`
	Language *Language `json:"language"`
}

// Database represents the type of a Simple Hosting database
type Database struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Language represents the type of a Simple Hosting database
type Language struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Instance struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Size     string    `json:"size"`
	Status   string    `json:"status"`
	Database *Database `json:"database"`
	Language *Language `json:"language"`
}

type InstanceType struct {
	Database *Database `json:"database"`
	Language *Language `json:"language"`
}

type CreateInstanceRequest struct {
	Location string        `json:"location"`
	Type     *InstanceType `json:"type"`
	Name     string        `json:"name"`
	Size     string        `json:"size"`
}

type ErrorResponse struct {
	Cause   string `json:"cause,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Object  string `json:"object,omitempty"`
}
