package main

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func initPages() {
	// Die Template Dateien werden eingelesen
	var err error
	templates, err = template.ParseFS(webFiles, "web/*.html")
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Templates: %v", err)
	}

	// Statische Dateien bedienen
	cssFiles, err := fs.Sub(webFiles, "web/css")
	if err != nil {
		log.Fatalf("Fehler beim Zugriff auf CSS-Dateien: %v", err)
	}

	// Statische Dateien bedienen
	jsFiles, err := fs.Sub(webFiles, "web/js")
	if err != nil {
		log.Fatalf("Fehler beim Zugriff auf CSS-Dateien: %v", err)
	}

	// Stellt die Seiten bereit
	http.HandleFunc("/", rootpage)

	// Stellt die CSS sowie die JS Dateien bereit
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.FS(cssFiles))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.FS(jsFiles))))
}

func getPageDataConfig() *PageConfigData {
	return &PageConfigData{
		Lang:                   german,
		WebsiteName:            "Fluffel's Place",
		ModalTitle:             "Logon",
		BsThemeDark:            "dark",
		HasThirdPartyProviders: false,
	}
}
