package internal

import "context"

// PageResult is one page of a TMDB list endpoint's results, read from the
// JSON body (page/total_pages) — TMDB does not expose pagination via headers.
type PageResult[T any] struct {
	Results    []T
	Page       int
	TotalPages int
}

// FetchAllPages calls fetch for page 1, 2, ... until either TMDB reports no
// more pages (Page >= TotalPages) or pagesLimit is reached (0 = unlimited),
// accumulating every page's Results into a single slice.
func FetchAllPages[T any](ctx context.Context, pagesLimit int, fetch func(ctx context.Context, page int) (PageResult[T], error)) ([]T, error) {
	var all []T
	page := 1
	for {
		pr, err := fetch(ctx, page)
		if err != nil {
			return nil, err
		}
		all = append(all, pr.Results...)

		if pr.TotalPages <= page || (pagesLimit > 0 && page >= pagesLimit) {
			break
		}
		page++
	}
	return all, nil
}
