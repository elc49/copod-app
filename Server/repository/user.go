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
		r.log.WithError(err).WithFields(logrus.Fields{"user": args}).Errorf("repository: CreateUser")
		return nil, err
	}

	return &model.User{
		ID:            u.ID,
		Email:         u.Email,
		WalletAddress: u.WalletAddress,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
}

func (r *User) GetUser(ctx context.Context, walletAddress string) (*model.User, error) {
	u, err := r.sql.GetUser(ctx, walletAddress)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"walletAddress": walletAddress}).Errorf("repository: GetUser")
		return nil, err
	}

	return &model.User{
		ID:            u.ID,
		Email:         u.Email,
		WalletAddress: u.WalletAddress,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
}
