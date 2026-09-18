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

func TestTVEpisodesGetAccountStates(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/account_states", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "sess" {
			t.Errorf("session_id = %s, want sess", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":    63056,
			"rated": false,
		})
	})

	states, _, err := client.TVEpisodes.GetAccountStates(context.Background(), 1399, 1, 1, "sess")
	if err != nil {
		t.Fatalf("GetAccountStates() error = %v", err)
	}
	if states.ID != 63056 {
		t.Errorf("GetAccountStates() = %+v, want ID=63056", states)
	}
}
