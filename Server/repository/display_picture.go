package repository

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DisplayPicture struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *DisplayPicture) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *DisplayPicture) CreateDisplayPicture(ctx context.Context, args sql.CreateDisplayPictureParams) (*model.DisplayPicture, error) {
	d, err := r.sql.CreateDisplayPicture(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: CreateDisplayPicture")
		return nil, err
	}

	return &model.DisplayPicture{
		ID:        d.ID,
		Email:     d.Email,
		URL:       d.Url,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}, nil
}

func (r *DisplayPicture) GetDisplayPictureByID(ctx context.Context, id uuid.UUID) (*model.DisplayPicture, error) {
	d, err := r.sql.GetDisplayPictureByID(ctx, id)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"id": id}).Errorf("repository: GetDisplayPictureByID")
		return nil, err
	}

	return &model.DisplayPicture{
		ID:        d.ID,
		Email:     d.Email,
		URL:       d.Url,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}, nil
}

func (r *DisplayPicture) UpdateDisplayPictureByID(ctx context.Context, args sql.UpdateDisplayPictureByIDParams) (*model.DisplayPicture, error) {
	d, err := r.sql.UpdateDisplayPictureByID(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateDisplayPictureByID")
		return nil, err
	}

	return &model.DisplayPicture{
		ID:        d.ID,
		Email:     d.Email,
		URL:       d.Url,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}, nil
}
