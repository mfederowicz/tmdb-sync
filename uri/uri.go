// Package uri provides URL/query-string helpers shared by internal services.
package uri

import (
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/mfederowicz/tmdb-sync/consts"
)

// SanitizeURL redacts the api_key query parameter, if present, so it never
// leaks into logs or error messages.
func SanitizeURL(u *url.URL) *url.URL {
	if u == nil {
		return u
	}
	q := u.Query()
	if q.Get("api_key") != "" {
		q.Set("api_key", "REDACTED")
	}
	uCopy := *u
	uCopy.RawQuery = q.Encode()
	return &uCopy
}

// AccountOptions carries the session_id query parameter required by every
// account (🔒) endpoint.
type AccountOptions struct {
	SessionID string `url:"session_id,omitempty"`
}

// AccountListOptions carries the session_id and page query parameters
// required by paginated account (🔒) list endpoints.
type AccountListOptions struct {
	SessionID string `url:"session_id,omitempty"`
	Page      int    `url:"page,omitempty"`
}

// AccountV4Options carries the query parameters of the v4 account list
// endpoints; which of Language/SortBy an endpoint accepts varies.
type AccountV4Options struct {
	Page     int    `url:"page,omitempty"`
	Language string `url:"language,omitempty"`
	SortBy   string `url:"sort_by,omitempty"`
}

// RatingOptions carries the session_id or guest_session_id query parameter
// accepted by movie/tv/tv-episode add-rating (🔒) endpoints, which take
// either an account session or a guest session.
type RatingOptions struct {
	SessionID      string `url:"session_id,omitempty"`
	GuestSessionID string `url:"guest_session_id,omitempty"`
}

// ChangesOptions carries the optional query parameters accepted by the
// movie/tv/person change-list endpoints.
type ChangesOptions struct {
	StartDate string `url:"start_date,omitempty"`
	EndDate   string `url:"end_date,omitempty"`
	Page      int    `url:"page,omitempty"`
}

// ListOptions specifies the optional query parameters accepted by TMDB list
// endpoints. Fields are added as modules need them rather than mirroring the
// full TMDB param surface up front.
type ListOptions struct {
	Page         int    `url:"page,omitempty"`
	Language     string `url:"language,omitempty"`
	Region       string `url:"region,omitempty"`
	IncludeAdult bool   `url:"include_adult,omitempty"`
}

// ListItemStatusOptions carries the query parameters accepted by
// GET /list/{list_id}/item_status.
type ListItemStatusOptions struct {
	MovieID  int64  `url:"movie_id,omitempty"`
	Language string `url:"language,omitempty"`
}

// ListMutationOptions carries the session_id query parameter required by
// the list create/add-item/remove-item/delete (🔒) endpoints.
type ListMutationOptions struct {
	SessionID string `url:"session_id,omitempty"`
}

// ListClearOptions carries the query parameters required by
// POST /list/{list_id}/clear (🔒) — TMDB requires an explicit confirm=true
// alongside the session_id to guard against accidental clears.
type ListClearOptions struct {
	SessionID string `url:"session_id,omitempty"`
	Confirm   bool   `url:"confirm,omitempty"`
}

// MovieAlternativeTitlesOptions carries the optional query parameter
// accepted by GET /movie/{movie_id}/alternative_titles.
type MovieAlternativeTitlesOptions struct {
	Country string `url:"country,omitempty"`
}

// SearchCollectionOptions carries the query parameters accepted by
// GET /search/collection.
type SearchCollectionOptions struct {
	Query        string `url:"query,omitempty"`
	Page         int    `url:"page,omitempty"`
	Language     string `url:"language,omitempty"`
	IncludeAdult bool   `url:"include_adult,omitempty"`
	Region       string `url:"region,omitempty"`
}

// SearchCompanyOptions carries the query parameters accepted by
// GET /search/company.
type SearchCompanyOptions struct {
	Query string `url:"query,omitempty"`
	Page  int    `url:"page,omitempty"`
}

// SearchKeywordOptions carries the query parameters accepted by
// GET /search/keyword.
type SearchKeywordOptions struct {
	Query string `url:"query,omitempty"`
	Page  int    `url:"page,omitempty"`
}

// SearchMovieOptions carries the query parameters accepted by
// GET /search/movie.
type SearchMovieOptions struct {
	Query              string `url:"query,omitempty"`
	Page               int    `url:"page,omitempty"`
	Language           string `url:"language,omitempty"`
	IncludeAdult       bool   `url:"include_adult,omitempty"`
	Region             string `url:"region,omitempty"`
	Year               int    `url:"year,omitempty"`
	PrimaryReleaseYear int    `url:"primary_release_year,omitempty"`
}

// SearchMultiOptions carries the query parameters accepted by
// GET /search/multi.
type SearchMultiOptions struct {
	Query        string `url:"query,omitempty"`
	Page         int    `url:"page,omitempty"`
	Language     string `url:"language,omitempty"`
	IncludeAdult bool   `url:"include_adult,omitempty"`
}

// SearchPersonOptions carries the query parameters accepted by
// GET /search/person.
type SearchPersonOptions struct {
	Query        string `url:"query,omitempty"`
	Page         int    `url:"page,omitempty"`
	Language     string `url:"language,omitempty"`
	IncludeAdult bool   `url:"include_adult,omitempty"`
}

