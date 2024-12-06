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
	titleNo := "title/403d/rix4/q"

	t.Run("create_title", func(t *testing.T) {
		args := sql.CreateTitleParams{
			Url:          docUri,
			Email:        email,
			Title:        titleNo,
			SupportDocID: supportdoc.ID,
		}
		title, err = tc.CreateTitle(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, title.URL, docUri)
	})

	t.Run("update_title_verification", func(t *testing.T) {
		args := sql.UpdateTitleVerificationByIdParams{
			ID:           title.ID,
			Verification: model.VerificationVerified.String(),
		}
		title, err := tc.UpdateTitleVerificationById(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, title.Verified, model.VerificationVerified)
	})

	t.Run("get_titles_by_email", func(t *testing.T) {
		titles, err := tc.GetTitlesByEmail(ctx, email)

		assert.Nil(t, err)
		assert.True(t, len(titles) > 0)
	})
}
