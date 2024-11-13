package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Title_Controller(t *testing.T) {
	ctx := context.Background()

	t.Run("create_title", func(t *testing.T) {
		tc := controller.GetTitleController()
		args := sql.CreateTitleParams{
			Title:         docUri,
			Email:         superUserEmail,
			WalletAddress: superUserWallet,
		}
		title, err := tc.CreateTitle(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, title.Title, docUri)
	})
}
