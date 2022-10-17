package log

import (
	"github.com/sdslabs/nymeria/config"
	"github.com/sirupsen/logrus"
)

func getLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	if config.NymeriaConfig.Env == "prod" {
		logger.SetLevel(logrus.WarnLevel)
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}

var (
	Logger = getLogger()
)
