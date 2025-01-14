package controller

import (
	"context"

	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
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
	return c.r.CreateEarlySignup(ctx, email)
}
