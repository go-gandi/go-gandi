package config

// Config manages common config for all Gandi API types
type Config struct {
	// APIKey is available from https://account.gandi.net/en/
	APIKey string
	// SharingID is the Organization ID, available from the Organization API
	SharingID string
	// Debug enables verbose debugging of HTTP calls
	Debug bool
	// DryRun prevents the API from making changes. Only certain API calls support it.
	DryRun bool
	// APIURL is the Gandi API URL. By default, it fallbacks to
	// https://api.gandi.net.
	APIURL string
}

const (
	// APIURL is the default Config.APIURL value
	APIURL = "https://api.gandi.net"
	// SandboxAPIURL is the URL of the Gandi Sandbox API
	SandboxAPIURL = "https://api.sandbox.gandi.net"
)
