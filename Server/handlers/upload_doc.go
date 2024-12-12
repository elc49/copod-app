package handlers

import (
	"net/http"

	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/tigris"
	"github.com/elc49/copod/util"
	"github.com/sirupsen/logrus"
)

const (
	MAX_MEMORY = int64(6000000)
)

func UploadDoc() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		tigris := tigris.T
		err := r.ParseMultipartForm(MAX_MEMORY)
		if err != nil {
			log.WithError(err).Errorf("handlers: ParseMultipartForm")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.WithError(err).WithFields(logrus.Fields{"file": file}).Errorf("handlers: r.FormFile")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer file.Close()
		url, err := tigris.Upload(r.Context(), file, fileHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := util.WriteHttp(w, struct {
			ImageUri string `json:"imageUri"`
		}{
			ImageUri: *url,
		}, http.StatusCreated); err != nil {
			log.WithError(err).Errorf("handlers: util.WriteHttp")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
