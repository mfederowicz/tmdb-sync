package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetMovieCertifications(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/certification/movie/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"certifications": map[string]any{
				"US": []map[string]any{
					{"certification": "G", "meaning": "General audiences.", "order": 1},
				},
			},
		})
	})

	certifications, _, err := client.Certifications.GetMovieCertifications(context.Background())
	if err != nil {
		t.Fatalf("GetMovieCertifications() error = %v", err)
	}
	us, ok := certifications.Certifications["US"]
	if !ok || len(us) != 1 || us[0].Certification != "G" {
		t.Errorf("Certifications[US] = %+v, want one entry with Certification=G", us)
	}
}

func TestGetTVCertifications(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/certification/tv/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"certifications": map[string]any{
				"US": []map[string]any{
					{"certification": "TV-G", "meaning": "General audience.", "order": 1},
				},
			},
		})
	})

	certifications, _, err := client.Certifications.GetTVCertifications(context.Background())
	if err != nil {
		t.Fatalf("GetTVCertifications() error = %v", err)
	}
	us, ok := certifications.Certifications["US"]
	if !ok || len(us) != 1 || us[0].Certification != "TV-G" {
		t.Errorf("Certifications[US] = %+v, want one entry with Certification=TV-G", us)
	}
}
