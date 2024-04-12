package rpc

import (
	"net/http"

	"github.com/0xPolygon/cdk-rpc/log"
	"go.uber.org/zap"
)

func WithHealthHandler(h http.Handler) Option {
	return func(s *Server) {
		s.healthHandler = h
	}
}

func WithLogger(zapLogger *zap.SugaredLogger) Option {
	return func(s *Server) {
		log.SetLogger(zapLogger)
	}
}
