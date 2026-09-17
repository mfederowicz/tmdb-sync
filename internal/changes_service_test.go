package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetMovieChanges(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/changes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("start_date"); got != "2024-01-01" {
			t.Errorf("start_date query param = %q, want %q", got, "2024-01-01")
		}
		if got := r.URL.Query().Get("end_date"); got != "2024-01-10" {
			t.Errorf("end_date query param = %q, want %q", got, "2024-01-10")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 550, "adult": false},
			},
		})
	})

	changes, err := client.Changes.GetMovieChanges(context.Background(), "2024-01-01", "2024-01-10", 0)
	if err != nil {
		t.Fatalf("GetMovieChanges() error = %v", err)
	}
	if len(changes) != 1 || changes[0].ID != 550 {
		t.Errorf("GetMovieChanges() = %+v, want one change with ID=550", changes)
	}
}

func TestGetTVChanges(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/changes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 1396, "adult": false},
			},
		})
	})

	changes, err := client.Changes.GetTVChanges(context.Background(), "", "", 0)
	if err != nil {
		t.Fatalf("GetTVChanges() error = %v", err)
	}
	if len(changes) != 1 || changes[0].ID != 1396 {
		t.Errorf("GetTVChanges() = %+v, want one change with ID=1396", changes)
	}
}

func TestGetPersonChanges(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/person/changes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 287, "adult": false},
			},
		})
	})

	changes, err := client.Changes.GetPersonChanges(context.Background(), "", "", 0)
	if err != nil {
		t.Fatalf("GetPersonChanges() error = %v", err)
	}
	if len(changes) != 1 || changes[0].ID != 287 {
		t.Errorf("GetPersonChanges() = %+v, want one change with ID=287", changes)
	}
}
