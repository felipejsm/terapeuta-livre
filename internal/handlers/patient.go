package main

import (
	"html/template"
	"net/http"
)

type Patient struct {
	PatientHeader []Header
	PatientRow    []Row
}
type Header struct {
	Patient string
	Status  string
	Date    string
}
type Row struct {
	Name   string
	Email  string
	Status template.HTML
	Date   string
}

func main() {
	indexTmpl := template.Must(template.ParseFiles("index.html"))
	layoutTmpl := template.Must(template.ParseFiles("template.html"))

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HanleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTmpl.Execute(w, nil)
	})
	http.ListenAndServe(":21067", nil)

}
