package logger

import (
	"time"

	"github.com/elc49/copod/config"
	"github.com/getsentry/sentry-go"
	logrusSentry "github.com/getsentry/sentry-go/logrus"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func GetLogger() *logrus.Logger {
	// Attach sentry hook - only in prod/staging
	if config.IsProd() {
		// Error levels to report
		eLevels := []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		}
		hook, err := logrusSentry.New(eLevels, sentry.ClientOptions{
			Dsn:              config.C.Sentry.Dsn,
			AttachStacktrace: true,
		})
		if err != nil {
			log.WithError(err).Fatalln("logger: logrusSentry.New")
		}

		// Register sentry hook
		log.AddHook(hook)

		defer hook.Flush(5 * time.Second)
		logrus.RegisterExitHandler(func() { hook.Flush(5 * time.Second) })
	}

	return log
}
