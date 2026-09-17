package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
)

func TestGetTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   1399,
			"name": "Game of Thrones",
		})
	})

	tv, _, err := client.TV.GetTV(context.Background(), 1399)
	if err != nil {
		t.Fatalf("GetTV() error = %v", err)
	}
	if tv.ID != 1399 || tv.Name != "Game of Thrones" {
		t.Errorf("GetTV() = %+v, want ID=1399 Name=%q", tv, "Game of Thrones")
	}
}

func TestTVGetAccountStates(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/account_states", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "sess" {
			t.Errorf("session_id = %s, want sess", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":        1399,
			"favorite":  true,
			"rated":     map[string]any{"value": 8},
			"watchlist": false,
		})
	})

	states, _, err := client.TV.GetAccountStates(context.Background(), 1399, "sess")
	if err != nil {
		t.Fatalf("GetAccountStates() error = %v", err)
	}
	if states.ID != 1399 || !states.Favorite || states.Watchlist {
		t.Errorf("GetAccountStates() = %+v, want ID=1399 Favorite=true Watchlist=false", states)
	}
}

func TestGetAggregateCredits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/aggregate_credits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1399,
			"cast": []map[string]any{
				{
					"id":   1223792,
					"name": "Emilia Clarke",
					"roles": []map[string]any{
						{"credit_id": "abc", "character": "Daenerys Targaryen", "episode_count": 62},
					},
					"total_episode_count": 62,
				},
			},
			"crew": []map[string]any{
				{
					"id":   1,
					"name": "David Benioff",
					"jobs": []map[string]any{
						{"credit_id": "def", "job": "Executive Producer", "episode_count": 73},
					},
					"total_episode_count": 73,
				},
			},
		})
	})

	credits, _, err := client.TV.GetAggregateCredits(context.Background(), 1399, "")
	if err != nil {
		t.Fatalf("GetAggregateCredits() error = %v", err)
	}
	if credits.ID != 1399 || len(credits.Cast) != 1 || len(credits.Cast[0].Roles) != 1 || len(credits.Crew) != 1 || len(credits.Crew[0].Jobs) != 1 {
		t.Errorf("GetAggregateCredits() = %+v, want ID=1399 with 1 cast (1 role) and 1 crew (1 job)", credits)
	}
}

func TestGetTVAlternativeTitles(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/alternative_titles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1399,
			"results": []map[string]any{
				{"iso_3166_1": "US", "title": "Game of Thrones", "type": ""},
			},
		})
	})

	titles, _, err := client.TV.GetAlternativeTitles(context.Background(), 1399)
	if err != nil {
		t.Fatalf("GetAlternativeTitles() error = %v", err)
	}
	if titles.ID != 1399 || len(titles.Results) != 1 || titles.Results[0].Title != "Game of Thrones" {
		t.Errorf("GetAlternativeTitles() = %+v, want ID=1399 with 1 title %q", titles, "Game of Thrones")
	}
}

func TestGetContentRatings(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/content_ratings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1399,
			"results": []map[string]any{
				{"iso_3166_1": "US", "rating": "TV-MA"},
			},
		})
	})

	ratings, _, err := client.TV.GetContentRatings(context.Background(), 1399)
	if err != nil {
		t.Fatalf("GetContentRatings() error = %v", err)
	}
	if ratings.ID != 1399 || len(ratings.Results) != 1 || ratings.Results[0].Rating != "TV-MA" {
		t.Errorf("GetContentRatings() = %+v, want ID=1399 with 1 rating %q", ratings, "TV-MA")
	}
}

func TestGetTVCredits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/credits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1399,
			"cast": []map[string]any{
				{"id": 22970, "name": "Peter Dinklage", "character": "Tyrion Lannister", "credit_id": "abc"},
			},
			"crew": []map[string]any{
				{"id": 1, "name": "David Benioff", "job": "Executive Producer", "credit_id": "def"},
			},
		})
	})

	credits, _, err := client.TV.GetCredits(context.Background(), 1399, "")
	if err != nil {
		t.Fatalf("GetCredits() error = %v", err)
	}
	if credits.ID != 1399 || len(credits.Cast) != 1 || len(credits.Crew) != 1 {
		t.Errorf("GetCredits() = %+v, want ID=1399 with 1 cast and 1 crew", credits)
	}
}

