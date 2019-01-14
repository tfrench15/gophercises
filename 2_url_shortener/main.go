package main

import (
	"fmt"
	"net/http"

	"github.com/tfrench15/gophercises/2_url_shortener/urlshort"
)

func main() {
	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	_, err := urlshort.YAMLHandler([]byte(yml), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func placeholder() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yml := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	  url: http://github.com/gophercises/urlshort/tree/solution`

	ymlHandler, err := urlshort.YAMLHandler([]byte(yml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8000")
	http.ListenAndServe(":8000", ymlHandler)
}
