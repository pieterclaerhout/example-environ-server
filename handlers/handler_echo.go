package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pieterclaerhout/go-log"
	"github.com/twixlmedia/example-environ-server/response"
)

const headerContentType = "Content-Type"

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError).Write(w, r)
		return
	}

	contentType := r.Header.Get(headerContentType)

	if contentType == "application/json" {
		body = formatJSON(body)
	}

	w.Header().Set(headerContentType, contentType)
	w.Write(body)

	log.Info(string(body))

}

func formatJSON(in []byte) []byte {
	var out bytes.Buffer
	if err := json.Indent(&out, in, "", "    "); err == nil {
		return out.Bytes()
	}
	return in

}
