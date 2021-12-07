package certificate

import (
	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Certificate API client
func New(config config.Config) *Certificate {
	client := client.New(config.APIKey, config.APIURL, config.SharingID, config.Debug, config.DryRun)
	client.SetEndpoint("certificate/")
	return &Certificate{client: *client}
}

// ListCertificates requests the list of issued certificates
func (g *Certificate) ListCertificates() (certificates []CertificateType, err error) {
	_, err = g.client.Get("issued-certs", nil, &certificates)
	return
}

// GetCertificate request details of an issued certificates
func (g *Certificate) GetCertificate(certificateId string) (certificate CertificateType, err error) {
	_, err = g.client.Get("issued-certs/"+certificateId, nil, &certificate)
	return
}

// CreateCertificate creates a certificate
func (g *Certificate) CreateCertificate(req CreateCertificateRequest) (response CreateCertificateResponse, err error) {
	_, err = g.client.Post("issued-certs", req, &response)
	return
}

// DeleteCertificate revokes a certificate
func (g *Certificate) DeleteCertificate(certificateId string) (response ErrorResponse, err error) {
	_, err = g.client.Delete("issued-certs/"+certificateId, nil, &response)
	return
}

// ListPackages lists certificate package types
func (g *Certificate) ListPackages() (packages []Package, err error) {
	_, err = g.client.Get("packages", nil, &packages)
	return
}
