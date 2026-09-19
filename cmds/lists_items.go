package cmds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mfederowicz/tmdb-sync/str"
)

// itemFlag collects a repeatable -item <movie|tv>:<id>[:<comment>] flag; the
// comment (which may itself contain colons) is only used by update-items.
type itemFlag []str.ListMediaV4

func (f *itemFlag) String() string {
	parts := make([]string, 0, len(*f))
	for _, item := range *f {
		value := fmt.Sprintf("%s:%d", item.MediaType, item.MediaID)
		if item.Comment != "" {
			value += ":" + item.Comment
		}
		parts = append(parts, value)
	}
	return strings.Join(parts, ",")
}

// Set parses one <movie|tv>:<id>[:<comment>] value.
func (f *itemFlag) Set(value string) error {
	mediaType, rest, ok := strings.Cut(value, ":")
	rawID, comment, _ := strings.Cut(rest, ":")
	if !ok || (mediaType != "movie" && mediaType != "tv") {
		return fmt.Errorf("invalid item %q, want movie:<id>[:<comment>] or tv:<id>[:<comment>]", value)
	}
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil || id <= 0 {
		return fmt.Errorf("invalid item %q, id must be a positive number", value)
	}
	*f = append(*f, str.ListMediaV4{MediaType: mediaType, MediaID: id, Comment: comment})
	return nil
}
