package repository

import (
	"context"
	db "database/sql"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type SupportingDoc struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *SupportingDoc) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *SupportingDoc) CreateSupportDoc(ctx context.Context, args sql.CreateSupportDocParams) (*model.SupportingDoc, error) {
	s, err := r.sql.CreateSupportDoc(ctx, args)
	if err != nil {
		r.log.WithError(err).Error("repository: CreateSupportDoc")
		return nil, err
	}

	return &model.SupportingDoc{
		ID:        s.ID,
		GovtID:    s.GovtID,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}, nil
}

func (r *SupportingDoc) GetSupportDocByEmail(ctx context.Context, email string) (*model.SupportingDoc, error) {
	s, err := r.sql.GetSupportDocByEmail(ctx, email)
	switch {
	case err != nil && err == db.ErrNoRows:
		return nil, nil
	case err != nil:
		r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: GetSupportDocByEmail")
		return nil, err
	}

	return &model.SupportingDoc{
		ID:        s.ID,
		GovtID:    s.GovtID,
		Verified:  model.Verification(s.Verification),
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}, nil
}

func (r *SupportingDoc) UpdateSupportDocByEmail(ctx context.Context, args sql.UpdateSupportDocByEmailParams) (*model.SupportingDoc, error) {
	u, err := r.sql.UpdateSupportDocByEmail(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateSupportDocByEmail")
		return nil, err
	}

	return &model.SupportingDoc{
		ID:        u.ID,
		GovtID:    u.GovtID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
