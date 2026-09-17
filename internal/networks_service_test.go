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

func TestGetNetworkAlternativeNames(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/network/1/alternative_names", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1,
			"results": []map[string]any{
				{"name": "Home Box Office", "type": ""},
			},
		})
	})

	names, _, err := client.Networks.GetNetworkAlternativeNames(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetNetworkAlternativeNames() error = %v", err)
	}
	if len(names.Results) != 1 || names.Results[0].Name != "Home Box Office" {
		t.Errorf("GetNetworkAlternativeNames() = %+v, want one result named Home Box Office", names)
	}
}

func TestGetNetworkImages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/network/1/images", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":    1,
			"logos": []map[string]any{{"file_path": "/logo.png"}},
		})
	})

	images, _, err := client.Networks.GetNetworkImages(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetNetworkImages() error = %v", err)
	}
	if len(images.Logos) != 1 {
		t.Errorf("GetNetworkImages() = %+v, want one logo", images)
	}
}
