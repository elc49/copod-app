package repository

import (
	"context"

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
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
	}, nil
}
