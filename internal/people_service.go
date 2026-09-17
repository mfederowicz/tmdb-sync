package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
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

// GetPersonTVCredits fetches the TV credits for a single person by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/person-tv-credits
func (s *PeopleService) GetPersonTVCredits(ctx context.Context, personID int64) (*str.PersonTVCredits, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("person/%d/tv_credits", personID), nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.PersonTVCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}

// GetPersonTranslations fetches the translations for a single person by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/person-translations
func (s *PeopleService) GetPersonTranslations(ctx context.Context, personID int64) (*str.PersonTranslations, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("person/%d/translations", personID), nil)
	if err != nil {
		return nil, nil, err
	}

	translations := new(str.PersonTranslations)
	resp, err := s.client.Do(ctx, req, translations)
	if err != nil {
		return nil, resp, err
	}

	return translations, resp, nil
}

// getPopularPersonsPage fetches a single page of the popular-people list.
func (s *PeopleService) getPopularPersonsPage(ctx context.Context, opts *uri.ListOptions) (*str.PopularPersons, *str.Response, error) {
	urlStr, err := uri.AddQuery("person/popular", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	persons := new(str.PopularPersons)
	resp, err := s.client.Do(ctx, req, persons)
	if err != nil {
		return nil, resp, err
	}

	return persons, resp, nil
}

// GetPopularPeople returns the current popular-people list, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/person-popular-list
func (s *PeopleService) GetPopularPeople(ctx context.Context, pagesLimit int) ([]str.PopularPerson, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.PopularPerson], error) {
		persons, _, err := s.getPopularPersonsPage(ctx, &uri.ListOptions{Page: page})
		if err != nil {
			return PageResult[str.PopularPerson]{}, err
		}
		return PageResult[str.PopularPerson]{
			Results:    persons.Results,
			Page:       persons.Page,
			TotalPages: persons.TotalPages,
		}, nil
	})
}
