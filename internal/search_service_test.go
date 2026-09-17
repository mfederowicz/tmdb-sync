package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
)

func TestSearchCollections(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/collection", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "harry potter" {
			t.Errorf("query = %q, want %q", got, "harry potter")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 1241, "name": "Harry Potter Collection"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Search.SearchCollections(context.Background(), uri.SearchCollectionOptions{Query: "harry potter"}, 0)
	if err != nil {
		t.Fatalf("SearchCollections() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "Harry Potter Collection" {
		t.Errorf("SearchCollections() = %+v, want one result named Harry Potter Collection", results)
	}
}

func TestSearchCompanies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/company", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "pixar" {
			t.Errorf("query = %q, want %q", got, "pixar")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 3, "name": "Pixar"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Search.SearchCompanies(context.Background(), uri.SearchCompanyOptions{Query: "pixar"}, 0)
	if err != nil {
		t.Fatalf("SearchCompanies() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "Pixar" {
		t.Errorf("SearchCompanies() = %+v, want one result named Pixar", results)
	}
}

func TestSearchKeywords(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/keyword", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "space" {
			t.Errorf("query = %q, want %q", got, "space")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 9882, "name": "space"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Search.SearchKeywords(context.Background(), uri.SearchKeywordOptions{Query: "space"}, 0)
	if err != nil {
		t.Fatalf("SearchKeywords() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "space" {
		t.Errorf("SearchKeywords() = %+v, want one result named space", results)
	}
}

func TestSearchMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/movie", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "matrix" {
			t.Errorf("query = %q, want %q", got, "matrix")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 603, "title": "The Matrix"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Search.SearchMovies(context.Background(), uri.SearchMovieOptions{Query: "matrix"}, 0)
	if err != nil {
		t.Fatalf("SearchMovies() error = %v", err)
	}
	if len(results) != 1 || results[0].Title != "The Matrix" {
		t.Errorf("SearchMovies() = %+v, want one result titled The Matrix", results)
	}
}

func TestSearchMulti(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/multi", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "matrix" {
			t.Errorf("query = %q, want %q", got, "matrix")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 603, "title": "The Matrix", "media_type": "movie"},
				{"id": 6384, "name": "Keanu Reeves", "media_type": "person"},
			},
			"total_pages":   1,
			"total_results": 2,
		})
	})

	results, err := client.Search.SearchMulti(context.Background(), uri.SearchMultiOptions{Query: "matrix"}, 0)
	if err != nil {
		t.Fatalf("SearchMulti() error = %v", err)
	}
	if len(results) != 2 || results[0].MediaType != "movie" || results[1].MediaType != "person" {
		t.Errorf("SearchMulti() = %+v, want one movie and one person result", results)
	}
}

func TestSearchPeople(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/person", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "keanu" {
			t.Errorf("query = %q, want %q", got, "keanu")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 6384, "name": "Keanu Reeves"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Search.SearchPeople(context.Background(), uri.SearchPersonOptions{Query: "keanu"}, 0)
	if err != nil {
		t.Fatalf("SearchPeople() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "Keanu Reeves" {
		t.Errorf("SearchPeople() = %+v, want one result named Keanu Reeves", results)
	}
}

func TestSearchTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/tv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "breaking bad" {
			t.Errorf("query = %q, want %q", got, "breaking bad")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 1396, "name": "Breaking Bad"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Search.SearchTV(context.Background(), uri.SearchTVOptions{Query: "breaking bad"}, 0)
	if err != nil {
		t.Fatalf("SearchTV() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "Breaking Bad" {
		t.Errorf("SearchTV() = %+v, want one result named Breaking Bad", results)
	}
}
