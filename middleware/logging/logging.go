package logging

import (
	"fmt"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/pieterclaerhout/go-log"
)

// Logger returns the logging middleware
func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		ri := &requestInfo{
			ts:        time.Now(),
			proto:     r.Proto,
			method:    r.Method,
			uri:       r.URL.String(),
			referer:   r.Header.Get("Referer"),
			userAgent: r.Header.Get("User-Agent"),
			ipaddr:    requestGetRemoteAddress(r),
		}

		m := httpsnoop.CaptureMetrics(next, w, r)

		ri.code = m.Code
		ri.size = m.Written
		ri.duration = m.Duration

		log.Info(
			fmt.Sprintf("%16v", ri.duration),
			"|",
			ri.ipaddr,
			"\""+ri.method,
			ri.uri,
			ri.proto+"\"",
			ri.code,
			ri.size,
			"\""+ri.userAgent+"\"",
		)

	}

	return http.HandlerFunc(fn)

}