// SearchTVOptions carries the query parameters accepted by GET /search/tv.
type SearchTVOptions struct {
	Query            string `url:"query,omitempty"`
	Page             int    `url:"page,omitempty"`
	Language         string `url:"language,omitempty"`
	IncludeAdult     bool   `url:"include_adult,omitempty"`
	FirstAirDateYear int    `url:"first_air_date_year,omitempty"`
	Year             int    `url:"year,omitempty"`
}

// TrendingOptions carries the query parameters accepted by the
// GET /trending/{media_type}/{time_window} endpoints.
type TrendingOptions struct {
	Page int `url:"page,omitempty"`
}

// LanguageOptions carries the optional language query parameter accepted by
// several per-resource endpoints (e.g. movie/{id}/credits, /videos, ...).
type LanguageOptions struct {
	Language string `url:"language,omitempty"`
}

// ImagesOptions carries the optional query parameters accepted by
// per-resource images endpoints.
type ImagesOptions struct {
	Language             string `url:"language,omitempty"`
	IncludeImageLanguage string `url:"include_image_language,omitempty"`
}

// DiscoverMovieOptions carries the optional query parameters accepted by
// GET /discover/movie. Fields are added as modules need them rather than
// mirroring TMDB's full discover-movie filter surface up front.
type DiscoverMovieOptions struct {
	Page               int     `url:"page,omitempty"`
	Language           string  `url:"language,omitempty"`
	Region             string  `url:"region,omitempty"`
	SortBy             string  `url:"sort_by,omitempty"`
	IncludeAdult       bool    `url:"include_adult,omitempty"`
	PrimaryReleaseYear int     `url:"primary_release_year,omitempty"`
	WithGenres         string  `url:"with_genres,omitempty"`
	VoteAverageGTE     float64 `url:"vote_average.gte,omitempty"`
	VoteAverageLTE     float64 `url:"vote_average.lte,omitempty"`
	WithWatchProviders string  `url:"with_watch_providers,omitempty"`
	WatchRegion        string  `url:"watch_region,omitempty"`
}

// DiscoverTVOptions carries the optional query parameters accepted by
// GET /discover/tv. Fields are added as modules need them rather than
// mirroring TMDB's full discover-tv filter surface up front.
type DiscoverTVOptions struct {
	Page               int     `url:"page,omitempty"`
	Language           string  `url:"language,omitempty"`
	SortBy             string  `url:"sort_by,omitempty"`
	IncludeAdult       bool    `url:"include_adult,omitempty"`
	FirstAirDateYear   int     `url:"first_air_date_year,omitempty"`
	WithGenres         string  `url:"with_genres,omitempty"`
	VoteAverageGTE     float64 `url:"vote_average.gte,omitempty"`
	VoteAverageLTE     float64 `url:"vote_average.lte,omitempty"`
	WithWatchProviders string  `url:"with_watch_providers,omitempty"`
	WatchRegion        string  `url:"watch_region,omitempty"`
}

// WatchProviderListOptions carries the optional query parameters accepted by
// GET /watch/providers/movie and GET /watch/providers/tv.
type WatchProviderListOptions struct {
	Language    string `url:"language,omitempty"`
	WatchRegion string `url:"watch_region,omitempty"`
}

// GuestSessionListOptions carries the optional query parameters accepted by
// the /guest_session/{guest_session_id}/rated/... endpoints.
type GuestSessionListOptions struct {
	Language string `url:"language,omitempty"`
	SortBy   string `url:"sort_by,omitempty"`
	Page     int    `url:"page,omitempty"`
}

// FindOptions carries the query parameters accepted by GET /find/{external_id}.
type FindOptions struct {
	ExternalSource string `url:"external_source,omitempty"`
	Language       string `url:"language,omitempty"`
}

// AddQuery adds opts' non-zero fields to s as URL query parameters, sorted by
// key. opts may be a struct or pointer to struct (including nil, which adds
// nothing).
func AddQuery(s string, opts any) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs := url.Values{}
	if err := flatOptsStruct(reflect.ValueOf(opts), &qs); err != nil {
		return "", err
	}
	u.RawQuery = encodeParams(qs)
	return u.String(), nil
}

func flatOptsStruct(v reflect.Value, qs *url.Values) error {
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	for i := consts.ZeroValue; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldTag := v.Type().Field(i).Tag.Get("url")
		if fieldTag == consts.EmptyString {
			continue
		}
		fieldTag = strings.Split(fieldTag, ",")[consts.ZeroValue]
		addFieldValue(qs, fieldTag, fieldValue)
	}
	return nil
}

func addFieldValue(qs *url.Values, fieldTag string, fieldValue reflect.Value) {
	if isEmptyValue(fieldValue) {
		return
	}

	var value string
	switch fieldValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = strconv.FormatInt(fieldValue.Int(), 10)
	case reflect.Float32, reflect.Float64:
		value = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
	case reflect.Bool:
		value = strconv.FormatBool(fieldValue.Bool())
	case reflect.String:
		value = fieldValue.String()
	default:
		return
	}
	qs.Add(fieldTag, value)
}

// isEmptyValue reports whether v is the zero value of its type.
func isEmptyValue(v reflect.Value) bool {
	zero := reflect.Zero(v.Type())
	return reflect.DeepEqual(v.Interface(), zero.Interface())
}

// encodeParams encodes values sorted by key.
func encodeParams(values url.Values) string {
	if len(values) == consts.ZeroValue {
		return consts.EmptyString
	}
	keys := make([]string, consts.ZeroValue, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf strings.Builder
	for _, k := range keys {
		for _, v := range values[k] {
			if buf.Len() > consts.ZeroValue {
				buf.WriteByte('&')
			}
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}
