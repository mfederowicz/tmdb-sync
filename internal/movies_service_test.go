package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
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

func TestGetImages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/images", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":        550,
			"backdrops": []map[string]any{{"file_path": "/a.jpg"}},
			"posters":   []map[string]any{{"file_path": "/b.jpg"}},
		})
	})

	images, _, err := client.Movies.GetImages(context.Background(), 550, &uri.ImagesOptions{})
	if err != nil {
		t.Fatalf("GetImages() error = %v", err)
	}
	if images.ID != 550 || len(images.Backdrops) != 1 || len(images.Posters) != 1 {
		t.Errorf("GetImages() = %+v, want ID=550 with 1 backdrop and 1 poster", images)
	}
}

func TestGetKeywords(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/keywords", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":       550,
			"keywords": []map[string]any{{"id": 1701, "name": "hero"}},
		})
	})

	keywords, _, err := client.Movies.GetKeywords(context.Background(), 550)
	if err != nil {
		t.Fatalf("GetKeywords() error = %v", err)
	}
	if keywords.ID != 550 || len(keywords.Keywords) != 1 || keywords.Keywords[0].Name != "hero" {
		t.Errorf("GetKeywords() = %+v, want ID=550 with 1 keyword %q", keywords, "hero")
	}
}

func TestGetLatest(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/latest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":    1000000,
			"title": "brand new movie",
		})
	})

	movie, _, err := client.Movies.GetLatest(context.Background())
	if err != nil {
		t.Fatalf("GetLatest() error = %v", err)
	}
	if movie.ID != 1000000 || movie.Title != "brand new movie" {
		t.Errorf("GetLatest() = %+v, want ID=1000000 Title=%q", movie, "brand new movie")
	}
}

func TestGetMovieLists(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/lists", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "name": "list"}}
		json.NewEncoder(w).Encode(map[string]any{
			"id":            550,
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	lists, err := client.Movies.GetLists(context.Background(), 550, "", 0)
	if err != nil {
		t.Fatalf("GetLists() error = %v", err)
	}
	if len(lists) != 2 {
		t.Fatalf("len(lists) = %d, want 2", len(lists))
	}
}

func TestGetNowPlayingMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/now_playing", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "title": "movie"}}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	movies, err := client.Movies.GetNowPlayingMovies(context.Background(), 0)
	if err != nil {
		t.Fatalf("GetNowPlayingMovies() error = %v", err)
	}
	if len(movies) != 2 {
		t.Fatalf("len(movies) = %d, want 2", len(movies))
	}
}

func TestGetTopRatedMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/top_rated", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "title": "movie"}}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	movies, err := client.Movies.GetTopRatedMovies(context.Background(), 0)
	if err != nil {
		t.Fatalf("GetTopRatedMovies() error = %v", err)
	}
	if len(movies) != 2 {
		t.Fatalf("len(movies) = %d, want 2", len(movies))
	}
}

func TestGetTranslations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/translations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 550,
			"translations": []map[string]any{
				{"iso_3166_1": "US", "iso_639_1": "en", "name": "", "english_name": "English"},
			},
		})
	})

	translations, _, err := client.Movies.GetTranslations(context.Background(), 550)
	if err != nil {
		t.Fatalf("GetTranslations() error = %v", err)
	}
	if translations.ID != 550 || len(translations.Translations) != 1 {
		t.Errorf("GetTranslations() = %+v, want ID=550 with 1 translation", translations)
	}
}

func TestGetReleaseDates(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/release_dates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 550,
			"results": []map[string]any{
				{
					"iso_3166_1": "US",
					"release_dates": []map[string]any{
						{"certification": "R", "release_date": "1999-10-15T00:00:00.000Z", "type": 3},
					},
				},
			},
		})
	})

	dates, _, err := client.Movies.GetReleaseDates(context.Background(), 550)
	if err != nil {
		t.Fatalf("GetReleaseDates() error = %v", err)
	}
	if dates.ID != 550 || len(dates.Results) != 1 || dates.Results[0].Iso31661 != "US" {
		t.Errorf("GetReleaseDates() = %+v, want ID=550 with 1 result for US", dates)
	}
}

func TestGetReviews(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/reviews", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": "abc", "author": "someone", "content": "great movie"}}
		json.NewEncoder(w).Encode(map[string]any{
			"id":            550,
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	reviews, err := client.Movies.GetReviews(context.Background(), 550, "", 0)
	if err != nil {
		t.Fatalf("GetReviews() error = %v", err)
	}
	if len(reviews) != 2 {
		t.Fatalf("len(reviews) = %d, want 2", len(reviews))
	}
}

func TestGetSimilar(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/similar", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "title": "movie"}}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	movies, err := client.Movies.GetSimilar(context.Background(), 550, "", 0)
	if err != nil {
		t.Fatalf("GetSimilar() error = %v", err)
	}
	if len(movies) != 2 {
		t.Fatalf("len(movies) = %d, want 2", len(movies))
	}
}

func TestGetRecommendations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/550/recommendations", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "title": "movie"}}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	movies, err := client.Movies.GetRecommendations(context.Background(), 550, "", 0)
	if err != nil {
		t.Fatalf("GetRecommendations() error = %v", err)
	}
	if len(movies) != 2 {
		t.Fatalf("len(movies) = %d, want 2", len(movies))
	}
}
