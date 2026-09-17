package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// NetworksImagesHandler handles `networks -a images -i <network_id>`.
type NetworksImagesHandler struct {
	NetworkID int64
}

// Handle fetches the images for a single network.
func (h NetworksImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.Networks.GetNetworkImages(ctx, h.NetworkID)
	if err != nil {
		return nil, err
	}
	return images, nil
}
