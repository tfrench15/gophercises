package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type chapter struct {
	Title   string
	Story   []string
	Options []option
}

type option struct {
	Text string
	Arc  string
}

func openAndReadFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func parseJSON(b []byte) (map[string]chapter, error) {
	s := make(map[string]chapter)
	err := json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
