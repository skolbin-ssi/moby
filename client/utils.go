package client // import "github.com/docker/docker/client"

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/errdefs"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

var headerRegexp = regexp.MustCompile(`\ADocker/.+\s\((.+)\)\z`)

// getDockerOS returns the operating system based on the server header from the daemon.
func getDockerOS(serverHeader string) string {
	var osType string
	matches := headerRegexp.FindStringSubmatch(serverHeader)
	if len(matches) > 0 {
		osType = matches[1]
	}
	return osType
}

// getFiltersQuery returns a url query with "filters" query term, based on the
// filters provided.
func getFiltersQuery(f filters.Args) (url.Values, error) {
	query := url.Values{}
	if f.Len() > 0 {
		filterJSON, err := filters.ToJSON(f)
		if err != nil {
			return query, err
		}
		query.Set("filters", filterJSON)
	}
	return query, nil
}

// encodePlatforms marshals the given platform(s) to JSON format, to
// be used for query-parameters for filtering / selecting platforms.
func encodePlatforms(platform ...ocispec.Platform) ([]string, error) {
	if len(platform) == 0 {
		return []string{}, nil
	}
	if len(platform) == 1 {
		p, err := encodePlatform(&platform[0])
		if err != nil {
			return nil, err
		}
		return []string{p}, nil
	}

	seen := make(map[string]struct{}, len(platform))
	out := make([]string, 0, len(platform))
	for i := range platform {
		p, err := encodePlatform(&platform[i])
		if err != nil {
			return nil, err
		}
		if _, ok := seen[p]; !ok {
			out = append(out, p)
			seen[p] = struct{}{}
		}
	}
	return out, nil
}

// encodePlatforms marshals the given platform to JSON format, to
// be used for query-parameters for filtering / selecting platforms. It
// is used as a helper for encodePlatforms,
func encodePlatform(platform *ocispec.Platform) (string, error) {
	p, err := json.Marshal(platform)
	if err != nil {
		return "", errdefs.InvalidParameter(fmt.Errorf("invalid platform: %v", err))
	}
	return string(p), nil
}
