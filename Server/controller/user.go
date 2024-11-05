package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var UC *User

type UserController interface {
	CreateUser(context.Context, sql.CreateUserParams) (*model.User, error)
	GetUser(context.Context, string) (*model.User, error)
}

var _ UserController = (*User)(nil)

type User struct {
	r *repository.User
}

func (c *User) Init(sql *sql.Queries) {
	r := &repository.User{}
	r.Init(sql)
	c.r = r
	UC = c
}

func (c *User) CreateUser(ctx context.Context, args sql.CreateUserParams) (*model.User, error) {
	return c.r.CreateUser(ctx, args)
}

func (c *User) GetUser(ctx context.Context, walletAddress string) (*model.User, error) {
	return c.r.GetUser(ctx, walletAddress)
}
