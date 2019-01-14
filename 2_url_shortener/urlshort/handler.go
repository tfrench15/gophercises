package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
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
	ym, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	m := buildMap(ym)
	return MapHandler(m, fallback), nil
}

func parseYAML(yml []byte) ([]yamlMap, error) {
	var ym []yamlMap
	err := yaml.Unmarshal(yml, &ym)
	if err != nil {
		return nil, err
	}
	return ym, nil
}

func buildMap(ym []yamlMap) map[string]string {
	m := make(map[string]string)
	for _, item := range ym {
		m[item.Path] = item.URL
	}
	return m
}

type yamlMap struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
