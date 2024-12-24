package repository

import (
	"context"
	db "database/sql"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Onboarding struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *Onboarding) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *Onboarding) CreateOnboarding(ctx context.Context, args sql.CreateOnboardingParams) (*model.Onboarding, error) {
	o, err := r.sql.CreateOnboarding(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: CreateOnboarding")
		return nil, err
	}

	return &model.Onboarding{
		ID:               o.ID,
		TitleID:          o.TitleID,
		SupportDocID:     o.SupportDocID,
		DisplayPictureID: o.DisplayPictureID,
		CreatedAt:        o.CreatedAt,
		UpdatedAt:        o.UpdatedAt,
	}, nil
}

func (r *Onboarding) GetOnboardingByEmailAndVerification(ctx context.Context, args sql.GetOnboardingByEmailAndVerificationParams) (*model.Onboarding, error) {
	o, err := r.sql.GetOnboardingByEmailAndVerification(ctx, args)
	switch {
	case err != nil && err == db.ErrNoRows:
		return nil, nil
	case err != nil:
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: GetOnboardingByEmailAndVerification")
		return nil, err
	default:
		return &model.Onboarding{
			ID:               o.ID,
			Email:            o.Email,
			TitleID:          o.TitleID,
			SupportDocID:     o.SupportDocID,
			DisplayPictureID: o.DisplayPictureID,
			CreatedAt:        o.CreatedAt,
			UpdatedAt:        o.UpdatedAt,
		}, nil
	}
}

func (r *Onboarding) GetOnboardingByID(ctx context.Context, id uuid.UUID) (*model.Onboarding, error) {
	o, err := r.sql.GetOnboardingByID(ctx, id)
	if err != nil && err == db.ErrNoRows {
		return nil, nil
	} else if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"id": id}).Errorf("repository: GetOnboardingByID")
		return nil, err
	}

	return &model.Onboarding{
		ID:               o.ID,
		TitleID:          o.TitleID,
		DisplayPictureID: o.DisplayPictureID,
		SupportDocID:     o.SupportDocID,
		CreatedAt:        o.CreatedAt,
		UpdatedAt:        o.UpdatedAt,
	}, nil
}
