package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetMovieList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/genre/movie/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"genres": []map[string]any{
				{"id": 28, "name": "Action"},
			},
		})
	})

	genres, _, err := client.Genre.GetMovieList(context.Background())
	if err != nil {
		t.Fatalf("GetMovieList() error = %v", err)
	}
	if len(genres.Genres) != 1 || genres.Genres[0].Name != "Action" {
		t.Errorf("Genres = %+v, want one entry with Name=Action", genres.Genres)
	}
}

func TestGetTVList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/genre/tv/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"genres": []map[string]any{
				{"id": 10759, "name": "Action & Adventure"},
			},
		})
	})

	genres, _, err := client.Genre.GetTVList(context.Background())
	if err != nil {
		t.Fatalf("GetTVList() error = %v", err)
	}
	if len(genres.Genres) != 1 || genres.Genres[0].Name != "Action & Adventure" {
		t.Errorf("Genres = %+v, want one entry with Name=Action & Adventure", genres.Genres)
	}
}
