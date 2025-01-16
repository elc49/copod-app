package test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/elc49/copod/controller"
	db "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_User_Controller(t *testing.T) {
	t.Run("create_user", func(t *testing.T) {
		assert.Equal(t, superUser.Email, email)
	})

	t.Run("don't_recreate_existing_user", func(t *testing.T) {
		u, err := controller.GetUserController().CreateUser(context.Background(), db.CreateUserParams{
			Email:     email,
			Firstname: "John",
			Lastname:  "Doe",
		})
		assert.Nil(t, err)

		c, cErr := q.CountUsers(context.Background())
		assert.Nil(t, cErr)

		assert.Equal(t, u.Email, email)
		assert.Equal(t, int(c), 1)
	})

	t.Run("get_user_by_email", func(t *testing.T) {
		u, err := q.GetUserByEmail(context.Background(), email)

		assert.Nil(t, err)
		assert.Equal(t, u.Email, email)
	})

	t.Run("get_not_found_user", func(t *testing.T) {
		_, err := q.GetUserByEmail(context.Background(), "e@email.com")

		assert.True(t, errors.Is(err, sql.ErrNoRows))
	})
}
