package test

import (
	"context"
	db "database/sql"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Onboarding_Controller(t *testing.T) {
	ctx := context.Background()
	oc := controller.GetOnboardingController()

	t.Run("create_new_onboarding", func(t *testing.T) {
		o, err := oc.CreateOnboarding(ctx, model.CreateOnboardingInput{
			TitleURL:          "https://title.url",
			DisplayPictureURL: "https://dp.url",
			SupportdocURL:     "https://supp.doc",
			Email:             email,
		})

		assert.Nil(t, err)
		assert.Equal(t, o.Verification, model.VerificationOnboarding)
	})

	t.Run("update_existing_onboarding", func(t *testing.T) {
		o, err := oc.CreateOnboarding(ctx, model.CreateOnboardingInput{
			TitleURL:          "https://title.url",
			DisplayPictureURL: "https://dp.url",
			SupportdocURL:     "https://supp.doc",
			Email:             email,
		})

		assert.Nil(t, err)
		assert.NotNil(t, o)
	})

	t.Run("get_onboardings_by_verification_and_payment_status", func(t *testing.T) {
		obs, err := oc.GetOnboardingByVerificationAndPaymentStatus(ctx, sql.GetOnboardingByVerificationAndPaymentStatusParams{
			Verification:  model.VerificationOnboarding.String(),
			PaymentStatus: db.NullString{String: model.PaymentStatusSuccess.String(), Valid: true},
		})

		assert.Nil(t, err)
		assert.True(t, len(obs) == 0)
	})
}
