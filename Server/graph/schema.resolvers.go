package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"time"

	"github.com/elc49/copod/cache"
	"github.com/elc49/copod/contracts"
	"github.com/elc49/copod/contracts/land"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/paystack"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Registration is the resolver for the registration field.
func (r *landDetailsResolver) Registration(ctx context.Context, obj *land.LandDetails) (string, error) {
	t := time.UnixMilli(obj.Registration.Int64())
	dateString := t.Format("January 2, 2006")
	return dateString, nil
}

// Size is the resolver for the size field.
func (r *landDetailsResolver) Size(ctx context.Context, obj *land.LandDetails) (int, error) {
	v := obj.Size.Int64()
	return int(v), nil
}

// ChargeMpesa is the resolver for the chargeMpesa field.
func (r *mutationResolver) ChargeMpesa(ctx context.Context, input model.PayWithMpesaInput) (*string, error) {
	charge := paystack.MpesaCharge{
		Email:    input.Email,
		Reason:   input.Reason.String(),
		Currency: input.Currency,
	}

	res, err := r.paystack.ChargeMpesa(ctx, input.PaymentFor, charge)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"charge": charge}).Errorf("graph resolvers: ChargeMpesa")
		return nil, err
	}

	return &res.Data.Reference, nil
}

// CreateOnboarding is the resolver for the createOnboarding field.
func (r *mutationResolver) CreateOnboarding(ctx context.Context, input model.CreateOnboardingInput) (*model.Onboarding, error) {
	return r.onboardingController.CreateOnboarding(ctx, input)
}

// UpdateTitleVerificationByID is the resolver for the updateTitleVerificationById field.
func (r *mutationResolver) UpdateTitleVerificationByID(ctx context.Context, input model.UpdateTitleVerificationByIDInput) (*model.Title, error) {
	args := sql.UpdateTitleVerificationByIDParams{
		ID:           input.TitleID,
		Verification: input.Verification.String(),
	}
	return r.titleController.UpdateTitleVerificationByID(ctx, args)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	args := sql.CreateUserParams{
		Email:     input.Email,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
	}
	_, err := r.supportDocController.UpdateSupportDocVerificationByID(ctx, sql.UpdateSupportDocVerificationByIDParams{
		ID:           input.SupportDocID,
		Verification: input.SupportDocVerification.String(),
	})
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"input": input}).Errorf("resolver: UpdateSupportDocVerificationByID")
		return nil, err
	}
	return r.userController.CreateUser(ctx, args)
}

// UpdateDisplayPictureVerificationByID is the resolver for the updateDisplayPictureVerificationById field.
func (r *mutationResolver) UpdateDisplayPictureVerificationByID(ctx context.Context, input model.UpdateDisplayPictureVerificationByIDInput) (*model.DisplayPicture, error) {
	args := sql.UpdateDisplayPictureVerificationByIDParams{
		ID:           input.DisplayPictureID,
		Verification: input.Verification.String(),
	}
	return r.displayPictureController.UpdateDisplayPictureVerificationByID(ctx, args)
}

// Title is the resolver for the title field.
func (r *onboardingResolver) Title(ctx context.Context, obj *model.Onboarding) (*model.Title, error) {
	return r.titleController.GetTitleByID(ctx, obj.TitleID)
}

// SupportingDoc is the resolver for the supportingDoc field.
func (r *onboardingResolver) SupportingDoc(ctx context.Context, obj *model.Onboarding) (*model.SupportingDoc, error) {
	return r.supportDocController.GetSupportingDocByID(ctx, obj.SupportDocID)
}

// DisplayPicture is the resolver for the displayPicture field.
func (r *onboardingResolver) DisplayPicture(ctx context.Context, obj *model.Onboarding) (*model.DisplayPicture, error) {
	return r.displayPictureController.GetDisplayPictureByID(ctx, obj.DisplayPictureID)
}

// Onboarding is the resolver for the onboarding field.
func (r *paymentResolver) Onboarding(ctx context.Context, obj *model.Payment) (*model.Onboarding, error) {
	return r.onboardingController.GetOnboardingByID(ctx, obj.OnboardingID)
}

