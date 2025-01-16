package middleware

import (
	"net/http"

	"github.com/elc49/copod/config"
	sentryHttp "github.com/getsentry/sentry-go/http"
)

func Sentry(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.IsProd() {
			sentryHttp.New(sentryHttp.Options{}).Handle(next).ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
