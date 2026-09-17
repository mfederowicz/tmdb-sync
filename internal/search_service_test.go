package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
)

func TestSearchCollections(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/search/collection", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("query"); got != "harry potter" {
			t.Errorf("query = %q, want %q", got, "harry potter")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page": 1,
			"results": []map[string]any{
				{"id": 1241, "name": "Harry Potter Collection"},
			},
			"total_pages":   1,
			"total_results": 1,
		})
	})

	results, err := client.Search.SearchCollections(context.Background(), uri.SearchCollectionOptions{Query: "harry potter"}, 0)
	if err != nil {
		t.Fatalf("SearchCollections() error = %v", err)
	}
	if len(results) != 1 || results[0].Name != "Harry Potter Collection" {
		t.Errorf("SearchCollections() = %+v, want one result named Harry Potter Collection", results)
	}
}
