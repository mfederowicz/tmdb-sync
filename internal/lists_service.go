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
// The V4-suffixed methods use the v4 API and the v4 user access token.
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

// GetListV4 fetches a v4 list, walking the item pages until TMDB reports no
// more (total_pages) or pagesLimit is reached (0 = unlimited). The user
// access token is optional: without it only public lists are readable.
//
// Api docs: https://developer.themoviedb.org/v4/reference/list-details
func (s *ListsService) GetListV4(ctx context.Context, accessToken, listID string, pagesLimit int, opts uri.ListV4Options) (*str.ListV4, error) {
	var reqOpts []RequestOption
	if accessToken != "" {
		reqOpts = append(reqOpts, withUserToken(accessToken))
	}

	list := new(str.ListV4)
	_, err := FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.ListItemV4], error) {
		opts.Page = page
		urlStr, err := uri.AddQuery(fmt.Sprintf("list/%s", listID), &opts)
		if err != nil {
			return PageResult[str.ListItemV4]{}, err
		}

		req, err := s.client.NewRequestV4(http.MethodGet, urlStr, nil, reqOpts...)
		if err != nil {
			return PageResult[str.ListItemV4]{}, err
		}

		body := new(str.ListV4)
		if _, err := s.client.Do(ctx, req, body); err != nil {
			return PageResult[str.ListItemV4]{}, err
		}
		results := body.Results
		if page == 1 {
			*list = *body
			list.Results = nil
		}
		list.Results = append(list.Results, results...)
		return PageResult[str.ListItemV4]{Results: results, Page: body.Page, TotalPages: body.TotalPages}, nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

// CreateListV4 creates a list owned by the v4 user access token's account.
//
// Api docs: https://developer.themoviedb.org/v4/reference/list-create
func (s *ListsService) CreateListV4(ctx context.Context, accessToken string, body *str.ListCreateRequestV4) (*str.ListCreateResponseV4, error) {
	req, err := s.client.NewRequestV4(http.MethodPost, "list", body, withUserToken(accessToken))
	if err != nil {
		return nil, err
	}

	created := new(str.ListCreateResponseV4)
	if _, err := s.client.Do(ctx, req, created); err != nil {
		return nil, err
	}

	return created, nil
}

// UpdateListV4 updates a list's name, description, visibility, sort order or
// backdrop; only the fields set on body are changed.
//
// Api docs: https://developer.themoviedb.org/v4/reference/list-update
func (s *ListsService) UpdateListV4(ctx context.Context, accessToken, listID string, body *str.ListUpdateRequestV4) (*str.ListStatusV4, error) {
	req, err := s.client.NewRequestV4(http.MethodPut, fmt.Sprintf("list/%s", listID), body, withUserToken(accessToken))
	if err != nil {
		return nil, err
	}

	status := new(str.ListStatusV4)
	if _, err := s.client.Do(ctx, req, status); err != nil {
		return nil, err
	}

	return status, nil
}
