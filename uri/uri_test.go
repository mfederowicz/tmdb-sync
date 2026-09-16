package uri

import "testing"

func TestAddQuery_OmitsZeroValues(t *testing.T) {
	got, err := AddQuery("movie/popular", &ListOptions{Page: 2})
	if err != nil {
		t.Fatalf("AddQuery() error = %v", err)
	}
	want := "movie/popular?page=2"
	if got != want {
		t.Fatalf("AddQuery() = %q, want %q", got, want)
	}
}

func TestAddQuery_SortsKeys(t *testing.T) {
	got, err := AddQuery("movie/popular", &ListOptions{Page: 2, Region: "US", Language: "en-US"})
	if err != nil {
		t.Fatalf("AddQuery() error = %v", err)
	}
	want := "movie/popular?language=en-US&page=2&region=US"
	if got != want {
		t.Fatalf("AddQuery() = %q, want %q", got, want)
	}
}

func TestAddQuery_NilOpts(t *testing.T) {
	got, err := AddQuery("movie/popular", (*ListOptions)(nil))
	if err != nil {
		t.Fatalf("AddQuery() error = %v", err)
	}
	want := "movie/popular"
	if got != want {
		t.Fatalf("AddQuery() = %q, want %q", got, want)
	}
}

func TestAddQuery_IncludeAdultTrue(t *testing.T) {
	got, err := AddQuery("movie/popular", &ListOptions{IncludeAdult: true})
	if err != nil {
		t.Fatalf("AddQuery() error = %v", err)
	}
	want := "movie/popular?include_adult=true"
	if got != want {
		t.Fatalf("AddQuery() = %q, want %q", got, want)
	}
}
