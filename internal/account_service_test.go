package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/str"
)

func TestGetAccountDetails(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":       123,
			"username": "testuser",
		})
	})

	account, _, err := client.Account.GetDetails(context.Background(), 123, "abc123")
	if err != nil {
		t.Fatalf("GetDetails() error = %v", err)
	}
	if account.ID != 123 || account.Username != "testuser" {
		t.Errorf("GetDetails() = %+v, want ID=123 Username=testuser", account)
	}
}

func TestGetAccountDetails_SelfResolve(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":       456,
			"username": "selfuser",
		})
	})

	account, _, err := client.Account.GetDetails(context.Background(), 0, "abc123")
	if err != nil {
		t.Fatalf("GetDetails() error = %v", err)
	}
	if account.ID != 456 || account.Username != "selfuser" {
		t.Errorf("GetDetails() = %+v, want ID=456 Username=selfuser", account)
	}
}

func TestAddToWatchlist(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/watchlist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		var body str.AccountWatchlistRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.MediaType != "movie" || body.MediaID != 550 || !body.Watchlist {
			t.Errorf("body = %+v, want MediaType=movie MediaID=550 Watchlist=true", body)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Account.AddToWatchlist(context.Background(), 123, "abc123", &str.AccountWatchlistRequest{
		MediaType: "movie",
		MediaID:   550,
		Watchlist: true,
	})
	if err != nil {
		t.Fatalf("AddToWatchlist() error = %v", err)
	}
	if !status.Success {
		t.Errorf("AddToWatchlist() Success = %v, want true", status.Success)
	}
}

func TestAddRemoveFavorite(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/favorite", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		var body str.AccountFavoriteRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.MediaType != "movie" || body.MediaID != 550 || !body.Favorite {
			t.Errorf("body = %+v, want MediaType=movie MediaID=550 Favorite=true", body)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Account.AddRemoveFavorite(context.Background(), 123, "abc123", &str.AccountFavoriteRequest{
		MediaType: "movie",
		MediaID:   550,
		Favorite:  true,
	})
	if err != nil {
		t.Fatalf("AddRemoveFavorite() error = %v", err)
	}
	if !status.Success {
		t.Errorf("AddRemoveFavorite() Success = %v, want true", status.Success)
	}
}

func TestGetFavoriteMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/favorite/movies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 550, "title": "Fight Club"},
			},
		})
	})

	movies, err := client.Account.GetFavoriteMovies(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetFavoriteMovies() error = %v", err)
	}
	if len(movies) != 1 || movies[0].ID != 550 {
		t.Errorf("GetFavoriteMovies() = %+v, want one movie with ID=550", movies)
	}
}

func TestGetFavoriteTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/favorite/tv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 1396, "name": "Breaking Bad"},
			},
		})
	})

	shows, err := client.Account.GetFavoriteTV(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetFavoriteTV() error = %v", err)
	}
	if len(shows) != 1 || shows[0].ID != 1396 {
		t.Errorf("GetFavoriteTV() = %+v, want one show with ID=1396", shows)
	}
}

func TestGetLists(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/lists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 42, "name": "My List", "item_count": 5},
			},
		})
	})

	lists, err := client.Account.GetLists(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetLists() error = %v", err)
	}
	if len(lists) != 1 || lists[0].ID != 42 || lists[0].Name != "My List" {
		t.Errorf("GetLists() = %+v, want one list with ID=42 Name=\"My List\"", lists)
	}
}

func TestGetRatedMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/rated/movies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 550, "title": "Fight Club", "rating": 8.5},
			},
		})
	})

	movies, err := client.Account.GetRatedMovies(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetRatedMovies() error = %v", err)
	}
	if len(movies) != 1 || movies[0].ID != 550 || movies[0].Rating != 8.5 {
		t.Errorf("GetRatedMovies() = %+v, want one movie with ID=550 Rating=8.5", movies)
	}
}

func TestGetRatedTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/rated/tv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 1396, "name": "Breaking Bad", "rating": 9.5},
			},
		})
	})

	shows, err := client.Account.GetRatedTV(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetRatedTV() error = %v", err)
	}
	if len(shows) != 1 || shows[0].ID != 1396 || shows[0].Rating != 9.5 {
		t.Errorf("GetRatedTV() = %+v, want one show with ID=1396 Rating=9.5", shows)
	}
}

func TestGetRatedTVEpisodes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/rated/tv/episodes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 64782, "name": "The Workplace Proximity", "show_id": 1418, "rating": 8.0},
			},
		})
	})

	episodes, err := client.Account.GetRatedTVEpisodes(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetRatedTVEpisodes() error = %v", err)
	}
	if len(episodes) != 1 || episodes[0].ID != 64782 || episodes[0].ShowID != 1418 || episodes[0].Rating != 8.0 {
		t.Errorf("GetRatedTVEpisodes() = %+v, want one episode with ID=64782 ShowID=1418 Rating=8.0", episodes)
	}
}

func TestGetWatchlistMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/watchlist/movies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 550, "title": "Fight Club"},
			},
		})
	})

	movies, err := client.Account.GetWatchlistMovies(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetWatchlistMovies() error = %v", err)
	}
	if len(movies) != 1 || movies[0].ID != 550 {
		t.Errorf("GetWatchlistMovies() = %+v, want one movie with ID=550", movies)
	}
}

func TestGetWatchlistTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/watchlist/tv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 1396, "name": "Breaking Bad"},
			},
		})
	})

	shows, err := client.Account.GetWatchlistTV(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetWatchlistTV() error = %v", err)
	}
	if len(shows) != 1 || shows[0].ID != 1396 {
		t.Errorf("GetWatchlistTV() = %+v, want one show with ID=1396", shows)
	}
}

func TestAccountGetListsV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/account/acc123/lists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user access token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		page := r.URL.Query().Get("page")
		id := 1
		if page == "2" {
			id = 2
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          id,
			"total_pages":   2,
			"total_results": 2,
			"results": []map[string]any{
				{"id": id, "name": "list", "public": 1, "number_of_items": 3, "created_at": "2024-01-01 00:00:00 UTC"},
			},
		})
	})

	lists, err := client.Account.GetListsV4(context.Background(), "usertok", "acc123", 0)
	if err != nil {
		t.Fatalf("GetListsV4() error = %v", err)
	}
	if len(lists) != 2 || lists[0].ID != 1 || lists[1].ID != 2 || lists[0].NumberOfItems != 3 || lists[0].Public != 1 {
		t.Errorf("lists = %+v, want two pages of results", lists)
	}
}

func TestAccountGetListsV4_PagesLimit(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	calls := 0
	mux.HandleFunc("/account/acc123/lists", func(w http.ResponseWriter, _ *http.Request) {
		calls++
		json.NewEncoder(w).Encode(map[string]any{"page": calls, "total_pages": 5, "results": []map[string]any{{"id": calls}}})
	})

	lists, err := client.Account.GetListsV4(context.Background(), "usertok", "acc123", 1)
	if err != nil {
		t.Fatalf("GetListsV4() error = %v", err)
	}
	if calls != 1 || len(lists) != 1 {
		t.Errorf("calls = %d, lists = %d, want 1 and 1", calls, len(lists))
	}
}

func TestAccountGetListsV4_NeedsReadToken(t *testing.T) {
	client, _, teardown := setupV4()
	defer teardown()

	if _, err := client.Account.GetListsV4(context.Background(), "usertok", "acc123", 1); err == nil {
		t.Error("GetListsV4() error = nil, want error without read_access_token")
	}
}

func TestAccountGetFavoriteMoviesV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/account/acc123/movie/favorites", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user access token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		page := 1
		if r.URL.Query().Get("page") == "2" {
			page = 2
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":        page,
			"total_pages": 2,
			"results":     []map[string]any{{"id": 550 + page, "title": "Fight Club"}},
		})
	})

	movies, err := client.Account.GetFavoriteMoviesV4(context.Background(), "usertok", "acc123", 0)
	if err != nil {
		t.Fatalf("GetFavoriteMoviesV4() error = %v", err)
	}
	if len(movies) != 2 || movies[0].ID != 551 || movies[1].ID != 552 {
		t.Errorf("movies = %+v, want two pages of results", movies)
	}
}

func TestAccountGetFavoriteTVV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/account/acc123/tv/favorites", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user access token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":        1,
			"total_pages": 1,
			"results":     []map[string]any{{"id": 1399, "name": "Game of Thrones"}},
		})
	})

	shows, err := client.Account.GetFavoriteTVV4(context.Background(), "usertok", "acc123", 0)
	if err != nil {
		t.Fatalf("GetFavoriteTVV4() error = %v", err)
	}
	if len(shows) != 1 || shows[0].ID != 1399 {
		t.Errorf("shows = %+v, want Game of Thrones", shows)
	}
}

func TestAccountGetRatedMoviesV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/account/acc123/movie/rated", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user access token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":        1,
			"total_pages": 1,
			"results": []map[string]any{{
				"id": 550, "title": "Fight Club",
				"account_rating": map[string]any{"created_at": "2024-01-01 00:00:00 UTC", "value": 9},
			}},
		})
	})

	movies, err := client.Account.GetRatedMoviesV4(context.Background(), "usertok", "acc123", 0)
	if err != nil {
		t.Fatalf("GetRatedMoviesV4() error = %v", err)
	}
	if len(movies) != 1 || movies[0].ID != 550 || movies[0].AccountRating == nil || movies[0].AccountRating.Value != 9 {
		t.Errorf("movies = %+v, want Fight Club rated 9", movies)
	}
}

func TestAccountGetRatedTVV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/account/acc123/tv/rated", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user access token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":        1,
			"total_pages": 1,
			"results": []map[string]any{{
				"id": 1399, "name": "Game of Thrones",
				"account_rating": map[string]any{"created_at": "2024-01-01 00:00:00 UTC", "value": 10},
			}},
		})
	})

	shows, err := client.Account.GetRatedTVV4(context.Background(), "usertok", "acc123", 0)
	if err != nil {
		t.Fatalf("GetRatedTVV4() error = %v", err)
	}
	if len(shows) != 1 || shows[0].ID != 1399 || shows[0].AccountRating == nil || shows[0].AccountRating.Value != 10 {
		t.Errorf("shows = %+v, want Game of Thrones rated 10", shows)
	}
}

func TestAccountGetRecommendedMoviesV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/account/acc123/movie/recommendations", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user access token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":        1,
			"total_pages": 1,
			"results":     []map[string]any{{"id": 680, "title": "Pulp Fiction"}},
		})
	})

	movies, err := client.Account.GetRecommendedMoviesV4(context.Background(), "usertok", "acc123", 0)
	if err != nil {
		t.Fatalf("GetRecommendedMoviesV4() error = %v", err)
	}
	if len(movies) != 1 || movies[0].ID != 680 {
		t.Errorf("movies = %+v, want Pulp Fiction", movies)
	}
}
