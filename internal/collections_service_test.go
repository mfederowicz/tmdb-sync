package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
)

func TestGetCollection(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/collection/10", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   10,
			"name": "Star Wars Collection",
		})
	})

	collection, _, err := client.Collections.GetCollection(context.Background(), 10)
	if err != nil {
		t.Fatalf("GetCollection() error = %v", err)
	}
	if collection.ID != 10 || collection.Name != "Star Wars Collection" {
		t.Errorf("GetCollection() = %+v, want ID=10 Name=%q", collection, "Star Wars Collection")
	}
}

func TestGetCollectionImages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/collection/10/images", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("include_image_language"); got != "en,null" {
			t.Errorf("include_image_language query param = %q, want %q", got, "en,null")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":        10,
			"backdrops": []map[string]any{{"file_path": "/backdrop.jpg"}},
			"posters":   []map[string]any{{"file_path": "/poster.jpg"}},
		})
	})

	images, _, err := client.Collections.GetCollectionImages(context.Background(), 10, &uri.ImagesOptions{IncludeImageLanguage: "en,null"})
	if err != nil {
		t.Fatalf("GetCollectionImages() error = %v", err)
	}
	if len(images.Backdrops) != 1 || len(images.Posters) != 1 {
		t.Errorf("GetCollectionImages() = %+v, want one backdrop and one poster", images)
	}
}

func TestGetCollectionTranslations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/collection/10/translations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 10,
			"translations": []map[string]any{
				{"iso_3166_1": "US", "iso_639_1": "en", "name": "United States", "english_name": "English"},
			},
		})
	})

	translations, _, err := client.Collections.GetCollectionTranslations(context.Background(), 10)
	if err != nil {
		t.Fatalf("GetCollectionTranslations() error = %v", err)
	}
	if len(translations.Translations) != 1 || translations.Translations[0].Iso6391 != "en" {
		t.Errorf("GetCollectionTranslations() = %+v, want one translation with Iso6391=en", translations)
	}
}
