package main

import (
	"fmt"
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
}
