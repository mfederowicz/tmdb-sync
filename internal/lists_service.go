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

// CreateList creates a new list owned by the session's account.
//
// Api docs: https://developer.themoviedb.org/reference/list-create
func (s *ListsService) CreateList(ctx context.Context, sessionID string, body *str.ListCreateRequest) (*str.ListCreateResponse, *str.Response, error) {
	urlStr, err := uri.AddQuery("list", &uri.ListMutationOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, nil, err
	}

	created := new(str.ListCreateResponse)
	resp, err := s.client.Do(ctx, req, created)
	if err != nil {
		return nil, resp, err
	}

	return created, resp, nil
}

// AddMovie adds a movie to a list.
//
// Api docs: https://developer.themoviedb.org/reference/list-add-movie
func (s *ListsService) AddMovie(ctx context.Context, listID string, sessionID string, movieID int64) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("list/%s/add_item", listID), &uri.ListMutationOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, &str.ListItemRequest{MediaID: movieID})
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

// RemoveMovie removes a movie from a list.
//
// Api docs: https://developer.themoviedb.org/reference/list-remove-movie
func (s *ListsService) RemoveMovie(ctx context.Context, listID string, sessionID string, movieID int64) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("list/%s/remove_item", listID), &uri.ListMutationOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, &str.ListItemRequest{MediaID: movieID})
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

// ClearList removes all items from a list.
//
// Api docs: https://developer.themoviedb.org/reference/list-clear
func (s *ListsService) ClearList(ctx context.Context, listID string, sessionID string) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("list/%s/clear", listID), &uri.ListClearOptions{SessionID: sessionID, Confirm: true})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, nil)
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

// DeleteList deletes a list.
//
// Api docs: https://developer.themoviedb.org/reference/list-delete
func (s *ListsService) DeleteList(ctx context.Context, listID string, sessionID string) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("list/%s", listID), &uri.ListMutationOptions{SessionID: sessionID})
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
