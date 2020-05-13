package handlers

import (
	"encoding/xml"
	"net/http"
	"os"
	"strings"

	"github.com/twixlmedia/example-environ-server/response"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	type rootResponse struct {
		XMLName     xml.Name `json:"-" xml:"environment"`
		Environment []string `json:"environment" xml:"variable"`
	}

	resp := rootResponse{
		Environment: []string{},
	}

	for _, item := range os.Environ() {
		if strings.HasPrefix(item, "TWX_") {
			resp.Environment = append(resp.Environment, item)
		}
	}

	response.OK(resp).Write(w, r)

}
