package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var paymentController *Payment

type PaymentController interface {
	CreatePayment(context.Context, sql.CreatePaymentParams) (*model.Payment, error)
	GetPaymentByReferenceID(context.Context, string) (*model.Payment, error)
	UpdatePaymentStatus(context.Context, sql.UpdatePaymentStatusParams) (*model.Payment, error)
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

func (c *Payment) CreatePayment(ctx context.Context, args sql.CreatePaymentParams) (*model.Payment, error) {
	return c.r.CreatePayment(ctx, args)
}

func (c *Payment) GetPaymentByReferenceID(ctx context.Context, referenceID string) (*model.Payment, error) {
	return c.r.GetPaymentByReferenceID(ctx, referenceID)
}

func (c *Payment) UpdatePaymentStatus(ctx context.Context, args sql.UpdatePaymentStatusParams) (*model.Payment, error) {
	return c.r.UpdatePaymentStatus(ctx, args)
}
