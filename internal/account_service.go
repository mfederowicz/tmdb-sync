package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountService handles communication with the /account 🔒 endpoints of the
// TMDB API. Every method requires a valid v3 session id.
type AccountService Service

// GetDetails fetches an account's details.
//
// Api docs: https://developer.themoviedb.org/reference/account-details
func (s *AccountService) GetDetails(ctx context.Context, accountID int64, sessionID string) (*str.Account, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("account/%d", accountID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(str.Account)
	resp, err := s.client.Do(ctx, req, account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}
