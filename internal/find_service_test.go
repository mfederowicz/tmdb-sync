package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
)

func TestFind(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/find/tt0137523", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("external_source"); got != "imdb_id" {
			t.Errorf("external_source = %q, want %q", got, "imdb_id")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"movie_results": []map[string]any{
				{"id": 550, "title": "Fight Club"},
			},
			"person_results":     []map[string]any{},
			"tv_results":         []map[string]any{},
			"tv_episode_results": []map[string]any{},
			"tv_season_results":  []map[string]any{},
		})
	})

	results, _, err := client.Find.Find(context.Background(), "tt0137523", &uri.FindOptions{ExternalSource: "imdb_id"})
	if err != nil {
		t.Fatalf("Find() error = %v", err)
	}
	if len(results.MovieResults) != 1 || results.MovieResults[0].Title != "Fight Club" {
		t.Errorf("Find() = %+v, want one movie result titled Fight Club", results)
	}
}
