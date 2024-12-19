package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var onboardingController *Onboarding

type OnboardingController interface {
	CreateOnboarding(context.Context, model.CreateOnboardingInput) (*model.Onboarding, error)
	GetOnboardingByID(context.Context, uuid.UUID) (*model.Onboarding, error)
	GetOnboardingByEmail(context.Context, string) (*model.Onboarding, error)
}

type Onboarding struct {
	r   *repository.Onboarding
	sql *sql.Queries
	log *logrus.Logger
}

func (c *Onboarding) Init(sql *sql.Queries) {
	r := &repository.Onboarding{}
	r.Init(sql)
	c.r = r
	onboardingController = c
	c.sql = sql
	c.log = logger.GetLogger()
}

func GetOnboardingController() OnboardingController {
	return onboardingController
}

func (c *Onboarding) CreateOnboarding(ctx context.Context, input model.CreateOnboardingInput) (*model.Onboarding, error) {
	var supportDoc *model.SupportingDoc
	// Check existing pending onboarding
	onboarding, oErr := c.r.GetOnboardingByEmail(ctx, input.Email)
	oArgs := sql.CreateOnboardingParams{
		Email: input.Email,
	}

	switch {
	case oErr == nil && onboarding == nil:
		// Create new onboarding
		// New support doc
		sArgs := sql.CreateSupportDocParams{
			Email: input.Email,
			Url:   input.SupportdocURL,
		}
		if s, sErr := c.sql.CreateSupportDoc(ctx, sArgs); sErr != nil {
			c.log.WithError(sErr).WithFields(logrus.Fields{"args": sArgs}).Errorf("controller: CreateOnboarding: CreateSupportDoc")
			return nil, sErr
		} else {
			oArgs.SupportDocID = s.ID
			supportDoc = &model.SupportingDoc{
				ID: s.ID,
			}
		}

		// New title
		tArgs := sql.CreateTitleParams{
			Url:          input.TitleURL,
			Email:        input.Email,
			SupportDocID: supportDoc.ID,
		}
		if t, tErr := c.sql.CreateTitle(ctx, tArgs); tErr != nil {
			c.log.WithError(tErr).WithFields(logrus.Fields{"args": tArgs}).Errorf("controller: CreateOnboarding: CreateTitle")
			return nil, tErr
		} else {
			oArgs.TitleID = t.ID
		}

		// New display picture
		dArgs := sql.CreateDisplayPictureParams{
			Email:        input.Email,
			SupportDocID: supportDoc.ID,
			Url:          input.DisplayPictureURL,
		}
		if dp, dErr := c.sql.CreateDisplayPicture(ctx, dArgs); dErr != nil {
			c.log.WithError(dErr).WithFields(logrus.Fields{"args": dArgs}).Errorf("controller: CreateOnboarding: CreateDisplayPicture")
			return nil, dErr
		} else {
			oArgs.DisplayPictureID = dp.ID
		}

		return c.r.CreateOnboarding(ctx, oArgs)
	case oErr == nil && onboarding != nil:
		// Update existing with new incoming onboarding data
		// Update Support doc
		suArgs := sql.UpdateSupportDocByIDParams{
			ID:           onboarding.SupportDocID,
			Url:          input.SupportdocURL,
			Verification: model.VerificationOnboarding.String(),
		}
		if _, err := c.sql.UpdateSupportDocByID(ctx, suArgs); err != nil {
			return nil, err
		}

		// Update Title
		tuArgs := sql.UpdateTitleByIDParams{
			ID:           onboarding.TitleID,
			Url:          input.TitleURL,
			Verification: model.VerificationOnboarding.String(),
		}
		if _, err := c.sql.UpdateTitleByID(ctx, tuArgs); err != nil {
			c.log.WithError(err).WithFields(logrus.Fields{"args": tuArgs}).Errorf("controller: CreateOnboarding: UpdateTitleByID")
			return nil, err
		}

		// Update Display picture
		udArgs := sql.UpdateDisplayPictureByIDParams{
			ID:           onboarding.DisplayPictureID,
			Url:          input.DisplayPictureURL,
			Verification: model.VerificationOnboarding.String(),
		}
		if _, err := c.sql.UpdateDisplayPictureByID(ctx, udArgs); err != nil {
			c.log.WithError(err).WithFields(logrus.Fields{"args": udArgs}).Errorf("controller: CreateOnboarding: UpdateDisplayPictureByID")
			return nil, err
		}

		return onboarding, nil
	default:
		return nil, oErr
	}
}

func (c *Onboarding) GetOnboardingByID(ctx context.Context, id uuid.UUID) (*model.Onboarding, error) {
	return c.r.GetOnboardingByID(ctx, id)
}

func (c *Onboarding) GetOnboardingByEmail(ctx context.Context, email string) (*model.Onboarding, error) {
	return c.r.GetOnboardingByEmail(ctx, email)
}
