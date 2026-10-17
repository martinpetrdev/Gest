package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/martinpetrdev/Gest/internal/logger"
	"github.com/martinpetrdev/Gest/internal/url"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

type Router struct {
	mux    *http.ServeMux
	prefix string

	logger *logger.Logger
}

type RouteHandler = func(ctx *RouteContext) string

func New(mux *http.ServeMux) *Router {
	return &Router{mux: mux, logger: logger.New("Router"), prefix: "/"}
}

func (r *Router) getMuxRoute(method HttpMethod, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}

func (r *Router) registerRoute(method HttpMethod, path string, handler RouteHandler) {
	fullUrl := url.Join(r.prefix, path)
	r.logger.Log(fmt.Sprintf("Mapped {%s, %s} route", "GET", fullUrl))

	r.mux.HandleFunc(r.getMuxRoute(method, fullUrl), func(writer http.ResponseWriter, request *http.Request) {
		res := handler(NewRouteContext(request))

		_, err := writer.Write([]byte(res))
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (r *Router) Prefix(prefix string) {
	r.prefix = prefix
}

func (r *Router) Get(path string, handler RouteHandler) {
	r.registerRoute(GET, path, handler)
}

func (r *Router) Post(path string, handler RouteHandler) {
	r.registerRoute(POST, path, handler)
}

func (r *Router) Put(path string, handler RouteHandler) {
	r.registerRoute(PUT, path, handler)
}

func (r *Router) Patch(path string, handler RouteHandler) {
	r.registerRoute(PATCH, path, handler)
}

func (r *Router) Delete(path string, handler RouteHandler) {
	r.registerRoute(DELETE, path, handler)
}
