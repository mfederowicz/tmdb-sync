package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetCredit(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/credit/52542282760ee313280017f9", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":          "52542282760ee313280017f9",
			"credit_type": "cast",
			"department":  "Acting",
			"job":         "Actor",
			"media_type":  "movie",
			"media": map[string]any{
				"id":        550,
				"title":     "Fight Club",
				"character": "The Narrator",
			},
			"person": map[string]any{
				"id":   819,
				"name": "Edward Norton",
			},
		})
	})

	credit, _, err := client.Credits.GetCredit(context.Background(), "52542282760ee313280017f9")
	if err != nil {
		t.Fatalf("GetCredit() error = %v", err)
	}
	if credit.ID != "52542282760ee313280017f9" || credit.Person.Name != "Edward Norton" {
		t.Errorf("GetCredit() = %+v, want ID=%q Person.Name=%q", credit, "52542282760ee313280017f9", "Edward Norton")
	}
}
