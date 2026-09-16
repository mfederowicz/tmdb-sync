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
