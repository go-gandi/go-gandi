package email

import "github.com/go-gandi/go-gandi/internal/client"

// Email is the API client to the Gandi v5 Email API
type Email struct {
	client client.Gandi
}

// New returns an instance of the Email API client
func New(apikey string, sharingid string, debug bool, dryRun bool) *Email {
	client := client.New(apikey, sharingid, debug, dryRun)
	client.SetEndpoint("email/")
	return &Email{client: *client}
}

// NewFromClient returns an instance of the Email API client
func NewFromClient(g client.Gandi) *Email {
	g.SetEndpoint("email/")
	return &Email{client: g}
}

// ListMailboxResponse describes mailbox
type ListMailboxResponse struct {
	Address     string `json:"address"`
	Domain      string `json:"domain"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	Login       string `json:"login"`
	MailboxType string `json:"mailbox_type"`
	QuataUsed   int    `json:"quota_used"`
}

// MailboxResponse mailbox parameters
type MailboxResponse struct {
	Domain    string `json:"domain"`
	Responder struct {
		Message string `json:"message"`
		Enabled bool   `json:"enabled"`
	} `json:"responder"`
	MailboxType string   `json:"mailbox_type"`
	Login       string   `json:"login"`
	QuotaUsed   int      `json:"quota_used"`
	Aliases     []string `json:"aliases"`
	Address     string   `json:"address"`
	Href        string   `json:"href"`
	ID          string   `json:"id"`
}

// CreateEmailRequest create mailbox request
type CreateEmailRequest struct {
	Login       string   `json:"login"`
	MailboxType string   `json:"mailbox_type"`
	Password    string   `json:"password"`
	Aliases     []string `json:"aliases,omitempty"`
}

// UpdateEmailRequest update mailbox request
type UpdateEmailRequest struct {
	Login    string   `json:"login,omitempty"`
	Password string   `json:"password,omitempty"`
	Aliases  []string `json:"aliases"`
}

// CreateForwardRequest structure for forwarding request
type CreateForwardRequest struct {
	Source       string   `json:"source"`
	Destinations []string `json:"destinations"`
}

// GetForwardRequest structure for forwarding responses
type GetForwardRequest struct {
	Source       string   `json:"source"`
	Destinations []string `json:"destinations"`
	Href         string   `json:"href"`
}

// UpdateForwardRequest structure for updating forwarding
type UpdateForwardRequest struct {
	Destinations []string `json:"destinations"`
}

// ListMailboxes list mailboxes attached to domain
func (e *Email) ListMailboxes(domain string) (mailboxes []ListMailboxResponse, err error) {
	_, err = e.client.Get("/mailboxes/"+domain, nil, &mailboxes)
	return
}

// GetMailbox returns all the parameters linked to a specific mailbox
func (e *Email) GetMailbox(domain, mailbox_id string) (mailbox MailboxResponse, err error) {
	_, err = e.client.Get("/mailboxes/"+domain+"/"+mailbox_id, nil, &mailbox)
	return
}

// CreateEmail creates a new mailbox for the given domain
func (e *Email) CreateEmail(domain string, req CreateEmailRequest) (err error) {
	_, err = e.client.Post("/mailboxes/"+domain, req, nil)
	return
}

// UpdateEmail update mailbox parameters
func (e *Email) UpdateEmail(domain, mailbox_id string, req UpdateEmailRequest) (err error) {
	_, err = e.client.Patch("/mailboxes/"+domain+"/"+mailbox_id, req, nil)
	return
}

// DeleteEmail remove mailbox
func (e *Email) DeleteEmail(domain, mailbox_id string) (err error) {
	_, err = e.client.Delete("/mailboxes/"+domain+"/"+mailbox_id, nil, nil)
	return
}

// CreateForward creates forwarding
func (e *Email) CreateForward(domain string, req CreateForwardRequest) (err error) {
	_, err = e.client.Post("/forwards/"+domain, req, nil)
	return
}

// GetForwards retrieves all forwardings for domain
func (e *Email) GetForwards(domain string) (forwards []GetForwardRequest, err error) {
	_, err = e.client.Get("/forwards/"+domain, nil, &forwards)
	return
}

// GetForward retrieve single forwarding
func (e *Email) GetForward(domain, source string) (forward GetForwardRequest, err error) {
	_, err = e.client.Get("/forwards/"+domain+"/"+source, nil, &forward)
	return
}

// UpdateForward update forwarding
func (e *Email) UpdateForward(domain, source string, req UpdateForwardRequest) (err error) {
	_, err = e.client.Put("/forwards/"+domain+"/"+source, req, nil)
	return
}

// DeleteForward delete forwarding
func (e *Email) DeleteForward(domain, source string) (err error) {
	_, err = e.client.Delete("/forwards/"+domain+"/"+source, nil, nil)
	return
}
