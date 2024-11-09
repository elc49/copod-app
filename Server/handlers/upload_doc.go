package handlers

import (
	"net/http"

	"github.com/elc49/copod/tigris"
	"github.com/elc49/copod/util"
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

		if err := util.WriteHttp(w, struct {
			ImageUri string `json:"imageUri"`
		}{
			ImageUri: *url,
		}, http.StatusCreated); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
