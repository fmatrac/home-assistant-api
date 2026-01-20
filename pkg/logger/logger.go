package logger

import (
	"github.com/fmatrac/home-assistant-api/config"
	"github.com/sirupsen/logrus"
)

func Setup(cfg *config.Config) *logrus.Logger {
	var level logrus.Level

	switch cfg.Logging.Level {
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warn":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	default:
		level = logrus.InfoLevel
	}

	logger := logrus.New()
	logger.SetLevel(level)
	return logger
}
