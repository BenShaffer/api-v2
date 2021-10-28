package config

import (
	"api/internal/infrastructure"
	"api/pkg/log"
	"os"
	"time"
)

type ApiConfig struct {
	Log    log.Config
	Server Config
	DB     infrastructure.Config
}

func LoadConfiguration() *ApiConfig {
	return &ApiConfig{
		Log: log.Config{
			Environment: os.Getenv("ENVIRONMENT"),
			ServiceName: os.Getenv("SERVICE_NAME"),
			AppVersion:  os.Getenv("APP_VERSION"),
			HostName:    getHostName(),
		},
		Server: Config{
			Port:             os.Getenv("PORT"),
			ReadWriteTimeout: time.Second * 15,
			IdleTimeout:      time.Second * 60,
		},
		DB: infrastructure.Config{
			ConnectionString: os.Getenv("SQL_CONN"),
		},
	}
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "N/A"
	} else {
		return hostname
	}
}
