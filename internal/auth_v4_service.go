// Package internal used for client and services
package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// AuthV4Service handles communication with the TMDB v4 authentication
// endpoints (request token + access token, https://developer.themoviedb.org/v4/reference/auth-create-request-token).
type AuthV4Service Service

// CreateRequestToken requests a new, unapproved v4 request token. redirectTo
// is optional: when set, TMDB redirects the browser there after approval.
//
// Api docs: https://developer.themoviedb.org/v4/reference/auth-create-request-token
func (s *AuthV4Service) CreateRequestToken(ctx context.Context, redirectTo string) (*str.RequestTokenV4, *str.Response, error) {
	var body any
	if redirectTo != "" {
		body = struct {
			RedirectTo string `json:"redirect_to"`
		}{RedirectTo: redirectTo}
	}

	req, err := s.client.NewRequestV4(http.MethodPost, "auth/request_token", body)
	if err != nil {
		return nil, nil, err
	}

	token := new(str.RequestTokenV4)
	resp, err := s.client.Do(ctx, req, token)
	if err != nil {
		return nil, resp, err
	}

	return token, resp, nil
}

// CreateAccessToken exchanges an approved v4 request token for a user access
// token and account id.
//
// Api docs: https://developer.themoviedb.org/v4/reference/auth-create-access-token
func (s *AuthV4Service) CreateAccessToken(ctx context.Context, requestToken string) (*str.AccessTokenV4, *str.Response, error) {
	body := struct {
		RequestToken string `json:"request_token"`
	}{RequestToken: requestToken}

	req, err := s.client.NewRequestV4(http.MethodPost, "auth/access_token", body)
	if err != nil {
		return nil, nil, err
	}

	accessToken := new(str.AccessTokenV4)
	resp, err := s.client.Do(ctx, req, accessToken)
	if err != nil {
		return nil, resp, err
	}

	return accessToken, resp, nil
}

// DeleteAccessToken logs out by invalidating a v4 user access token.
//
// Api docs: https://developer.themoviedb.org/v4/reference/auth-delete-access-token
func (s *AuthV4Service) DeleteAccessToken(ctx context.Context, accessToken string) (*str.AuthStatus, *str.Response, error) {
	body := struct {
		AccessToken string `json:"access_token"`
	}{AccessToken: accessToken}

	req, err := s.client.NewRequestV4(http.MethodDelete, "auth/access_token", body)
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
