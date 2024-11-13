package repository

import (
	"context"

	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type Payment struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *Payment) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *Payment) CreatePayment(ctx context.Context, args sql.CreatePaymentParams) (*bool, error) {
	b := false
	_, err := r.sql.CreatePayment(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: CreatePayment")
		return nil, err
	}
	b = true

	return &b, nil
}
