package config

import (
	"api/pkg/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	SetRoutes(router *gin.Engine)
}

type Config struct {
	Port             string
	ReadWriteTimeout time.Duration
	IdleTimeout      time.Duration
}

type ApiServer struct {
	logger log.Logger
	*http.Server
}

func NewApiServer(logger log.Logger, resolver *ApiResolver, config Config) *ApiServer {
	server := &ApiServer{
		logger: logger,
		Server: &http.Server{
			Addr:         ":" + config.Port,
			WriteTimeout: config.ReadWriteTimeout,
			ReadTimeout:  config.ReadWriteTimeout,
			IdleTimeout:  config.IdleTimeout,
		}}

	server.setHandler(resolver)
	return server
}

func (as *ApiServer) Run() (err error) {
	as.logger.Trace("Starting Server...")
	time.AfterFunc(3*time.Second, func() {
		as.logger.Debugf("Server running on port %s", as.Server.Addr)
	})

	return as.Server.ListenAndServe()
}

func (as *ApiServer) setHandler(resolver *ApiResolver) {
	router := gin.New()
	as.configureMiddleware(router)

	handlers := []Handler{
		resolver.ResolveHealthHandler("/health"),
	}

	for _, handler := range handlers {
		handler.SetRoutes(router)
	}

	as.Server.Handler = router
}

func (as *ApiServer) configureMiddleware(router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)

	router.Use(as.requestLogger())
	router.Use(gin.Recovery())
}

func (as *ApiServer) requestLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		as.logger.Tracef("%s %s %d - %s",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
		return ""
	})
}
