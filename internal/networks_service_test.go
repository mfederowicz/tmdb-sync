package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetNetwork(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/network/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   1,
			"name": "HBO",
		})
	})

	network, _, err := client.Networks.GetNetwork(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetNetwork() error = %v", err)
	}
	if network.ID != 1 || network.Name != "HBO" {
		t.Errorf("GetNetwork() = %+v, want ID=1 Name=%q", network, "HBO")
	}
}
