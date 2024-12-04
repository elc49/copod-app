package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
)

var supportingDocController *SupportingDoc

type SupportingDocController interface {
	CreateSupportingDoc(context.Context, sql.CreateSupportDocParams) (*model.SupportingDoc, error)
	GetSupportingDocsByVerification(context.Context, model.Verification) ([]*model.SupportingDoc, error)
	GetSupportingDocByID(context.Context, uuid.UUID) (*model.SupportingDoc, error)
	UpdateSupportingDocByEmail(context.Context, sql.UpdateUserSupportDocByEmailParams) (*model.SupportingDoc, error)
}

type SupportingDoc struct {
	r *repository.SupportingDoc
}

func (c *SupportingDoc) Init(sql *sql.Queries) {
	r := &repository.SupportingDoc{}
	r.Init(sql)
	c.r = r
	supportingDocController = c
}

func GetSupportingDocController() SupportingDocController {
	return supportingDocController
}

func (c *SupportingDoc) CreateSupportingDoc(ctx context.Context, args sql.CreateSupportDocParams) (*model.SupportingDoc, error) {
	s, err := c.r.GetSupportDocByEmail(ctx, args.Email)
	switch {
	case s != nil && err == nil:
		switch s.Verified {
		case model.VerificationRejected:
			// update don't recreate
			args := sql.UpdateUserSupportDocByEmailParams{
				Email:        args.Email,
				GovtID:       args.GovtID,
				Verification: model.VerificationOnboarding.String(),
			}
			return c.r.UpdateUserSupportDocByEmail(ctx, args)
		}
		return s, nil
	case err != nil:
		return nil, err
	default:
		return c.r.CreateSupportDoc(ctx, args)
	}
}

func (c *SupportingDoc) GetSupportingDocsByVerification(ctx context.Context, verification model.Verification) ([]*model.SupportingDoc, error) {
	return c.r.GetSupportingDocsByVerification(ctx, verification)
}

func (c *SupportingDoc) GetSupportingDocByID(ctx context.Context, id uuid.UUID) (*model.SupportingDoc, error) {
	return c.r.GetSupportingDocByID(ctx, id)
}

func (c *SupportingDoc) UpdateSupportingDocByEmail(ctx context.Context, args sql.UpdateUserSupportDocByEmailParams) (*model.SupportingDoc, error) {
	return c.r.UpdateUserSupportDocByEmail(ctx, args)
}
