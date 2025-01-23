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

var displayPictureController DisplayPictureController

type DisplayPictureController interface {
	CreateDisplayPicture(context.Context, sql.CreateDisplayPictureParams) (*model.DisplayPicture, error)
	GetDisplayPictureByID(context.Context, uuid.UUID) (*model.DisplayPicture, error)
	UpdateDisplayPictureByID(context.Context, sql.UpdateDisplayPictureByIDParams) (*model.DisplayPicture, error)
	UpdateDisplayPictureVerificationByID(context.Context, string, sql.UpdateDisplayPictureVerificationByIDParams) (*model.DisplayPicture, error)
}

type DisplayPicture struct {
	r            *repository.DisplayPicture
	emailService email.Resend
}

func (c *DisplayPicture) Init(sql *sql.Queries) {
	r := &repository.DisplayPicture{}
	r.Init(sql)
	c.r = r
	c.emailService = email.GetResendEmailService()
	displayPictureController = c
}

func GetDisplayPictureController() DisplayPictureController {
	return displayPictureController
}

func (c *DisplayPicture) CreateDisplayPicture(ctx context.Context, args sql.CreateDisplayPictureParams) (*model.DisplayPicture, error) {
	return c.r.CreateDisplayPicture(ctx, args)
}

func (c *DisplayPicture) GetDisplayPictureByID(ctx context.Context, id uuid.UUID) (*model.DisplayPicture, error) {
	return c.r.GetDisplayPictureByID(ctx, id)
}

func (c *DisplayPicture) UpdateDisplayPictureByID(ctx context.Context, args sql.UpdateDisplayPictureByIDParams) (*model.DisplayPicture, error) {
	return c.r.UpdateDisplayPictureByID(ctx, args)
}

func (c *DisplayPicture) UpdateDisplayPictureVerificationByID(ctx context.Context, email string, args sql.UpdateDisplayPictureVerificationByIDParams) (*model.DisplayPicture, error) {
	u, err := c.r.UpdateDisplayPictureVerificationByID(ctx, args)
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
					Subject: "Copod- Document Verification Status",
					Html:    "Display picture document submitted could not be verified or is not valid. Please re-submit again using the app.</p>",
				}
				if err := c.emailService.Send(context.Background(), req); err != nil {
					return
				}
			}
		}()
	}

	return u, nil
}
