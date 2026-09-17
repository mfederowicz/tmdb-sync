package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/str"
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

func TestCreateList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "sess" {
			t.Errorf("session_id = %s, want sess", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
			"list_id":        42,
		})
	})

	created, _, err := client.Lists.CreateList(context.Background(), "sess", &str.ListCreateRequest{Name: "watch later"})
	if err != nil {
		t.Fatalf("CreateList() error = %v", err)
	}
	if created.ListID != 42 || !created.Success {
		t.Errorf("CreateList() = %+v, want ListID=42 Success=true", created)
	}
}

func TestAddMovie(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/1/add_item", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Lists.AddMovie(context.Background(), "1", "sess", 100)
	if err != nil {
		t.Fatalf("AddMovie() error = %v", err)
	}
	if !status.Success {
		t.Errorf("AddMovie() = %+v, want Success=true", status)
	}
}

func TestRemoveMovie(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/1/remove_item", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Lists.RemoveMovie(context.Background(), "1", "sess", 100)
	if err != nil {
		t.Fatalf("RemoveMovie() error = %v", err)
	}
	if !status.Success {
		t.Errorf("RemoveMovie() = %+v, want Success=true", status)
	}
}

func TestClearList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/1/clear", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.URL.Query().Get("confirm"); got != "true" {
			t.Errorf("confirm = %s, want true", got)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Lists.ClearList(context.Background(), "1", "sess")
	if err != nil {
		t.Fatalf("ClearList() error = %v", err)
	}
	if !status.Success {
		t.Errorf("ClearList() = %+v, want Success=true", status)
	}
}

func TestDeleteList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Lists.DeleteList(context.Background(), "1", "sess")
	if err != nil {
		t.Fatalf("DeleteList() error = %v", err)
	}
	if !status.Success {
		t.Errorf("DeleteList() = %+v, want Success=true", status)
	}
}
