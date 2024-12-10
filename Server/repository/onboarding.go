package repository

import (
	"context"
	db "database/sql"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type Onboarding struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (o *Onboarding) Init(sql *sql.Queries) {
	o.sql = sql
	o.log = logger.GetLogger()
}

func (o *Onboarding) CreateOnboarding(ctx context.Context, args sql.CreateOnboardingParams) (*model.Onboarding, error) {
	r, err := o.sql.CreateOnboarding(ctx, args)
	if err != nil {
		o.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: CreateOnboarding")
		return nil, err
	}

	return &model.Onboarding{
		ID:               r.ID,
		TitleID:          r.TitleID,
		SupportDocID:     r.SupportDocID,
		DisplayPictureID: r.DisplayPictureID,
		Verification:     model.Verification(r.Verification),
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
	}, nil
}

func (o *Onboarding) GetOnboardingByEmailAndVerification(ctx context.Context, args sql.GetOnboardingByEmailAndVerificationParams) (*model.Onboarding, error) {
	r, err := o.sql.GetOnboardingByEmailAndVerification(ctx, args)
	switch {
	case err != nil && err == db.ErrNoRows:
		return nil, nil
	case err != nil:
		o.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: GetOnboardingByEmailAndVerification")
		return nil, err
	default:
		return &model.Onboarding{
			ID:               r.ID,
			TitleID:          r.TitleID,
			SupportDocID:     r.SupportDocID,
			DisplayPictureID: r.DisplayPictureID,
			Verification:     model.Verification(r.Verification),
			CreatedAt:        r.CreatedAt,
			UpdatedAt:        r.UpdatedAt,
		}, nil
	}
}

func (o *Onboarding) UpdateOnboardingVerificationByID(ctx context.Context, args sql.UpdateOnboardingVerificationByIDParams) (*model.Onboarding, error) {
	u, err := o.sql.UpdateOnboardingVerificationByID(ctx, args)
	if err != nil {
		o.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateOnboardingVerificationByID")
		return nil, err
	}

	return &model.Onboarding{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (o *Onboarding) GetOnboardingByVerificationAndPaymentStatus(ctx context.Context, verification model.Verification, status model.PaymentStatus) ([]*model.Onboarding, error) {
	var onboardings []*model.Onboarding
	return onboardings, nil
}
