package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
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

func TestTVSeasonsGetAggregateCredits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/aggregate_credits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   3624,
			"cast": []map[string]any{{"id": 22970, "name": "Emilia Clarke"}},
			"crew": []map[string]any{{"id": 9813, "name": "David Benioff"}},
		})
	})

	credits, _, err := client.TVSeasons.GetAggregateCredits(context.Background(), 1399, 1, "")
	if err != nil {
		t.Fatalf("GetAggregateCredits() error = %v", err)
	}
	if credits.ID != 3624 || len(credits.Cast) != 1 || len(credits.Crew) != 1 {
		t.Errorf("GetAggregateCredits() = %+v, want ID=3624 with 1 cast and 1 crew", credits)
	}
}

func TestTVSeasonsGetCredits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/credits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   3624,
			"cast": []map[string]any{{"id": 22970, "name": "Emilia Clarke"}},
			"crew": []map[string]any{{"id": 9813, "name": "David Benioff"}},
		})
	})

	credits, _, err := client.TVSeasons.GetCredits(context.Background(), 1399, 1, "")
	if err != nil {
		t.Fatalf("GetCredits() error = %v", err)
	}
	if credits.ID != 3624 || len(credits.Cast) != 1 || len(credits.Crew) != 1 {
		t.Errorf("GetCredits() = %+v, want ID=3624 with 1 cast and 1 crew", credits)
	}
}

func TestTVSeasonsGetExternalIDs(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/external_ids", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":          3624,
			"tvdb_id":     3436,
			"wikidata_id": "Q3629848",
		})
	})

	ids, _, err := client.TVSeasons.GetExternalIDs(context.Background(), 1399, 1)
	if err != nil {
		t.Fatalf("GetExternalIDs() error = %v", err)
	}
	if ids.ID != 3624 || ids.TvdbID != 3436 || ids.WikidataID != "Q3629848" {
		t.Errorf("GetExternalIDs() = %+v, want ID=3624 TvdbID=3436 WikidataID=Q3629848", ids)
	}
}

func TestTVSeasonsGetImages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/images", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 3624,
			"posters": []map[string]any{
				{"file_path": "/poster.jpg"},
			},
		})
	})

	images, _, err := client.TVSeasons.GetImages(context.Background(), 1399, 1, &uri.ImagesOptions{})
	if err != nil {
		t.Fatalf("GetImages() error = %v", err)
	}
	if images.ID != 3624 || len(images.Posters) != 1 {
		t.Errorf("GetImages() = %+v, want ID=3624 with 1 poster", images)
	}
}

func TestTVSeasonsGetTranslations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/translations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 3624,
			"translations": []map[string]any{
				{"iso_3166_1": "US", "iso_639_1": "en", "name": "", "english_name": "English"},
			},
		})
	})

	translations, _, err := client.TVSeasons.GetTranslations(context.Background(), 1399, 1)
	if err != nil {
		t.Fatalf("GetTranslations() error = %v", err)
	}
	if translations.ID != 3624 || len(translations.Translations) != 1 {
		t.Errorf("GetTranslations() = %+v, want ID=3624 with 1 translation", translations)
	}
}
