package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/stretchr/testify/assert"
)

func Test_Early_Signup_Controller(t *testing.T) {
	ec := controller.GetEarlySignupController()

	t.Run("create_early_signup", func(t *testing.T) {
		e, err := ec.CreateEarlySignup(context.Background(), email)

		assert.Nil(t, err)
		assert.Equal(t, *e, email)

	})

	t.Run("onboard_early_signup", func(t *testing.T) {
		e, err := q.OnboardEarlySignup(context.Background(), email)

		assert.Nil(t, err)
		assert.True(t, e.Onboarded.Valid)
	})

	t.Run("don't_recreate_early_signup", func(t *testing.T) {
		e, err := ec.CreateEarlySignup(context.Background(), email)

		assert.Nil(t, err)
		assert.Equal(t, *e, email)
	})
}
