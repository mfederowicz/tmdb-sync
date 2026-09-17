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
