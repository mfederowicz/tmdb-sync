package str

// AuthStatus is the response of GET /authentication (validate key) and of
// DELETE /authentication/session (delete session) - both just confirm success.
type AuthStatus struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
