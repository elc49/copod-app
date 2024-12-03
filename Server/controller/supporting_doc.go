package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var supportingDocController *SupportingDoc

type SupportingDocController interface {
	CreateSupportingDoc(context.Context, sql.CreateSupportDocParams) (*model.SupportingDoc, error)
	GetSupportingDocsByVerification(context.Context, model.Verification) ([]*model.SupportingDoc, error)
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
			return c.r.UpdateSupportDocByEmail(ctx, sql.UpdateSupportDocByEmailParams{
				Email:        args.Email,
				GovtID:       args.GovtID,
				Verification: model.VerificationOnboarding.String(),
			})
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
