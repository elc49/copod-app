package logger

import "github.com/sirupsen/logrus"

var log = logrus.New()

func GetLogger() *logrus.Logger {
	// TODO attach sentry hook - only in prod
	return log
}
