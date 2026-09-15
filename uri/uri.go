// Package uri provides URL/query-string helpers shared by internal services.
package uri

import "net/url"

// SanitizeURL redacts the api_key query parameter, if present, so it never
// leaks into logs or error messages.
func SanitizeURL(u *url.URL) *url.URL {
	if u == nil {
		return u
	}
	q := u.Query()
	if q.Get("api_key") != "" {
		q.Set("api_key", "REDACTED")
	}
	uCopy := *u
	uCopy.RawQuery = q.Encode()
	return &uCopy
}
