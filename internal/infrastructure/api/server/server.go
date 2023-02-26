package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct{
	srv http.Server
}

func NewServer(addr string, h http.Handler) *Server{
	return &Server{
		srv: http.Server{
			Addr: addr,
			Handler: h,
			ReadTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			ReadHeaderTimeout: 30 * time.Second,
		},
	}
}

func (s *Server) Start(ctx context.Context) error{
	go func() {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := s.srv.Shutdown(stopCtx); err != nil {

			fmt.Errorf("Server Shutdown Failed: %w", err)
		}
	}()

	return s.srv.ListenAndServe()
}