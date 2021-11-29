package certificate

import (
	"github.com/go-gandi/go-gandi/internal/client"
)

type Certificate struct {
	client client.Gandi
}

type Package struct {
	Name string `json:"name"`
}

type CertificateType struct {
	ID      string   `json:"id"`
	CN      string   `json:"cn"`
	Package *Package `json:"package"`
	Status  string   `json:"status"`
}

type CreateCertificateRequest struct {
	CN      string `json:"cn"`
	Package string `json:"package"`
}

type CreateCertificateResponse struct {
	Href    string `json:"href"`
	ID      string `json:"id"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Cause   string `json:"cause,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Object  string `json:"object,omitempty"`
}
