package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// NetworksDetailsHandler handles `networks -a details -i <network_id>`.
type NetworksDetailsHandler struct {
	NetworkID int64
}

// Handle fetches details for a single network.
func (h NetworksDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	network, _, err := client.Networks.GetNetwork(ctx, h.NetworkID)
	if err != nil {
		return nil, err
	}
	return network, nil
}
