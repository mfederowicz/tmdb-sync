// Package internal used for client and services
package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

type contextKey string

// basic consts for client
const (
	TimezoneKey        contextKey = "timezone"
	skipRateLimitCheck contextKey = "skipRateLimitCheck"
	// BaseURL is the TMDB v3 API base URL.
	BaseURL = "https://api.themoviedb.org/3/"
	// BaseURLV4 is the TMDB v4 API base URL.
	BaseURLV4 = "https://api.themoviedb.org/4/"
	// APIKeyParam is the query string parameter name used for v3 API key auth.
	APIKeyParam = "api_key"
)

var errNonNilContext = errors.New("context must be non-nil")
var errV4NeedsReadToken = errors.New("tmdb: v4 endpoints need read_access_token in the config (api_key alone is not enough)")
var emptyReader = strings.NewReader("")

// RequestOption represents an option that can modify an http.Request.
type RequestOption func(req *http.Request)

// A Client manages communication with the themoviedb.org API.
type Client struct {
	client         *http.Client
	BaseURL        *url.URL
	BaseURLV4      *url.URL
	AuthURL        *url.URL
	headers        map[string]any
	common         Service
	Account        *AccountService
	Auth           *AuthService
	AuthV4         *AuthV4Service
	Certifications *CertificationsService
	Changes        *ChangesService
	Collections    *CollectionsService
	Companies      *CompaniesService
	Configuration  *ConfigurationService
	Credits        *CreditsService
	Discover       *DiscoverService
	Find           *FindService
	Genre          *GenreService
	GuestSessions  *GuestSessionsService
	Keywords       *KeywordsService
	Lists          *ListsService
	Movies         *MoviesService
	Networks       *NetworksService
	People         *PeopleService
	Reviews        *ReviewsService
	Search         *SearchService
	Trending       *TrendingService
	TV             *TVService
	TVSeasons      *TVSeasonsService
	TVEpisodes     *TVEpisodesService
	TVEpisodeGroup *TVEpisodeGroupService
	WatchProviders *WatchProvidersService
	rateMu         sync.Mutex
	RateLimitReset time.Time
	// Debug, when true, prints the full request URL for each API call to stdout.
	Debug bool
}

// GetTimezone to get timezone from ctx object
func (*Client) GetTimezone(ctx context.Context) *time.Location {
	loc, ok := ctx.Value(TimezoneKey).(*time.Location)
	if !ok {
		return time.UTC
	}
	return loc
}

// UpdateHeaders is for update client headers map
func (c *Client) UpdateHeaders(headers map[string]any) {
	c.headers = headers
}

// GetHeaders is for get headers map
func (c *Client) GetHeaders() map[string]any {
	return c.headers
}

