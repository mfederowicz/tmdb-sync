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

// ImagesOptions carries the optional query parameters accepted by
// per-resource images endpoints.
type ImagesOptions struct {
	Language             string `url:"language,omitempty"`
	IncludeImageLanguage string `url:"include_image_language,omitempty"`
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
