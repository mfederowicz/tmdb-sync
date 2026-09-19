package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountService handles communication with the /account 🔒 endpoints of the
// TMDB API. Methods require a valid v3 session id, except the V4-suffixed ones,
// which use the v4 user access token instead of a session_id.
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

func (s *AccountService) getRatedMoviesPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.RatedMovies, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/rated/movies", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	movies := new(str.RatedMovies)
	resp, err := s.client.Do(ctx, req, movies)
	if err != nil {
		return nil, resp, err
	}

	return movies, resp, nil
}

// GetRatedMovies returns an account's rated movies, walking pages until TMDB
// reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-rated-movies
func (s *AccountService) GetRatedMovies(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.RatedMovie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.RatedMovie], error) {
		movies, _, err := s.getRatedMoviesPage(ctx, accountID, sessionID, page)
		if err != nil {
			return PageResult[str.RatedMovie]{}, err
		}
		return PageResult[str.RatedMovie]{
			Results:    movies.Results,
			Page:       movies.Page,
			TotalPages: movies.TotalPages,
		}, nil
	})
}

func (s *AccountService) getRatedTVPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.RatedTVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/rated/tv", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	shows := new(str.RatedTVShows)
	resp, err := s.client.Do(ctx, req, shows)
	if err != nil {
		return nil, resp, err
	}

	return shows, resp, nil
}

// GetRatedTV returns an account's rated TV shows, walking pages until TMDB
// reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-rated-tv
func (s *AccountService) GetRatedTV(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.RatedTV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.RatedTV], error) {
		shows, _, err := s.getRatedTVPage(ctx, accountID, sessionID, page)
		if err != nil {
			return PageResult[str.RatedTV]{}, err
		}
		return PageResult[str.RatedTV]{
			Results:    shows.Results,
			Page:       shows.Page,
			TotalPages: shows.TotalPages,
		}, nil
	})
}

func (s *AccountService) getRatedTVEpisodesPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.RatedTVEpisodes, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/rated/tv/episodes", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	episodes := new(str.RatedTVEpisodes)
	resp, err := s.client.Do(ctx, req, episodes)
	if err != nil {
		return nil, resp, err
	}

	return episodes, resp, nil
}

// GetRatedTVEpisodes returns an account's rated TV episodes, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-rated-tv-episodes
func (s *AccountService) GetRatedTVEpisodes(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.RatedTVEpisode, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.RatedTVEpisode], error) {
		episodes, _, err := s.getRatedTVEpisodesPage(ctx, accountID, sessionID, page)
		if err != nil {
			return PageResult[str.RatedTVEpisode]{}, err
		}
		return PageResult[str.RatedTVEpisode]{
			Results:    episodes.Results,
			Page:       episodes.Page,
			TotalPages: episodes.TotalPages,
		}, nil
	})
}

func (s *AccountService) getWatchlistMoviesPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/watchlist/movies", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
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

// GetWatchlistMovies returns an account's watchlisted movies, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-watchlist-movies
func (s *AccountService) GetWatchlistMovies(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getWatchlistMoviesPage(ctx, accountID, sessionID, page)
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

func (s *AccountService) getWatchlistTVPage(ctx context.Context, accountID int64, sessionID string, page int) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d/watchlist/tv", accountID), &uri.AccountListOptions{SessionID: sessionID, Page: page})
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

// GetWatchlistTV returns an account's watchlisted TV shows, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/account-watchlist-tv
func (s *AccountService) GetWatchlistTV(ctx context.Context, accountID int64, sessionID string, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		shows, _, err := s.getWatchlistTVPage(ctx, accountID, sessionID, page)
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

// withUserToken replaces the read access token with the v4 user access token.
func withUserToken(accessToken string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
}

// pagedV4 is the envelope shared by v4 account list endpoints.
type pagedV4[T any] struct {
	Page       int `json:"page"`
	Results    []T `json:"results"`
	TotalPages int `json:"total_pages"`
}

// fetchAccountV4 walks every page (up to pagesLimit, 0 = unlimited) of the v4
// account endpoint at path, authenticated with the user access token.
func fetchAccountV4[T any](ctx context.Context, s *AccountService, accessToken, path string, pagesLimit int) ([]T, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[T], error) {
		urlStr, err := uri.AddQuery(path, &uri.PageOptions{Page: page})
		if err != nil {
			return PageResult[T]{}, err
		}

		req, err := s.client.NewRequestV4(http.MethodGet, urlStr, nil, withUserToken(accessToken))
		if err != nil {
			return PageResult[T]{}, err
		}

		body := new(pagedV4[T])
		if _, err := s.client.Do(ctx, req, body); err != nil {
			return PageResult[T]{}, err
		}
		return PageResult[T]{Results: body.Results, Page: body.Page, TotalPages: body.TotalPages}, nil
	})
}

