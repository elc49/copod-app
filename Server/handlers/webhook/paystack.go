package webhook

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/paystack"
	"github.com/sirupsen/logrus"
)

func Paystack() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		res := &paystack.PaystackWebhook{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorf("webhook: paystack io.ReadAll")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &res); err != nil {
			log.WithError(err).Errorf("webhook: paystack: json.Unmarshal")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.WithFields(logrus.Fields{"data": res}).Infoln("Paystack webhook")
		w.WriteHeader(http.StatusOK)
	})
}
