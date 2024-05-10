package weather

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.srv = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	fmt.Printf("server started on %s port\n", port)

	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
