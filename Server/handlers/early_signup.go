package handlers

import (
	"io"
	"net/http"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/util"
	"github.com/sirupsen/logrus"
)

func EarlySignup() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		ec := controller.GetEarlySignupController()
		var email string

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorf("handlers: EarlySignup: io.ReadAll")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err := util.DecodeJson(body, &email); err != nil {
			log.WithError(err).Errorf("handlers: EarlySignup: util.DecodeJson")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := ec.CreateEarlySignup(r.Context(), email); err != nil {
			log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("handlers: CreateEarlySignup")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		util.WriteHttp(w, nil, http.StatusOK)
	})
}
