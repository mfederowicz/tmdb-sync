package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// MoviesService handles communication with the /movie endpoints of the TMDB API.
type MoviesService Service

// GetMovie fetches details for a single movie by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/movie-details
func (s *MoviesService) GetMovie(ctx context.Context, movieID int64) (*str.Movie, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	movie := new(str.Movie)
	resp, err := s.client.Do(ctx, req, movie)
	if err != nil {
		return nil, resp, err
	}

	return movie, resp, nil
}

// GetAccountStates fetches an account's favorite/rated/watchlist status for
// a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-account-states
func (s *MoviesService) GetAccountStates(ctx context.Context, movieID int64, sessionID string) (*str.MovieAccountStates, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/account_states", movieID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	states := new(str.MovieAccountStates)
	resp, err := s.client.Do(ctx, req, states)
	if err != nil {
		return nil, resp, err
	}

	return states, resp, nil
}

// GetAlternativeTitles fetches the alternative titles for a single movie,
// optionally filtered to a single country (ISO 3166-1).
//
// Api docs: https://developer.themoviedb.org/reference/movie-alternative-titles
func (s *MoviesService) GetAlternativeTitles(ctx context.Context, movieID int64, country string) (*str.MovieAlternativeTitles, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/alternative_titles", movieID), &uri.MovieAlternativeTitlesOptions{Country: country})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	titles := new(str.MovieAlternativeTitles)
	resp, err := s.client.Do(ctx, req, titles)
	if err != nil {
		return nil, resp, err
	}

	return titles, resp, nil
}

// GetCredits fetches the cast and crew for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-credits
func (s *MoviesService) GetCredits(ctx context.Context, movieID int64, language string) (*str.MovieCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/credits", movieID), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.MovieCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}

// GetExternalIDs fetches a single movie's external ids (IMDb, Wikidata, ...).
//
// Api docs: https://developer.themoviedb.org/reference/movie-external-ids
func (s *MoviesService) GetExternalIDs(ctx context.Context, movieID int64) (*str.MovieExternalIDs, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d/external_ids", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	ids := new(str.MovieExternalIDs)
	resp, err := s.client.Do(ctx, req, ids)
	if err != nil {
		return nil, resp, err
	}

	return ids, resp, nil
}

// GetImages fetches the images for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-images
func (s *MoviesService) GetImages(ctx context.Context, movieID int64, opts *uri.ImagesOptions) (*str.Images, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/images", movieID), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	images := new(str.Images)
	resp, err := s.client.Do(ctx, req, images)
	if err != nil {
		return nil, resp, err
	}

	return images, resp, nil
}

// GetKeywords fetches the keywords for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-keywords
func (s *MoviesService) GetKeywords(ctx context.Context, movieID int64) (*str.MovieKeywords, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d/keywords", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	keywords := new(str.MovieKeywords)
	resp, err := s.client.Do(ctx, req, keywords)
	if err != nil {
		return nil, resp, err
	}

	return keywords, resp, nil
}

// GetLatest fetches the most recently created movie on TMDB.
//
// Api docs: https://developer.themoviedb.org/reference/movie-latest-id
func (s *MoviesService) GetLatest(ctx context.Context) (*str.Movie, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "movie/latest", nil)
	if err != nil {
		return nil, nil, err
	}

	movie := new(str.Movie)
	resp, err := s.client.Do(ctx, req, movie)
	if err != nil {
		return nil, resp, err
	}

	return movie, resp, nil
}

// getMovieListsPage fetches a single page of the lists a movie belongs to.
func (s *MoviesService) getMovieListsPage(ctx context.Context, movieID int64, opts *uri.ListOptions) (*str.MovieLists, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/lists", movieID), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	lists := new(str.MovieLists)
	resp, err := s.client.Do(ctx, req, lists)
	if err != nil {
		return nil, resp, err
	}

	return lists, resp, nil
}

