package repository

import (
	"context"

	"github.com/elc49/copod/graph/model"
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

func (r *Payment) CreatePayment(ctx context.Context, args sql.CreatePaymentParams) (*model.Payment, error) {
	p, err := r.sql.CreatePayment(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: CreatePayment")
		return nil, err
	}

	return &model.Payment{
		ID:          p.ID,
		ReferenceID: p.ReferenceID,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}, nil
}

func (r *Payment) GetPaymentByReferenceID(ctx context.Context, referenceID string) (*model.Payment, error) {
	p, err := r.sql.GetPaymentByReferenceID(ctx, referenceID)
	if err != nil {
		r.log.WithError(err).Errorf("repository: GetPaymentByReferenceID")
		return nil, err
	}

	return &model.Payment{
		ID:          p.ID,
		Status:      p.Status,
		ReferenceID: p.ReferenceID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}, nil
}

func (r *Payment) UpdatePaymentStatus(ctx context.Context, args sql.UpdatePaymentStatusParams) (*model.Payment, error) {
	u, err := r.sql.UpdatePaymentStatus(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdatePaymentStatus")
		return nil, err
	}

	return &model.Payment{
		ID:          u.ID,
		Status:      u.Status,
		ReferenceID: u.ReferenceID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}, nil
}
