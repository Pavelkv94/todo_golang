package core_http_response

import "net/http"

var (
	StatusCodeUninitialized = -1
)

type HTTPResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewHTTPResponseWriter(rw http.ResponseWriter) *HTTPResponseWriter {
	return &HTTPResponseWriter{ResponseWriter: rw}
}

func (w *HTTPResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (rw *HTTPResponseWriter) GetStatusCode() int {
	if rw.statusCode == StatusCodeUninitialized {
		panic("status code is not initialized")
	}
	return rw.statusCode
}
