package core_http_server

import "net/http"

type APIVersionRouter struct {
	mux *http.ServeMux
}

func NewAPIVersionRouter() *APIVersionRouter {
	return &APIVersionRouter{mux: http.NewServeMux()}
}

func (r *APIVersionRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.mux.ServeHTTP(w, r)
}
