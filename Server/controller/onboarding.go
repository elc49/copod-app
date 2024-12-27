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
	GetOnboardingByEmailAndVerification(context.Context, sql.GetOnboardingByEmailAndVerificationParams) (*model.Onboarding, error)
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
	oArgs := sql.CreateOnboardingParams{
		Email: input.Email,
	}
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
}

func (c *Onboarding) GetOnboardingByID(ctx context.Context, id uuid.UUID) (*model.Onboarding, error) {
	return c.r.GetOnboardingByID(ctx, id)
}

func (c *Onboarding) GetOnboardingByEmailAndVerification(ctx context.Context, args sql.GetOnboardingByEmailAndVerificationParams) (*model.Onboarding, error) {
	// Get onboarding
	o, err := c.r.GetOnboardingByEmailAndVerification(ctx, args)
	if err == nil && o == nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	// Check if docs are verified
	if c.areDocsOnboarded(ctx, o) {
		// Finish onboarding
		_, err := c.sql.UpdateOnboardingVerificationByID(ctx, sql.UpdateOnboardingVerificationByIDParams{
			ID:           o.ID,
			Verification: model.VerificationVerified.String(),
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return o, nil
}

func (c *Onboarding) areDocsOnboarded(ctx context.Context, o *model.Onboarding) bool {
	// Check if title is verified
	if o.Title.Verified != model.VerificationVerified {
		return false
	}
	// Check if display picture is verified
	if o.DisplayPicture.Verified != model.VerificationVerified {
		return false
	}
	// Check if supporting doc is verified
	if o.SupportingDoc.Verified != model.VerificationVerified {
		return false
	}
	return true
}
