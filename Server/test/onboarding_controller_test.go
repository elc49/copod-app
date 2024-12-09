package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Onboarding_Controller(t *testing.T) {
	ctx := context.Background()
	tc := controller.GetTitleController()
	oc := controller.GetOnboardingController()
	dc := controller.GetDisplayPictureController()
	titleNo := "tuos/490s/39"

	title, err := tc.CreateTitle(ctx, sql.CreateTitleParams{
		Url:          docUri,
		Email:        email,
		Title:        titleNo,
		SupportDocID: supportdoc.ID,
	})
	assert.Nil(t, err)

	display, err := dc.CreateDisplayPicture(ctx, sql.CreateDisplayPictureParams{
		Email:        email,
		Url:          docUri,
		SupportDocID: supportdoc.ID,
	})
	assert.Nil(t, err)

	t.Run("create_onboarding", func(t *testing.T) {
		o, err := oc.CreateOnboarding(ctx, sql.CreateOnboardingParams{
			TitleID:          title.ID,
			DisplayPictureID: display.ID,
			SupportDocID:     supportdoc.ID,
		})

		assert.Nil(t, err)
		assert.Equal(t, title.ID, o.TitleID)
	})
}
