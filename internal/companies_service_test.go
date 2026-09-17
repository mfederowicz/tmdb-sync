package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetCompany(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/company/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   1,
			"name": "Lucasfilm Ltd.",
		})
	})

	company, _, err := client.Companies.GetCompany(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetCompany() error = %v", err)
	}
	if company.ID != 1 || company.Name != "Lucasfilm Ltd." {
		t.Errorf("GetCompany() = %+v, want ID=1 Name=%q", company, "Lucasfilm Ltd.")
	}
}

func TestGetCompanyAlternativeNames(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/company/1/alternative_names", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1,
			"results": []map[string]any{
				{"name": "Lucasfilm", "type": ""},
			},
		})
	})

	names, _, err := client.Companies.GetCompanyAlternativeNames(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetCompanyAlternativeNames() error = %v", err)
	}
	if len(names.Results) != 1 || names.Results[0].Name != "Lucasfilm" {
		t.Errorf("GetCompanyAlternativeNames() = %+v, want one result named Lucasfilm", names)
	}
}

func TestGetCompanyImages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/company/1/images", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":    1,
			"logos": []map[string]any{{"file_path": "/logo.png"}},
		})
	})

	images, _, err := client.Companies.GetCompanyImages(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetCompanyImages() error = %v", err)
	}
	if len(images.Logos) != 1 {
		t.Errorf("GetCompanyImages() = %+v, want one logo", images)
	}
}
