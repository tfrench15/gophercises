package main

import (
	"html/template"
	"net/http"
)

// Adventure contains the HTML templates and logic governing the
// story.
type Adventure struct {
	tmpl template.Template
}

// ServeHTTP implements the HTTP Handler interface.
func (a *Adventure) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO...
}
