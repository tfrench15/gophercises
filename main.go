package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func main() {
	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	ym, err := parseYAML([]byte(yml))
	if err != nil {
		fmt.Println(err)
	}
	m := buildMap(ym)
	fmt.Println(m)
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
