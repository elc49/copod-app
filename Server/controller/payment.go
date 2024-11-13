package controller

import (
	"context"

	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var paymentController *Payment

type PaymentController interface {
	CreatePayment(context.Context, sql.CreatePaymentParams) (*bool, error)
}

type Payment struct {
	r *repository.Payment
}

func GetPaymentController() PaymentController {
	return paymentController
}

func (c *Payment) Init(sql *sql.Queries) {
	r := &repository.Payment{}
	r.Init(sql)
	c.r = r
	paymentController = c
}

func (c *Payment) CreatePayment(ctx context.Context, args sql.CreatePaymentParams) (*bool, error) {
	return c.r.CreatePayment(ctx, args)
}
