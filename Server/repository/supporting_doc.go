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

func (r *SupportingDoc) UpdateUserSupportDocByEmail(ctx context.Context, args sql.UpdateUserSupportDocByEmailParams) (*model.SupportingDoc, error) {
	u, err := r.sql.UpdateUserSupportDocByEmail(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateUserSupportDocByEmail")
		return nil, err
	}

	return &model.SupportingDoc{
		ID:        u.ID,
		GovtID:    u.GovtID,
		Verified:  model.Verification(u.Verification),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *SupportingDoc) GetSupportingDocsByVerification(ctx context.Context, verification model.Verification) ([]*model.SupportingDoc, error) {
	var docs []*model.SupportingDoc
	d, err := r.sql.GetSupportingDocsByVerification(ctx, verification.String())
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"verification": verification}).Errorf("repository: GetSupportingDocsByVerification")
		return nil, err
	}

	for _, item := range d {
		doc := &model.SupportingDoc{
			ID:        item.ID,
			Email:     item.Email,
			GovtID:    item.GovtID,
			Verified:  model.Verification(item.Verification),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		docs = append(docs, doc)
	}

	return docs, nil
}

func (r *SupportingDoc) GetSupportingDocByID(ctx context.Context, id uuid.UUID) (*model.SupportingDoc, error) {
	d, err := r.sql.GetSupportingDocById(ctx, id)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"id": id}).Errorf("repository: GetSupportingDocById")
		return nil, err
	}

	return &model.SupportingDoc{
		ID:        d.ID,
		GovtID:    d.GovtID,
		Verified:  model.Verification(d.Verification),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}, nil
}
