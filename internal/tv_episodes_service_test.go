package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetTVEpisode(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":             63056,
			"name":           "Winter Is Coming",
			"episode_number": 1,
			"season_number":  1,
		})
	})

	episode, _, err := client.TVEpisodes.GetTVEpisode(context.Background(), 1399, 1, 1)
	if err != nil {
		t.Fatalf("GetTVEpisode() error = %v", err)
	}
	if episode.ID != 63056 || episode.Name != "Winter Is Coming" || episode.EpisodeNumber != 1 || episode.SeasonNumber != 1 {
		t.Errorf("GetTVEpisode() = %+v, want ID=63056 Name=%q EpisodeNumber=1 SeasonNumber=1", episode, "Winter Is Coming")
	}
}
