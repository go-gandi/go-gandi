package config

// Config manages common config for all Gandi API types
type Config struct {
	// SharingID is the Organization ID, available from the Organization API
	SharingID string
	// Debug enables verbose debugging of HTTP calls
	Debug bool
	// DryRun prevents the API from making changes. Only certain API calls support it.
	DryRun bool
}
