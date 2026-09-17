package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":             1,
			"name":           "watch later",
			"description":    "movies to watch",
			"created_by":     "someone",
			"favorite_count": 0,
			"item_count":     2,
			"iso_639_1":      "en",
			"items":          []map[string]any{{"id": 100, "title": "movie a"}},
		})
	})

	list, _, err := client.Lists.GetList(context.Background(), "1")
	if err != nil {
		t.Fatalf("GetList() error = %v", err)
	}
	if list.ID != 1 || list.Name != "watch later" || len(list.Items) != 1 {
		t.Errorf("GetList() = %+v, want ID=1 Name=%q with 1 item", list, "watch later")
	}
}

func TestGetItemStatus(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/1/item_status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("movie_id"); got != "100" {
			t.Errorf("movie_id = %s, want 100", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":           1,
			"item_present": true,
		})
	})

	status, _, err := client.Lists.GetItemStatus(context.Background(), "1", 100)
	if err != nil {
		t.Fatalf("GetItemStatus() error = %v", err)
	}
	if status.ID != 1 || !status.ItemPresent {
		t.Errorf("GetItemStatus() = %+v, want ID=1 ItemPresent=true", status)
	}
}
