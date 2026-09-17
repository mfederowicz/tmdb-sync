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
