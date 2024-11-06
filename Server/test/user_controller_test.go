package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_User_Controller(t *testing.T) {
	defer func() {
		q.ClearTestUsers(context.Background())
	}()

	t.Run("create_user", func(t *testing.T) {
		uc := controller.UC
		email := RandomEmailAddress()
		wallet := RandomWalletAddress()
		u, err := uc.CreateUser(context.Background(), sql.CreateUserParams{
			Email:         email,
			WalletAddress: wallet,
			GovtID:        RandomGovtID(),
		})

		assert.Equal(t, u.WalletAddress, wallet)
		assert.Equal(t, u.Email, email)
		assert.Nil(t, err)
	})
}
