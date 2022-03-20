package server

import (
	"context"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func NewHTTPServer(port string, handler http.Handler) *Server {
	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	return &Server{srv: server}
}
