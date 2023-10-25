package main

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	slog.Info("Starting server!")

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, r.Method+": Hello world")
	}

	fileHandler := func(w http.ResponseWriter, r *http.Request) {

		data := map[string][]Film{
			"Films": {
				{Title: "Film1", Director: "Director1"},
				{Title: "Film2", Director: "Director2"},
				{Title: "Film3", Director: "Director3"},
			},
		}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)

	}

	http.HandleFunc("/", fileHandler)
	http.HandleFunc("/io/write-string", homeHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
