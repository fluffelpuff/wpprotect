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

type ThirdPartyProvider struct {
}

type LangSpeficData struct {
	WebsiteLang                      string
	LangEMailInputFieldHiddenText    string
	LangPasswordInputFieldHiddenText string
	LangLogonButtonText              string
}

type PageData struct {
	Lang                   *LangSpeficData
	PageTitel              string
	ModalTitle             string
	BsThemeDark            string
	WebsiteName            string
	HasThirdPartyProviders bool
	ThirdPartyProviders    []*ThirdPartyProvider
}

func initPages() {
	var err error
	templates, err = template.ParseFS(webFiles, "web/*.html")
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Templates: %v", err)
	}

	// Seiten Daten
	data := PageData{
		ThirdPartyProviders:    make([]*ThirdPartyProvider, 0),
		WebsiteName:            "Fluffel's SecureZone",
		ModalTitle:             "Anmeldung",
		HasThirdPartyProviders: false,
		BsThemeDark:            "auto",
		Lang: &LangSpeficData{
			WebsiteLang:                   "de",
			LangEMailInputFieldHiddenText: "E-Mail Adresse",
		},
	}

	// Stellt die Seiten bereit
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmplName string

		// Bestimmen, welche Seite gerendert werden soll
		switch r.URL.Path {
		case "/logon", "/logon.html":
			tmplName = "logon.html"
			data.PageTitel = "Anmeldung"
		case "/register", "/register.html":
			tmplName = "logon.html"
			data.PageTitel = "Registrieren"
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
