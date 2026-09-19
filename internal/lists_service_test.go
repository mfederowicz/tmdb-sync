package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"

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

func TestListsGetListV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/list/8", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user token", got)
		}
		if got := r.URL.Query().Get("language"); got != "en-US" {
			t.Errorf("language = %q, want en-US", got)
		}
		page := r.URL.Query().Get("page")
		item := map[string]any{"id": 1, "media_type": "movie", "title": "one"}
		if page == "2" {
			item = map[string]any{"id": 2, "media_type": "tv", "name": "two"}
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 8, "name": "mine", "page": json.Number(page), "total_pages": 2, "total_results": 2,
			"created_by": map[string]any{"username": "me"},
			"results":    []any{item},
		})
	})

	list, err := client.Lists.GetListV4(context.Background(), "usertok", "8", 0, uri.ListV4Options{Language: "en-US"})
	if err != nil {
		t.Fatalf("GetListV4() error = %v", err)
	}
	if list.ID != 8 || list.Name != "mine" || list.CreatedBy.Username != "me" {
		t.Errorf("GetListV4() = %+v", list)
	}
	if len(list.Results) != 2 || list.Results[0].Title != "one" || list.Results[1].Name != "two" {
		t.Errorf("Results = %+v, want both pages merged", list.Results)
	}
}

func TestListsGetListV4_PublicWithoutUserToken(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/list/8", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer read" {
			t.Errorf("Authorization = %q, want read token", got)
		}
		json.NewEncoder(w).Encode(map[string]any{"id": 8, "page": 1, "total_pages": 1})
	})

	if _, err := client.Lists.GetListV4(context.Background(), "", "8", 0, uri.ListV4Options{}); err != nil {
		t.Fatalf("GetListV4() error = %v", err)
	}
}

func TestListsCreateListV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user token", got)
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body["name"] != "mine" || body["iso_639_1"] != "en" || body["iso_3166_1"] != "US" || body["public"] != true {
			t.Errorf("body = %v", body)
		}
		json.NewEncoder(w).Encode(map[string]any{"id": 77, "success": true, "status_code": 1})
	})

	created, err := client.Lists.CreateListV4(context.Background(), "usertok", &str.ListCreateRequestV4{Name: "mine", ISO6391: "en", ISO31661: "US", Public: true})
	if err != nil {
		t.Fatalf("CreateListV4() error = %v", err)
	}
	if created.ID != 77 || !created.Success {
		t.Errorf("CreateListV4() = %+v", created)
	}
}

func TestListsUpdateListV4(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/list/8", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user token", got)
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if len(body) != 2 || body["name"] != "renamed" || body["public"] != false {
			t.Errorf("body = %v, want only name and public=false", body)
		}
		json.NewEncoder(w).Encode(map[string]any{"success": true, "status_code": 1})
	})

	public := false
	status, err := client.Lists.UpdateListV4(context.Background(), "usertok", "8", &str.ListUpdateRequestV4{Name: "renamed", Public: &public})
	if err != nil {
		t.Fatalf("UpdateListV4() error = %v", err)
	}
	if !status.Success {
		t.Errorf("UpdateListV4() = %+v", status)
	}
}
