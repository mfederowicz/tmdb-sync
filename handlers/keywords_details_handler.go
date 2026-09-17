package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// KeywordsDetailsHandler handles `keywords -a details -i <keyword_id>`.
type KeywordsDetailsHandler struct {
	KeywordID string
}

// Handle fetches details for a single keyword.
func (h KeywordsDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	keyword, _, err := client.Keywords.GetKeyword(ctx, h.KeywordID)
	if err != nil {
		return nil, err
	}
	return keyword, nil
}
