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
	UpdateSupportDocVerificationById(context.Context, sql.UpdateSupportDocVerificationByIdParams) (*model.SupportingDoc, error)
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
	return c.r.CreateUser(ctx, args)
}

func (c *User) UpdateSupportDocVerificationById(ctx context.Context, args sql.UpdateSupportDocVerificationByIdParams) (*model.SupportingDoc, error) {
	return c.r.UpdateSupportDocVerificationById(ctx, args)
}
