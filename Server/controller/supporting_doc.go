package controller

import (
	"context"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/email"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
	"github.com/resend/resend-go/v2"
)

var supportingDocController SupportingDocController

type SupportingDocController interface {
	CreateSupportingDoc(context.Context, sql.CreateSupportDocParams) (*model.SupportingDoc, error)
	GetSupportingDocsByVerification(context.Context, model.Verification) ([]*model.SupportingDoc, error)
	GetSupportingDocByID(context.Context, uuid.UUID) (*model.SupportingDoc, error)
	UpdateSupportDocByID(context.Context, sql.UpdateSupportDocByIDParams) (*model.SupportingDoc, error)
	UpdateSupportDocVerificationByID(context.Context, string, sql.UpdateSupportDocVerificationByIDParams) (*model.SupportingDoc, error)
}

type SupportingDoc struct {
	r            *repository.SupportingDoc
	emailService email.Resend
}

func (c *SupportingDoc) Init(sql *sql.Queries) {
	r := &repository.SupportingDoc{}
	r.Init(sql)
	c.r = r
	c.emailService = email.GetResendEmailService()
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
			args := sql.UpdateSupportDocByIDParams{
				ID:           s.ID,
				Verification: model.VerificationOnboarding.String(),
			}
			return c.r.UpdateSupportDocByID(ctx, args)
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

func (c *SupportingDoc) UpdateSupportDocByID(ctx context.Context, args sql.UpdateSupportDocByIDParams) (*model.SupportingDoc, error) {
	return c.r.UpdateSupportDocByID(ctx, args)
}

func (c *SupportingDoc) UpdateSupportDocVerificationByID(ctx context.Context, email string, args sql.UpdateSupportDocVerificationByIDParams) (*model.SupportingDoc, error) {
	u, err := c.r.UpdateSupportDocVerificationByID(ctx, args)
	if err != nil {
		return nil, err
	}

	// Comms rejected doc
	if config.IsProd() || config.IsDev() {
		go func() {
			switch u.Verified {
			case model.VerificationRejected:
				req := &resend.SendEmailRequest{
					From:    "Chanzu <chanzu@info.copodap.com>",
					To:      []string{email},
					Subject: "Copod- Documents Verification Status",
					Html:    "<p>Land title supporting document could not be verified or is not valid. Please re-submit again using the app.<p>",
				}
				if err := c.emailService.Send(context.Background(), req); err != nil {
					return
				}
			}
		}()
	}

	return u, nil
}
