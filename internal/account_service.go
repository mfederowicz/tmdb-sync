package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountService handles communication with the /account 🔒 endpoints of the
// TMDB API. Every method requires a valid v3 session id.
type AccountService Service

// GetDetails fetches an account's details. If accountID is 0, it resolves
// the account belonging to sessionID (TMDB's session-only GET /account form)
// instead of requiring the id up front.
//
// Api docs: https://developer.themoviedb.org/reference/account-details
func (s *AccountService) GetDetails(ctx context.Context, accountID int64, sessionID string) (*str.Account, *str.Response, error) {
	path := "account"
	if accountID != 0 {
		path = fmt.Sprintf("account/%d", accountID)
	}

	urlStr, err := uri.AddQuery(path, &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(str.Account)
	resp, err := s.client.Do(ctx, req, account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}

// AddToWatchlist adds or removes a movie/TV show from an account's watchlist.
//
// Api docs: https://developer.themoviedb.org/reference/account-add-to-watchlist
func (s *AccountService) AddToWatchlist(ctx context.Context, accountID int64, sessionID string, body *str.AccountWatchlistRequest) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/watchlist", accountID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, nil, err
	}

	status := new(str.AuthStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}

func (s *AccountService) getFavoriteMoviesPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/favorite/movies", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	movies := new(str.Movies)
	resp, err := s.client.Do(ctx, req, movies)
	if err != nil {
		return nil, resp, err
	}

	return movies, resp, nil
}

// GetFavoriteMovies returns an account's favorited movies, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-favorite-movies
func (s *AccountService) GetFavoriteMovies(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getFavoriteMoviesPage(ctx, accountID, sessionID, page)
		if err != nil {
			return PageResult[str.Movie]{}, err
		}
		return PageResult[str.Movie]{
			Results:    movies.Results,
			Page:       movies.Page,
			TotalPages: movies.TotalPages,
		}, nil
	})
}

func (s *AccountService) getFavoriteTVPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/favorite/tv", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	shows := new(str.TVShows)
	resp, err := s.client.Do(ctx, req, shows)
	if err != nil {
		return nil, resp, err
	}

	return shows, resp, nil
}

// GetFavoriteTV returns an account's favorited TV shows, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-favorite-tv
func (s *AccountService) GetFavoriteTV(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		shows, _, err := s.getFavoriteTVPage(ctx, accountID, sessionID, page)
		if err != nil {
			return PageResult[str.TV]{}, err
		}
		return PageResult[str.TV]{
			Results:    shows.Results,
			Page:       shows.Page,
			TotalPages: shows.TotalPages,
		}, nil
	})
}

func (s *AccountService) getListsPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.AccountLists, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/lists", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	lists := new(str.AccountLists)
	resp, err := s.client.Do(ctx, req, lists)
	if err != nil {
		return nil, resp, err
	}

	return lists, resp, nil
}

// GetLists returns an account's custom lists, walking pages until TMDB
// reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-lists
func (s *AccountService) GetLists(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.AccountList, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.AccountList], error) {
		lists, _, err := s.getListsPage(ctx, accountID, sessionID, page)
		if err != nil {
			return PageResult[str.AccountList]{}, err
		}
		return PageResult[str.AccountList]{
			Results:    lists.Results,
			Page:       lists.Page,
			TotalPages: lists.TotalPages,
		}, nil
	})
}

// AddRemoveFavorite marks or unmarks a movie/TV show as one of an account's favorites.
//
// Api docs: https://developer.themoviedb.org/reference/account-add-favorite
func (s *AccountService) AddRemoveFavorite(ctx context.Context, accountID int64, sessionID string, body *str.AccountFavoriteRequest) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/favorite", accountID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, nil, err
	}

	status := new(str.AuthStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
