package log

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Environment string
	ServiceName string
	AppVersion  string
	HostName    string
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

type ApiLogger struct {
	*zap.SugaredLogger
}

func NewLogger(config Config) *ApiLogger {
	var logConfig zap.Config
	if strings.EqualFold(config.Environment, "Development") {
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05")

	} else {
		logConfig = zap.NewProductionConfig()
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logConfig.InitialFields = map[string]interface{}{
			"env":      config.Environment,
			"service":  config.ServiceName,
			"version":  config.AppVersion,
			"hostname": config.HostName,
		}
	}

	base, _ := logConfig.Build()
	logger := &ApiLogger{
		SugaredLogger: base.Sugar(),
	}

	return logger
}
