package buuurst_dev

import (
	"bytes"
	"io"
	"net/http"
)

func MiddlewareFunc(config *BuuurstDevConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bodyBytes, err := io.ReadAll(r.Body)
			r.Body.Close()
			if err != nil {
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			writer := NewResponseWriterWithStatus(w)
			next.ServeHTTP(writer, r)
			c := NewCollector(config)
			c.Collect(writer, r, bodyBytes)
		})
	}
}