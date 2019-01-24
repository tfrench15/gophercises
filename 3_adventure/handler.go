package main

import (
	"fmt"
	"net/http"
)

type server struct{}

func new() *server {
	return &server{}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
