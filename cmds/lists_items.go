package cmds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mfederowicz/tmdb-sync/str"
)

// itemFlag collects a repeatable -item <movie|tv>:<id> flag.
type itemFlag []str.ListMediaV4

func (f *itemFlag) String() string {
	parts := make([]string, 0, len(*f))
	for _, item := range *f {
		parts = append(parts, fmt.Sprintf("%s:%d", item.MediaType, item.MediaID))
	}
	return strings.Join(parts, ",")
}

// Set parses one <movie|tv>:<id> value.
func (f *itemFlag) Set(value string) error {
	mediaType, rawID, ok := strings.Cut(value, ":")
	if !ok || (mediaType != "movie" && mediaType != "tv") {
		return fmt.Errorf("invalid item %q, want movie:<id> or tv:<id>", value)
	}
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil || id <= 0 {
		return fmt.Errorf("invalid item %q, id must be a positive number", value)
	}
	*f = append(*f, str.ListMediaV4{MediaType: mediaType, MediaID: id})
	return nil
}
