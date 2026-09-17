package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// PeopleService handles communication with the /person endpoints of the TMDB API.
type PeopleService Service

// GetPerson fetches details for a single person by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/person-details
func (s *PeopleService) GetPerson(ctx context.Context, personID int64) (*str.PersonDetails, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("person/%d", personID), nil)
	if err != nil {
		return nil, nil, err
	}

	person := new(str.PersonDetails)
	resp, err := s.client.Do(ctx, req, person)
	if err != nil {
		return nil, resp, err
	}

	return person, resp, nil
}

// GetPersonCombinedCredits fetches the movie and TV credits for a single
// person by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/person-combined-credits
func (s *PeopleService) GetPersonCombinedCredits(ctx context.Context, personID int64) (*str.PersonCombinedCredits, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("person/%d/combined_credits", personID), nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.PersonCombinedCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}

// GetPersonExternalIDs fetches the external ids for a single person by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/person-external-ids
func (s *PeopleService) GetPersonExternalIDs(ctx context.Context, personID int64) (*str.PersonExternalIDs, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("person/%d/external_ids", personID), nil)
	if err != nil {
		return nil, nil, err
	}

	ids := new(str.PersonExternalIDs)
	resp, err := s.client.Do(ctx, req, ids)
	if err != nil {
		return nil, resp, err
	}

	return ids, resp, nil
}

// GetPersonImages fetches the profile images for a single person by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/person-images
func (s *PeopleService) GetPersonImages(ctx context.Context, personID int64) (*str.PersonImages, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("person/%d/images", personID), nil)
	if err != nil {
		return nil, nil, err
	}

	images := new(str.PersonImages)
	resp, err := s.client.Do(ctx, req, images)
	if err != nil {
		return nil, resp, err
	}

	return images, resp, nil
}

// GetLatest fetches the most recently created person entry.
//
// Api docs: https://developer.themoviedb.org/reference/person-latest-id
func (s *PeopleService) GetLatest(ctx context.Context) (*str.PersonDetails, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "person/latest", nil)
	if err != nil {
		return nil, nil, err
	}

	person := new(str.PersonDetails)
	resp, err := s.client.Do(ctx, req, person)
	if err != nil {
		return nil, resp, err
	}

	return person, resp, nil
}

// GetPersonMovieCredits fetches the movie credits for a single person by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/person-movie-credits
func (s *PeopleService) GetPersonMovieCredits(ctx context.Context, personID int64) (*str.PersonMovieCredits, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("person/%d/movie_credits", personID), nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.PersonMovieCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}