// GetLists returns the lists a movie belongs to, walking pages until TMDB
// reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/movie-lists
func (s *MoviesService) GetLists(ctx context.Context, movieID int64, language string, pagesLimit int) ([]str.AccountList, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.AccountList], error) {
		lists, _, err := s.getMovieListsPage(ctx, movieID, &uri.ListOptions{Page: page, Language: language})
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

// getNowPlayingMoviesPage fetches a single page of the now-playing-movies list.
func (s *MoviesService) getNowPlayingMoviesPage(ctx context.Context, opts *uri.ListOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery("movie/now_playing", opts)
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

// GetNowPlayingMovies returns the current now-playing-movies list, walking
// pages until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/movie-now-playing-list
func (s *MoviesService) GetNowPlayingMovies(ctx context.Context, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getNowPlayingMoviesPage(ctx, &uri.ListOptions{Page: page})
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

// getRecommendationsPage fetches a single page of a movie's recommendations.
func (s *MoviesService) getRecommendationsPage(ctx context.Context, movieID int64, opts *uri.ListOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/recommendations", movieID), opts)
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

// GetRecommendations returns movies recommended off a single movie, walking
// pages until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/movie-recommendations
func (s *MoviesService) GetRecommendations(ctx context.Context, movieID int64, language string, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getRecommendationsPage(ctx, movieID, &uri.ListOptions{Page: page, Language: language})
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

// GetReleaseDates fetches the per-country release dates and certifications
// for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-release-dates
func (s *MoviesService) GetReleaseDates(ctx context.Context, movieID int64) (*str.MovieReleaseDates, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d/release_dates", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	dates := new(str.MovieReleaseDates)
	resp, err := s.client.Do(ctx, req, dates)
	if err != nil {
		return nil, resp, err
	}

	return dates, resp, nil
}

// getReviewsPage fetches a single page of a movie's reviews.
func (s *MoviesService) getReviewsPage(ctx context.Context, movieID int64, opts *uri.ListOptions) (*str.MovieReviews, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/reviews", movieID), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	reviews := new(str.MovieReviews)
	resp, err := s.client.Do(ctx, req, reviews)
	if err != nil {
		return nil, resp, err
	}

	return reviews, resp, nil
}

// GetReviews returns a movie's reviews, walking pages until TMDB reports no
// more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/movie-reviews
func (s *MoviesService) GetReviews(ctx context.Context, movieID int64, language string, pagesLimit int) ([]str.MovieReview, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.MovieReview], error) {
		reviews, _, err := s.getReviewsPage(ctx, movieID, &uri.ListOptions{Page: page, Language: language})
		if err != nil {
			return PageResult[str.MovieReview]{}, err
		}
		return PageResult[str.MovieReview]{
			Results:    reviews.Results,
			Page:       reviews.Page,
			TotalPages: reviews.TotalPages,
		}, nil
	})
}

// getSimilarPage fetches a single page of a movie's similar-movies list.
func (s *MoviesService) getSimilarPage(ctx context.Context, movieID int64, opts *uri.ListOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/similar", movieID), opts)
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

// GetSimilar returns movies similar to a single movie, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/movie-similar
func (s *MoviesService) GetSimilar(ctx context.Context, movieID int64, language string, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getSimilarPage(ctx, movieID, &uri.ListOptions{Page: page, Language: language})
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

// getTopRatedMoviesPage fetches a single page of the top-rated-movies list.
func (s *MoviesService) getTopRatedMoviesPage(ctx context.Context, opts *uri.ListOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery("movie/top_rated", opts)
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

// GetTopRatedMovies returns the current top-rated-movies list, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/movie-top-rated-list
func (s *MoviesService) GetTopRatedMovies(ctx context.Context, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getTopRatedMoviesPage(ctx, &uri.ListOptions{Page: page})
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

// GetTranslations fetches the translated fields for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-translations
func (s *MoviesService) GetTranslations(ctx context.Context, movieID int64) (*str.Translations, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d/translations", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	translations := new(str.Translations)
	resp, err := s.client.Do(ctx, req, translations)
	if err != nil {
		return nil, resp, err
	}

	return translations, resp, nil
}

// getUpcomingMoviesPage fetches a single page of the upcoming-movies list.
func (s *MoviesService) getUpcomingMoviesPage(ctx context.Context, opts *uri.ListOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery("movie/upcoming", opts)
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

// GetUpcomingMovies returns the current upcoming-movies list, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/movie-upcoming-list
func (s *MoviesService) GetUpcomingMovies(ctx context.Context, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getUpcomingMoviesPage(ctx, &uri.ListOptions{Page: page})
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

// GetVideos fetches the videos (trailers, teasers, ...) for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-videos
func (s *MoviesService) GetVideos(ctx context.Context, movieID int64, language string) (*str.MovieVideos, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/videos", movieID), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	videos := new(str.MovieVideos)
	resp, err := s.client.Do(ctx, req, videos)
	if err != nil {
		return nil, resp, err
	}

	return videos, resp, nil
}

// GetWatchProviders fetches the per-region streaming/rental/purchase
// providers for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-watch-providers
func (s *MoviesService) GetWatchProviders(ctx context.Context, movieID int64) (*str.MovieWatchProviders, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d/watch/providers", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	providers := new(str.MovieWatchProviders)
	resp, err := s.client.Do(ctx, req, providers)
	if err != nil {
		return nil, resp, err
	}

	return providers, resp, nil
}

// AddRating rates a movie on behalf of the session's account, or a guest
// session when guestSessionID is set instead of sessionID.
//
// Api docs: https://developer.themoviedb.org/reference/movie-add-rating
func (s *MoviesService) AddRating(ctx context.Context, movieID int64, sessionID, guestSessionID string, value float64) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/rating", movieID), &uri.RatingOptions{SessionID: sessionID, GuestSessionID: guestSessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, &str.MovieRatingRequest{Value: value})
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

// DeleteRating removes the session's account's rating for a movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-delete-rating
func (s *MoviesService) DeleteRating(ctx context.Context, movieID int64, sessionID string) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/rating", movieID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodDelete, urlStr, nil)
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

// getPopularMoviesPage fetches a single page of the popular-movies list.
func (s *MoviesService) getPopularMoviesPage(ctx context.Context, opts *uri.ListOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery("movie/popular", opts)
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

// GetPopularMovies returns the current popular-movies list, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited). TMDB's page size is fixed at 20 by the API; pagesLimit is
// the only real lever over how much gets fetched.
//
// Api docs: https://developer.themoviedb.org/reference/movie-popular-list
func (s *MoviesService) GetPopularMovies(ctx context.Context, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getPopularMoviesPage(ctx, &uri.ListOptions{Page: page})
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
