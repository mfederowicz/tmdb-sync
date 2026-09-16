package internal

import (
	"fmt"
	"net/http"
	"time"
)

// AbuseRateLimitError occurs when TMDB returns 429 Too Many Requests, or when
// the client is locally short-circuiting requests before a known Retry-After.
type AbuseRateLimitError struct {
	Response   *http.Response
	Message    string
	RetryAfter *time.Duration
}

// Error implements the error interface.
func (r *AbuseRateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %s",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}
