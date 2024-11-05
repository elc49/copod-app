package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_User_Controller(t *testing.T) {
	t.Run("create_user", func(t *testing.T) {
		uc := controller.UC
		u, err := uc.CreateUser(context.Background(), sql.CreateUserParams{
			Email:         email,
			WalletAddress: walletAddress,
		})

		assert.Equal(t, u.Email, email)
		assert.Equal(t, u.WalletAddress, walletAddress)
		assert.Nil(t, err)
	})
}
