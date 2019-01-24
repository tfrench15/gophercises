package main

import (
	"html/template"
)

func placeholder() *template.Template {
	tmpl, err := template.ParseFiles("cyoa.html")
	if err != nil {
		panic(err)
	}
	return tmpl
}
