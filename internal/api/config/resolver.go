package config

import (
	"api/internal/api"
	"api/pkg/log"
)

type ApiResolver struct {
	config *ApiConfig

	// Singleton
	logger *log.ApiLogger
}

func NewResolver(config *ApiConfig) *ApiResolver {
	return &ApiResolver{config: config}
}

func (ar *ApiResolver) ResolveApiServer() *ApiServer {
	return NewApiServer(
		ar.ResolveLogger(),
		ar,
		ar.config.Server,
	)
}

func (ar *ApiResolver) ResolveLogger() *log.ApiLogger {
	if ar.logger == nil {
		ar.logger = log.NewLogger(ar.config.Log)
	}

	return ar.logger
}

func (ar *ApiResolver) ResolveHealthHandler(prefix string) api.HealthHandler {
	return api.NewHealthHandler(prefix)
}
