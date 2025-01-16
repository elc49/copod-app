package repository

import (
	"context"
	db "database/sql"
	"errors"

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

func (r *User) CreateUser(ctx context.Context, args sql.CreateUserParams) (*model.User, error) {
	u, err := r.sql.CreateUser(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: CreateUser")
		return nil, err
	}

	return &model.User{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := r.sql.GetUserByEmail(ctx, email)
	switch {
	case err != nil && errors.Is(err, db.ErrNoRows):
		return nil, nil
	case err != nil:
		r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: GetUserByEmail")
		return nil, err
	}

	return &model.User{
		ID:             u.ID,
		Email:          u.Email,
		EmailOnboarded: u.EmailOnboarded,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}, nil
}

func (r *User) UpdateUserEmailOnboardByID(ctx context.Context, args sql.UpdateUserEmailOnboardByIDParams) (*model.User, error) {
	u, err := r.sql.UpdateUserEmailOnboardByID(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateUserEmailOnboardByID")
		return nil, err
	}

	return &model.User{
		ID:             u.ID,
		Email:          u.Email,
		EmailOnboarded: u.EmailOnboarded,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}, nil
}
