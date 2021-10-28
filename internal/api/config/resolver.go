package config

import (
	"api/internal/api"
	"api/internal/application"
	"api/internal/infrastructure"
	"api/pkg/log"
)

type ApiResolver struct {
	config *ApiConfig

	// Singleton
	logger *log.ApiLogger
	db     *infrastructure.SQLDatabase
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

func (ar *ApiResolver) ResolveSQLDatabase() *infrastructure.SQLDatabase {
	if ar.db == nil {
		ar.db = infrastructure.NewSQLDatabase(
			ar.ResolveLogger(),
			ar.config.DB,
		)
	}

	return ar.db
}

func (ar *ApiResolver) ResolveHealthHandler(prefix string) *api.HealthHandler {
	return api.NewHealthHandler(prefix)
}

func (ar *ApiResolver) ResolvePersonRepo() *infrastructure.PersonRepo {
	return infrastructure.NewPersonRepo(
		ar.ResolveLogger(),
		ar.ResolveSQLDatabase(),
	)
}

func (ar *ApiResolver) ResolvePersonService() *application.PersonService {
	return application.NewPersonService(
		ar.ResolveLogger(),
		ar.ResolvePersonRepo(),
	)
}

func (ar *ApiResolver) ResolvePersonHandler(prefix string) *api.PersonHandler {
	return api.NewPersonHandler(
		prefix,
		ar.ResolveLogger(),
		ar.ResolvePersonService(),
	)
}
