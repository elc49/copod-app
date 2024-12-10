package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Display_Picture_Controller(t *testing.T) {
	ctx := context.Background()
	dc := controller.GetDisplayPictureController()
	avatar := "https://avatar.com/user"
	var displayPicture *model.DisplayPicture
	var err error

	t.Run("create_display_picture", func(t *testing.T) {
		displayPicture, err = dc.CreateDisplayPicture(ctx, sql.CreateDisplayPictureParams{
			Url:          avatar,
			SupportDocID: supportdoc.ID,
			Email:        email,
		})

		assert.Nil(t, err)
		assert.Equal(t, displayPicture.URL, avatar)
	})

	t.Run("get_display_picture_by_id", func(t *testing.T) {
		dp, err := dc.GetDisplayPictureByID(ctx, displayPicture.ID)

		assert.Nil(t, err)
		assert.Equal(t, displayPicture.ID, dp.ID)
	})

	t.Run("update_display_picture_by_id", func(t *testing.T) {
		dp, err := dc.UpdateDisplayPictureByID(ctx, sql.UpdateDisplayPictureByIDParams{
			ID:           displayPicture.ID,
			Url:          docUri,
			Verification: model.VerificationOnboarding.String(),
		})

		assert.Nil(t, err)
		assert.NotEqual(t, dp.URL, displayPicture.URL)
	})
}
