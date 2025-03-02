package repository

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
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
		ID:           p.ID,
		OnboardingID: p.OnboardingID.UUID,
		ReferenceID:  p.ReferenceID,
		Status:       p.Status,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}, nil
}

func (r *Payment) GetPaymentByReferenceID(ctx context.Context, referenceID string) (*model.Payment, error) {
	p, err := r.sql.GetPaymentByReferenceID(ctx, referenceID)
	if err != nil {
		r.log.WithError(err).Errorf("repository: GetPaymentByReferenceID")
		return nil, err
	}

	return &model.Payment{
		ID:           p.ID,
		Status:       p.Status,
		ReferenceID:  p.ReferenceID,
		OnboardingID: p.OnboardingID.UUID,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}, nil
}

func (r *Payment) UpdatePaymentStatus(ctx context.Context, args sql.UpdatePaymentStatusParams) (*model.Payment, error) {
	u, err := r.sql.UpdatePaymentStatus(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdatePaymentStatus")
		return nil, err
	}

	return &model.Payment{
		ID:           u.ID,
		Status:       u.Status,
		ReferenceID:  u.ReferenceID,
		OnboardingID: u.OnboardingID.UUID,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}, nil
}

func (r *Payment) GetPaymentOnboardingByID(ctx context.Context, onboardingID uuid.UUID) (*model.Onboarding, error) {
	t, err := r.sql.GetPaymentOnboardingByID(ctx, onboardingID)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"onboarding_id": onboardingID}).Errorf("repository: GetPaymentOnboardingByID")
		return nil, err
	}

	return &model.Onboarding{
		ID:        t.ID,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

func (r *Payment) GetPaymentsByStatus(ctx context.Context, status string) ([]*model.Payment, error) {
	var payments []*model.Payment
	p, err := r.sql.GetPaymentsByStatus(ctx, status)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"status": status}).Errorf("repository: GetPaymentsByStatus")
		return nil, err
	}

	for _, i := range p {
		payment := &model.Payment{
			ID:           i.ID,
			Status:       i.Status,
			Email:        i.Email,
			ReferenceID:  i.ReferenceID,
			OnboardingID: i.OnboardingID.UUID,
			CreatedAt:    i.CreatedAt,
			UpdatedAt:    i.UpdatedAt,
		}

		payments = append(payments, payment)
	}

	return payments, nil
}
