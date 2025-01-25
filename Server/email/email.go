package email

import (
	"context"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/resend/resend-go/v2"
	"github.com/sirupsen/logrus"
)

var r Resend

type Resend interface {
	Send(context.Context, *resend.SendEmailRequest) error
}

type rClient struct {
	client *resend.Client
	log    *logrus.Logger
	sql    *sql.Queries
}

func NewResend(sql *sql.Queries) {
	log := logger.GetLogger()
	client := resend.NewClient(config.AppConfig().Resend.ApiKey)

	r = &rClient{client, log, sql}
	log.Infoln("emailservice: Connected")
}

func GetResendEmailService() Resend {
	return r
}

func (r *rClient) Send(ctx context.Context, params *resend.SendEmailRequest) error {
	if _, err := r.client.Emails.Send(params); err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"params": params}).Errorf("email: Send")
		return err
	}

	return nil
}
