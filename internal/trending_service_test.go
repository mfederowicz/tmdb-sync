package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestTrendingAll(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/trending/all/day", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 550, "media_type": "movie", "title": "Fight Club"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Trending.TrendingAll(context.Background(), "day", 0)
	if err != nil {
		t.Fatalf("TrendingAll() error = %v", err)
	}
	if len(results) != 1 || results[0].Title != "Fight Club" {
		t.Errorf("TrendingAll() = %+v, want one result titled Fight Club", results)
	}
}

func TestTrendingMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/trending/movie/week", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 550, "title": "Fight Club"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Trending.TrendingMovies(context.Background(), "week", 0)
	if err != nil {
		t.Fatalf("TrendingMovies() error = %v", err)
	}
	if len(results) != 1 || results[0].Title != "Fight Club" {
		t.Errorf("TrendingMovies() = %+v, want one result titled Fight Club", results)
	}
}

func TestTrendingTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/trending/tv/day", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 1399, "name": "Game of Thrones"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Trending.TrendingTV(context.Background(), "day", 0)
	if err != nil {
		t.Fatalf("TrendingTV() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "Game of Thrones" {
		t.Errorf("TrendingTV() = %+v, want one result named Game of Thrones", results)
	}
}

func TestTrendingPeople(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/trending/person/week", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 287, "name": "Brad Pitt"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Trending.TrendingPeople(context.Background(), "week", 0)
	if err != nil {
		t.Fatalf("TrendingPeople() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "Brad Pitt" {
		t.Errorf("TrendingPeople() = %+v, want one result named Brad Pitt", results)
	}
}
