package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/elc49/copod/logger"
)

func Favicon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		templatesDir, err := filepath.Abs("static")
		if err != nil {
			log.WithError(err).Errorf("handlers: filepath.Abs")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.ServeFile(w, r, filepath.Join(templatesDir, "favicon.ico"))
	})
}
