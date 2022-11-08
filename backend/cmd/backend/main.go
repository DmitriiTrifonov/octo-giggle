package main

import (
	"log"
	"net/http"

	"github.com/DmitriiTrifonov/octo-giggle/backend/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/handle-data", &handlers.DataHandler{})
	// todo: make it with context
	if err := http.ListenAndServe(":8080", mux); err != nil { //nolint:gosec // no need for timeout now
		log.Fatal(err)
	}
}
