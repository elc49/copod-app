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

func (r *User) UpdateUserSupportDocByEmail(ctx context.Context, args sql.UpdateUserSupportDocByEmailParams) (*model.SupportingDoc, error) {
	u, err := r.sql.UpdateUserSupportDocByEmail(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"args": args}).Errorf("repository: UpdateUserSupportDocByEmail")
		return nil, err
	}

	return &model.SupportingDoc{
		ID:        u.ID,
		GovtID:    u.GovtID,
		Verified:  model.Verification(u.Verification),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
