package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Payment_Controller(t *testing.T) {
	ctx := context.Background()
	pc := controller.GetPaymentController()
	oc := controller.GetOnboardingController()
	refId := RandomStringByLength(10)
	var p *model.Payment
	var err error

	t.Run("create_payment", func(t *testing.T) {
		ob, err := oc.CreateOnboarding(ctx, model.CreateOnboardingInput{
			TitleURL:          "https://title.url",
			DisplayPictureURL: "https://dp.url",
			SupportdocURL:     "https://supp.doc",
			Email:             email,
		})
		p, err = pc.CreatePayment(ctx, sql.CreatePaymentParams{
			Email:        email,
			ReferenceID:  refId,
			Status:       "pay_offline",
			Reason:       model.PaymentReasonLandRegistry.String(),
			Amount:       1500,
			Currency:     "KES",
			OnboardingID: uuid.NullUUID{UUID: ob.ID, Valid: true},
		})

		assert.Nil(t, err)
		assert.NotNil(t, p)
	})

	t.Run("get_payment_by_reference_id", func(t *testing.T) {
		p, err = pc.GetPaymentByReferenceID(ctx, p.ReferenceID)

		assert.Nil(t, err)
		assert.Equal(t, p.ReferenceID, refId)
		assert.Equal(t, p.Status, "pay_offline")
	})

	t.Run("update_payment_status", func(t *testing.T) {
		p, err = pc.UpdatePaymentStatus(ctx, sql.UpdatePaymentStatusParams{
			ReferenceID: p.ReferenceID,
			Status:      "success",
		})

		assert.Nil(t, err)
		assert.Equal(t, p.Status, "success")
	})

	t.Run("get_payment_title", func(t *testing.T) {
		title, err := pc.GetPaymentOnboardingByID(ctx, p.OnboardingID)

		assert.Nil(t, err)
		assert.NotNil(t, title)
		assert.Equal(t, title.ID, p.OnboardingID)
	})

	t.Run("get_payment_by_status", func(t *testing.T) {
		payments, err := pc.GetPaymentsByStatus(ctx, model.PaymentStatusSuccess.String())

		assert.Nil(t, err)
		assert.True(t, len(payments) > 0)
	})
}
