package controller

import (
	"context"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/email"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
	"github.com/resend/resend-go/v2"
	"github.com/sirupsen/logrus"
)

var onboardingController *Onboarding

type OnboardingController interface {
	CreateOnboarding(context.Context, model.CreateOnboardingInput) (*model.Onboarding, error)
	GetOnboardingByID(context.Context, uuid.UUID) (*model.Onboarding, error)
	GetOnboardingByEmailAndVerification(context.Context, sql.GetOnboardingByEmailAndVerificationParams) (*model.Onboarding, error)
	GetOnboardingsByStatus(context.Context, model.Verification) ([]*model.Onboarding, error)
	UpdateOnboardingVerificationByID(context.Context, sql.UpdateOnboardingVerificationByIDParams) (*model.Onboarding, error)
}

type Onboarding struct {
	r            *repository.Onboarding
	sql          *sql.Queries
	log          *logrus.Logger
	emailService email.Resend
}

func (c *Onboarding) Init(sql *sql.Queries) {
	r := &repository.Onboarding{}
	r.Init(sql)
	c.r = r
	c.sql = sql
	c.log = logger.GetLogger()
	c.emailService = email.GetResendEmailService()
	onboardingController = c
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

	// Comms submission success
	if config.IsProd() || config.IsDev() {
		go func() {
			req := &resend.SendEmailRequest{
				From:    "Chanzu <chanzu@info.copodap.com>",
				To:      []string{input.Email},
				Subject: "Copod- Documents Received",
				Html:    "<p>We have received your documents. We'll update you about their verification status.",
			}
			if err := c.emailService.Send(context.Background(), req); err != nil {
				return
			}
		}()
	}

	return c.r.CreateOnboarding(ctx, oArgs)
}

func (c *Onboarding) GetOnboardingByID(ctx context.Context, id uuid.UUID) (*model.Onboarding, error) {
	return c.r.GetOnboardingByID(ctx, id)
}

func (c *Onboarding) GetOnboardingByEmailAndVerification(ctx context.Context, args sql.GetOnboardingByEmailAndVerificationParams) (*model.Onboarding, error) {
	return c.r.GetOnboardingByEmailAndVerification(ctx, args)
}

func (c *Onboarding) GetOnboardingsByStatus(ctx context.Context, status model.Verification) ([]*model.Onboarding, error) {
	return c.r.GetOnboardingsByStatus(ctx, status)
}

// TODO only verify after all the docs are verified
func (c *Onboarding) UpdateOnboardingVerificationByID(ctx context.Context, args sql.UpdateOnboardingVerificationByIDParams) (*model.Onboarding, error) {
	u, err := c.r.UpdateOnboardingVerificationByID(ctx, args)
	if err != nil {
		return nil, err
	}

	// Comms onboarding status
	switch u.Verification {
	case model.VerificationVerified:
		if config.IsProd() || config.IsDev() {
			go func() {
				req := &resend.SendEmailRequest{
					From:    "Chanzu <chanzu@info.copodap.com>",
					To:      []string{u.Email},
					Subject: "Copod- Documents Verification Status",
					Html:    "<strong>Congratulation!</strong><br><p>Your land is verified and registered to the blockchain. You can now view your land in the app.</p>",
				}
				if err := c.emailService.Send(context.Background(), req); err != nil {
					return
				}
			}()
		}
	}

	return c.r.UpdateOnboardingVerificationByID(ctx, args)
}
