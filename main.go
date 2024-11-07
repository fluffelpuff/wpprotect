package main

import (
	"log"
	"net/http"
)

func main() {
	initPages()
	log.Println("Server startet auf Port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
	}
}
