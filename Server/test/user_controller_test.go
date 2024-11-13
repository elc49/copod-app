package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_User_Controller(t *testing.T) {
	t.Run("create_user", func(t *testing.T) {
		assert.Equal(t, superUser.WalletAddress, superUserWallet)
		assert.Equal(t, superUser.Email, superUserEmail)
	})
}
