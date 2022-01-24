package email

import "github.com/go-gandi/go-gandi/internal/client"

// Email is the API client to the Gandi v5 Email API
type Email struct {
	client client.Gandi
}

// ListMailboxResponse describes mailbox
type ListMailboxResponse struct {
	Address     string `json:"address"`
	Domain      string `json:"domain"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	Login       string `json:"login"`
	MailboxType string `json:"mailbox_type"`
	QuotaUsed   int    `json:"quota_used"`
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
