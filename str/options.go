package str

// Options represents runtime app options derived from config + flags.
type Options struct {
	Headers      map[string]any
	Session      *Session
	Account      *Account
	GuestSession *GuestSession
	Verbose      bool
}
