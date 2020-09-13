package main

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlURL struct {
	Path string `yaml: "path"`
	URL  string `yaml: "url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		dest, isPresent := pathsToUrls[path]
		if isPresent {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathsToUrls := yamlToMap(pathURLs)

	return MapHandler(pathsToUrls, fallback), nil
}

func parseYAML(yml []byte) ([]yamlURL, error) {
	var pathURLs []yamlURL
	err := yaml.Unmarshal(yml, &pathURLs)
	return pathURLs, err
}

func yamlToMap(pathURLs []yamlURL) map[string]string {
	pathsToUrls := map[string]string{}
	for _, el := range pathURLs {
		pathsToUrls[el.Path] = el.URL
	}
	return pathsToUrls
}
