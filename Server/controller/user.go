package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var userController *User

type UserController interface {
	CreateUser(context.Context, string) (*model.User, error)
	GetUser(context.Context, string) (*model.User, error)
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

func (c *User) CreateUser(ctx context.Context, email string) (*model.User, error) {
	return c.r.CreateUser(ctx, email)
}

func (c *User) GetUser(ctx context.Context, email string) (*model.User, error) {
	return c.r.GetUser(ctx, email)
}
