package main

import (
	"net/http"
)

func main() {
	initPages()
	http.ListenAndServe("0.0.0.0:8080", nil)
}
