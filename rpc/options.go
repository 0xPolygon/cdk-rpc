package rpc

import "net/http"

func WithHealthHandler(h http.Handler) Option {
	return func(s *Server) {
		s.healthHandler = h
	}
}
