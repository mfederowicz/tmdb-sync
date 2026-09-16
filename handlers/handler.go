// Package handlers used to handle module actions
package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// Handler is implemented by every module action.
type Handler interface {
	Handle(ctx context.Context, client *internal.Client) (any, error)
}
