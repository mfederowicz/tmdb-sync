package cmds

import (
	"testing"

	"github.com/mfederowicz/tmdb-sync/str"
)

func TestResolveRatingGuestSessionID(t *testing.T) {
	account := &str.Session{SessionID: "acc"}
	guest := &str.GuestSession{GuestSessionID: "cached"}

	tests := []struct {
		name    string
		action  string
		flagID  string
		guest   bool
		options *str.Options
		want    string
		wantErr bool
	}{
		{"other action", "details", "x", true, &str.Options{GuestSession: guest}, "", false},
		{"flag id wins", "add-rating", "flag", true, &str.Options{Session: account, GuestSession: guest}, "flag", false},
		{"guest overrides account session", "add-rating", "", true, &str.Options{Session: account, GuestSession: guest}, "cached", false},
		{"guest without cache", "add-rating", "", true, &str.Options{Session: account}, "", true},
		{"account session wins by default", "add-rating", "", false, &str.Options{Session: account, GuestSession: guest}, "", false},
		{"cached guest fallback", "add-rating", "", false, &str.Options{GuestSession: guest}, "cached", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolveRatingGuestSessionID(tt.action, tt.flagID, tt.guest, tt.options)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
