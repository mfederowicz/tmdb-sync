package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// ListsService handles communication with the /list endpoints of the TMDB
// API. Read endpoints (Details, Check Item Status) are public; mutation
// endpoints (Create, Add/Remove Movie, Clear, Delete) require a v3 session.
type ListsService Service

// GetList fetches details for a single list by TMDB list id.
//
// Api docs: https://developer.themoviedb.org/reference/list-details
func (s *ListsService) GetList(ctx context.Context, listID string) (*str.List, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("list/%s", listID), nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(str.List)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, nil
}

// GetItemStatus reports whether a movie is present on a list.
//
// Api docs: https://developer.themoviedb.org/reference/list-check-item-status
func (s *ListsService) GetItemStatus(ctx context.Context, listID string, movieID int64) (*str.ListItemStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("list/%s/item_status", listID), &uri.ListItemStatusOptions{MovieID: movieID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	status := new(str.ListItemStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
