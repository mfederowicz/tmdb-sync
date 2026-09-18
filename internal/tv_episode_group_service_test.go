package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetTVEpisodeGroup(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tv/episode_group/5acf98d40e0a26346a0389c3", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":            "5acf98d40e0a26346a0389c3",
			"name":          "Specials",
			"description":   "Bonus episodes",
			"episode_count": 1,
			"group_count":   1,
			"type":          6,
			"network": map[string]any{
				"id":   49,
				"name": "HBO",
			},
			"groups": []map[string]any{
				{
					"id":     "5acf98d40e0a26346a0389c4",
					"name":   "Season 1",
					"order":  1,
					"locked": true,
					"episodes": []map[string]any{
						{
							"id":             63056,
							"name":           "Winter Is Coming",
							"episode_number": 1,
							"season_number":  1,
							"order":          0,
						},
					},
				},
			},
		})
	})

	group, _, err := client.TVEpisodeGroup.GetTVEpisodeGroup(context.Background(), "5acf98d40e0a26346a0389c3")
	if err != nil {
		t.Fatalf("GetTVEpisodeGroup() error = %v", err)
	}
	if group.ID != "5acf98d40e0a26346a0389c3" || group.Name != "Specials" {
		t.Errorf("GetTVEpisodeGroup() = %+v, want ID=%q Name=%q", group, "5acf98d40e0a26346a0389c3", "Specials")
	}
	if len(group.Groups) != 1 || len(group.Groups[0].Episodes) != 1 || group.Groups[0].Episodes[0].Name != "Winter Is Coming" {
		t.Errorf("GetTVEpisodeGroup() groups = %+v, want one group with one episode named %q", group.Groups, "Winter Is Coming")
	}
}
