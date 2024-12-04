package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_User_Controller(t *testing.T) {
	uc := controller.GetUserController()
	ctx := context.Background()

	t.Run("create_user", func(t *testing.T) {
		assert.Equal(t, superUser.Email, email)
	})

	t.Run("should_update_user_support_doc_verification_status", func(t *testing.T) {
		args := sql.UpdateUserSupportDocByEmailParams{
			GovtID:       "3849#",
			Verification: model.VerificationVerified.String(),
			Email:        email,
		}
		doc, err := uc.UpdateUserSupportDocByEmail(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, doc.Verified, model.VerificationVerified)
	})
}
