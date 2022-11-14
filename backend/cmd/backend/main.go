package main

import (
	"context"
	"github.com/DmitriiTrifonov/octo-giggle/backend/internal/model"
	"log"
	"net/http"

	"github.com/DmitriiTrifonov/octo-giggle/backend/internal/handlers"
)

func main() {
	dataChan := make(chan model.Request, 0)
	ctx := context.Background()
	mux := http.NewServeMux()
	mux.Handle("/handle-data", &handlers.DataHandler{Out: dataChan})
	mux.Handle("/", handlers.NewWS(ctx, dataChan))
	// todo: make it with context
	if err := http.ListenAndServe(":8080", mux); err != nil { //nolint:gosec // no need for timeout now
		log.Fatal(err)
	}
}
