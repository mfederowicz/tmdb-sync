package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetAvailableRegions(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watch/providers/regions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"results": []map[string]any{
				{"iso_3166_1": "US", "english_name": "United States", "native_name": "United States"},
			},
		})
	})

	regions, _, err := client.WatchProviders.GetAvailableRegions(context.Background(), "")
	if err != nil {
		t.Fatalf("GetAvailableRegions() error = %v", err)
	}
	if len(regions.Results) != 1 || regions.Results[0].ISO31661 != "US" {
		t.Errorf("GetAvailableRegions() = %+v, want one result with ISO31661=US", regions)
	}
}

func TestGetMovieProviders(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watch/providers/movie", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"results": []map[string]any{
				{"provider_id": 8, "provider_name": "Netflix", "display_priority": 0, "display_priorities": map[string]any{"US": 0}},
			},
		})
	})

	providers, _, err := client.WatchProviders.GetMovieProviders(context.Background(), "", "")
	if err != nil {
		t.Fatalf("GetMovieProviders() error = %v", err)
	}
	if len(providers.Results) != 1 || providers.Results[0].ProviderName != "Netflix" {
		t.Errorf("GetMovieProviders() = %+v, want one result named Netflix", providers)
	}
}
