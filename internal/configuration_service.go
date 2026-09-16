package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// ConfigurationService handles the TMDB /configuration endpoint.
type ConfigurationService Service

// GetAPIConfiguration returns TMDB's image base URLs/sizes and change keys.
//
// Api docs: https://developer.themoviedb.org/reference/configuration-details
func (s *ConfigurationService) GetAPIConfiguration(ctx context.Context) (*str.Configuration, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "configuration", nil)
	if err != nil {
		return nil, nil, err
	}

	config := new(str.Configuration)
	resp, err := s.client.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

// GetCountries returns the countries used throughout TMDB.
//
// Api docs: https://developer.themoviedb.org/reference/configuration-countries
func (s *ConfigurationService) GetCountries(ctx context.Context) ([]*str.Country, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "configuration/countries", nil)
	if err != nil {
		return nil, nil, err
	}

	var countries []*str.Country
	resp, err := s.client.Do(ctx, req, &countries)
	if err != nil {
		return nil, resp, err
	}

	return countries, resp, nil
}

// GetJobs returns the jobs and departments used by TMDB for movie credits.
//
// Api docs: https://developer.themoviedb.org/reference/configuration-jobs
func (s *ConfigurationService) GetJobs(ctx context.Context) ([]*str.JobDepartment, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "configuration/jobs", nil)
	if err != nil {
		return nil, nil, err
	}

	var jobs []*str.JobDepartment
	resp, err := s.client.Do(ctx, req, &jobs)
	if err != nil {
		return nil, resp, err
	}

	return jobs, resp, nil
}

// GetLanguages returns the languages used throughout TMDB.
//
// Api docs: https://developer.themoviedb.org/reference/configuration-languages
func (s *ConfigurationService) GetLanguages(ctx context.Context) ([]*str.Language, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "configuration/languages", nil)
	if err != nil {
		return nil, nil, err
	}

	var languages []*str.Language
	resp, err := s.client.Do(ctx, req, &languages)
	if err != nil {
		return nil, resp, err
	}

	return languages, resp, nil
}

// GetPrimaryTranslations returns the primary translation locales used by TMDB.
//
// Api docs: https://developer.themoviedb.org/reference/configuration-primary-translations
func (s *ConfigurationService) GetPrimaryTranslations(ctx context.Context) ([]string, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "configuration/primary_translations", nil)
	if err != nil {
		return nil, nil, err
	}

	var translations []string
	resp, err := s.client.Do(ctx, req, &translations)
	if err != nil {
		return nil, resp, err
	}

	return translations, resp, nil
}

// GetTimezones returns the timezones used throughout TMDB.
//
// Api docs: https://developer.themoviedb.org/reference/configuration-timezones
func (s *ConfigurationService) GetTimezones(ctx context.Context) ([]*str.TimezoneRegion, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "configuration/timezones", nil)
	if err != nil {
		return nil, nil, err
	}

	var timezones []*str.TimezoneRegion
	resp, err := s.client.Do(ctx, req, &timezones)
	if err != nil {
		return nil, resp, err
	}

	return timezones, resp, nil
}
