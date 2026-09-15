package str

// GuestSession is returned by GET /authentication/guest_session/new.
type GuestSession struct {
	Success        bool   `json:"success"`
	GuestSessionID string `json:"guest_session_id"`
	ExpiresAt      string `json:"expires_at"`
}
