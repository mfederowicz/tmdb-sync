package cmds

import "errors"

// apiVersion is the TMDB API version selected for one invocation.
type apiVersion string

const (
	apiV3 apiVersion = "v3"
	apiV4 apiVersion = "v4"
)

// resolveAPIVersion maps the -v3/-v4 flags to a version: v3 is the default,
// and giving both is an error.
func resolveAPIVersion(v3, v4 bool) (apiVersion, error) {
	if v3 && v4 {
		return "", errors.New("-v3 and -v4 are mutually exclusive")
	}
	if v4 {
		return apiV4, nil
	}
	return apiV3, nil
}