// GetListsV4 returns a v4 account's lists, walking pages until TMDB reports no
// more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-lists
func (s *AccountService) GetListsV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.AccountListV4, error) {
	return fetchAccountV4[str.AccountListV4](ctx, s, accessToken, fmt.Sprintf("account/%s/lists", accountID), pagesLimit)
}

// GetFavoriteMoviesV4 returns a v4 account's favorited movies, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-favorite-movies
func (s *AccountService) GetFavoriteMoviesV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.Movie, error) {
	return fetchAccountV4[str.Movie](ctx, s, accessToken, fmt.Sprintf("account/%s/movie/favorites", accountID), pagesLimit)
}

// GetFavoriteTVV4 returns a v4 account's favorited TV shows, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-favorite-tv
func (s *AccountService) GetFavoriteTVV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.TV, error) {
	return fetchAccountV4[str.TV](ctx, s, accessToken, fmt.Sprintf("account/%s/tv/favorites", accountID), pagesLimit)
}

// GetRatedMoviesV4 returns a v4 account's rated movies, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-rated-movies
func (s *AccountService) GetRatedMoviesV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.RatedMovieV4, error) {
	return fetchAccountV4[str.RatedMovieV4](ctx, s, accessToken, fmt.Sprintf("account/%s/movie/rated", accountID), pagesLimit)
}

// GetRatedTVV4 returns a v4 account's rated TV shows, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-rated-tv
func (s *AccountService) GetRatedTVV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.RatedTVV4, error) {
	return fetchAccountV4[str.RatedTVV4](ctx, s, accessToken, fmt.Sprintf("account/%s/tv/rated", accountID), pagesLimit)
}

// GetRecommendedMoviesV4 returns movie recommendations for a v4 account,
// walking pages until TMDB reports no more (total_pages) or pagesLimit is
// reached (0 = unlimited). There is no v3 equivalent.
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-recommended-movies
func (s *AccountService) GetRecommendedMoviesV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.Movie, error) {
	return fetchAccountV4[str.Movie](ctx, s, accessToken, fmt.Sprintf("account/%s/movie/recommendations", accountID), pagesLimit)
}

// GetRecommendedTVV4 returns TV show recommendations for a v4 account,
// walking pages until TMDB reports no more (total_pages) or pagesLimit is
// reached (0 = unlimited). There is no v3 equivalent.
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-recommended-tv
func (s *AccountService) GetRecommendedTVV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.TV, error) {
	return fetchAccountV4[str.TV](ctx, s, accessToken, fmt.Sprintf("account/%s/tv/recommendations", accountID), pagesLimit)
}

// GetWatchlistMoviesV4 returns a v4 account's watchlisted movies, walking
// pages until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/v4/reference/account-watchlist-movies
func (s *AccountService) GetWatchlistMoviesV4(ctx context.Context, accessToken, accountID string, pagesLimit int) ([]str.Movie, error) {
	return fetchAccountV4[str.Movie](ctx, s, accessToken, fmt.Sprintf("account/%s/movie/watchlist", accountID), pagesLimit)
}
