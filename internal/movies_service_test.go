package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

func TestGetPopularMovies_WalksEveryPage(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/popular", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "title": "movie"}}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          pageNum,
			"results":       results,
			"total_pages":   3,
			"total_results": 3,
		})
	})

	movies, err := client.Movies.GetPopularMovies(context.Background(), 0)
	if err != nil {
		t.Fatalf("GetPopularMovies() error = %v", err)
	}
	if len(movies) != 3 {
		t.Fatalf("len(movies) = %d, want 3", len(movies))
	}
}

func TestGetPopularMovies_RespectsPagesLimit(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/popular", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"results":       []map[string]any{{"id": 1, "title": "movie"}},
			"total_pages":   50,
			"total_results": 50,
		})
	})

	movies, err := client.Movies.GetPopularMovies(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetPopularMovies() error = %v", err)
	}
	if len(movies) != 1 {
		t.Fatalf("len(movies) = %d, want 1 (pagesLimit=1)", len(movies))
	}
}

func TestGetAccountStates(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/account_states", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "sess" {
			t.Errorf("session_id = %s, want sess", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":        550,
			"favorite":  true,
			"rated":     map[string]any{"value": 8},
			"watchlist": false,
		})
	})

	states, _, err := client.Movies.GetAccountStates(context.Background(), 550, "sess")
	if err != nil {
		t.Fatalf("GetAccountStates() error = %v", err)
	}
	if states.ID != 550 || !states.Favorite || states.Watchlist {
		t.Errorf("GetAccountStates() = %+v, want ID=550 Favorite=true Watchlist=false", states)
	}
}

func TestGetAlternativeTitles(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/alternative_titles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("country"); got != "US" {
			t.Errorf("country = %s, want US", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 550,
			"titles": []map[string]any{
				{"iso_3166_1": "US", "title": "Fight Club", "type": ""},
			},
		})
	})

	titles, _, err := client.Movies.GetAlternativeTitles(context.Background(), 550, "US")
	if err != nil {
		t.Fatalf("GetAlternativeTitles() error = %v", err)
	}
	if titles.ID != 550 || len(titles.Titles) != 1 || titles.Titles[0].Title != "Fight Club" {
		t.Errorf("GetAlternativeTitles() = %+v, want ID=550 with 1 title %q", titles, "Fight Club")
	}
}

func TestGetCredits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/credits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 550,
			"cast": []map[string]any{
				{"id": 819, "name": "Edward Norton", "character": "The Narrator", "credit_id": "abc"},
			},
			"crew": []map[string]any{
				{"id": 7467, "name": "David Fincher", "job": "Director", "credit_id": "def"},
			},
		})
	})

	credits, _, err := client.Movies.GetCredits(context.Background(), 550, "")
	if err != nil {
		t.Fatalf("GetCredits() error = %v", err)
	}
	if credits.ID != 550 || len(credits.Cast) != 1 || len(credits.Crew) != 1 {
		t.Errorf("GetCredits() = %+v, want ID=550 with 1 cast and 1 crew", credits)
	}
}

func TestGetExternalIDs(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/external_ids", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":      550,
			"imdb_id": "tt0137523",
		})
	})

	ids, _, err := client.Movies.GetExternalIDs(context.Background(), 550)
	if err != nil {
		t.Fatalf("GetExternalIDs() error = %v", err)
	}
	if ids.ID != 550 || ids.ImdbID != "tt0137523" {
		t.Errorf("GetExternalIDs() = %+v, want ID=550 ImdbID=%q", ids, "tt0137523")
	}
}
