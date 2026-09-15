package uri

import (
	"fmt"
	"net/url"
	"strconv"
)

// AddPage appends a TMDB "page" query parameter to urlStr.
func AddPage(urlStr string, page int) (string, error) {
	if page <= 0 {
		return urlStr, nil
	}
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("invalid url %q: %w", urlStr, err)
	}
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()
	return u.String(), nil
}
