package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// NetworksAlternativeNamesHandler handles `networks -a alternative-names -i <network_id>`.
type NetworksAlternativeNamesHandler struct {
	NetworkID int64
}

// Handle fetches the alternative names for a single network.
func (h NetworksAlternativeNamesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	names, _, err := client.Networks.GetNetworkAlternativeNames(ctx, h.NetworkID)
	if err != nil {
		return nil, err
	}
	return names, nil
}
