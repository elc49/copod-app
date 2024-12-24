package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Onboarding_Controller(t *testing.T) {
	ctx := context.Background()
	oc := controller.GetOnboardingController()
	var ob *model.Onboarding

	t.Run("create_new_onboarding", func(t *testing.T) {
		o, err := oc.CreateOnboarding(ctx, model.CreateOnboardingInput{
			TitleURL:          "https://title.url",
			DisplayPictureURL: "https://dp.url",
			SupportdocURL:     "https://supp.doc",
			Email:             email,
		})
		ob = o

		assert.Nil(t, err)
		assert.NotNil(t, o)
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

	t.Run("get_onboarding_by_id", func(t *testing.T) {
		o, err := oc.GetOnboardingByID(ctx, ob.ID)

		assert.Nil(t, err)
		assert.Equal(t, ob.ID.String(), o.ID.String())
	})

	t.Run("get_onboarding_by_email", func(t *testing.T) {
		o, err := oc.GetOnboardingByEmailAndVerification(ctx, sql.GetOnboardingByEmailAndVerificationParams{
			Email:        email,
			Verification: model.VerificationOnboarding.String(),
		})

		assert.Nil(t, err)
		assert.Equal(t, email, o.Email)
	})
}
