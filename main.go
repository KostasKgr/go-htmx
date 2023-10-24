package main

import (
	"io"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Starting server")

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, r.Method+": Hello world")
	}

	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
