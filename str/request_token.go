package str

// RequestToken is returned by GET /authentication/token/new.
type RequestToken struct {
	Success      bool   `json:"success"`
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
}
