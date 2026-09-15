package str

import "fmt"

// ErrorResponse is the body TMDB returns on non-2xx responses.
type ErrorResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// Error implements the error interface.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("tmdb: %d %s", e.StatusCode, e.StatusMessage)
}
