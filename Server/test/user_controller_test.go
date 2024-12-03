package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/stretchr/testify/assert"
)

func Test_User_Controller(t *testing.T) {
	uc := controller.GetUserController()
	ctx := context.Background()

	t.Run("create_user", func(t *testing.T) {
		assert.Equal(t, superUser.Email, email)
	})

	t.Run("get_user_by_email", func(t *testing.T) {
		u, e := uc.GetUser(ctx, email)

		assert.Nil(t, e)
		assert.Equal(t, u.Email, email)
	})
}
