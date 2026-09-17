package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetTVSeason(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":            3624,
			"name":          "Season 1",
			"season_number": 1,
		})
	})

	season, _, err := client.TVSeasons.GetTVSeason(context.Background(), 1399, 1)
	if err != nil {
		t.Fatalf("GetTVSeason() error = %v", err)
	}
	if season.ID != 3624 || season.Name != "Season 1" || season.SeasonNumber != 1 {
		t.Errorf("GetTVSeason() = %+v, want ID=3624 Name=%q SeasonNumber=1", season, "Season 1")
	}
}

func TestTVSeasonsGetAccountStates(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/account_states", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "sess" {
			t.Errorf("session_id = %s, want sess", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 3624,
			"results": []map[string]any{
				{"id": 63056, "episode_number": 1, "rated": false},
			},
		})
	})

	states, _, err := client.TVSeasons.GetAccountStates(context.Background(), 1399, 1, "sess")
	if err != nil {
		t.Fatalf("GetAccountStates() error = %v", err)
	}
	if states.ID != 3624 || len(states.Results) != 1 || states.Results[0].EpisodeNumber != 1 {
		t.Errorf("GetAccountStates() = %+v, want ID=3624 with 1 result EpisodeNumber=1", states)
	}
}
