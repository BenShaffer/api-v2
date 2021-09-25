package config

import (
	"api/pkg/log"
	"os"
	"time"
)

type ApiConfig struct {
	Log    log.Config
	Server Config
}

func LoadConfiguration() *ApiConfig {
	return &ApiConfig{
		Log: log.Config{
			Level:       os.Getenv("LOG_LEVEL"),
			Environment: os.Getenv("ENVIRONMENT"),
			ServiceName: os.Getenv("SERVICE_NAME"),
			AppVersion:  os.Getenv("APP_VERSION"),
			HostName:    getHostName(),
			LogCaller:   false,
		},
		Server: Config{
			Port:             os.Getenv("PORT"),
			ReadWriteTimeout: time.Second * 15,
			IdleTimeout:      time.Second * 60,
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
