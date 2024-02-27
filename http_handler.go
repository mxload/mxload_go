package mxload

import (
	"bytes"
	"io"
	"net/http"
)

func MiddlewareFunc(config *MxloadConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Responsibility for error handling is deferred to the parent application.
			bodyBytes, _ := io.ReadAll(r.Body)
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			writer := NewResponseWriterWithStatus(w)
			next.ServeHTTP(writer, r)
			c := NewCollector(config)
			c.Collect(writer, r, bodyBytes)
		})
	}
}
