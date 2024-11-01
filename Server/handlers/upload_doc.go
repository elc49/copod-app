package handlers

import (
	"net/http"

	"github.com/elc49/copod/logger"
)

const (
	MAX_MEMORY = int64(6000000)
)

func UploadDoc() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		err := r.ParseMultipartForm(MAX_MEMORY)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer file.Close()
		log.Infoln(fileHeader)
		// TODO upload image to tigris and return uri
	})
}
