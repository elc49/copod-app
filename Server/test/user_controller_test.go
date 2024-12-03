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
	var u *model.User
	var e error

	t.Run("create_user", func(t *testing.T) {
		assert.Equal(t, superUser.Email, email)
	})

	t.Run("get_user_by_email", func(t *testing.T) {
		u, e = uc.GetUser(ctx, email)

		assert.Nil(t, e)
		assert.Equal(t, u.Email, email)
	})

	t.Run("update_user_details", func(t *testing.T) {
		args := sql.UpdateUserByEmailParams{
			Email:     email,
			Firstname: "John",
			Lastname:  "Doe",
			GovtID:    "3209",
		}
		u, e = uc.UpdateUserByEmail(ctx, args)

		assert.Nil(t, e)
		assert.Equal(t, u.GovtID, "3209")
	})
}
