package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var userController *User

type UserController interface {
	CreateUser(context.Context, sql.CreateUserParams) (*model.User, error)
}

type User struct {
	r *repository.User
}

func (c *User) Init(sql *sql.Queries) {
	r := &repository.User{}
	r.Init(sql)
	c.r = r
	userController = c
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
		// TODO email onboard
		return newU, err
	default:
		return u, err
	}
}
