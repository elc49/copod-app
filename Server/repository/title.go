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
		URL:       t.Url,
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
		URL:       t.Url,
		Verified:  model.Verification(t.Verification),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

func (r *Title) UpdateTitleByID(ctx context.Context, args sql.UpdateTitleByIDParams) (*model.Title, error) {
	u, err := r.sql.UpdateTitleByID(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateTitleByID")
		return nil, err
	}

	return &model.Title{
		ID:        u.ID,
		URL:       u.Url,
		Verified:  model.Verification(u.Verification),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *Title) GetTitlesByEmailAndVerification(ctx context.Context, args sql.GetTitlesByEmailAndVerificationParams) ([]*model.Title, error) {
	var titles []*model.Title
	t, err := r.sql.GetTitlesByEmailAndVerification(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: GetTitlesByEmailAndVerification")
		return nil, err
	}

	for _, item := range t {
		title := &model.Title{
			ID:        item.ID,
			URL:       item.Url,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		titles = append(titles, title)
	}

	return titles, nil
}

func (r *Title) GetTitleByID(ctx context.Context, id uuid.UUID) (*model.Title, error) {
	t, err := r.sql.GetTitleByID(ctx, id)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"id": id}).Errorf("repository: GetTitleByID")
		return nil, err
	}

	return &model.Title{
		ID:        t.ID,
		URL:       t.Url,
		Verified:  model.Verification(t.Verification),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

func (r *Title) UpdateTitleVerificationByID(ctx context.Context, args sql.UpdateTitleVerificationByIDParams) (*model.Title, error) {
	t, err := r.sql.UpdateTitleVerificationByID(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateTitleVerificationByID")
		return nil, err
	}

	return &model.Title{
		ID:        t.ID,
		URL:       t.Url,
		Verified:  model.Verification(t.Verification),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}
