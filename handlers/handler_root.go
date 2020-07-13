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
		XMLName     xml.Name          `json:"-" xml:"environment"`
		Environment []string          `json:"environment" xml:"variable"`
		Headers     map[string]string `json:"headers" xml:"headers"`
		Host        string            `json:"host" xml:"host"`
	}

	resp := rootResponse{
		Environment: []string{},
		Headers:     map[string]string{},
		Host:        r.Host,
	}

	for key, val := range r.Header {
		if len(val) > 0 {
			resp.Headers[key] = val[0]
		}
	}

	for _, item := range os.Environ() {
		if strings.HasPrefix(item, "TWX_") {
			resp.Environment = append(resp.Environment, item)
		}
	}

	response.OK(resp).Write(w, r)

}
