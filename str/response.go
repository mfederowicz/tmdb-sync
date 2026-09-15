// Package str holds shared data structures used across the app.
package str

import "net/http"

// Response wraps http.Response, exposing rate-limit info parsed from headers.
type Response struct {
	*http.Response
}
