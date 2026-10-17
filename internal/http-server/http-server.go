package http_server

import "net/http"

type HttpServer struct {
}

func Create() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Listen(httpInterface string) error {
	err := http.ListenAndServe(
		httpInterface, nil,
	)

	return err
}
