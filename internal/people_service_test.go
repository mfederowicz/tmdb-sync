package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetPerson(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/person/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":   1,
			"name": "George Clooney",
		})
	})

	person, _, err := client.People.GetPerson(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetPerson() error = %v", err)
	}
	if person.ID != 1 || person.Name != "George Clooney" {
		t.Errorf("GetPerson() = %+v, want ID=1 Name=%q", person, "George Clooney")
	}
}

func TestGetPersonCombinedCredits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/person/1/combined_credits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id": 1,
			"cast": []map[string]any{
				{"id": 100, "media_type": "movie", "title": "Ocean's Eleven", "character": "Danny Ocean"},
			},
			"crew": []map[string]any{
				{"id": 200, "media_type": "tv", "name": "Unscripted", "job": "Director"},
			},
		})
	})

	credits, _, err := client.People.GetPersonCombinedCredits(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetPersonCombinedCredits() error = %v", err)
	}
	if len(credits.Cast) != 1 || credits.Cast[0].Title != "Ocean's Eleven" {
		t.Errorf("GetPersonCombinedCredits() cast = %+v, want one entry titled Ocean's Eleven", credits.Cast)
	}
	if len(credits.Crew) != 1 || credits.Crew[0].Job != "Director" {
		t.Errorf("GetPersonCombinedCredits() crew = %+v, want one entry job Director", credits.Crew)
	}
}