// initialize sets default values and initializes services.
func (c *Client) initialize() {
	if c.client == nil {
		c.client = &http.Client{}
	}
	if c.BaseURL == nil {
		c.BaseURL, _ = url.Parse(BaseURL)
	}
	if c.BaseURLV4 == nil {
		c.BaseURLV4, _ = url.Parse(BaseURLV4)
	}
	c.common.client = c
	c.Account = (*AccountService)(&c.common)
	c.Auth = (*AuthService)(&c.common)
	c.AuthV4 = (*AuthV4Service)(&c.common)
	c.Certifications = (*CertificationsService)(&c.common)
	c.Changes = (*ChangesService)(&c.common)
	c.Collections = (*CollectionsService)(&c.common)
	c.Companies = (*CompaniesService)(&c.common)
	c.Configuration = (*ConfigurationService)(&c.common)
	c.Credits = (*CreditsService)(&c.common)
	c.Discover = (*DiscoverService)(&c.common)
	c.Find = (*FindService)(&c.common)
	c.Genre = (*GenreService)(&c.common)
	c.GuestSessions = (*GuestSessionsService)(&c.common)
	c.Keywords = (*KeywordsService)(&c.common)
	c.Lists = (*ListsService)(&c.common)
	c.Movies = (*MoviesService)(&c.common)
	c.Networks = (*NetworksService)(&c.common)
	c.People = (*PeopleService)(&c.common)
	c.Reviews = (*ReviewsService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Trending = (*TrendingService)(&c.common)
	c.TV = (*TVService)(&c.common)
	c.TVSeasons = (*TVSeasonsService)(&c.common)
	c.TVEpisodes = (*TVEpisodesService)(&c.common)
	c.TVEpisodeGroup = (*TVEpisodeGroupService)(&c.common)
	c.WatchProviders = (*WatchProvidersService)(&c.common)
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body any, opts ...RequestOption) (*http.Request, error) {
	return c.newRequest(c.BaseURL, true, method, urlStr, body, opts...)
}

// NewRequestV4 creates a TMDB v4 API request: it targets BaseURLV4 and
// authenticates with the bearer token only, never appending api_key.
func (c *Client) NewRequestV4(method, urlStr string, body any, opts ...RequestOption) (*http.Request, error) {
	if c.headers["Authorization"] == nil {
		return nil, errV4NeedsReadToken
	}
	return c.newRequest(c.BaseURLV4, false, method, urlStr, body, opts...)
}

func (c *Client) newRequest(base *url.URL, withAPIKey bool, method, urlStr string, body any, opts ...RequestOption) (*http.Request, error) {
	if !strings.HasSuffix(base.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", base)
	}
	u, err := base.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if apiKey, ok := c.headers[APIKeyParam]; ok && withAPIKey {
		q := u.Query()
		q.Set(APIKeyParam, fmt.Sprintf("%v", apiKey))
		u.RawQuery = q.Encode()
	}

	if c.Debug {
		printer.Println("tmdb-sync: request URL:", uri.SanitizeURL(u).String())
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req = c.requestSetHeaders(req, body)

	for _, opt := range opts {
		opt(req)
	}

	if c.Debug {
		printRequestHeaders(req)
	}

	return req, nil
}

// printRequestHeaders prints the request headers in sorted order, redacting the Authorization value.
func printRequestHeaders(req *http.Request) {
	names := make([]string, 0, len(req.Header))
	for name := range req.Header {
		names = append(names, name)
	}
	sort.Strings(names)

	printer.Println("tmdb-sync: request headers:")
	for _, name := range names {
		value := strings.Join(req.Header[name], ", ")
		if http.CanonicalHeaderKey(name) == "Authorization" {
			value = "<redacted>"
		}
		printer.Println("  " + name + ": " + value)
	}
}

func (c *Client) requestSetHeaders(r *http.Request, body any) *http.Request {
	if body != nil {
		r.Header.Set("Content-Type", "application/json")
	}

	if c.headers["Authorization"] != nil {
		r.Header.Set("Authorization", c.headers["Authorization"].(string))
	}

	return r
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred
func (c *Client) Do(ctx context.Context, req *http.Request, v any) (*str.Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		if v != nil && resp != nil && resp.StatusCode > 399 {
			json.NewDecoder(resp.Body).Decode(v)
		}
		return resp, err
	}
	defer resp.Body.Close()
	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return resp, err
}

// BareDo sends an API request and lets you handle the api response.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*str.Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	req = c.WithContext(ctx, req)

	skipResp, skipErr := c.skipCheck(ctx, req)
	if skipErr != nil {
		return skipResp, skipErr
	}

	httpResp, err := c.client.Do(req)
	if err != nil {
		return handleBareDoError(ctx, err)
	}

	resp := c.NewResponse(httpResp)

	if err := CheckResponse(httpResp); err != nil {
		return resp, err
	}

	return resp, nil
}

func handleBareDoError(ctx context.Context, err error) (*str.Response, error) {
	// If the context has been canceled, return its error.
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// Try to sanitize the URL in the error and return the sanitized error if successful.
	if sanitizedErr := sanitizeURL(err); sanitizedErr != nil {
		return nil, sanitizedErr
	}

	// Return the original error if URL sanitization fails.
	return nil, err
}

func (c *Client) skipCheck(ctx context.Context, req *http.Request) (*str.Response, error) {
	if skip := ctx.Value(skipRateLimitCheck); skip == nil {
		// don't make further requests before Retry After.
		if err := c.CheckRetryAfter(req); err != nil {
			return &str.Response{
				Response: err.Response,
			}, err
		}
	}
	return nil, nil
}

func sanitizeURL(err error) error {
	if e, ok := err.(*url.Error); ok {
		if u, err := url.Parse(e.URL); err == nil {
			e.URL = uri.SanitizeURL(u).String()
			return e
		}
	}
	return nil
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	if r.StatusCode == http.StatusTooManyRequests {
		return &AbuseRateLimitError{
			Response: r,
			Message:  "API rate limit exceeded",
		}
	}

	errorResponse := &str.ErrorResponse{}
	if err := json.NewDecoder(r.Body).Decode(errorResponse); err == nil {
		errorResponse.StatusCode = r.StatusCode
		return errorResponse
	}

	return fmt.Errorf("tmdb: unexpected status %s", r.Status)
}

// CheckRetryAfter check Retry After header.
func (c *Client) CheckRetryAfter(req *http.Request) *AbuseRateLimitError {
	c.rateMu.Lock()
	reset := c.RateLimitReset
	c.rateMu.Unlock()
	if !reset.IsZero() && time.Now().Before(reset) {
		// Create a fake response.
		resp := &http.Response{
			Status:     http.StatusText(http.StatusForbidden),
			StatusCode: http.StatusForbidden,
			Request:    req,
			Header:     make(http.Header),
			Body:       io.NopCloser(emptyReader),
		}

		retryAfter := time.Until(reset)
		return &AbuseRateLimitError{
			Response:   resp,
			Message:    fmt.Sprintf("API rate limit exceeded until %v, not making remote request.", reset),
			RetryAfter: &retryAfter,
		}
	}

	return nil
}

// WithContext pass context to request
func (*Client) WithContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

// NewResponse creates a new Response for the provided http.Response.
// r must not be nil.
func (*Client) NewResponse(r *http.Response) *str.Response {
	response := &str.Response{Response: r}
	return response
}

// NewClient returns a new API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient2 := *httpClient
	c := &Client{client: &httpClient2}
	c.initialize()
	return c
}
