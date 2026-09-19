package str

// RequestTokenV4 is returned by POST /4/auth/request_token.
type RequestTokenV4 struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	RequestToken  string `json:"request_token"`
}
