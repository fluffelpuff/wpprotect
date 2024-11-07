package main

import (
	"embed"
	"net/http"
)

var webFiles embed.FS

func initPages() {
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

	http.HandleFunc("/logon", func(w http.ResponseWriter, r *http.Request) {
		// Datei lesen
		data, err := webFiles.ReadFile("web/logon.html")
		if err != nil {
			http.Error(w, "Datei nicht gefunden", http.StatusNotFound)
			return
		}

		// Content-Type setzen und Datei senden
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	http.HandleFunc("/css/modal.css", func(w http.ResponseWriter, r *http.Request) {
		// Datei lesen
		data, err := webFiles.ReadFile("web/css/modal.css")
		if err != nil {
			http.Error(w, "Datei nicht gefunden", http.StatusNotFound)
			return
		}

		// Content-Type setzen und Datei senden
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})
}
