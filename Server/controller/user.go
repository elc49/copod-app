package controller

import (
	"context"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/email"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/resend/resend-go/v2"
	"github.com/sirupsen/logrus"
)

var userController *User

type UserController interface {
	CreateUser(context.Context, sql.CreateUserParams) (*model.User, error)
}

type User struct {
	r   *repository.User
	log *logrus.Logger
}

func (c *User) Init(sql *sql.Queries) {
	r := &repository.User{}
	r.Init(sql)
	c.r = r
	userController = c
	c.log = logger.GetLogger()
}

func GetUserController() UserController {
	return userController
}

func (c *User) CreateUser(ctx context.Context, args sql.CreateUserParams) (*model.User, error) {
	// Check if user exists
	u, err := c.r.GetUserByEmail(ctx, args.Email)
	switch {
	case u == nil && err == nil:
		// Create new user and email onboard
		newU, err := c.r.CreateUser(ctx, args)
		if err != nil {
			return nil, err
		}
		// Email onboard - this will run only in production
		if config.IsProd() {
			go func() {
				if err := c.emailOnboard(context.Background(), *newU); err != nil {
					c.log.WithError(err).WithFields(logrus.Fields{"user": *newU}).Errorf("controller: goroutine: emailOnboard")
					return
				}
			}()
		}

		return newU, nil
	case err != nil:
		return nil, err
	default:
		return u, nil
	}
}

func (c *User) emailOnboard(ctx context.Context, user model.User) error {
	params := &resend.SendEmailRequest{
		From:    "Chanzu <chanzu@info.copodap.com>",
		To:      []string{user.Email},
		Subject: "Account created",
		Html:    "<p>Hello,</p><p>Welcome to Copod!</p><p>This the first step towards the future of land ownership.</p><p>I can't wait to share with you exciting features/updates.</p><br><strong>Best,</strong><p>Edwin Chanzu.</p>",
	}
	if err := email.GetResendEmailService().Send(ctx, params); err != nil {
		return err
	}

	// Update user email onboard async
	args := sql.UpdateUserEmailOnboardByIDParams{
		ID:             user.ID,
		EmailOnboarded: true,
	}
	_, err := c.r.UpdateUserEmailOnboardByID(ctx, args)
	if err != nil {
		c.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("controller: UpdateUserEmailOnboardByID: emailOnboard")
		return err
	}

	return nil
}
