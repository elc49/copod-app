package webhook

import (
	"io"
	"net/http"

	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/paystack"
	"github.com/elc49/copod/util"
	"github.com/sirupsen/logrus"
)

func Paystack() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		res := &paystack.PaystackWebhook{}
		p := paystack.GetPaystackService()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorf("webhook: paystack io.ReadAll")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := util.DecodeJson(body, &res); err != nil {
			log.WithError(err).Errorf("webhook: paystack: json.Unmarshal")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := p.HappyPaystack(r.Context(), res); err != nil {
			log.WithError(err).WithFields(logrus.Fields{"update": res}).Errorf("webhook: HappyPaystack")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.WithFields(logrus.Fields{"data": res}).Infoln("Paystack webhook")
		w.WriteHeader(http.StatusOK)
	})
}
