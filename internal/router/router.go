package router

import "net/http"

type Router struct {
	mux    *http.ServeMux
	prefix string
}

type RouteHandler = func() string

func New(mux *http.ServeMux) *Router {
	return &Router{mux: mux}
}

func (r *Router) Prefix(prefix string) {
	r.prefix = prefix
}

func (r *Router) Get(path string, handler RouteHandler) {
	r.mux.HandleFunc("GET "+r.prefix+path, func(writer http.ResponseWriter, request *http.Request) {
		res := handler()

		writer.Write([]byte(res))
	})
}
