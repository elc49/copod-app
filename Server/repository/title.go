package repository

import (
	"context"
	db "database/sql"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type Title struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *Title) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *Title) CreateTitle(ctx context.Context, args sql.CreateTitleParams) (*model.Title, error) {
	t, err := r.sql.CreateTitle(ctx, args)
	if err != nil {
		r.log.WithError(err).Errorf("repository: CreateTitle")
		return nil, err
	}

	return &model.Title{
		ID:        t.ID,
		Title:     t.Title,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

func (r *Title) GetTitleByEmail(ctx context.Context, email string) (*model.Title, error) {
	t, err := r.sql.GetTitleByEmail(ctx, email)
	switch {
	case err != nil && err == db.ErrNoRows:
		return nil, nil
	case err != nil:
		r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: GetTitleByEmail")
		return nil, err
	}

	return &model.Title{
		ID:        t.ID,
		Title:     t.Title,
		Verified:  model.Verification(t.Verification),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

func (r *Title) UpdateTitleByEmail(ctx context.Context, args sql.UpdateTitleByEmailParams) (*model.Title, error) {
	u, err := r.sql.UpdateTitleByEmail(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateTitleByEmail")
		return nil, err
	}

	return &model.Title{
		ID:        u.ID,
		Title:     u.Title,
		Verified:  model.Verification(u.Verification),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *Title) UpdateTitleVerificationById(ctx context.Context, args sql.UpdateTitleVerificationByIdParams) (*model.Title, error) {
	u, err := r.sql.UpdateTitleVerificationById(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateTitleVerificationById")
		return nil, err
	}

	return &model.Title{
		ID:        u.ID,
		Verified:  model.Verification(u.Verification),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
