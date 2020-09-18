package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pieterclaerhout/go-log"
	"github.com/twixlmedia/example-environ-server/handlers"
	"github.com/twixlmedia/example-environ-server/middleware/logging"
)

func main() {

	log.PrintColors = true
	log.PrintTimestamp = true

	r := mux.NewRouter()

	r.Use(logging.Logger)

	r.HandleFunc("/", handlers.RootHandler).Methods(http.MethodGet)
	r.HandleFunc("/error", handlers.ErrorHandler).Methods(http.MethodGet)
	r.HandleFunc("/echo", handlers.EchoHandler)

	log.Info("Starting HTTP server on :8080")
	err := http.ListenAndServe(":8080", r)
	log.CheckError(err)

}
