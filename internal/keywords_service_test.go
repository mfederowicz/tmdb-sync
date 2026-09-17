package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetKeyword(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/keyword/1701", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   1701,
			"name": "hero",
		})
	})

	keyword, _, err := client.Keywords.GetKeyword(context.Background(), "1701")
	if err != nil {
		t.Fatalf("GetKeyword() error = %v", err)
	}
	if keyword.ID != 1701 || keyword.Name != "hero" {
		t.Errorf("GetKeyword() = %+v, want ID=1701 Name=%q", keyword, "hero")
	}
}
