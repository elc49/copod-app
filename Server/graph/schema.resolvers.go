package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/elc49/copod/cache"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/paystack"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UploadLandTitle is the resolver for the uploadLandTitle field.
func (r *mutationResolver) UploadLandTitle(ctx context.Context, input model.DocUploadInput) (*model.Title, error) {
	args := sql.CreateTitleParams{
		Email: input.Email,
		Title: input.URL,
	}

	return r.titleController.CreateTitle(ctx, args)
}

// UploadSupportingDoc is the resolver for the uploadSupportingDoc field.
func (r *mutationResolver) UploadSupportingDoc(ctx context.Context, input model.DocUploadInput) (*model.SupportingDoc, error) {
	args := sql.CreateSupportDocParams{
		Email:  input.Email,
		GovtID: input.URL,
	}

	return r.supportDocController.CreateSupportingDoc(ctx, args)
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

// UpdateTitleVerificationByID is the resolver for the updateTitleVerificationById field.
func (r *mutationResolver) UpdateTitleVerificationByID(ctx context.Context, input model.UpdateTitleVerificationInput) (*model.Title, error) {
	args := sql.UpdateTitleVerificationByIdParams{
		ID:           input.ID,
		Verification: input.Verification.String(),
	}
	return r.titleController.UpdateTitleVerificationById(ctx, args)
}

// Title is the resolver for the title field.
func (r *paymentResolver) Title(ctx context.Context, obj *model.Payment) (*model.Title, error) {
	return r.paymentController.GetPaymentTitleByID(ctx, obj.TitleID)
}

// GetLocalLands is the resolver for the getLocalLands field.
func (r *queryResolver) GetLocalLands(ctx context.Context) ([]*model.Land, error) {
	return make([]*model.Land, 0), nil
}

// GetUserLands is the resolver for the getUserLands field.
func (r *queryResolver) GetUserLands(ctx context.Context, email string) ([]*model.Land, error) {
	return make([]*model.Land, 0), nil
}

// GetPaymentsByStatus is the resolver for the getPaymentsByStatus field.
func (r *queryResolver) GetPaymentsByStatus(ctx context.Context, status model.PaymentStatus) ([]*model.Payment, error) {
	return r.paymentController.GetPaymentsByStatus(ctx, status.String())
}

// GetPaymentDetailsByID is the resolver for the getPaymentDetailsById field.
func (r *queryResolver) GetPaymentDetailsByID(ctx context.Context, id uuid.UUID) (*model.Payment, error) {
	return r.paymentController.GetPaymentDetailsByID(ctx, id)
}

// GetSupportingDocsByVerification is the resolver for the getSupportingDocsByVerification field.
func (r *queryResolver) GetSupportingDocsByVerification(ctx context.Context, verification model.Verification) ([]*model.SupportingDoc, error) {
	return r.supportDocController.GetSupportingDocsByVerification(ctx, verification)
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

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Payment returns PaymentResolver implementation.
func (r *Resolver) Payment() PaymentResolver { return &paymentResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type paymentResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
