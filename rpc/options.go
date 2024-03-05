package rpc

import "net/http"

func WithHealthcheckHandler(h http.Handler) Option {
	return func(s *Server) {
		s.healthHandler = h
	}
}
