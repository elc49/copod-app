package handlers

import (
	"net/http"

	"github.com/elc49/copod/ip"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/util"
	"github.com/sirupsen/logrus"
)

func Ipinfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		info, err := ip.GetIpService().GetIpinfo(r.Context(), r.RemoteAddr)
		if err != nil {
			log.WithError(err).WithFields(logrus.Fields{"ip": r.RemoteAddr}).Errorf("handlers: Ipinfo")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := util.WriteHttp(w, info, http.StatusOK); err != nil {
			log.WithError(err).WithFields(logrus.Fields{"info": info}).Errorf("handlers: write to http")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
