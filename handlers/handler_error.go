package handlers

import (
	"net/http"

	"github.com/twixlmedia/example-environ-server/response"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	response.Error("the error message", http.StatusInternalServerError).Write(w, r)
}
