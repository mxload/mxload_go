package buuurst_dev

import (
	"net/http"
)

type ResponseWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriterWithStatus(w http.ResponseWriter) *ResponseWriterWithStatus {
	return &ResponseWriterWithStatus{ResponseWriter: w}
}

func (r *ResponseWriterWithStatus) Status() int {
	return r.statusCode
}

func (r *ResponseWriterWithStatus) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.statusCode = statusCode
}
