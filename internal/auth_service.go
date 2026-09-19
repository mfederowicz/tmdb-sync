// Package internal used for client and services
package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// AuthService handles communication with the TMDB authentication endpoints:
// v3 (request token + session, https://developer.themoviedb.org/reference/authentication-create-request-token)
// and, in the V4-suffixed methods, v4 (request token + user access token).
type AuthService Service

// CreateRequestToken requests a new, unapproved request token.
//
// Api docs: https://developer.themoviedb.org/reference/authentication-create-request-token
func (s *AuthService) CreateRequestToken(ctx context.Context) (*str.RequestToken, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "authentication/token/new", nil)
	if err != nil {
		return nil, nil, err
	}

	token := new(str.RequestToken)
	resp, err := s.client.Do(ctx, req, token)
	if err != nil {
		return nil, resp, err
	}

	return token, resp, nil
}

// CreateSession exchanges an approved request token for a session id.
//
// Api docs: https://developer.themoviedb.org/reference/authentication-create-session
func (s *AuthService) CreateSession(ctx context.Context, requestToken string) (*str.Session, *str.Response, error) {
	body := struct {
		RequestToken string `json:"request_token"`
	}{RequestToken: requestToken}

	req, err := s.client.NewRequest(http.MethodPost, "authentication/session/new", body)
	if err != nil {
		return nil, nil, err
	}

	session := new(str.Session)
	resp, err := s.client.Do(ctx, req, session)
	if err != nil {
		return nil, resp, err
	}

	return session, resp, nil
}

// ValidateKey confirms the configured api_key/read_access_token is valid.
//
// Api docs: https://developer.themoviedb.org/reference/authentication-valid-key
func (s *AuthService) ValidateKey(ctx context.Context) (*str.AuthStatus, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "authentication", nil)
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

// CreateGuestSession requests a new guest session (no login required, used
// for temporary rating without a full account session).
//
// Api docs: https://developer.themoviedb.org/reference/authentication-create-guest-session
func (s *AuthService) CreateGuestSession(ctx context.Context) (*str.GuestSession, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "authentication/guest_session/new", nil)
	if err != nil {
		return nil, nil, err
	}

	guestSession := new(str.GuestSession)
	resp, err := s.client.Do(ctx, req, guestSession)
	if err != nil {
		return nil, resp, err
	}

	return guestSession, resp, nil
}

// DeleteSession logs out by invalidating a session id.
//
// Api docs: https://developer.themoviedb.org/reference/authentication-delete-session
func (s *AuthService) DeleteSession(ctx context.Context, sessionID string) (*str.AuthStatus, *str.Response, error) {
	body := struct {
		SessionID string `json:"session_id"`
	}{SessionID: sessionID}

	req, err := s.client.NewRequest(http.MethodDelete, "authentication/session", body)
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

// CreateRequestTokenV4 requests a new, unapproved v4 request token. redirectTo
// is optional: when set, TMDB redirects the browser there after approval.
//
// Api docs: https://developer.themoviedb.org/v4/reference/auth-create-request-token
func (s *AuthService) CreateRequestTokenV4(ctx context.Context, redirectTo string) (*str.RequestTokenV4, *str.Response, error) {
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

// CreateAccessTokenV4 exchanges an approved v4 request token for a user access
// token and account id.
//
// Api docs: https://developer.themoviedb.org/v4/reference/auth-create-access-token
func (s *AuthService) CreateAccessTokenV4(ctx context.Context, requestToken string) (*str.AccessTokenV4, *str.Response, error) {
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

// DeleteAccessTokenV4 logs out by invalidating a v4 user access token.
//
// Api docs: https://developer.themoviedb.org/v4/reference/auth-delete-access-token
func (s *AuthService) DeleteAccessTokenV4(ctx context.Context, accessToken string) (*str.AuthStatus, *str.Response, error) {
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
