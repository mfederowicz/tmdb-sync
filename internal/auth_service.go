// Package internal used for client and services
package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// AuthService handles communication with the TMDB v3 authentication
// endpoints (request token + session, https://developer.themoviedb.org/reference/authentication-create-request-token).
type AuthService Service

// CreateRequestToken requests a new, unapproved request token.
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
