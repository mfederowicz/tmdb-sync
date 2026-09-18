package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
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

func TestTVEpisodesGetCredits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/credits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"cast": []map[string]any{
				{"id": 22970, "name": "Sean Bean"},
			},
			"crew": []map[string]any{},
		})
	})

	credits, _, err := client.TVEpisodes.GetCredits(context.Background(), 1399, 1, 1, "")
	if err != nil {
		t.Fatalf("GetCredits() error = %v", err)
	}
	if len(credits.Cast) != 1 || credits.Cast[0].Name != "Sean Bean" {
		t.Errorf("GetCredits() = %+v, want Cast[0].Name=Sean Bean", credits)
	}
}

func TestTVEpisodesGetExternalIDs(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/external_ids", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":      63056,
			"imdb_id": "tt1480055",
		})
	})

	ids, _, err := client.TVEpisodes.GetExternalIDs(context.Background(), 1399, 1, 1)
	if err != nil {
		t.Fatalf("GetExternalIDs() error = %v", err)
	}
	if ids.ID != 63056 || ids.IMDbID != "tt1480055" {
		t.Errorf("GetExternalIDs() = %+v, want ID=63056 IMDbID=tt1480055", ids)
	}
}

func TestTVEpisodesGetImages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/images", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"stills": []map[string]any{
				{"file_path": "/still.jpg"},
			},
		})
	})

	images, _, err := client.TVEpisodes.GetImages(context.Background(), 1399, 1, 1, &uri.ImagesOptions{})
	if err != nil {
		t.Fatalf("GetImages() error = %v", err)
	}
	if len(images.Stills) != 1 || images.Stills[0].FilePath != "/still.jpg" {
		t.Errorf("GetImages() = %+v, want Stills[0].FilePath=/still.jpg", images)
	}
}

func TestTVEpisodesGetTranslations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/translations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 63056,
			"translations": []map[string]any{
				{"iso_639_1": "en"},
			},
		})
	})

	translations, _, err := client.TVEpisodes.GetTranslations(context.Background(), 1399, 1, 1)
	if err != nil {
		t.Fatalf("GetTranslations() error = %v", err)
	}
	if len(translations.Translations) != 1 || translations.Translations[0].Iso6391 != "en" {
		t.Errorf("GetTranslations() = %+v, want Translations[0].Iso6391=en", translations)
	}
}

func TestTVEpisodesGetVideos(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/videos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 63056,
			"results": []map[string]any{
				{"key": "abc123", "site": "YouTube"},
			},
		})
	})

	videos, _, err := client.TVEpisodes.GetVideos(context.Background(), 1399, 1, 1, "")
	if err != nil {
		t.Fatalf("GetVideos() error = %v", err)
	}
	if len(videos.Results) != 1 || videos.Results[0].Key != "abc123" {
		t.Errorf("GetVideos() = %+v, want Results[0].Key=abc123", videos)
	}
}

func TestTVEpisodesAddRating(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/rating", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "sess" {
			t.Errorf("session_id = %s, want sess", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.TVEpisodes.AddRating(context.Background(), 1399, 1, 1, "sess", "", 8.5)
	if err != nil {
		t.Fatalf("AddRating() error = %v", err)
	}
	if !status.Success {
		t.Errorf("AddRating() = %+v, want Success=true", status)
	}
}

func TestTVEpisodesDeleteRating(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/season/1/episode/1/rating", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.TVEpisodes.DeleteRating(context.Background(), 1399, 1, 1, "sess")
	if err != nil {
		t.Fatalf("DeleteRating() error = %v", err)
	}
	if !status.Success {
		t.Errorf("DeleteRating() = %+v, want Success=true", status)
	}
}
