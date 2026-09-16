package internal

import (
	"context"
	"errors"
	"testing"
)

var errTest = errors.New("boom")

func TestFetchAllPages_StopsAtTotalPages(t *testing.T) {
	var callsMade int
	fetch := func(_ context.Context, page int) (PageResult[int], error) {
		callsMade++
		return PageResult[int]{Results: []int{page}, Page: page, TotalPages: 3}, nil
	}

	got, err := FetchAllPages(context.Background(), 0, fetch)
	if err != nil {
		t.Fatalf("FetchAllPages() error = %v", err)
	}
	if callsMade != 3 {
		t.Errorf("callsMade = %d, want 3", callsMade)
	}
	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i, v := range want {
		if got[i] != v {
			t.Errorf("got[%d] = %d, want %d", i, got[i], v)
		}
	}
}

func TestFetchAllPages_RespectsPagesLimit(t *testing.T) {
	var callsMade int
	fetch := func(_ context.Context, page int) (PageResult[int], error) {
		callsMade++
		return PageResult[int]{Results: []int{page}, Page: page, TotalPages: 100}, nil
	}

	got, err := FetchAllPages(context.Background(), 2, fetch)
	if err != nil {
		t.Fatalf("FetchAllPages() error = %v", err)
	}
	if callsMade != 2 {
		t.Errorf("callsMade = %d, want 2 (pagesLimit)", callsMade)
	}
	if len(got) != 2 {
		t.Errorf("len(got) = %d, want 2", len(got))
	}
}

func TestFetchAllPages_PropagatesError(t *testing.T) {
	wantErr := errTest
	fetch := func(_ context.Context, _ int) (PageResult[int], error) {
		return PageResult[int]{}, wantErr
	}

	if _, err := FetchAllPages(context.Background(), 0, fetch); err != wantErr {
		t.Errorf("FetchAllPages() error = %v, want %v", err, wantErr)
	}
}
