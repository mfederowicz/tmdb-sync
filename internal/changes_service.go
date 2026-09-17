package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// ChangesService handles communication with the /movie/changes, /tv/changes,
// and /person/changes endpoints of the TMDB API.
type ChangesService Service

func (s *ChangesService) getMovieChangesPage(ctx context.Context, startDate, endDate string, page int) (*str.Changes, *str.Response, error) {
	urlStr, err := uri.AddQuery("movie/changes", &uri.ChangesOptions{StartDate: startDate, EndDate: endDate, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	changes := new(str.Changes)
	resp, err := s.client.Do(ctx, req, changes)
	if err != nil {
		return nil, resp, err
	}

	return changes, resp, nil
}

// GetMovieChanges returns movie ids changed between startDate and endDate
// (both optional, YYYY-MM-DD; TMDB defaults to the last 24 hours and caps
// the range at 14 days), walking pages until TMDB reports no more
// (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/changes-movie-list
func (s *ChangesService) GetMovieChanges(ctx context.Context, startDate, endDate string, pagesLimit int) ([]str.ChangeItem, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.ChangeItem], error) {
		changes, _, err := s.getMovieChangesPage(ctx, startDate, endDate, page)
		if err != nil {
			return PageResult[str.ChangeItem]{}, err
		}
		return PageResult[str.ChangeItem]{
			Results:    changes.Results,
			Page:       changes.Page,
			TotalPages: changes.TotalPages,
		}, nil
	})
}

func (s *ChangesService) getTVChangesPage(ctx context.Context, startDate, endDate string, page int) (*str.Changes, *str.Response, error) {
	urlStr, err := uri.AddQuery("tv/changes", &uri.ChangesOptions{StartDate: startDate, EndDate: endDate, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	changes := new(str.Changes)
	resp, err := s.client.Do(ctx, req, changes)
	if err != nil {
		return nil, resp, err
	}

	return changes, resp, nil
}

// GetTVChanges returns TV show ids changed between startDate and endDate
// (both optional, YYYY-MM-DD; TMDB defaults to the last 24 hours and caps
// the range at 14 days), walking pages until TMDB reports no more
// (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/changes-tv-list
func (s *ChangesService) GetTVChanges(ctx context.Context, startDate, endDate string, pagesLimit int) ([]str.ChangeItem, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.ChangeItem], error) {
		changes, _, err := s.getTVChangesPage(ctx, startDate, endDate, page)
		if err != nil {
			return PageResult[str.ChangeItem]{}, err
		}
		return PageResult[str.ChangeItem]{
			Results:    changes.Results,
			Page:       changes.Page,
			TotalPages: changes.TotalPages,
		}, nil
	})
}

func (s *ChangesService) getPersonChangesPage(ctx context.Context, startDate, endDate string, page int) (*str.Changes, *str.Response, error) {
	urlStr, err := uri.AddQuery("person/changes", &uri.ChangesOptions{StartDate: startDate, EndDate: endDate, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	changes := new(str.Changes)
	resp, err := s.client.Do(ctx, req, changes)
	if err != nil {
		return nil, resp, err
	}

	return changes, resp, nil
}

// GetPersonChanges returns person ids changed between startDate and endDate
// (both optional, YYYY-MM-DD; TMDB defaults to the last 24 hours and caps
// the range at 14 days), walking pages until TMDB reports no more
// (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/changes-people-list
func (s *ChangesService) GetPersonChanges(ctx context.Context, startDate, endDate string, pagesLimit int) ([]str.ChangeItem, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.ChangeItem], error) {
		changes, _, err := s.getPersonChangesPage(ctx, startDate, endDate, page)
		if err != nil {
			return PageResult[str.ChangeItem]{}, err
		}
		return PageResult[str.ChangeItem]{
			Results:    changes.Results,
			Page:       changes.Page,
			TotalPages: changes.TotalPages,
		}, nil
	})
}
