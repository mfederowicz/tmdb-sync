package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CreditsDetailsHandler handles `credits -a details -i <credit_id>`.
type CreditsDetailsHandler struct {
	CreditID string
}

// Handle fetches details for a single credit.
func (h CreditsDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credit, _, err := client.Credits.GetCredit(ctx, h.CreditID)
	if err != nil {
		return nil, err
	}
	return credit, nil
}
