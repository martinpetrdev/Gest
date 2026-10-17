package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	http_exception "github.com/martinpetrdev/Gest/internal/http-exception"
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
		defer func() {
			rec := recover()
			if rec == nil {
				return
			}

			ex, ok := rec.(http_exception.IHTTPException)
			if !ok {
				panic(rec) // not ours, rethrow
			}

			body, _ := json.Marshal(ex.ResponseBody())

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(ex.ResponseCode())
			_, _ = writer.Write(body)
		}()
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
