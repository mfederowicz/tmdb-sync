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
