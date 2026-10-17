package http_server

import "net/http"

type HttpServer struct {
	handler http.Handler
}

func Create(handler http.Handler) *HttpServer {
	return &HttpServer{
		handler: handler,
	}
}

func (s *HttpServer) Listen(httpInterface string) error {
	err := http.ListenAndServe(
		httpInterface, s.handler,
	)

	return err
}
