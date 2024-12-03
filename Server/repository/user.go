package repository

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type User struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *User) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *User) CreateUser(ctx context.Context, email string) (*model.User, error) {
	u, err := r.sql.CreateUser(ctx, email)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: CreateUser")
		return nil, err
	}

	return &model.User{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *User) GetUser(ctx context.Context, email string) (*model.User, error) {
	u, err := r.sql.GetUser(ctx, email)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: GetUser")
		return nil, err
	}

	return &model.User{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
