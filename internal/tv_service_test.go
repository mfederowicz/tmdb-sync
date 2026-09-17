package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
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
