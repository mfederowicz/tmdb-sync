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

// GetDetails fetches an account's details. If accountID is 0, it resolves
// the account belonging to sessionID (TMDB's session-only GET /account form)
// instead of requiring the id up front.
//
// Api docs: https://developer.themoviedb.org/reference/account-details
func (s *AccountService) GetDetails(ctx context.Context, accountID int64, sessionID string) (*str.Account, *str.Response, error) {
	path := "account"
	if accountID != 0 {
		path = fmt.Sprintf("account/%d", accountID)
	}

	urlStr, err := uri.AddQuery(path, &uri.AccountOptions{SessionID: sessionID})
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
