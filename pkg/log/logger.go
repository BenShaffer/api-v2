package log

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Level       string
	Environment string
	ServiceName string
	AppVersion  string
	HostName    string
	LogCaller   bool
}

type Logger interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
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
	*logrus.Entry
}

func NewLogger(config Config) *ApiLogger {
	base := logrus.New()
	base.SetReportCaller(config.LogCaller)

	logger := &ApiLogger{
		Entry: base.WithFields(logrus.Fields{
			"env":      config.Environment,
			"service":  config.ServiceName,
			"version":  config.AppVersion,
			"hostname": config.HostName,
		}),
	}

	logger.setLogLevel(config.Level)
	logger.setFormatter(config.Environment)
	return logger
}

func (al *ApiLogger) setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "trace":
		al.Logger.SetLevel(logrus.TraceLevel)
	case "debug":
		al.Logger.SetLevel(logrus.DebugLevel)
	case "info":
		al.Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		al.Logger.SetLevel(logrus.WarnLevel)
	case "error":
		al.Logger.SetLevel(logrus.ErrorLevel)
	default:
		al.Logger.SetLevel(logrus.InfoLevel)
	}
}

func (al *ApiLogger) setFormatter(env string) {
	if env == "Development" {
		al.Entry.Logger.Formatter = &logrus.TextFormatter{
			ForceColors: true,
		}
	} else {
		al.Entry.Logger.Formatter = &logrus.JSONFormatter{}
	}
}
