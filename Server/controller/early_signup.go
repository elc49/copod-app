package controller

import (
	"context"

	"github.com/elc49/copod/config"
	emailSvc "github.com/elc49/copod/email"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/resend/resend-go/v2"
	"github.com/sirupsen/logrus"
)

var earlySignupController *EarlySignup

type EarlySignupController interface {
	CreateEarlySignup(context.Context, string) (*string, error)
}

type EarlySignup struct {
	r   *repository.EarlySignup
	sql *sql.Queries
	log *logrus.Logger
}

func (c *EarlySignup) Init(sql *sql.Queries) {
	r := &repository.EarlySignup{}
	r.Init(sql)
	c.r = r
	c.sql = sql
	c.log = logger.GetLogger()
	earlySignupController = c
}

func GetEarlySignupController() EarlySignupController {
	return earlySignupController
}

func (c *EarlySignup) CreateEarlySignup(ctx context.Context, email string) (*string, error) {
	e, err := c.r.CreateEarlySignup(ctx, email)
	if err != nil {
		return nil, err
	}

	// Onboard early signup - this runs only in prod
	if config.IsProd() {
		go func() {
			if err := c.emailOnboard(context.Background(), email); err != nil {
				return
			}
		}()
	}

	return e, nil
}

func (c *EarlySignup) emailOnboard(ctx context.Context, email string) error {
	e, err := c.sql.GetEarlySignupByEmail(ctx, email)
	if err != nil {
		c.log.WithError(err).Errorf("controller: GetEarlySignupByEmail: emailOnboard")
		return err
	}
	if !e.Onboarded {
		emailRequest := &resend.SendEmailRequest{
			From:    "Chanzu <chanzu@info.copodap.com>",
			To:      []string{email},
			Html:    "<p>Hello,</p><p>I am building something great and I am happy to have you onboard to try it out.</p><p>Don't hesitate to reach out if something is not working/you have an idea of how I can make your experience on Copod better.</p><br /><strong>Best,</strong><p>Edwin Chanzu.</p><a href='https://x.com/gugachanzu' target='_blank'>X</a>",
			Subject: "Welcome onboard!",
		}

		if err := emailSvc.GetResendEmailService().Send(ctx, emailRequest); err != nil {
			return err
		}

		// Onboard early signup async
		args := sql.OnboardEarlySignupParams{
			Email:     email,
			Onboarded: true,
		}
		if _, err := c.sql.OnboardEarlySignup(ctx, args); err != nil {
			c.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("email: sql.OnboardEarlySignup: emailOnboard")
			return err
		}
	}

	return nil
}
