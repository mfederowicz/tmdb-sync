package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// WatchProvidersService handles communication with the /watch/providers
// endpoints of the TMDB API.
type WatchProvidersService Service

// GetAvailableRegions fetches the list of regions TMDB has watch provider
// (OTT/streaming) data for.
//
// Api docs: https://developer.themoviedb.org/reference/watch-providers-available-regions
func (s *WatchProvidersService) GetAvailableRegions(ctx context.Context, language string) (*str.WatchProviderRegions, *str.Response, error) {
	urlStr, err := uri.AddQuery("watch/providers/regions", &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	regions := new(str.WatchProviderRegions)
	resp, err := s.client.Do(ctx, req, regions)
	if err != nil {
		return nil, resp, err
	}

	return regions, resp, nil
}
