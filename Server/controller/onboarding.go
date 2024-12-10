package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

var onboardingController *Onboarding

type OnboardingController interface {
	CreateOnboarding(context.Context, model.CreateOnboardingInput) (*model.Onboarding, error)
	GetOnboardingByVerificationAndPaymentStatus(context.Context, model.Verification, model.PaymentStatus) ([]*model.Onboarding, error)
}

type Onboarding struct {
	r   *repository.Onboarding
	sql *sql.Queries
	log *logrus.Logger
}

func (o *Onboarding) Init(sql *sql.Queries) {
	r := &repository.Onboarding{}
	r.Init(sql)
	o.r = r
	onboardingController = o
	o.sql = sql
	o.log = logger.GetLogger()
}

func GetOnboardingController() OnboardingController {
	return onboardingController
}

func (o *Onboarding) CreateOnboarding(ctx context.Context, input model.CreateOnboardingInput) (*model.Onboarding, error) {
	var supportDoc *model.SupportingDoc
	// Check existing pending onboarding
	getArgs := sql.GetOnboardingByEmailAndVerificationParams{}
	onboarding, oErr := o.r.GetOnboardingByEmailAndVerification(ctx, getArgs)
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
		if s, sErr := o.sql.CreateSupportDoc(ctx, sArgs); sErr != nil {
			o.log.WithError(sErr).WithFields(logrus.Fields{"args": sArgs}).Errorf("controller: CreateOnboarding: CreateSupportDoc")
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
		if t, tErr := o.sql.CreateTitle(ctx, tArgs); tErr != nil {
			o.log.WithError(tErr).WithFields(logrus.Fields{"args": tArgs}).Errorf("controller: CreateOnboarding: CreateTitle")
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
		if dp, dErr := o.sql.CreateDisplayPicture(ctx, dArgs); dErr != nil {
			o.log.WithError(dErr).WithFields(logrus.Fields{"args": dArgs}).Errorf("controller: CreateOnboarding: CreateDisplayPicture")
			return nil, dErr
		} else {
			oArgs.DisplayPictureID = dp.ID
		}

		return o.r.CreateOnboarding(ctx, oArgs)
	case oErr == nil && onboarding != nil:
		// Update existing with new incoming onboarding data
		// Support doc
		suArgs := sql.UpdateSupportDocByIDParams{
			ID:           onboarding.SupportDocID,
			Url:          input.SupportdocURL,
			Verification: model.VerificationOnboarding.String(),
		}
		if _, err := o.sql.UpdateSupportDocByID(ctx, suArgs); err != nil {
			return nil, err
		}

		// Title
		tuArgs := sql.UpdateTitleByIDParams{
			ID:           onboarding.TitleID,
			Url:          input.TitleURL,
			Verification: model.VerificationOnboarding.String(),
		}
		if _, err := o.sql.UpdateTitleByID(ctx, tuArgs); err != nil {
			o.log.WithError(err).WithFields(logrus.Fields{"args": tuArgs}).Errorf("controller: CreateOnboarding: UpdateTitleByID")
			return nil, err
		}

		// Display picture
		udArgs := sql.UpdateDisplayPictureByIDParams{
			ID:           onboarding.DisplayPictureID,
			Url:          input.DisplayPictureURL,
			Verification: model.VerificationOnboarding.String(),
		}
		if _, err := o.sql.UpdateDisplayPictureByID(ctx, udArgs); err != nil {
			o.log.WithError(err).WithFields(logrus.Fields{"args": udArgs}).Errorf("controller: CreateOnboarding: UpdateDisplayPictureByID")
			return nil, err
		}

		return o.r.UpdateOnboardingVerificationByID(ctx, sql.UpdateOnboardingVerificationByIDParams{
			ID:           onboarding.ID,
			Verification: model.VerificationOnboarding.String(),
		})
	default:
		return nil, oErr
	}
}

func (o *Onboarding) GetOnboardingByVerificationAndPaymentStatus(ctx context.Context, verification model.Verification, status model.PaymentStatus) ([]*model.Onboarding, error) {
	return o.r.GetOnboardingByVerificationAndPaymentStatus(ctx, verification, status)
}
