package str

// Session is returned by GET /authentication/session/new and persisted to disk.
type Session struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}

// Valid reports whether a session was actually created.
func (s *Session) Valid() bool {
	return s != nil && s.Success && s.SessionID != ""
}

// AccessTokenV4 is returned by POST /4/auth/access_token and persisted to disk.
// AccountID is the v4 account_object_id used in v4 account paths.
type AccessTokenV4 struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	AccessToken   string `json:"access_token"`
	AccountID     string `json:"account_id"`
}

// Valid reports whether a user access token was actually created.
func (a *AccessTokenV4) Valid() bool {
	return a != nil && a.Success && a.AccessToken != ""
}
