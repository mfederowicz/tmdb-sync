package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// TVService handles communication with the /tv endpoints of the TMDB API.
type TVService Service

// GetTV fetches details for a single TV series by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-details
func (s *TVService) GetTV(ctx context.Context, seriesID int64) (*str.TV, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	tv := new(str.TV)
	resp, err := s.client.Do(ctx, req, tv)
	if err != nil {
		return nil, resp, err
	}

	return tv, resp, nil
}

// GetAccountStates fetches an account's favorite/rated/watchlist status for
// a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-account-states
func (s *TVService) GetAccountStates(ctx context.Context, seriesID int64, sessionID string) (*str.TVAccountStates, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/account_states", seriesID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	states := new(str.TVAccountStates)
	resp, err := s.client.Do(ctx, req, states)
	if err != nil {
		return nil, resp, err
	}

	return states, resp, nil
}

// GetAggregateCredits fetches the cast and crew for a single TV series,
// with roles/jobs aggregated across every episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-aggregate-credits
func (s *TVService) GetAggregateCredits(ctx context.Context, seriesID int64, language string) (*str.TVAggregateCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/aggregate_credits", seriesID), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.TVAggregateCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}

// GetEpisodeGroups fetches the episode groups for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-episode-groups
func (s *TVService) GetEpisodeGroups(ctx context.Context, seriesID int64) (*str.TVEpisodeGroups, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/episode_groups", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	groups := new(str.TVEpisodeGroups)
	resp, err := s.client.Do(ctx, req, groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}

// GetExternalIDs fetches a single TV series's external ids (IMDb, TVDB, ...).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-external-ids
func (s *TVService) GetExternalIDs(ctx context.Context, seriesID int64) (*str.TVExternalIDs, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/external_ids", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	ids := new(str.TVExternalIDs)
	resp, err := s.client.Do(ctx, req, ids)
	if err != nil {
		return nil, resp, err
	}

	return ids, resp, nil
}

// GetAlternativeTitles fetches the alternative titles for a single TV
// series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-alternative-titles
func (s *TVService) GetAlternativeTitles(ctx context.Context, seriesID int64) (*str.TVAlternativeTitles, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/alternative_titles", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	titles := new(str.TVAlternativeTitles)
	resp, err := s.client.Do(ctx, req, titles)
	if err != nil {
		return nil, resp, err
	}

	return titles, resp, nil
}

// GetContentRatings fetches the per-country content ratings for a single
// TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-content-ratings
func (s *TVService) GetContentRatings(ctx context.Context, seriesID int64) (*str.TVContentRatings, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/content_ratings", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	ratings := new(str.TVContentRatings)
	resp, err := s.client.Do(ctx, req, ratings)
	if err != nil {
		return nil, resp, err
	}

	return ratings, resp, nil
}

// GetCredits fetches the cast and crew for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-credits
func (s *TVService) GetCredits(ctx context.Context, seriesID int64, language string) (*str.TVCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/credits", seriesID), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.TVCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}

// GetImages fetches the images for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-images
func (s *TVService) GetImages(ctx context.Context, seriesID int64, opts *uri.ImagesOptions) (*str.Images, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/images", seriesID), opts)
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

// GetKeywords fetches the keywords for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-keywords
func (s *TVService) GetKeywords(ctx context.Context, seriesID int64) (*str.TVKeywords, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/keywords", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	keywords := new(str.TVKeywords)
	resp, err := s.client.Do(ctx, req, keywords)
	if err != nil {
		return nil, resp, err
	}

	return keywords, resp, nil
}

// getListsPage fetches a single page of the lists a TV series belongs to.
func (s *TVService) getListsPage(ctx context.Context, seriesID int64, opts *uri.ListOptions) (*str.TVLists, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/lists", seriesID), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	lists := new(str.TVLists)
	resp, err := s.client.Do(ctx, req, lists)
	if err != nil {
		return nil, resp, err
	}

	return lists, resp, nil
}

// GetLists returns the lists a TV series belongs to, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-lists
func (s *TVService) GetLists(ctx context.Context, seriesID int64, language string, pagesLimit int) ([]str.AccountList, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.AccountList], error) {
		lists, _, err := s.getListsPage(ctx, seriesID, &uri.ListOptions{Page: page, Language: language})
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

// getRecommendationsPage fetches a single page of a TV series'
// recommendations.
func (s *TVService) getRecommendationsPage(ctx context.Context, seriesID int64, opts *uri.ListOptions) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/recommendations", seriesID), opts)
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

// GetRecommendations returns TV series recommended off a single series,
// walking pages until TMDB reports no more (total_pages) or pagesLimit is
// reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-recommendations
func (s *TVService) GetRecommendations(ctx context.Context, seriesID int64, language string, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		shows, _, err := s.getRecommendationsPage(ctx, seriesID, &uri.ListOptions{Page: page, Language: language})
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

// getReviewsPage fetches a single page of a TV series' reviews.
func (s *TVService) getReviewsPage(ctx context.Context, seriesID int64, opts *uri.ListOptions) (*str.TVReviews, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/reviews", seriesID), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	reviews := new(str.TVReviews)
	resp, err := s.client.Do(ctx, req, reviews)
	if err != nil {
		return nil, resp, err
	}

	return reviews, resp, nil
}

// GetReviews returns a TV series' reviews, walking pages until TMDB reports
// no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-reviews
func (s *TVService) GetReviews(ctx context.Context, seriesID int64, language string, pagesLimit int) ([]str.TVReview, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TVReview], error) {
		reviews, _, err := s.getReviewsPage(ctx, seriesID, &uri.ListOptions{Page: page, Language: language})
		if err != nil {
			return PageResult[str.TVReview]{}, err
		}
		return PageResult[str.TVReview]{
			Results:    reviews.Results,
			Page:       reviews.Page,
			TotalPages: reviews.TotalPages,
		}, nil
	})
}

// GetScreenedTheatrically fetches the episodes of a TV series that had a
// theatrical screening.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-screened-theatrically
func (s *TVService) GetScreenedTheatrically(ctx context.Context, seriesID int64) (*str.TVScreenedTheatrically, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/screened_theatrically", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	screened := new(str.TVScreenedTheatrically)
	resp, err := s.client.Do(ctx, req, screened)
	if err != nil {
		return nil, resp, err
	}

	return screened, resp, nil
}

// getSimilarPage fetches a single page of a TV series' similar-shows list.
func (s *TVService) getSimilarPage(ctx context.Context, seriesID int64, opts *uri.ListOptions) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/similar", seriesID), opts)
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

// GetSimilar returns TV series similar to a single series, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-similar
func (s *TVService) GetSimilar(ctx context.Context, seriesID int64, language string, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		shows, _, err := s.getSimilarPage(ctx, seriesID, &uri.ListOptions{Page: page, Language: language})
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

// GetTranslations fetches the translated fields for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-translations
func (s *TVService) GetTranslations(ctx context.Context, seriesID int64) (*str.Translations, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/translations", seriesID), nil)
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

// GetVideos fetches the videos (trailers, teasers, ...) for a single TV
// series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-videos
func (s *TVService) GetVideos(ctx context.Context, seriesID int64, language string) (*str.TVVideos, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/videos", seriesID), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	videos := new(str.TVVideos)
	resp, err := s.client.Do(ctx, req, videos)
	if err != nil {
		return nil, resp, err
	}

	return videos, resp, nil
}

// GetWatchProviders fetches the per-region streaming/rental/purchase
// providers for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-watch-providers
func (s *TVService) GetWatchProviders(ctx context.Context, seriesID int64) (*str.TVWatchProviders, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/watch/providers", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	providers := new(str.TVWatchProviders)
	resp, err := s.client.Do(ctx, req, providers)
	if err != nil {
		return nil, resp, err
	}

	return providers, resp, nil
}

// getPopularTVPage fetches a single page of the popular-tv-series list.
func (s *TVService) getPopularTVPage(ctx context.Context, opts *uri.ListOptions) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery("tv/popular", opts)
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

// GetPopularTV returns the current popular-tv-series list, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-popular-list
func (s *TVService) GetPopularTV(ctx context.Context, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		shows, _, err := s.getPopularTVPage(ctx, &uri.ListOptions{Page: page})
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

// getTopRatedTVPage fetches a single page of the top-rated-tv-series list.
func (s *TVService) getTopRatedTVPage(ctx context.Context, opts *uri.ListOptions) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery("tv/top_rated", opts)
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

// GetTopRatedTV returns the current top-rated-tv-series list, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-top-rated-list
func (s *TVService) GetTopRatedTV(ctx context.Context, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		shows, _, err := s.getTopRatedTVPage(ctx, &uri.ListOptions{Page: page})
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

// GetLatest fetches the most recently created TV series on TMDB.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-latest-id
func (s *TVService) GetLatest(ctx context.Context) (*str.TV, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "tv/latest", nil)
	if err != nil {
		return nil, nil, err
	}

	tv := new(str.TV)
	resp, err := s.client.Do(ctx, req, tv)
	if err != nil {
		return nil, resp, err
	}

	return tv, resp, nil
}
