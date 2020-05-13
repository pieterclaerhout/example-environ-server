package response

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/markusthoemmes/goautoneg"
)

// Response is used for a HTTP response
type Response struct {
	XMLName    xml.Name `xml:"response"`
	Body       interface{}
	StatusCode int
}

// ErrorResponse is used for a HTTP error response
type ErrorResponse struct {
	XMLName xml.Name `json:"-" xml:"error"`
	Error   string   `json:"error,omitempty" xml:"message,omitempty"`
}

// Write write the response in the appropriate format based on the request header
func (resp Response) Write(w http.ResponseWriter, r *http.Request) error {

	accepts := goautoneg.ParseAccept(r.Header.Get("Accept"))
	for _, accept := range accepts {
		switch accept.SubType {
		case "xml":
			return resp.ToXML(w)
		case "html", "xhtml", "xhtml+xml":
			return resp.ToHTML(w)
		case "text":
			return resp.ToText(w)
		}
	}

	return resp.ToJSON(w)

}

// ToText returns the response as plain text
func (resp Response) ToText(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(resp.StatusCode)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	enc.SetIndent("", "    ")
	return enc.Encode(resp.Body)
}

// ToHTML returns the response as HTML
func (resp Response) ToHTML(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte("<pre>"))
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	enc.SetIndent("", "    ")
	if err := enc.Encode(resp.Body); err != nil {
		return err
	}
	w.Write([]byte("</pre>"))
	return nil
}

// ToJSON returns the response as JSON
func (resp Response) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	return json.NewEncoder(w).Encode(resp.Body)
}

// ToXML returns the response as XML
func (resp Response) ToXML(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(resp.StatusCode)
	return xml.NewEncoder(w).Encode(resp.Body)
}

// OK is used to send a HTTP 200 response
func OK(body interface{}) *Response {
	return &Response{
		Body:       body,
		StatusCode: http.StatusOK,
	}
}

// Error is used to send a HTTP response with a custom status code
func Error(message string, statusCode int) *Response {
	return &Response{
		Body: ErrorResponse{
			Error: message,
		},
		StatusCode: http.StatusOK,
	}
}
