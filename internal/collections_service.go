package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// CollectionsService handles communication with the /collection endpoints of the TMDB API.
type CollectionsService Service

// GetCollection fetches details for a single collection by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/collection-details
func (s *CollectionsService) GetCollection(ctx context.Context, collectionID int64) (*str.Collection, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("collection/%d", collectionID), nil)
	if err != nil {
		return nil, nil, err
	}

	collection := new(str.Collection)
	resp, err := s.client.Do(ctx, req, collection)
	if err != nil {
		return nil, resp, err
	}

	return collection, resp, nil
}

// GetCollectionImages fetches the backdrops/posters/logos for a single collection by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/collection-images
func (s *CollectionsService) GetCollectionImages(ctx context.Context, collectionID int64, opts *uri.ImagesOptions) (*str.Images, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("collection/%d/images", collectionID), opts)
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

// GetCollectionTranslations fetches the translated fields for a single collection by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/collection-translations
func (s *CollectionsService) GetCollectionTranslations(ctx context.Context, collectionID int64) (*str.Translations, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("collection/%d/translations", collectionID), nil)
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
