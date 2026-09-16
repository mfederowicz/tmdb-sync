package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetAPIConfiguration(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/configuration", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"images": map[string]any{
				"base_url":        "http://image.tmdb.org/t/p/",
				"secure_base_url": "https://image.tmdb.org/t/p/",
			},
			"change_keys": []string{"adult"},
		})
	})

	config, _, err := client.Configuration.GetAPIConfiguration(context.Background())
	if err != nil {
		t.Fatalf("GetAPIConfiguration() error = %v", err)
	}
	if config.Images.BaseURL != "http://image.tmdb.org/t/p/" {
		t.Errorf("Images.BaseURL = %q, want http://image.tmdb.org/t/p/", config.Images.BaseURL)
	}
}

func TestGetCountries(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/configuration/countries", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode([]map[string]any{
			{"iso_3166_1": "US", "english_name": "United States of America", "native_name": "United States"},
		})
	})

	countries, _, err := client.Configuration.GetCountries(context.Background())
	if err != nil {
		t.Fatalf("GetCountries() error = %v", err)
	}
	if len(countries) != 1 || countries[0].ISO31661 != "US" {
		t.Errorf("Countries = %+v, want one entry with ISO31661=US", countries)
	}
}

func TestGetJobs(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/configuration/jobs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode([]map[string]any{
			{"department": "Directing", "jobs": []string{"Director"}},
		})
	})

	jobs, _, err := client.Configuration.GetJobs(context.Background())
	if err != nil {
		t.Fatalf("GetJobs() error = %v", err)
	}
	if len(jobs) != 1 || jobs[0].Department != "Directing" {
		t.Errorf("Jobs = %+v, want one entry with Department=Directing", jobs)
	}
}

func TestGetLanguages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/configuration/languages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode([]map[string]any{
			{"iso_639_1": "en", "english_name": "English", "name": "English"},
		})
	})

	languages, _, err := client.Configuration.GetLanguages(context.Background())
	if err != nil {
		t.Fatalf("GetLanguages() error = %v", err)
	}
	if len(languages) != 1 || languages[0].ISO6391 != "en" {
		t.Errorf("Languages = %+v, want one entry with ISO6391=en", languages)
	}
}

func TestGetPrimaryTranslations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/configuration/primary_translations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode([]string{"en-US", "fr-FR"})
	})

	translations, _, err := client.Configuration.GetPrimaryTranslations(context.Background())
	if err != nil {
		t.Fatalf("GetPrimaryTranslations() error = %v", err)
	}
	if len(translations) != 2 || translations[0] != "en-US" {
		t.Errorf("PrimaryTranslations = %+v, want [en-US fr-FR]", translations)
	}
}

func TestGetTimezones(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/configuration/timezones", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode([]map[string]any{
			{"iso_3166_1": "US", "zones": []string{"America/New_York"}},
		})
	})

	timezones, _, err := client.Configuration.GetTimezones(context.Background())
	if err != nil {
		t.Fatalf("GetTimezones() error = %v", err)
	}
	if len(timezones) != 1 || timezones[0].ISO31661 != "US" {
		t.Errorf("Timezones = %+v, want one entry with ISO31661=US", timezones)
	}
}
