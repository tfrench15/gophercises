package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := pathsToUrls[r.URL.Path]
		if path == "" {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, path, http.StatusFound)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return an http.HandlerFunc (which
// also implements http.Handler) that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the fallback http.Handler will be
// called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	ymlMap := yamlMap{}
	err := yaml.Unmarshal([]byte(yml), &ymlMap)
	if err != nil {
		return nil, err
	}
	fmt.Println(ymlMap)
	return MapHandler(nil, fallback), nil
}

type yamlMap struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
