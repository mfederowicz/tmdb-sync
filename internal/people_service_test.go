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
