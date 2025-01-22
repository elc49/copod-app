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
		// Get title verification
		title, err := r.sql.GetTitleByID(ctx, o.TitleID)
		if err != nil && err == db.ErrNoRows {
			return nil, nil
		} else if err != nil {
			r.log.WithError(err).WithFields(logrus.Fields{"id": o.TitleID}).Errorf("repository: GetOnboardingByEmailAndVerification: GetTitleByID")
			return nil, err
		}
		// Get support doc verification
		supportDoc, err := r.sql.GetSupportDocByID(ctx, o.SupportDocID)
		if err != nil && err == db.ErrNoRows {
			return nil, nil
		} else if err != nil {
			r.log.WithError(err).WithFields(logrus.Fields{"id": o.SupportDocID}).Errorf("repository: GetOnboardingByEmailAndVerification: GetSupportDocByID")
			return nil, err
		}
		// Get display picture verification
		dp, err := r.sql.GetDisplayPictureByID(ctx, o.DisplayPictureID)
		if err != nil && err == db.ErrNoRows {
			return nil, nil
		} else if err != nil {
			r.log.WithError(err).WithFields(logrus.Fields{"id": o.DisplayPictureID}).Errorf("repository: GetOnboardingByEmailAndVerification: GetDisplayPictureByID")
			return nil, err
		}
		return &model.Onboarding{
			ID:               o.ID,
			Email:            o.Email,
			Title:            &model.Title{Verified: model.Verification(title.Verification)},
			TitleID:          o.TitleID,
			SupportingDoc:    &model.SupportingDoc{Verified: model.Verification(supportDoc.Verification)},
			SupportDocID:     o.SupportDocID,
			DisplayPicture:   &model.DisplayPicture{Verified: model.Verification(dp.Verification)},
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

func (r *Onboarding) GetOnboardingsByStatus(ctx context.Context, status model.Verification) ([]*model.Onboarding, error) {
	var onboardings []*model.Onboarding

	obs, err := r.sql.GetOnboardingsByStatus(ctx, status.String())
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"status": status}).Errorf("repository: GetOnboardingsByStatus")
		return nil, err
	}

	for _, item := range obs {
		ob := &model.Onboarding{
			ID:               item.ID,
			TitleID:          item.TitleID,
			SupportDocID:     item.SupportDocID,
			DisplayPictureID: item.DisplayPictureID,
			Verification:     model.Verification(item.Verification),
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
		}

		onboardings = append(onboardings, ob)
	}
	return onboardings, nil
}

func (r *Onboarding) UpdateOnboardingVerificationByID(ctx context.Context, args sql.UpdateOnboardingVerificationByIDParams) (*model.Onboarding, error) {
	ob, err := r.sql.UpdateOnboardingVerificationByID(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateOnboardingVerificationByID")
		return nil, err
	}

	return &model.Onboarding{
		ID:               ob.ID,
		TitleID:          ob.TitleID,
		SupportDocID:     ob.SupportDocID,
		DisplayPictureID: ob.DisplayPictureID,
		Verification:     model.Verification(ob.Verification),
		CreatedAt:        ob.CreatedAt,
		UpdatedAt:        ob.UpdatedAt,
	}, nil
}
