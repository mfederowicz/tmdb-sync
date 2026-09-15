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
