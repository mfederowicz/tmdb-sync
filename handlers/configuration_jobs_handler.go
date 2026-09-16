package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ConfigurationJobsHandler handles `configuration -a jobs`.
type ConfigurationJobsHandler struct{}

// Handle fetches TMDB's list of jobs/departments.
func (ConfigurationJobsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	jobs, _, err := client.Configuration.GetJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
