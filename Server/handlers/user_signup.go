package handlers

import (
	"io"
	"net/http"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/util"
	"github.com/sirupsen/logrus"
)

func UserSignup() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		var req struct {
			Email     string `json:"email"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorf("handler: io.ReadAll: UserSignup")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := util.DecodeJson(body, &req); err != nil {
			log.WithError(err).Errorf("handlers: util.DecodeJson: UserSignup")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		args := sql.CreateUserParams{
			Email:     req.Email,
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
		}
		user, err := controller.GetUserController().CreateUser(r.Context(), args)
		if err != nil {
			log.WithError(err).WithFields(logrus.Fields{"user": req}).Errorf("handlers: controller.GetUserController.CreateUser: UserSignup")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		util.WriteHttp(w, user, http.StatusOK)
	})
}
