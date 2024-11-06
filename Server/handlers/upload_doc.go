package handlers

import (
	"net/http"

	"github.com/elc49/copod/tigris"
)

const (
	MAX_MEMORY = int64(6000000)
)

func UploadDoc() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tigris := tigris.T
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
		url, err := tigris.Upload(r.Context(), file, fileHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := writeJSON(w, struct {
			ImageUri string `json:"image_uri"`
		}{
			ImageUri: *url,
		}, http.StatusCreated); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
