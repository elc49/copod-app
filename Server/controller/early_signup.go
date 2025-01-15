package controller

import (
	"context"

	emailSvc "github.com/elc49/copod/email"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/resend/resend-go/v2"
)

var earlySignupController *EarlySignup

type EarlySignupController interface {
	CreateEarlySignup(context.Context, string) (*string, error)
}

type EarlySignup struct {
	r *repository.EarlySignup
}

func (c *EarlySignup) Init(sql *sql.Queries) {
	r := &repository.EarlySignup{}
	r.Init(sql)
	c.r = r
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

	// TODO enable this only in staging/prod
	rs := emailSvc.GetResendEmailService()
	emailRequest := &resend.SendEmailRequest{
		From:    "Chanzu <chanzu@copodap.com>",
		To:      []string{email},
		Html:    "<p>Hello,</p><p>I am building something great and I am happy to have you onboard to try it out.</p><p>Don't hesitate to reach out if something is not working/you have an idea of how I can make your experience on Copod better.</p><br /><strong>Best,</strong><p>Edwin Chanzu</p><a href='https://x.com/gugachanzu' target='_blank'>X</a>",
		Subject: "Welcome onboard!",
	}

	if err := rs.Send(ctx, emailRequest); err != nil {
		return nil, err
	}

	return e, nil
}