func TestGetEpisodeGroups(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/episode_groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1399,
			"results": []map[string]any{
				{
					"id":            "abc123",
					"name":          "Seasons",
					"description":   "",
					"episode_count": 73,
					"group_count":   8,
					"type":          6,
					"network":       map[string]any{"id": 49, "name": "HBO"},
				},
			},
		})
	})

	groups, _, err := client.TV.GetEpisodeGroups(context.Background(), 1399)
	if err != nil {
		t.Fatalf("GetEpisodeGroups() error = %v", err)
	}
	if groups.ID != 1399 || len(groups.Results) != 1 || groups.Results[0].Name != "Seasons" {
		t.Errorf("GetEpisodeGroups() = %+v, want ID=1399 with 1 group %q", groups, "Seasons")
	}
}

func TestGetTVExternalIDs(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/external_ids", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":      1399,
			"imdb_id": "tt0944947",
		})
	})

	ids, _, err := client.TV.GetExternalIDs(context.Background(), 1399)
	if err != nil {
		t.Fatalf("GetExternalIDs() error = %v", err)
	}
	if ids.ID != 1399 || ids.ImdbID != "tt0944947" {
		t.Errorf("GetExternalIDs() = %+v, want ID=1399 ImdbID=%q", ids, "tt0944947")
	}
}

func TestGetTVImages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/images", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1399,
			"backdrops": []map[string]any{
				{"file_path": "/backdrop.jpg", "width": 1920, "height": 1080},
			},
			"posters": []map[string]any{
				{"file_path": "/poster.jpg", "width": 500, "height": 750},
			},
		})
	})

	images, _, err := client.TV.GetImages(context.Background(), 1399, &uri.ImagesOptions{})
	if err != nil {
		t.Fatalf("GetImages() error = %v", err)
	}
	if len(images.Backdrops) != 1 || len(images.Posters) != 1 {
		t.Errorf("GetImages() = %+v, want 1 backdrop and 1 poster", images)
	}
}

func TestGetTVKeywords(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/keywords", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1399,
			"results": []map[string]any{
				{"id": 818, "name": "based on novel or book"},
			},
		})
	})

	keywords, _, err := client.TV.GetKeywords(context.Background(), 1399)
	if err != nil {
		t.Fatalf("GetKeywords() error = %v", err)
	}
	if keywords.ID != 1399 || len(keywords.Results) != 1 || keywords.Results[0].Name != "based on novel or book" {
		t.Errorf("GetKeywords() = %+v, want ID=1399 with 1 keyword %q", keywords, "based on novel or book")
	}
}

func TestGetTVLatest(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/latest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   1000000,
			"name": "brand new series",
		})
	})

	tv, _, err := client.TV.GetLatest(context.Background())
	if err != nil {
		t.Fatalf("GetLatest() error = %v", err)
	}
	if tv.ID != 1000000 || tv.Name != "brand new series" {
		t.Errorf("GetLatest() = %+v, want ID=1000000 Name=%q", tv, "brand new series")
	}
}

func TestGetTVLists(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/lists", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "name": "a list"}}
		json.NewEncoder(w).Encode(map[string]any{
			"id":            1399,
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	lists, err := client.TV.GetLists(context.Background(), 1399, "", 0)
	if err != nil {
		t.Fatalf("GetLists() error = %v", err)
	}
	if len(lists) != 2 {
		t.Fatalf("len(lists) = %d, want 2", len(lists))
	}
}

func TestGetTVRecommendations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/recommendations", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": pageNum, "name": "a show"}}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	shows, err := client.TV.GetRecommendations(context.Background(), 1399, "", 0)
	if err != nil {
		t.Fatalf("GetRecommendations() error = %v", err)
	}
	if len(shows) != 2 {
		t.Fatalf("len(shows) = %d, want 2", len(shows))
	}
}

func TestGetTVReviews(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/1399/reviews", func(w http.ResponseWriter, r *http.Request) {
		pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageNum = 1
		}
		results := []map[string]any{{"id": "abc", "author": "someone", "content": "great series"}}
		json.NewEncoder(w).Encode(map[string]any{
			"id":            1399,
			"page":          pageNum,
			"results":       results,
			"total_pages":   2,
			"total_results": 2,
		})
	})

	reviews, err := client.TV.GetReviews(context.Background(), 1399, "", 0)
	if err != nil {
		t.Fatalf("GetReviews() error = %v", err)
	}
	if len(reviews) != 2 {
		t.Fatalf("len(reviews) = %d, want 2", len(reviews))
	}
}