// GetUserLands is the resolver for the getUserLands field.
func (r *queryResolver) GetUserLands(ctx context.Context, input model.GetUserLandsInput) ([]*model.Title, error) {
	args := sql.GetTitlesByEmailAndVerificationParams{
		Email:        input.Email,
		Verification: input.Verification.String(),
	}
	return r.titleController.GetTitlesByEmailAndVerification(ctx, args)
}

// GetPaymentsByStatus is the resolver for the getPaymentsByStatus field.
func (r *queryResolver) GetPaymentsByStatus(ctx context.Context, status model.PaymentStatus) ([]*model.Payment, error) {
	return r.paymentController.GetPaymentsByStatus(ctx, status.String())
}

// GetTitleByID is the resolver for the getTitleById field.
func (r *queryResolver) GetTitleByID(ctx context.Context, id uuid.UUID) (*model.Title, error) {
	return r.titleController.GetTitleByID(ctx, id)
}

// GetSupportingDocByID is the resolver for the getSupportingDocById field.
func (r *queryResolver) GetSupportingDocByID(ctx context.Context, id uuid.UUID) (*model.SupportingDoc, error) {
	return r.supportDocController.GetSupportingDocByID(ctx, id)
}

// GetDisplayPictureByID is the resolver for the getDisplayPictureById field.
func (r *queryResolver) GetDisplayPictureByID(ctx context.Context, id uuid.UUID) (*model.DisplayPicture, error) {
	return r.displayPictureController.GetDisplayPictureByID(ctx, id)
}

// GetOnboardingByEmailAndVerification is the resolver for the getOnboardingByEmailAndVerification field.
func (r *queryResolver) GetOnboardingByEmailAndVerification(ctx context.Context, input model.GetOnboardingByEmailAndVerificationInput) (*model.Onboarding, error) {
	args := sql.GetOnboardingByEmailAndVerificationParams{
		Email:        input.Email,
		Verification: input.Verification.String(),
	}
	return r.onboardingController.GetOnboardingByEmailAndVerification(ctx, args)
}

// GetIsTitleVerified is the resolver for the getIsTitleVerified field.
func (r *queryResolver) GetIsTitleVerified(ctx context.Context, titleNo string) (bool, error) {
	value, err := r.ethBackend.GetRegistryContract().GetLandERC721Contract(nil, titleNo)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"title_no": titleNo}).Errorf("resolver: GetIsTitleVerified: GetLandERC721Contract")
		return false, err
	}

	if value.String() == contracts.ZERO_ADDRESS {
		return false, nil
	}

	return true, nil
}

// GetLandTitleDetails is the resolver for the getLandTitleDetails field.
func (r *queryResolver) GetLandTitleDetails(ctx context.Context, titleNo string) (*land.LandDetails, error) {
	details, err := r.ethBackend.GetLandTitleDetails(titleNo)
	if err != nil {
		return nil, err
	}

	return &land.LandDetails{
		TitleNo:      details.TitleNo,
		Registration: details.Registration,
		Size:         details.Size,
		Symbol:       details.Symbol,
	}, nil
}

// PaymentUpdate is the resolver for the paymentUpdate field.
func (r *subscriptionResolver) PaymentUpdate(ctx context.Context, email string) (<-chan *model.PaymentUpdate, error) {
	ch := make(chan *model.PaymentUpdate)
	pubsub := r.redis.Subscribe(context.Background(), cache.PAYMENT_UPDATED_CHANNEL)

	go func() {
		for msg := range pubsub.Channel() {
			var result *model.PaymentUpdate
			if err := util.DecodeJson([]byte(msg.Payload), &result); err != nil {
				r.log.WithError(err).WithFields(logrus.Fields{"payload": msg.Payload}).Errorf("resolvers: DecodeJson msg.Payload")
				return
			}

			if result.Email == email {
				ch <- result
			}
		}
	}()

	return ch, nil
}

// LandDetails returns LandDetailsResolver implementation.
func (r *Resolver) LandDetails() LandDetailsResolver { return &landDetailsResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Onboarding returns OnboardingResolver implementation.
func (r *Resolver) Onboarding() OnboardingResolver { return &onboardingResolver{r} }

// Payment returns PaymentResolver implementation.
func (r *Resolver) Payment() PaymentResolver { return &paymentResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type landDetailsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type onboardingResolver struct{ *Resolver }
type paymentResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
