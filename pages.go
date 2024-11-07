package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed web/*
var webFiles embed.FS

func initPages() {
	var err error
	tmpl, err := template.ParseFS(webFiles, "web/layout.html", "web/index.html", "web/logon.html")
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Templates: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/")
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
		data := struct {
			Title string
		}{
			Title: "Login",
		}
		// Nutzt layout.html und logon.html
		err := tmpl.ExecuteTemplate(w, "layout.html", data)
		if err != nil {
			http.Error(w, "Fehler beim Rendern des Templates", http.StatusInternalServerError)
		}
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
