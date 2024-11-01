package urlshortener

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Route defines a single mapping from a path to a URL in the YAML configuration.
type Route struct {
	Path string `yaml:"path"` // Incoming request path
	URL  string `yaml:"url"`  // Target URL for redirection
}

// MapHandler returns an http.HandlerFunc that attempts to redirect based on
// the given pathsToUrls map. If a path isn't found, it calls the fallback handler.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Look up the requested path in the map
		if dest, found := pathsToUrls[r.URL.Path]; found {
			http.Redirect(w, r, dest, http.StatusFound) // Redirect if found
			return
		}
		// Fallback to the default handler if path is not mapped
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler parses YAML data and returns an http.HandlerFunc that maps paths to URLs.
// If the path is not found in the parsed data, it defaults to the fallback handler.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse the YAML into a slice of Route structs
	routes, err := parseYAML(yamlData)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %w", err)
	}

	// Convert the parsed routes into a map for fast lookup
	pathMap := buildPathMap(routes)

	// Return a MapHandler that uses pathMap and falls back to the provided handler
	return MapHandler(pathMap, fallback), nil
}

// parseYAML unmarshals YAML data into a slice of Route structs.
// Returns an error if the data cannot be parsed.
func parseYAML(data []byte) ([]Route, error) {
	var routes []Route
	if err := yaml.Unmarshal(data, &routes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	return routes, nil
}

// buildPathMap takes a slice of Route structs and constructs a map where the
// key is the path and the value is the target URL. This map enables efficient lookups.
func buildPathMap(routes []Route) map[string]string {
	pathMap := make(map[string]string, len(routes))
	for _, route := range routes {
		pathMap[route.Path] = route.URL
	}
	return pathMap
}
