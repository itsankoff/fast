// Package api defines basic operations for speed test APIs
package api

// API interface defines operations for speed test API
type API interface {
	GetDownloadURLs(count int, secure bool) ([]string, error)
}
