package main

import (
	"log"
	"net/http"
)

func rootpage(w http.ResponseWriter, r *http.Request) {
	// Die Seitendaten sowie die Konfiguration werden abgerufen
	siteDataAndConfigs := getPageDataConfig()

	// Bestimmen, welche Seite gerendert werden soll
	var tmplName string
	switch r.URL.Path {
	case "/logon", "/logon.html":
		tmplName = "logon.html"
		siteDataAndConfigs.PageTitel = "Anmeldung"
	case "/register", "/register.html":
		tmplName = "logon.html"
		siteDataAndConfigs.PageTitel = "Registrieren"
	default:
		http.NotFound(w, r)
		return
	}

	// Content-Type setzen
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Template ausführen
	err := templates.ExecuteTemplate(w, tmplName, siteDataAndConfigs)
	if err != nil {
		log.Printf("Fehler beim Ausführen des Templates %s: %v", tmplName, err)
		http.Error(w, "Fehler beim Ausführen des Templates: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
