// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sql

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateDisplayPicture(ctx context.Context, arg CreateDisplayPictureParams) (DisplayPicture, error)
	CreateOnboarding(ctx context.Context, arg CreateOnboardingParams) (Onboarding, error)
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error)
	CreateSupportDoc(ctx context.Context, arg CreateSupportDocParams) (SupportDoc, error)
	CreateTitle(ctx context.Context, arg CreateTitleParams) (TitleDeed, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetDisplayPictureByID(ctx context.Context, id uuid.UUID) (DisplayPicture, error)
	GetOnboardingByEmail(ctx context.Context, email string) (Onboarding, error)
	GetOnboardingByID(ctx context.Context, id uuid.UUID) (Onboarding, error)
	GetPaymentByReferenceID(ctx context.Context, referenceID string) (Payment, error)
	GetPaymentOnboardingByID(ctx context.Context, id uuid.UUID) (Onboarding, error)
	GetPaymentsByStatus(ctx context.Context, status string) ([]Payment, error)
	GetSupportDocByEmail(ctx context.Context, email string) (SupportDoc, error)
	GetSupportDocByID(ctx context.Context, id uuid.UUID) (SupportDoc, error)
	GetSupportingDocsByVerification(ctx context.Context, verification string) ([]SupportDoc, error)
	GetTitleByEmail(ctx context.Context, email string) (TitleDeed, error)
	GetTitleByID(ctx context.Context, id uuid.UUID) (TitleDeed, error)
	GetTitlesByEmailAndVerification(ctx context.Context, arg GetTitlesByEmailAndVerificationParams) ([]TitleDeed, error)
	UpdateDisplayPictureByID(ctx context.Context, arg UpdateDisplayPictureByIDParams) (DisplayPicture, error)
	UpdateDisplayPictureVerificationByID(ctx context.Context, arg UpdateDisplayPictureVerificationByIDParams) (DisplayPicture, error)
	UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (Payment, error)
	UpdateSupportDocByID(ctx context.Context, arg UpdateSupportDocByIDParams) (SupportDoc, error)
	UpdateSupportDocVerificationByID(ctx context.Context, arg UpdateSupportDocVerificationByIDParams) (SupportDoc, error)
	UpdateTitleByID(ctx context.Context, arg UpdateTitleByIDParams) (TitleDeed, error)
	UpdateTitleVerificationByID(ctx context.Context, arg UpdateTitleVerificationByIDParams) (TitleDeed, error)
}

var _ Querier = (*Queries)(nil)
