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
		Verification:     model.Verification(o.Verification),
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
			TitleID:          o.TitleID,
			SupportDocID:     o.SupportDocID,
			DisplayPictureID: o.DisplayPictureID,
			Verification:     model.Verification(o.Verification),
			CreatedAt:        o.CreatedAt,
			UpdatedAt:        o.UpdatedAt,
		}, nil
	}
}

func (r *Onboarding) UpdateOnboardingVerificationByID(ctx context.Context, args sql.UpdateOnboardingVerificationByIDParams) (*model.Onboarding, error) {
	u, err := r.sql.UpdateOnboardingVerificationByID(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateOnboardingVerificationByID")
		return nil, err
	}

	return &model.Onboarding{
		ID:           u.ID,
		Verification: model.Verification(u.Verification),
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}, nil
}

func (r *Onboarding) GetOnboardingByVerificationAndPaymentStatus(ctx context.Context, args sql.GetOnboardingByVerificationAndPaymentStatusParams) ([]*model.Onboarding, error) {
	var onboardings []*model.Onboarding
	obs, err := r.sql.GetOnboardingByVerificationAndPaymentStatus(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: GetOnboardingByVerificationAndPaymentStatus")
		return nil, err
	}

	for _, item := range obs {
		ob := &model.Onboarding{
			ID:               item.ID,
			TitleID:          item.TitleID,
			DisplayPictureID: item.DisplayPictureID,
			SupportDocID:     item.SupportDocID,
			Verification:     model.Verification(item.Verification),
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
		}

		onboardings = append(onboardings, ob)
	}

	return onboardings, nil
}

func (r *Onboarding) GetOnboardingByEmail(ctx context.Context, email string) (*model.Onboarding, error) {
	ob, err := r.sql.GetOnboardingByEmail(ctx, email)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: GetOnboardingByEmail")
		return nil, err
	}

	return &model.Onboarding{
		ID:               ob.ID,
		TitleID:          ob.TitleID,
		DisplayPictureID: ob.DisplayPictureID,
		SupportDocID:     ob.SupportDocID,
		Verification:     model.Verification(ob.Verification),
		CreatedAt:        ob.CreatedAt,
		UpdatedAt:        ob.UpdatedAt,
	}, nil
}
