package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed web/*
var webFiles embed.FS

var templates *template.Template

func initPages() {
	// Templates initialisieren
	templates = template.Must(template.ParseFS(webFiles, "web/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmplName string

		// Bestimmen, welche Seite gerendert werden soll
		switch r.URL.Path {
		case "/":
			tmplName = "index.html"
		case "/logon":
			tmplName = "logon.html"
		default:
			http.NotFound(w, r)
			return
		}

		// Template ausführen
		err := templates.ExecuteTemplate(w, tmplName, nil)
		if err != nil {
			http.Error(w, "Fehler beim Ausführen des Templates: "+err.Error(), http.StatusInternalServerError)
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
		w.Header().Set("Content-Type", "text/css")
		w.Write(data)
	})
}
