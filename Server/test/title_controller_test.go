package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Title_Controller(t *testing.T) {
	ctx := context.Background()
	var title *model.Title
	var err error
	tc := controller.GetTitleController()

	t.Run("create_title", func(t *testing.T) {
		args := sql.CreateTitleParams{
			Url:          docUri,
			Email:        email,
			SupportDocID: supportdoc.ID,
		}
		title, err = tc.CreateTitle(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, title.URL, docUri)
	})

	t.Run("update_title_verification", func(t *testing.T) {
		args := sql.UpdateTitleByIDParams{
			ID:           title.ID,
			Url:          docUri,
			Verification: model.VerificationVerified.String(),
		}
		title, err := tc.UpdateTitleByID(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, title.Verified, model.VerificationVerified)
	})

	t.Run("get_titles_by_email", func(t *testing.T) {
		titles, err := tc.GetTitlesByEmail(ctx, email)

		assert.Nil(t, err)
		assert.True(t, len(titles) > 0)
	})
}
