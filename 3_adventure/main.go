package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	b, err := openAndReadFile("gopher.json")
	if err != nil {
		fmt.Println(err)
	}

	s, err := parseJSON(b)
	if err != nil {
		fmt.Println(err)
	}
	ch := s["sean-kelly"]
	fmt.Println(ch)

	tmpl := placeholder()
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}

	srv := new()
	http.ListenAndServe(":7000", srv)
}
