package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetReview(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/review/abc123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":      "abc123",
			"author":  "Jane",
			"content": "Great movie.",
		})
	})

	review, _, err := client.Reviews.GetReview(context.Background(), "abc123")
	if err != nil {
		t.Fatalf("GetReview() error = %v", err)
	}
	if review.ID != "abc123" || review.Author != "Jane" {
		t.Errorf("GetReview() = %+v, want ID=%q Author=%q", review, "abc123", "Jane")
	}
}
