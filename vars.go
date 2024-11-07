package main

import (
	"embed"
	"html/template"
)

//go:embed web/*
var webFiles embed.FS

var templates *template.Template

var german = &LangSpeficData{
	WebsiteLang:                       "de",
	EMailUsernameInputFieldHiddenText: "E-Mail Adresse / Benutzername",
	HelpTextUnderLogonButton:          template.HTML(`Indem Sie auf „Anmelden“ klicken, stimmen Sie den <a href="test">Nutzungsbedingungen</a> zu.`),
	PasswordInputFieldHiddenText:      "Passwort",
	LogonButtonText:                   "Anmelden",
}
