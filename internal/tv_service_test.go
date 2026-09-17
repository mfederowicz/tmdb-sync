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
