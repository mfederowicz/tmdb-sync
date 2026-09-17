package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// NetworksService handles communication with the /network endpoints of the TMDB API.
type NetworksService Service

// GetNetwork fetches details for a single network by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/network-details
func (s *NetworksService) GetNetwork(ctx context.Context, networkID int64) (*str.Network, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("network/%d", networkID), nil)
	if err != nil {
		return nil, nil, err
	}

	network := new(str.Network)
	resp, err := s.client.Do(ctx, req, network)
	if err != nil {
		return nil, resp, err
	}

	return network, resp, nil
}

// GetNetworkAlternativeNames fetches the alternative names for a single network by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/network-alternative-names
func (s *NetworksService) GetNetworkAlternativeNames(ctx context.Context, networkID int64) (*str.NetworkAlternativeNames, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("network/%d/alternative_names", networkID), nil)
	if err != nil {
		return nil, nil, err
	}

	names := new(str.NetworkAlternativeNames)
	resp, err := s.client.Do(ctx, req, names)
	if err != nil {
		return nil, resp, err
	}

	return names, resp, nil
}
