package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// TVEpisodeGroupService handles communication with the /tv/episode_group
// endpoint of the TMDB API.
type TVEpisodeGroupService Service

// GetTVEpisodeGroup fetches details for a single TV episode group by id.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-group-details
func (s *TVEpisodeGroupService) GetTVEpisodeGroup(ctx context.Context, id string) (*str.TVEpisodeGroupDetails, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/episode_group/%s", id), nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(str.TVEpisodeGroupDetails)
	resp, err := s.client.Do(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}
