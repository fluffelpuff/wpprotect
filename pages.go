package main

import (
	"embed"
	"net/http"
)

//go:embed web/*
var webFiles embed.FS

func initPages() {
	// Handler für die Startseite (root path /) - lädt die index.html manuell
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Datei lesen
		data, err := webFiles.ReadFile("web/index.html")
		if err != nil {
			http.Error(w, "Datei nicht gefunden", http.StatusNotFound)
			return
		}

		// Content-Type setzen und Datei senden
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})
}
