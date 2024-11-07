package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/*
var webFiles embed.FS

var templates *template.Template

type PageData struct {
	Title       string
	Lang        string
	BsThemeDark string
}

func initPages() {
	var err error
	templates, err = template.ParseFS(webFiles, "web/*.html")
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Templates: %v", err)
	}

	// Seiten Daten
	data := PageData{
		BsThemeDark: "auto",
		Lang:        "de",
	}

	// Geladene Templates auflisten
	for _, tmpl := range templates.Templates() {
		log.Printf("Geladenes Template: %s", tmpl.Name())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmplName string

		// Bestimmen, welche Seite gerendert werden soll
		switch r.URL.Path {
		case "/logon", "/logon.html":
			tmplName = "logon.html"
			data.Title = "Anmeldung"
		case "/register", "/register.html":
			tmplName = "logon.html"
			data.Title = "Registrieren"
		default:
			http.NotFound(w, r)
			return
		}

		// Content-Type setzen
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Template ausführen
		err := templates.ExecuteTemplate(w, tmplName, data)
		if err != nil {
			log.Printf("Fehler beim Ausführen des Templates %s: %v", tmplName, err)
			http.Error(w, "Fehler beim Ausführen des Templates: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Statische Dateien bedienen
	cssFiles, err := fs.Sub(webFiles, "web/css")
	if err != nil {
		log.Fatalf("Fehler beim Zugriff auf CSS-Dateien: %v", err)
	}

	// Stellt die CSS Dateien bereit
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.FS(cssFiles))))
}
